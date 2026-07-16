package openapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/free5gc/openapi/models"
	"github.com/pkg/errors"
)

const (
	ServiceBaseURI_NNRF_NFM_v1  string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNRF_NFM) + "/v1"
	ServiceBaseURI_NNRF_DISC_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNRF_DISC) + "/v1"
	ServiceBaseURI_NNRF_OAUTH2  string = ""
	ServiceBaseURI_NNRF_OAM_v1  string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNRF_OAM) + "/v1"
	ServiceBaseURI_NNRF_CMI_v1  string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNRF_CMI) + "/v1"
	ServiceBaseURI_NUDM_SDM_v2  string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_SDM) + "/v2"
	ServiceBaseURI_NUDM_UECM_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_UECM) + "/v1"
	ServiceBaseURI_NUDM_UEAU_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_UEAU) + "/v1"
	ServiceBaseURI_NUDM_EE_v1   string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_EE) + "/v1"
	ServiceBaseURI_NUDM_PP_v1   string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_PP) + "/v1"
	// ServiceBaseURI_NUDM_NIDDAU string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_NIDDAU)
	// ServiceBaseURI_NUDM_MT     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_MT)
	ServiceBaseURI_NUDM_OAM_v1            string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_OAM) + "/v1"
	ServiceBaseURI_NUDM_CMI_v1            string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDM_CMI) + "/v1"
	ServiceBaseURI_NAMF_COMM_v1           string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAMF_COMM) + "/v1"
	ServiceBaseURI_NAMF_EVTS_v1           string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAMF_EVTS) + "/v1"
	ServiceBaseURI_NAMF_MT_v1             string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAMF_MT) + "/v1"
	ServiceBaseURI_NAMF_LOC_v1            string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAMF_LOC) + "/v1"
	ServiceBaseURI_NAMF_OAM_v1            string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAMF_OAM) + "/v1"
	ServiceBaseURI_NAMF_CMI_v1            string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAMF_CMI) + "/v1"
	ServiceBaseURI_NSMF_PDUSESSION_v1     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSMF_PDUSESSION) + "/v1"
	ServiceBaseURI_NSMF_EVENT_EXPOSURE_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSMF_EVENT_EXPOSURE) + "/v1"
	// ServiceBaseURI_NSMF_NIDD string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSMF_NIDD)
	ServiceBaseURI_NSMF_OAM_v1                 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSMF_OAM) + "/v1"
	ServiceBaseURI_NSMF_CMI_v1                 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSMF_CMI) + "/v1"
	ServiceBaseURI_NAUSF_AUTH_v1               string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAUSF_AUTH) + "/v1"
	ServiceBaseURI_NAUSF_SORPROTECTION_v1      string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAUSF_SORPROTECTION) + "/v1"
	ServiceBaseURI_NAUSF_UPUPROTECTION_v1      string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAUSF_UPUPROTECTION) + "/v1"
	ServiceBaseURI_NAUSF_OAM_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAUSF_OAM) + "/v1"
	ServiceBaseURI_NAUSF_CMI_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NAUSF_CMI) + "/v1"
	ServiceBaseURI_NNEF_PFDMANAGEMENT_v1       string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNEF_PFDMANAGEMENT) + "/v1"
	ServiceBaseURI_3GPP_AS_SESSION_WITH_QOS_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_3GPP_AS_SESSION_WITH_QOS) + "/v1"
	ServiceBaseURI_3GPP_PFD_MANAGEMENT_v1      string = "/" + string(models.Nrf_NFMgmt_ServiceName_3GPP_PFD_MANAGEMENT) + "/v1"
	ServiceBaseURI_3GPP_TRAFFIC_INFLUENCE_v1   string = "/" + string(models.Nrf_NFMgmt_ServiceName_3GPP_TRAFFIC_INFLUENCE) + "/v1"
	// ServiceBaseURI_NNEF_SMCONTEXT     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNEF_SMCONTEXT)
	// ServiceBaseURI_NNEF_EVENTEXPOSURE string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNEF_EVENTEXPOSURE)
	ServiceBaseURI_NNEF_OAM_v1                 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNEF_OAM) + "/v1"
	ServiceBaseURI_NNEF_CMI_v1                 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNEF_CMI) + "/v1"
	ServiceBaseURI_NPCF_AM_POLICY_CONTROL_v1   string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_AM_POLICY_CONTROL) + "/v1"
	ServiceBaseURI_NPCF_SMPOLICYCONTROL_v1     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_SMPOLICYCONTROL) + "/v1"
	ServiceBaseURI_NPCF_POLICYAUTHORIZATION_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_POLICYAUTHORIZATION) + "/v1"
	ServiceBaseURI_NPCF_BDTPOLICYCONTROL_v1    string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_BDTPOLICYCONTROL) + "/v1"
	// ServiceBaseURI_NPCF_EVENTEXPOSURE string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_EVENTEXPOSURE)
	ServiceBaseURI_NPCF_UE_POLICY_CONTROL_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_UE_POLICY_CONTROL) + "/v1"
	ServiceBaseURI_NPCF_OAM_v1               string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_OAM) + "/v1"
	ServiceBaseURI_NPCF_CMI_v1               string = "/" + string(models.Nrf_NFMgmt_ServiceName_NPCF_CMI) + "/v1"
	// ServiceBaseURI_NSMSF_SMS string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSMSF_SMS)
	ServiceBaseURI_NNSSF_NSSELECTION_v2       string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNSSF_NSSELECTION) + "/v2"
	ServiceBaseURI_NNSSF_NSSAIAVAILABILITY_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNSSF_NSSAIAVAILABILITY) + "/v1"
	ServiceBaseURI_NNSSF_OAM_v1               string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNSSF_OAM) + "/v1"
	ServiceBaseURI_NNSSF_CMI_v1               string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNSSF_CMI) + "/v1"
	ServiceBaseURI_NUDR_DR_v2                 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDR_DR) + "/v2"
	ServiceBaseURI_NUDR_IMS_DR_v1             string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDR_IMS_DR) + "/v1"
	ServiceBaseURI_NUDR_GROUP_ID_MAP_v1       string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDR_GROUP_ID_MAP) + "/v1"
	ServiceBaseURI_NUDR_OAM_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDR_OAM) + "/v1"
	ServiceBaseURI_NUDR_CMI_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDR_CMI) + "/v1"
	ServiceBaseURI_NLMF_LOC_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NLMF_LOC) + "/v1"
	ServiceBaseURI_NLMF_BROADCAST_v1          string = "/" + string(models.Nrf_NFMgmt_ServiceName_NLMF_BROADCAST) + "/v1"
	ServiceBaseURI_NLMF_OAM_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NLMF_OAM) + "/v1"
	ServiceBaseURI_NLMF_CMI_v1                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NLMF_CMI) + "/v1"
	// ServiceBaseURI_N5G_EIR_EIC                  string = "/" + string(models.Nrf_NFMgmt_ServiceName_N5G_EIR_EIC)
	// ServiceBaseURI_NBSF_MANAGEMENT              string = "/" + string(models.Nrf_NFMgmt_ServiceName_NBSF_MANAGEMENT)
	// ServiceBaseURI_NCHF_SPENDINGLIMITCONTROL    string = "/" + string(models.Nrf_NFMgmt_ServiceName_NCHF_SPENDINGLIMITCONTROL)
	// ServiceBaseURI_NCHF_CONVERGEDCHARGING       string = "/" + string(models.Nrf_NFMgmt_ServiceName_NCHF_CONVERGEDCHARGING)
	// ServiceBaseURI_NCHF_OFFLINEONLYCHARGING     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NCHF_OFFLINEONLYCHARGING)
	// ServiceBaseURI_NNWDAF_EVENTSSUBSCRIPTION    string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNWDAF_EVENTSSUBSCRIPTION)
	// ServiceBaseURI_NNWDAF_ANALYTICSINFO         string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNWDAF_ANALYTICSINFO)
	// ServiceBaseURI_NGMLC_LOC                    string = "/" + string(models.Nrf_NFMgmt_ServiceName_NGMLC_LOC)
	// ServiceBaseURI_NUCMF_PROVISIONING           string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUCMF_PROVISIONING)
	// ServiceBaseURI_NUCMF_UECAPABILITYMANAGEMENT string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUCMF_UECAPABILITYMANAGEMENT)
	// ServiceBaseURI_NHSS_SDM                     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_SDM)
	// ServiceBaseURI_NHSS_UECM                    string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_UECM)
	// ServiceBaseURI_NHSS_UEAU                    string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_UEAU)
	// ServiceBaseURI_NHSS_EE                      string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_EE)
	ServiceBaseURI_NHSS_IMS_SDM_v1  string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_IMS_SDM) + "/v1"
	ServiceBaseURI_NHSS_IMS_UECM_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_IMS_UECM) + "/v1"
	ServiceBaseURI_NHSS_IMS_UEAU_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NHSS_IMS_UEAU) + "/v1"
	// ServiceBaseURI_NSEPP_TELESCOPIC             string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSEPP_TELESCOPIC)
	// ServiceBaseURI_NSORAF_SOR                   string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSORAF_SOR)
	// ServiceBaseURI_NSPAF_SECURED_PACKET         string = "/" + string(models.Nrf_NFMgmt_ServiceName_NSPAF_SECURED_PACKET)
	// ServiceBaseURI_NUDSF_DR                     string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUDSF_DR)
	// ServiceBaseURI_NNSSAAF_NSSAA                string = "/" + string(models.Nrf_NFMgmt_ServiceName_NNSSAAF_NSSAA)
	ServiceBaseURI_NUPF_OAM_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUPF_OAM) + "/v1"
	ServiceBaseURI_NUPF_CMI_v1 string = "/" + string(models.Nrf_NFMgmt_ServiceName_NUPF_CMI) + "/v1"
)

var ServiceNfType map[models.Nrf_NFMgmt_ServiceName]models.Nrf_NFMgmt_NFType

func init() {
	ServiceNfType = make(map[models.Nrf_NFMgmt_ServiceName]models.Nrf_NFMgmt_NFType)
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNRF_NFM] = models.Nrf_NFMgmt_NFType_NRF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNRF_DISC] = models.Nrf_NFMgmt_NFType_NRF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNRF_OAUTH2] = models.Nrf_NFMgmt_NFType_NRF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNRF_OAM] = models.Nrf_NFMgmt_NFType_NRF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNRF_CMI] = models.Nrf_NFMgmt_NFType_NRF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_SDM] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_UECM] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_UEAU] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_EE] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_PP] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_NIDDAU] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_MT] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_OAM] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDM_CMI] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAMF_COMM] = models.Nrf_NFMgmt_NFType_AMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAMF_EVTS] = models.Nrf_NFMgmt_NFType_AMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAMF_MT] = models.Nrf_NFMgmt_NFType_AMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAMF_LOC] = models.Nrf_NFMgmt_NFType_AMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAMF_OAM] = models.Nrf_NFMgmt_NFType_AMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAMF_CMI] = models.Nrf_NFMgmt_NFType_AMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NSMF_PDUSESSION] = models.Nrf_NFMgmt_NFType_SMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NSMF_EVENT_EXPOSURE] = models.Nrf_NFMgmt_NFType_SMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NSMF_NIDD] = models.Nrf_NFMgmt_NFType_SMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NSMF_OAM] = models.Nrf_NFMgmt_NFType_SMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NSMF_CMI] = models.Nrf_NFMgmt_NFType_SMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAUSF_AUTH] = models.Nrf_NFMgmt_NFType_AUSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAUSF_SORPROTECTION] = models.Nrf_NFMgmt_NFType_AUSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAUSF_UPUPROTECTION] = models.Nrf_NFMgmt_NFType_AUSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAUSF_OAM] = models.Nrf_NFMgmt_NFType_AUSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NAUSF_CMI] = models.Nrf_NFMgmt_NFType_AUSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNEF_PFDMANAGEMENT] = models.Nrf_NFMgmt_NFType_NEF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNEF_SMCONTEXT] = models.Nrf_NFMgmt_NFType_NEF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNEF_EVENTEXPOSURE] = models.Nrf_NFMgmt_NFType_NEF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNEF_OAM] = models.Nrf_NFMgmt_NFType_NEF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNEF_CMI] = models.Nrf_NFMgmt_NFType_NEF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_AM_POLICY_CONTROL] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_SMPOLICYCONTROL] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_POLICYAUTHORIZATION] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_BDTPOLICYCONTROL] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_EVENTEXPOSURE] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_UE_POLICY_CONTROL] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_OAM] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NPCF_CMI] = models.Nrf_NFMgmt_NFType_PCF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NSMSF_SMS] = models.Nrf_NFMgmt_NFType_SMSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNSSF_NSSELECTION] = models.Nrf_NFMgmt_NFType_NSSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNSSF_NSSAIAVAILABILITY] = models.Nrf_NFMgmt_NFType_NSSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNSSF_OAM] = models.Nrf_NFMgmt_NFType_NSSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNSSF_CMI] = models.Nrf_NFMgmt_NFType_NSSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDR_DR] = models.Nrf_NFMgmt_NFType_UDR
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDR_IMS_DR] = models.Nrf_NFMgmt_NFType_UDR
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDR_GROUP_ID_MAP] = models.Nrf_NFMgmt_NFType_UDR
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDR_OAM] = models.Nrf_NFMgmt_NFType_UDR
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUDR_CMI] = models.Nrf_NFMgmt_NFType_UDR
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NLMF_LOC] = models.Nrf_NFMgmt_NFType_LMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NLMF_BROADCAST] = models.Nrf_NFMgmt_NFType_LMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NLMF_OAM] = models.Nrf_NFMgmt_NFType_LMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NLMF_CMI] = models.Nrf_NFMgmt_NFType_LMF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_N5G_EIR_EIC] = models.Nrf_NFMgmt_NFType_5_G_EIR
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NBSF_MANAGEMENT] = models.Nrf_NFMgmt_NFType_BSF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NCHF_SPENDINGLIMITCONTROL] = models.Nrf_NFMgmt_NFType_CHF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NCHF_CONVERGEDCHARGING] = models.Nrf_NFMgmt_NFType_CHF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NCHF_OFFLINEONLYCHARGING] = models.Nrf_NFMgmt_NFType_CHF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNWDAF_EVENTSSUBSCRIPTION] = models.Nrf_NFMgmt_NFType_NWDAF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NNWDAF_ANALYTICSINFO] = models.Nrf_NFMgmt_NFType_NWDAF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUPF_OAM] = models.Nrf_NFMgmt_NFType_UPF
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NUPF_CMI] = models.Nrf_NFMgmt_NFType_UPF
	// Now HSS service on UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NHSS_IMS_SDM] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NHSS_IMS_UECM] = models.Nrf_NFMgmt_NFType_UDM
	ServiceNfType[models.Nrf_NFMgmt_ServiceName_NHSS_IMS_UEAU] = models.Nrf_NFMgmt_NFType_UDM
}

func ServiceBaseUri(srvName models.Nrf_NFMgmt_ServiceName) string {
	return ServiceUri(srvName, "")
}

func ServiceUri(srvName models.Nrf_NFMgmt_ServiceName, apiPrefix string) string {
	suffix := ""
	switch srvName {
	case models.Nrf_NFMgmt_ServiceName_NNRF_NFM:
		suffix = ServiceBaseURI_NNRF_NFM_v1
	case models.Nrf_NFMgmt_ServiceName_NNRF_DISC:
		suffix = ServiceBaseURI_NNRF_DISC_v1
	case models.Nrf_NFMgmt_ServiceName_NNRF_OAUTH2:
		suffix = ServiceBaseURI_NNRF_OAUTH2
	case models.Nrf_NFMgmt_ServiceName_NNRF_OAM:
		suffix = ServiceBaseURI_NNRF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NNRF_CMI:
		suffix = ServiceBaseURI_NNRF_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NUDM_SDM:
		suffix = ServiceBaseURI_NUDM_SDM_v2
	case models.Nrf_NFMgmt_ServiceName_NUDM_UECM:
		suffix = ServiceBaseURI_NUDM_UECM_v1
	case models.Nrf_NFMgmt_ServiceName_NUDM_UEAU:
		suffix = ServiceBaseURI_NUDM_UEAU_v1
	case models.Nrf_NFMgmt_ServiceName_NUDM_EE:
		suffix = ServiceBaseURI_NUDM_EE_v1
	case models.Nrf_NFMgmt_ServiceName_NUDM_PP:
		suffix = ServiceBaseURI_NUDM_PP_v1
	// case models.Nrf_NFMgmt_ServiceName_NUDM_NIDDAU:
	// 	suffix = ServiceBaseURI_NUDM_NIDDAU
	// case models.Nrf_NFMgmt_ServiceName_NUDM_MT:
	// 	suffix = ServiceBaseURI_NUDM_MT
	case models.Nrf_NFMgmt_ServiceName_NUDM_OAM:
		suffix = ServiceBaseURI_NUDM_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NUDM_CMI:
		suffix = ServiceBaseURI_NUDM_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NAMF_COMM:
		suffix = ServiceBaseURI_NAMF_COMM_v1
	case models.Nrf_NFMgmt_ServiceName_NAMF_EVTS:
		suffix = ServiceBaseURI_NAMF_EVTS_v1
	case models.Nrf_NFMgmt_ServiceName_NAMF_MT:
		suffix = ServiceBaseURI_NAMF_MT_v1
	case models.Nrf_NFMgmt_ServiceName_NAMF_LOC:
		suffix = ServiceBaseURI_NAMF_LOC_v1
	case models.Nrf_NFMgmt_ServiceName_NAMF_OAM:
		suffix = ServiceBaseURI_NAMF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NAMF_CMI:
		suffix = ServiceBaseURI_NAMF_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NSMF_PDUSESSION:
		suffix = ServiceBaseURI_NSMF_PDUSESSION_v1
	case models.Nrf_NFMgmt_ServiceName_NSMF_EVENT_EXPOSURE:
		suffix = ServiceBaseURI_NSMF_EVENT_EXPOSURE_v1
	// case models.Nrf_NFMgmt_ServiceName_NSMF_NIDD:
	// 	suffix = ServiceBaseURI_NSMF_NIDD
	case models.Nrf_NFMgmt_ServiceName_NSMF_OAM:
		suffix = ServiceBaseURI_NSMF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NSMF_CMI:
		suffix = ServiceBaseURI_NSMF_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NAUSF_AUTH:
		suffix = ServiceBaseURI_NAUSF_AUTH_v1
	case models.Nrf_NFMgmt_ServiceName_NAUSF_SORPROTECTION:
		suffix = ServiceBaseURI_NAUSF_SORPROTECTION_v1
	case models.Nrf_NFMgmt_ServiceName_NAUSF_UPUPROTECTION:
		suffix = ServiceBaseURI_NAUSF_UPUPROTECTION_v1
	case models.Nrf_NFMgmt_ServiceName_NAUSF_OAM:
		suffix = ServiceBaseURI_NAUSF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NAUSF_CMI:
		suffix = ServiceBaseURI_NAUSF_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NNEF_PFDMANAGEMENT:
		suffix = ServiceBaseURI_NNEF_PFDMANAGEMENT_v1
	case models.Nrf_NFMgmt_ServiceName_3GPP_AS_SESSION_WITH_QOS:
		suffix = ServiceBaseURI_3GPP_AS_SESSION_WITH_QOS_v1
	case models.Nrf_NFMgmt_ServiceName_3GPP_PFD_MANAGEMENT:
		suffix = ServiceBaseURI_3GPP_PFD_MANAGEMENT_v1
	case models.Nrf_NFMgmt_ServiceName_3GPP_TRAFFIC_INFLUENCE:
		suffix = ServiceBaseURI_3GPP_TRAFFIC_INFLUENCE_v1
	// case models.Nrf_NFMgmt_ServiceName_NNEF_SMCONTEXT:
	// 	suffix = ServiceBaseURI_NNEF_SMCONTEXT
	// case models.Nrf_NFMgmt_ServiceName_NNEF_EVENTEXPOSURE:
	// 	suffix = ServiceBaseURI_NNEF_EVENTEXPOSURE
	case models.Nrf_NFMgmt_ServiceName_NNEF_OAM:
		suffix = ServiceBaseURI_NNEF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NNEF_CMI:
		suffix = ServiceBaseURI_NNEF_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NPCF_AM_POLICY_CONTROL:
		suffix = ServiceBaseURI_NPCF_AM_POLICY_CONTROL_v1
	case models.Nrf_NFMgmt_ServiceName_NPCF_SMPOLICYCONTROL:
		suffix = ServiceBaseURI_NPCF_SMPOLICYCONTROL_v1
	case models.Nrf_NFMgmt_ServiceName_NPCF_POLICYAUTHORIZATION:
		suffix = ServiceBaseURI_NPCF_POLICYAUTHORIZATION_v1
	case models.Nrf_NFMgmt_ServiceName_NPCF_BDTPOLICYCONTROL:
		suffix = ServiceBaseURI_NPCF_BDTPOLICYCONTROL_v1
	// case models.Nrf_NFMgmt_ServiceName_NPCF_EVENTEXPOSURE:
	// 	suffix = ServiceBaseURI_NPCF_EVENTEXPOSURE
	case models.Nrf_NFMgmt_ServiceName_NPCF_UE_POLICY_CONTROL:
		suffix = ServiceBaseURI_NPCF_UE_POLICY_CONTROL_v1
	case models.Nrf_NFMgmt_ServiceName_NPCF_OAM:
		suffix = ServiceBaseURI_NPCF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NPCF_CMI:
		suffix = ServiceBaseURI_NPCF_CMI_v1
	// case models.Nrf_NFMgmt_ServiceName_NSMSF_SMS:
	// 	suffix = ServiceBaseURI_NSMSF_SMS
	case models.Nrf_NFMgmt_ServiceName_NNSSF_NSSELECTION:
		suffix = ServiceBaseURI_NNSSF_NSSELECTION_v2
	case models.Nrf_NFMgmt_ServiceName_NNSSF_NSSAIAVAILABILITY:
		suffix = ServiceBaseURI_NNSSF_NSSAIAVAILABILITY_v1
	case models.Nrf_NFMgmt_ServiceName_NNSSF_OAM:
		suffix = ServiceBaseURI_NNSSF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NNSSF_CMI:
		suffix = ServiceBaseURI_NNSSF_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NUDR_DR:
		suffix = ServiceBaseURI_NUDR_DR_v2
	case models.Nrf_NFMgmt_ServiceName_NUDR_IMS_DR:
		suffix = ServiceBaseURI_NUDR_IMS_DR_v1
	// case models.Nrf_NFMgmt_ServiceName_NUDR_GROUP_ID_MAP:
	// 	suffix = ServiceBaseURI_NUDR_GROUP_ID_MAP
	case models.Nrf_NFMgmt_ServiceName_NUDR_OAM:
		suffix = ServiceBaseURI_NUDR_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NUDR_CMI:
		suffix = ServiceBaseURI_NUDR_CMI_v1
	case models.Nrf_NFMgmt_ServiceName_NLMF_LOC:
		suffix = ServiceBaseURI_NLMF_LOC_v1
	case models.Nrf_NFMgmt_ServiceName_NLMF_BROADCAST:
		suffix = ServiceBaseURI_NLMF_BROADCAST_v1
	case models.Nrf_NFMgmt_ServiceName_NLMF_OAM:
		suffix = ServiceBaseURI_NLMF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NLMF_CMI:
		suffix = ServiceBaseURI_NLMF_CMI_v1
	// case models.Nrf_NFMgmt_ServiceName_N5G_EIR_EIC:
	// 	suffix = ServiceBaseURI_N5G_EIR_EIC
	// case models.Nrf_NFMgmt_ServiceName_NBSF_MANAGEMENT:
	// 	suffix = ServiceBaseURI_NBSF_MANAGEMENT
	// case models.Nrf_NFMgmt_ServiceName_NCHF_SPENDINGLIMITCONTROL:
	// 	suffix = ServiceBaseURI_NCHF_SPENDINGLIMITCONTROL
	// case models.Nrf_NFMgmt_ServiceName_NCHF_CONVERGEDCHARGING:
	// 	suffix = ServiceBaseURI_NCHF_CONVERGEDCHARGING
	// case models.Nrf_NFMgmt_ServiceName_NCHF_OFFLINEONLYCHARGING:
	// 	suffix = ServiceBaseURI_NCHF_OFFLINEONLYCHARGING
	// case models.Nrf_NFMgmt_ServiceName_NNWDAF_EVENTSSUBSCRIPTION:
	// 	suffix = ServiceBaseURI_NNWDAF_EVENTSSUBSCRIPTION
	// case models.Nrf_NFMgmt_ServiceName_NNWDAF_ANALYTICSINFO:
	// 	suffix = ServiceBaseURI_NNWDAF_ANALYTICSINFO
	// case models.Nrf_NFMgmt_ServiceName_NGMLC_LOC:
	// 	suffix = ServiceBaseURI_NGMLC_LOC
	// case models.Nrf_NFMgmt_ServiceName_NUCMF_PROVISIONING:
	// 	suffix = ServiceBaseURI_NUCMF_PROVISIONING
	// case models.Nrf_NFMgmt_ServiceName_NUCMF_UECAPABILITYMANAGEMENT:
	// 	suffix = ServiceBaseURI_NUCMF_UECAPABILITYMANAGEMENT
	// case models.Nrf_NFMgmt_ServiceName_NHSS_SDM:
	// 	suffix = ServiceBaseURI_NHSS_SDM
	// case models.Nrf_NFMgmt_ServiceName_NHSS_UECM:
	// 	suffix = ServiceBaseURI_NHSS_UECM
	// case models.Nrf_NFMgmt_ServiceName_NHSS_UEAU:
	// 	suffix = ServiceBaseURI_NHSS_UEAU
	// case models.Nrf_NFMgmt_ServiceName_NHSS_EE:
	// 	suffix = ServiceBaseURI_NHSS_EE
	case models.Nrf_NFMgmt_ServiceName_NHSS_IMS_SDM:
		suffix = ServiceBaseURI_NHSS_IMS_SDM_v1
	case models.Nrf_NFMgmt_ServiceName_NHSS_IMS_UECM:
		suffix = ServiceBaseURI_NHSS_IMS_UECM_v1
	case models.Nrf_NFMgmt_ServiceName_NHSS_IMS_UEAU:
		suffix = ServiceBaseURI_NHSS_IMS_UEAU_v1
	// case models.Nrf_NFMgmt_ServiceName_NSEPP_TELESCOPIC:
	// 	suffix = ServiceBaseURI_NSEPP_TELESCOPIC
	// case models.Nrf_NFMgmt_ServiceName_NSORAF_SOR:
	// 	suffix = ServiceBaseURI_NSORAF_SOR
	// case models.Nrf_NFMgmt_ServiceName_NSPAF_SECURED_PACKET:
	// 	suffix = ServiceBaseURI_NSPAF_SECURED_PACKET
	// case models.Nrf_NFMgmt_ServiceName_NUDSF_DR:
	// 	suffix = ServiceBaseURI_NUDSF_DR
	// case models.Nrf_NFMgmt_ServiceName_NNSSAAF_NSSAA:
	// 	suffix = ServiceBaseURI_NNSSAAF_NSSAA
	case models.Nrf_NFMgmt_ServiceName_NUPF_OAM:
		suffix = ServiceBaseURI_NUPF_OAM_v1
	case models.Nrf_NFMgmt_ServiceName_NUPF_CMI:
		suffix = ServiceBaseURI_NUPF_CMI_v1
	default:
	}
	return apiPrefix + suffix
}

func ApiVersion(name models.Nrf_NFMgmt_ServiceName) string {
	s := ServiceBaseUri(name)
	if s != "" {
		return s[strings.LastIndex(s, "/")+1:]
	}
	return ""
}

func GetServiceNfProfileAndUri(
	nfInstances []models.Nrf_NFDisc_NFProfile,
	srvName models.Nrf_NFMgmt_ServiceName,
) (*models.Nrf_NFDisc_NFProfile, string, error) {
	// select the first ServiceName
	// TODO: select base on other info
	var nfProf *models.Nrf_NFDisc_NFProfile
	var uri string
	for i := range nfInstances {
		nfProf = &nfInstances[i]
		uri = GetNFServiceUri(nfProf, srvName)
		if uri != "" {
			break
		}
	}
	if uri == "" {
		return nil, "", fmt.Errorf("no uri for %s found", srvName)
	}
	return nfProf, uri, nil
}

// Returns NF Uri derived from NfProfile with corresponding service
func GetNFServiceUri(
	nfProf *models.Nrf_NFDisc_NFProfile,
	srvName models.Nrf_NFMgmt_ServiceName,
) string {
	if nfProf == nil {
		return ""
	}

	nfUri := ""
	for _, srv := range nfProf.NfServices {
		if srv.ServiceName == srvName &&
			srv.NfServiceStatus == models.Nrf_NFMgmt_NFServiceStatus_REGISTERED {
			if srv.ApiPrefix != "" {
				nfUri = srv.ApiPrefix
			} else if srv.Fqdn != "" {
				nfUri = string(srv.Scheme) + "://" + srv.Fqdn
			} else if nfProf.Fqdn != "" {
				nfUri = string(srv.Scheme) + "://" + nfProf.Fqdn
			} else if len(srv.IpEndPoints) != 0 {
				// Select the first IpEndPoint
				// TODO: select others when failure
				point := (srv.IpEndPoints)[0]
				if point.Ipv4Address != "" {
					nfUri = getUriFromIpEndPoint(srv.Scheme, point.Ipv4Address, point.Port)
				} else if len(nfProf.Ipv4Addresses) != 0 {
					nfUri = getUriFromIpEndPoint(srv.Scheme, nfProf.Ipv4Addresses[0], point.Port)
				}
			}
		}
		if nfUri != "" {
			break
		}
	}
	return nfUri
}

func getUriFromIpEndPoint(scheme models.UriScheme, ipv4Address string, port int32) string {
	uri := ""
	if port != 0 {
		uri = string(scheme) + "://" + ipv4Address + ":" + strconv.Itoa(int(port))
	} else {
		switch scheme {
		case models.UriScheme_HTTP:
			uri = string(scheme) + "://" + ipv4Address + ":80"
		case models.UriScheme_HTTPS:
			uri = string(scheme) + "://" + ipv4Address + ":443"
		}
	}
	return uri
}

func SnssaiEqualFold(s, t models.Snssai) bool {
	if s.Sst == t.Sst && sdEqualFold(s.Sd, t.Sd) {
		return true
	}
	return false
}

func ExtSnssaiEqualFold(s, t models.ExtSnssai) bool {
	if s.Sst == t.Sst && sdEqualFold(s.Sd, t.Sd) {
		return true
	}
	return false
}

// According to TS 23.003 Section 28.4.2 Format of the S-NSSAI:
// The SD field has a reserved value "no SD value associated with the SST"
// defined as hexadecimal FFFFFF.
// In certain protocols, the SD field is not included to
// indicate that no SD value is associated with the SST.
func sdEqualFold(s, t string) bool {
	if strings.EqualFold(s, t) {
		return true
	}
	if s == "" && strings.EqualFold(t, "ffffff") {
		return true
	}
	if strings.EqualFold(s, "ffffff") && t == "" {
		return true
	}
	return false
}

func PlmnIdEqual(s models.PlmnId, t models.PlmnId) bool {
	return (s.Mcc == t.Mcc) && (s.Mnc == t.Mnc)
}

func PlmnIdNidEqualFold(s models.PlmnIdNid, t models.PlmnIdNid) bool {
	return (s.Mcc == t.Mcc) &&
		(s.Mnc == t.Mnc) &&
		strings.EqualFold(s.Nid, t.Nid)
}

func GuamiEqualFold(s models.Guami, t models.Guami) bool {
	if s.PlmnId == nil || t.PlmnId == nil {
		// unexpected
		return false
	}
	return PlmnIdNidEqualFold(*s.PlmnId, *t.PlmnId) &&
		AmfIdEqualFold(s.AmfId, t.AmfId)
}

func AmfIdEqualFold(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func TaiEqualFold(s models.Tai, t models.Tai) bool {
	if s.PlmnId == nil || t.PlmnId == nil {
		// unexpected
		return false
	}
	return PlmnIdEqual(*s.PlmnId, *t.PlmnId) &&
		TacEqualFold(s.Tac, t.Tac) &&
		NidEqualFold(s.Nid, t.Nid)
}

func TacEqualFold(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func NidEqualFold(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func GnbIdEqualFold(s models.GNbId, t models.GNbId) bool {
	return (s.BitLength == t.BitLength) &&
		strings.EqualFold(s.GNBValue, t.GNBValue)
}

func SnssaiHexToModels(hexString string) (*models.Snssai, error) {
	hslen := len(hexString)
	if hslen != 2 && hslen != 8 {
		return nil, fmt.Errorf("hexString length(%d) should be 2 or 8", hslen)
	}
	sst, err := strconv.ParseInt(hexString[:2], 16, 32)
	if err != nil {
		return nil, err
	}
	sNssai := &models.Snssai{
		Sst: int32(sst), // #nosec G115
		Sd:  hexString[2:],
	}
	return sNssai, nil
}

func SnssaiModelsToHex(snssai models.Snssai) string {
	sst := fmt.Sprintf("%02x", snssai.Sst)
	return sst + snssai.Sd
}

func PlmnIdJsonToModels(plmnIdJson []byte) (*models.PlmnId, error) {
	var plmnId models.PlmnId
	err := json.Unmarshal(plmnIdJson, &plmnId)
	if err != nil {
		return nil, err
	}
	return &plmnId, nil
}

func ModelsGnbIdToUint(gnbId *models.GNbId) (uint64, error) {
	if gnbId == nil {
		return 0, errors.Errorf("ModelsGnbIdToUint(): gnbId is nil")
	}

	gnbIdVal, err := strconv.ParseUint(gnbId.GNBValue, 16, int(gnbId.BitLength))
	if err != nil {
		return 0, errors.Wrapf(err, "ModelsGnbIdToUint(): ParseUint(%s, 16, %d)",
			gnbId.GNBValue, gnbId.BitLength)
	}
	return gnbIdVal, nil
}

func ModelsPlmnIdNidToPlmnId(plmnidNid *models.PlmnIdNid) *models.PlmnId {
	if plmnidNid == nil {
		return nil
	}
	return &models.PlmnId{
		Mcc: plmnidNid.Mcc,
		Mnc: plmnidNid.Mnc,
	}
}

func ModelsPlmnIdToPlmnIdNid(plmnId *models.PlmnId) *models.PlmnIdNid {
	if plmnId == nil {
		return nil
	}
	return &models.PlmnIdNid{
		Mcc: plmnId.Mcc,
		Mnc: plmnId.Mnc,
		Nid: "",
	}
}

func getServiceNameFromUrl(path string) string {
	pathElements := strings.Split(path, "/")

	// Avoid a panic by checking if the returned slice is non-empty, should not happen
	if len(pathElements) == 0 {
		return ""
	}
	serviceName := pathElements[0]
	if len(pathElements) > 1 {
		serviceName = pathElements[1]
	}

	return serviceName
}

func getRespStatusCode(response *http.Response) int {
	if response != nil {
		return response.StatusCode
	}
	return 0
}
