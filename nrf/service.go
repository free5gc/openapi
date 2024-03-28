package nrf

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/nrf/NFDiscovery"
	"github.com/free5gc/openapi/nrf/NFManagement"
)

const (
	RetryRegisterNrfDuration = 2 * time.Second
	HBTimerInAdvanceDuration = 500 * time.Millisecond
)

var hbTickers sync.Map // key: nrfUri, value: *heartBeatTicker

type heartBeatTicker struct {
	ticker *time.Ticker
	stopCh chan struct{}
}

func newHBTicker(period int32) *heartBeatTicker {
	duration := time.Duration(period)*time.Second - HBTimerInAdvanceDuration
	return &heartBeatTicker{
		ticker: time.NewTicker(duration),
		stopCh: make(chan struct{}),
	}
}

func (t *heartBeatTicker) reset(period int32) {
	duration := time.Duration(period)*time.Second - HBTimerInAdvanceDuration
	t.ticker.Reset(duration)
}

func (t *heartBeatTicker) stop() {
	t.ticker.Stop()
	close(t.stopCh)
}

func startHBTimer(
	client *NFManagement.APIClient,
	nrfUri, nfInstID, nfType string,
	period int32,
	log *logrus.Entry,
) {
	err := resetHBTimer(nrfUri, period, log)
	if err == nil {
		return
	}
	if period <= 0 {
		return
	}

	// in advance to prevent from NRF expire before undergoing heart beat
	hbTicker := newHBTicker(period)
	hbTickers.Store(nrfUri, hbTicker)
	log.Debugf("new heartBeatTicker: %d", period)

	go func() {
		for {
			select {
			case <-hbTicker.stopCh:
				return
			case <-hbTicker.ticker.C:
				// use new traceID for each HeartBeat to debug easily
				traceCtx, span := otel.GetTracerProvider().Tracer(nfType).Start(
					context.Background(), "HeartBeat")
				defer span.End()
				newLog := log.WithContext(traceCtx)

				err := sendHeartBeat(traceCtx, client, nrfUri, nfInstID, newLog)
				if err != nil {
					newLog.Errorln(err)
				}
			}
		}
	}()
}

func resetHBTimer(nrfUri string, period int32, log *logrus.Entry) error {
	if period < 0 {
		return errors.New("period cannot be negative")
	}

	v, ok := hbTickers.Load(nrfUri)
	if !ok {
		return errors.New("no hbTicker found")
	}
	hbTicker, ok := v.(*heartBeatTicker)
	if !ok {
		return errors.New("Type is not heartBeatTicker")
	}

	if period == 0 {
		hbTickers.Delete(nrfUri)
		hbTicker.stop()
		log.Debugf("stop heartBeatTicker")
		return nil
	}

	hbTicker.reset(period)
	log.Debugf("reset heartBeatTicker: %d", period)
	return nil
}

func sendHeartBeat(
	traceCtx context.Context,
	client *NFManagement.APIClient,
	nrfUri, nfInstID string,
	log *logrus.Entry,
) error {
	var err error
	req := new(NFManagement.UpdateNFInstanceRequest)
	req.SetNfInstanceID(nfInstID)
	patchs := []models.PatchItem{
		{
			Op:    models.PatchOperation_REPLACE,
			Path:  "/nfStatus",
			Value: "REGISTERED",
		},
	}
	req.SetPatchItem(patchs)
	rsp, err := client.NFInstanceIDDocumentApi.UpdateNFInstance(traceCtx, req)
	if err != nil {
		return errors.Wrapf(err, "send heartbeat")
	}

	// NRF response 200 which has full NFProfile and NfInstanceId is mandatory.
	// 200 may have NFProfile.HeartBeatTimer=0 to disable heartbeat
	if rsp.NrfNfManagementNfProfile.NfInstanceId != "" {
		err = resetHBTimer(nrfUri, rsp.NrfNfManagementNfProfile.HeartBeatTimer, log)
		if err != nil {
			return errors.Wrapf(err, "set heartbeat")
		}
	}

	return nil
}

func RegisterNFInstance(
	traceCtx context.Context,
	client *NFManagement.APIClient,
	nrfUri string,
	nfProfile *models.NrfNfManagementNfProfile,
	log *logrus.Entry,
) error {
	if client == nil || nfProfile == nil || log == nil {
		return errors.Errorf("client/nfProfile/log is nil; Skip RegisterNFInstance")
	}

	req := &NFManagement.RegisterNFInstanceRequest{}
	req.SetNfInstanceID(nfProfile.NfInstanceId)
	req.SetNrfNfManagementNfProfile(*nfProfile)
	for {
		select {
		case <-traceCtx.Done():
			return errors.Errorf("Stop registration to NRF")
		default:
		}

		rsp, err := client.NFInstanceIDDocumentApi.RegisterNFInstance(traceCtx, req)
		if err != nil {
			apiError := openapi.GenericOpenAPIError{}
			if errors.As(err, &apiError) {
				errModels := apiError.ErrorModel.(NFManagement.RegisterNFInstanceError)
				switch apiError.ErrorStatus {
				case http.StatusTemporaryRedirect:
					return errors.Errorf(
						"RegisterNFInstance: redirect service cause by %s",
						errModels.RedirectResponse.Cause,
					)
				case http.StatusPermanentRedirect:
					return errors.Errorf(
						"RegisterNFInstance: redirect service cause by %s",
						errModels.RedirectResponse.Cause,
					)
				default:
					return errors.Errorf(
						"RegisterNFInstance: failed: problem[%#v]",
						errModels.ProblemDetails)
				}
			} else {
				log.Infof("Register to NRF Error[%v], sleep 2s and retry", err)
				time.Sleep(RetryRegisterNrfDuration)
				continue
			}
		}

		startHBTimer(client, nrfUri, nfProfile.NfInstanceId, string(nfProfile.NfType),
			rsp.NrfNfManagementNfProfile.HeartBeatTimer, log)

		if rsp.Location != "" {
			log.Infof("NFRegister Created")
			break
		} else {
			log.Infof("NFRegister Update")
			break
		}
	}
	return nil
}

func DeregisterNFInstance(
	traceCtx context.Context,
	client *NFManagement.APIClient,
	nrfUri, nfInstID string,
	log *logrus.Entry,
) error {
	if nrfUri == "" {
		log.Infof("No NRF URI; Skip DeregisterNFInstance")
		return nil
	}
	log.Infof("DeregisterNFInstance")

	v, ok := hbTickers.LoadAndDelete(nrfUri)
	if ok {
		hbTicker, ok := v.(*heartBeatTicker)
		if ok {
			hbTicker.stop()
		}
	}

	req := &NFManagement.DeregisterNFInstanceRequest{}
	req.SetNfInstanceID(nfInstID)
	_, err := client.NFInstanceIDDocumentApi.DeregisterNFInstance(traceCtx, req)
	if err != nil {
		apiError := &openapi.GenericOpenAPIError{}
		if errors.As(err, &apiError) {
			errModels := apiError.ErrorModel.(NFManagement.DeregisterNFInstanceError)
			switch apiError.ErrorStatus {
			case http.StatusTemporaryRedirect:
				return errors.Errorf(
					"redirect service cause by %s", errModels.RedirectResponse.Cause)
			case http.StatusPermanentRedirect:
				return errors.Errorf(
					"redirect service cause by %s", errModels.RedirectResponse.Cause)
			default:
				return errors.Errorf(
					"failed: problem[%#v]", errModels.ProblemDetails)
			}
		} else {
			return errors.Wrap(err, "deregister failed")
		}
	}
	return nil
}

func SearchNFInstances(
	traceCtx context.Context,
	client *NFDiscovery.APIClient,
	req *NFDiscovery.SearchNFInstancesRequest,
	log *logrus.Entry,
) (
	*models.NrfNfDiscoveryNfProfile, string, error,
) {
	if req == nil {
		return nil, "", errors.Errorf("SearchNFInstancesRequest is nil")
	}

	rsp, err := client.NFInstancesStoreApi.SearchNFInstances(traceCtx, req)
	if err != nil {
		apiError := &openapi.GenericOpenAPIError{}
		if errors.As(err, &apiError) {
			errModels := apiError.ErrorModel.(NFDiscovery.SearchNFInstancesError)
			switch apiError.ErrorStatus {
			case http.StatusTemporaryRedirect:
				return nil, "", errors.Errorf(
					"redirect service cause by %s", errModels.RedirectResponse.Cause)
			case http.StatusPermanentRedirect:
				return nil, "", errors.Errorf(
					"redirect service cause by %s", errModels.RedirectResponse.Cause)
			default:
				return nil, "", errors.Errorf(
					"failed: problem[%#v]", errModels.ProblemDetails)
			}
		} else {
			return nil, "", errors.Wrap(err, "search nf failed")
		}
	}

	nfProf, uri, err := openapi.GetServiceNfProfileAndUri(
		rsp.SearchResult.NfInstances, req.ServiceNames[0])
	if err != nil {
		return nil, "", err
	}
	return nfProf, uri, nil
}
