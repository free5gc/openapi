package openapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/free5gc/openapi/models"
)

const (
	ServiceBaseURI_NNRF_NFM_v1  string = "/" + string(models.ServiceName_NNRF_NFM) + "/v1"
	ServiceBaseURI_NNRF_DISC_v1 string = "/" + string(models.ServiceName_NNRF_DISC) + "/v1"
	ServiceBaseURI_NNRF_OAUTH2  string = ""
	ServiceBaseURI_NNRF_OAM_v1  string = "/" + string(models.ServiceName_NNRF_OAM) + "/v1"
	ServiceBaseURI_NNRF_CMI_v1  string = "/" + string(models.ServiceName_NNRF_CMI) + "/v1"
	ServiceBaseURI_NUDM_SDM_v2  string = "/" + string(models.ServiceName_NUDM_SDM) + "/v2"
	ServiceBaseURI_NUDM_UECM_v1 string = "/" + string(models.ServiceName_NUDM_UECM) + "/v1"
	ServiceBaseURI_NUDM_UEAU_v1 string = "/" + string(models.ServiceName_NUDM_UEAU) + "/v1"
	ServiceBaseURI_NUDM_EE_v1   string = "/" + string(models.ServiceName_NUDM_EE) + "/v1"
	ServiceBaseURI_NUDM_PP_v1   string = "/" + string(models.ServiceName_NUDM_PP) + "/v1"
	// ServiceBaseURI_NUDM_NIDDAU string = "/" + string(models.ServiceName_NUDM_NIDDAU)
	// ServiceBaseURI_NUDM_MT     string = "/" + string(models.ServiceName_NUDM_MT)
	ServiceBaseURI_NUDM_OAM_v1            string = "/" + string(models.ServiceName_NUDM_OAM) + "/v1"
	ServiceBaseURI_NUDM_CMI_v1            string = "/" + string(models.ServiceName_NUDM_CMI) + "/v1"
	ServiceBaseURI_NAMF_COMM_v1           string = "/" + string(models.ServiceName_NAMF_COMM) + "/v1"
	ServiceBaseURI_NAMF_EVTS_v1           string = "/" + string(models.ServiceName_NAMF_EVTS) + "/v1"
	ServiceBaseURI_NAMF_MT_v1             string = "/" + string(models.ServiceName_NAMF_MT) + "/v1"
	ServiceBaseURI_NAMF_LOC_v1            string = "/" + string(models.ServiceName_NAMF_LOC) + "/v1"
	ServiceBaseURI_NAMF_OAM_v1            string = "/" + string(models.ServiceName_NAMF_OAM) + "/v1"
	ServiceBaseURI_NAMF_CMI_v1            string = "/" + string(models.ServiceName_NAMF_CMI) + "/v1"
	ServiceBaseURI_NSMF_PDUSESSION_v1     string = "/" + string(models.ServiceName_NSMF_PDUSESSION) + "/v1"
	ServiceBaseURI_NSMF_EVENT_EXPOSURE_v1 string = "/" + string(models.ServiceName_NSMF_EVENT_EXPOSURE) + "/v1"
	// ServiceBaseURI_NSMF_NIDD string = "/" + string(models.ServiceName_NSMF_NIDD)
	ServiceBaseURI_NSMF_OAM_v1                 string = "/" + string(models.ServiceName_NSMF_OAM) + "/v1"
	ServiceBaseURI_NSMF_CMI_v1                 string = "/" + string(models.ServiceName_NSMF_CMI) + "/v1"
	ServiceBaseURI_NAUSF_AUTH_v1               string = "/" + string(models.ServiceName_NAUSF_AUTH) + "/v1"
	ServiceBaseURI_NAUSF_SORPROTECTION_v1      string = "/" + string(models.ServiceName_NAUSF_SORPROTECTION) + "/v1"
	ServiceBaseURI_NAUSF_UPUPROTECTION_v1      string = "/" + string(models.ServiceName_NAUSF_UPUPROTECTION) + "/v1"
	ServiceBaseURI_NAUSF_OAM_v1                string = "/" + string(models.ServiceName_NAUSF_OAM) + "/v1"
	ServiceBaseURI_NAUSF_CMI_v1                string = "/" + string(models.ServiceName_NAUSF_CMI) + "/v1"
	ServiceBaseURI_NNEF_PFDMANAGEMENT_v1       string = "/" + string(models.ServiceName_NNEF_PFDMANAGEMENT) + "/v1"
	ServiceBaseURI_3GPP_AS_SESSION_WITH_QOS_v1 string = "/" + string(models.ServiceName_3GPP_AS_SESSION_WITH_QOS) + "/v1"
	ServiceBaseURI_3GPP_PFD_MANAGEMENT_v1      string = "/" + string(models.ServiceName_3GPP_PFD_MANAGEMENT) + "/v1"
	ServiceBaseURI_3GPP_TRAFFIC_INFLUENCE_v1   string = "/" + string(models.ServiceName_3GPP_TRAFFIC_INFLUENCE) + "/v1"
	// ServiceBaseURI_NNEF_SMCONTEXT     string = "/" + string(models.ServiceName_NNEF_SMCONTEXT)
	// ServiceBaseURI_NNEF_EVENTEXPOSURE string = "/" + string(models.ServiceName_NNEF_EVENTEXPOSURE)
	ServiceBaseURI_NNEF_OAM_v1                 string = "/" + string(models.ServiceName_NNEF_OAM) + "/v1"
	ServiceBaseURI_NNEF_CMI_v1                 string = "/" + string(models.ServiceName_NNEF_CMI) + "/v1"
	ServiceBaseURI_NPCF_AM_POLICY_CONTROL_v1   string = "/" + string(models.ServiceName_NPCF_AM_POLICY_CONTROL) + "/v1"
	ServiceBaseURI_NPCF_SMPOLICYCONTROL_v1     string = "/" + string(models.ServiceName_NPCF_SMPOLICYCONTROL) + "/v1"
	ServiceBaseURI_NPCF_POLICYAUTHORIZATION_v1 string = "/" + string(models.ServiceName_NPCF_POLICYAUTHORIZATION) + "/v1"
	ServiceBaseURI_NPCF_BDTPOLICYCONTROL_v1    string = "/" + string(models.ServiceName_NPCF_BDTPOLICYCONTROL) + "/v1"
	// ServiceBaseURI_NPCF_EVENTEXPOSURE string = "/" + string(models.ServiceName_NPCF_EVENTEXPOSURE)
	ServiceBaseURI_NPCF_UE_POLICY_CONTROL_v1 string = "/" + string(models.ServiceName_NPCF_UE_POLICY_CONTROL) + "/v1"
	ServiceBaseURI_NPCF_OAM_v1               string = "/" + string(models.ServiceName_NPCF_OAM) + "/v1"
	ServiceBaseURI_NPCF_CMI_v1               string = "/" + string(models.ServiceName_NPCF_CMI) + "/v1"
	// ServiceBaseURI_NSMSF_SMS string = "/" + string(models.ServiceName_NSMSF_SMS)
	ServiceBaseURI_NNSSF_NSSELECTION_v2       string = "/" + string(models.ServiceName_NNSSF_NSSELECTION) + "/v2"
	ServiceBaseURI_NNSSF_NSSAIAVAILABILITY_v1 string = "/" + string(models.ServiceName_NNSSF_NSSAIAVAILABILITY) + "/v1"
	ServiceBaseURI_NNSSF_OAM_v1               string = "/" + string(models.ServiceName_NNSSF_OAM) + "/v1"
	ServiceBaseURI_NNSSF_CMI_v1               string = "/" + string(models.ServiceName_NNSSF_CMI) + "/v1"
	ServiceBaseURI_NUDR_DR_v2                 string = "/" + string(models.ServiceName_NUDR_DR) + "/v2"
	ServiceBaseURI_NUDR_GROUP_ID_MAP_v1       string = "/" + string(models.ServiceName_NUDR_GROUP_ID_MAP) + "/v1"
	ServiceBaseURI_NUDR_OAM_v1                string = "/" + string(models.ServiceName_NUDR_OAM) + "/v1"
	ServiceBaseURI_NUDR_CMI_v1                string = "/" + string(models.ServiceName_NUDR_CMI) + "/v1"
	// ServiceBaseURI_NLMF_LOC                     string = "/" + string(models.ServiceName_NLMF_LOC)
	// ServiceBaseURI_N5G_EIR_EIC                  string = "/" + string(models.ServiceName_N5G_EIR_EIC)
	// ServiceBaseURI_NBSF_MANAGEMENT              string = "/" + string(models.ServiceName_NBSF_MANAGEMENT)
	// ServiceBaseURI_NCHF_SPENDINGLIMITCONTROL    string = "/" + string(models.ServiceName_NCHF_SPENDINGLIMITCONTROL)
	// ServiceBaseURI_NCHF_CONVERGEDCHARGING       string = "/" + string(models.ServiceName_NCHF_CONVERGEDCHARGING)
	// ServiceBaseURI_NCHF_OFFLINEONLYCHARGING     string = "/" + string(models.ServiceName_NCHF_OFFLINEONLYCHARGING)
	// ServiceBaseURI_NNWDAF_EVENTSSUBSCRIPTION    string = "/" + string(models.ServiceName_NNWDAF_EVENTSSUBSCRIPTION)
	// ServiceBaseURI_NNWDAF_ANALYTICSINFO         string = "/" + string(models.ServiceName_NNWDAF_ANALYTICSINFO)
	// ServiceBaseURI_NGMLC_LOC                    string = "/" + string(models.ServiceName_NGMLC_LOC)
	// ServiceBaseURI_NUCMF_PROVISIONING           string = "/" + string(models.ServiceName_NUCMF_PROVISIONING)
	// ServiceBaseURI_NUCMF_UECAPABILITYMANAGEMENT string = "/" + string(models.ServiceName_NUCMF_UECAPABILITYMANAGEMENT)
	// ServiceBaseURI_NHSS_SDM                     string = "/" + string(models.ServiceName_NHSS_SDM)
	// ServiceBaseURI_NHSS_UECM                    string = "/" + string(models.ServiceName_NHSS_UECM)
	// ServiceBaseURI_NHSS_UEAU                    string = "/" + string(models.ServiceName_NHSS_UEAU)
	// ServiceBaseURI_NHSS_EE                      string = "/" + string(models.ServiceName_NHSS_EE)
	ServiceBaseURI_NHSS_IMS_SDM_v1  string = "/" + string(models.ServiceName_NHSS_IMS_SDM) + "/v1"
	ServiceBaseURI_NHSS_IMS_UECM_v1 string = "/" + string(models.ServiceName_NHSS_IMS_UECM) + "/v1"
	ServiceBaseURI_NHSS_IMS_UEAU_v1 string = "/" + string(models.ServiceName_NHSS_IMS_UEAU) + "/v1"
	// ServiceBaseURI_NSEPP_TELESCOPIC             string = "/" + string(models.ServiceName_NSEPP_TELESCOPIC)
	// ServiceBaseURI_NSORAF_SOR                   string = "/" + string(models.ServiceName_NSORAF_SOR)
	// ServiceBaseURI_NSPAF_SECURED_PACKET         string = "/" + string(models.ServiceName_NSPAF_SECURED_PACKET)
	// ServiceBaseURI_NUDSF_DR                     string = "/" + string(models.ServiceName_NUDSF_DR)
	// ServiceBaseURI_NNSSAAF_NSSAA                string = "/" + string(models.ServiceName_NNSSAAF_NSSAA)
	ServiceBaseURI_NUPF_OAM_v1 string = "/" + string(models.ServiceName_NUPF_OAM) + "/v1"
	ServiceBaseURI_NUPF_CMI_v1 string = "/" + string(models.ServiceName_NUPF_CMI) + "/v1"
)

var ServiceNfType map[models.ServiceName]models.NrfNfManagementNfType

func init() {
	ServiceNfType = make(map[models.ServiceName]models.NrfNfManagementNfType)
	ServiceNfType[models.ServiceName_NNRF_NFM] = models.NrfNfManagementNfType_NRF
	ServiceNfType[models.ServiceName_NNRF_DISC] = models.NrfNfManagementNfType_NRF
	ServiceNfType[models.ServiceName_NNRF_OAUTH2] = models.NrfNfManagementNfType_NRF
	ServiceNfType[models.ServiceName_NNRF_OAM] = models.NrfNfManagementNfType_NRF
	ServiceNfType[models.ServiceName_NNRF_CMI] = models.NrfNfManagementNfType_NRF
	ServiceNfType[models.ServiceName_NUDM_SDM] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_UECM] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_UEAU] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_EE] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_PP] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_NIDDAU] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_MT] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_OAM] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NUDM_CMI] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NAMF_COMM] = models.NrfNfManagementNfType_AMF
	ServiceNfType[models.ServiceName_NAMF_EVTS] = models.NrfNfManagementNfType_AMF
	ServiceNfType[models.ServiceName_NAMF_MT] = models.NrfNfManagementNfType_AMF
	ServiceNfType[models.ServiceName_NAMF_LOC] = models.NrfNfManagementNfType_AMF
	ServiceNfType[models.ServiceName_NAMF_OAM] = models.NrfNfManagementNfType_AMF
	ServiceNfType[models.ServiceName_NAMF_CMI] = models.NrfNfManagementNfType_AMF
	ServiceNfType[models.ServiceName_NSMF_PDUSESSION] = models.NrfNfManagementNfType_SMF
	ServiceNfType[models.ServiceName_NSMF_EVENT_EXPOSURE] = models.NrfNfManagementNfType_SMF
	ServiceNfType[models.ServiceName_NSMF_NIDD] = models.NrfNfManagementNfType_SMF
	ServiceNfType[models.ServiceName_NSMF_OAM] = models.NrfNfManagementNfType_SMF
	ServiceNfType[models.ServiceName_NSMF_CMI] = models.NrfNfManagementNfType_SMF
	ServiceNfType[models.ServiceName_NAUSF_AUTH] = models.NrfNfManagementNfType_AUSF
	ServiceNfType[models.ServiceName_NAUSF_SORPROTECTION] = models.NrfNfManagementNfType_AUSF
	ServiceNfType[models.ServiceName_NAUSF_UPUPROTECTION] = models.NrfNfManagementNfType_AUSF
	ServiceNfType[models.ServiceName_NAUSF_OAM] = models.NrfNfManagementNfType_AUSF
	ServiceNfType[models.ServiceName_NAUSF_CMI] = models.NrfNfManagementNfType_AUSF
	ServiceNfType[models.ServiceName_NNEF_PFDMANAGEMENT] = models.NrfNfManagementNfType_NEF
	ServiceNfType[models.ServiceName_NNEF_SMCONTEXT] = models.NrfNfManagementNfType_NEF
	ServiceNfType[models.ServiceName_NNEF_EVENTEXPOSURE] = models.NrfNfManagementNfType_NEF
	ServiceNfType[models.ServiceName_NNEF_OAM] = models.NrfNfManagementNfType_NEF
	ServiceNfType[models.ServiceName_NNEF_CMI] = models.NrfNfManagementNfType_NEF
	ServiceNfType[models.ServiceName_NPCF_AM_POLICY_CONTROL] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_SMPOLICYCONTROL] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_POLICYAUTHORIZATION] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_BDTPOLICYCONTROL] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_EVENTEXPOSURE] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_UE_POLICY_CONTROL] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_OAM] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NPCF_CMI] = models.NrfNfManagementNfType_PCF
	ServiceNfType[models.ServiceName_NSMSF_SMS] = models.NrfNfManagementNfType_SMSF
	ServiceNfType[models.ServiceName_NNSSF_NSSELECTION] = models.NrfNfManagementNfType_NSSF
	ServiceNfType[models.ServiceName_NNSSF_NSSAIAVAILABILITY] = models.NrfNfManagementNfType_NSSF
	ServiceNfType[models.ServiceName_NNSSF_OAM] = models.NrfNfManagementNfType_NSSF
	ServiceNfType[models.ServiceName_NNSSF_CMI] = models.NrfNfManagementNfType_NSSF
	ServiceNfType[models.ServiceName_NUDR_DR] = models.NrfNfManagementNfType_UDR
	ServiceNfType[models.ServiceName_NUDR_GROUP_ID_MAP] = models.NrfNfManagementNfType_UDR
	ServiceNfType[models.ServiceName_NUDR_OAM] = models.NrfNfManagementNfType_UDR
	ServiceNfType[models.ServiceName_NUDR_CMI] = models.NrfNfManagementNfType_UDR
	ServiceNfType[models.ServiceName_NLMF_LOC] = models.NrfNfManagementNfType_LMF
	ServiceNfType[models.ServiceName_N5G_EIR_EIC] = models.NrfNfManagementNfType__5_G_EIR
	ServiceNfType[models.ServiceName_NBSF_MANAGEMENT] = models.NrfNfManagementNfType_BSF
	ServiceNfType[models.ServiceName_NCHF_SPENDINGLIMITCONTROL] = models.NrfNfManagementNfType_CHF
	ServiceNfType[models.ServiceName_NCHF_CONVERGEDCHARGING] = models.NrfNfManagementNfType_CHF
	ServiceNfType[models.ServiceName_NCHF_OFFLINEONLYCHARGING] = models.NrfNfManagementNfType_CHF
	ServiceNfType[models.ServiceName_NNWDAF_EVENTSSUBSCRIPTION] = models.NrfNfManagementNfType_NWDAF
	ServiceNfType[models.ServiceName_NNWDAF_ANALYTICSINFO] = models.NrfNfManagementNfType_NWDAF
	ServiceNfType[models.ServiceName_NUPF_OAM] = models.NrfNfManagementNfType_UPF
	ServiceNfType[models.ServiceName_NUPF_CMI] = models.NrfNfManagementNfType_UPF
	// Now HSS service on UDM
	ServiceNfType[models.ServiceName_NHSS_IMS_SDM] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NHSS_IMS_UECM] = models.NrfNfManagementNfType_UDM
	ServiceNfType[models.ServiceName_NHSS_IMS_UEAU] = models.NrfNfManagementNfType_UDM
}

func ServiceBaseUri(srvName models.ServiceName) string {
	return ServiceUri(srvName, "")
}

func ServiceUri(srvName models.ServiceName, apiPrefix string) string {
	suffix := ""
	switch srvName {
	case models.ServiceName_NNRF_NFM:
		suffix = ServiceBaseURI_NNRF_NFM_v1
	case models.ServiceName_NNRF_DISC:
		suffix = ServiceBaseURI_NNRF_DISC_v1
	case models.ServiceName_NNRF_OAUTH2:
		suffix = ServiceBaseURI_NNRF_OAUTH2
	case models.ServiceName_NNRF_OAM:
		suffix = ServiceBaseURI_NNRF_OAM_v1
	case models.ServiceName_NNRF_CMI:
		suffix = ServiceBaseURI_NNRF_CMI_v1
	case models.ServiceName_NUDM_SDM:
		suffix = ServiceBaseURI_NUDM_SDM_v2
	case models.ServiceName_NUDM_UECM:
		suffix = ServiceBaseURI_NUDM_UECM_v1
	case models.ServiceName_NUDM_UEAU:
		suffix = ServiceBaseURI_NUDM_UEAU_v1
	case models.ServiceName_NUDM_EE:
		suffix = ServiceBaseURI_NUDM_EE_v1
	case models.ServiceName_NUDM_PP:
		suffix = ServiceBaseURI_NUDM_PP_v1
	// case models.ServiceName_NUDM_NIDDAU:
	// 	suffix = ServiceBaseURI_NUDM_NIDDAU
	// case models.ServiceName_NUDM_MT:
	// 	suffix = ServiceBaseURI_NUDM_MT
	case models.ServiceName_NUDM_OAM:
		suffix = ServiceBaseURI_NUDM_OAM_v1
	case models.ServiceName_NUDM_CMI:
		suffix = ServiceBaseURI_NUDM_CMI_v1
	case models.ServiceName_NAMF_COMM:
		suffix = ServiceBaseURI_NAMF_COMM_v1
	case models.ServiceName_NAMF_EVTS:
		suffix = ServiceBaseURI_NAMF_EVTS_v1
	case models.ServiceName_NAMF_MT:
		suffix = ServiceBaseURI_NAMF_MT_v1
	case models.ServiceName_NAMF_LOC:
		suffix = ServiceBaseURI_NAMF_LOC_v1
	case models.ServiceName_NAMF_OAM:
		suffix = ServiceBaseURI_NAMF_OAM_v1
	case models.ServiceName_NAMF_CMI:
		suffix = ServiceBaseURI_NAMF_CMI_v1
	case models.ServiceName_NSMF_PDUSESSION:
		suffix = ServiceBaseURI_NSMF_PDUSESSION_v1
	case models.ServiceName_NSMF_EVENT_EXPOSURE:
		suffix = ServiceBaseURI_NSMF_EVENT_EXPOSURE_v1
	// case models.ServiceName_NSMF_NIDD:
	// 	suffix = ServiceBaseURI_NSMF_NIDD
	case models.ServiceName_NSMF_OAM:
		suffix = ServiceBaseURI_NSMF_OAM_v1
	case models.ServiceName_NSMF_CMI:
		suffix = ServiceBaseURI_NSMF_CMI_v1
	case models.ServiceName_NAUSF_AUTH:
		suffix = ServiceBaseURI_NAUSF_AUTH_v1
	case models.ServiceName_NAUSF_SORPROTECTION:
		suffix = ServiceBaseURI_NAUSF_SORPROTECTION_v1
	case models.ServiceName_NAUSF_UPUPROTECTION:
		suffix = ServiceBaseURI_NAUSF_UPUPROTECTION_v1
	case models.ServiceName_NAUSF_OAM:
		suffix = ServiceBaseURI_NAUSF_OAM_v1
	case models.ServiceName_NAUSF_CMI:
		suffix = ServiceBaseURI_NAUSF_CMI_v1
	case models.ServiceName_NNEF_PFDMANAGEMENT:
		suffix = ServiceBaseURI_NNEF_PFDMANAGEMENT_v1
	case models.ServiceName_3GPP_AS_SESSION_WITH_QOS:
		suffix = ServiceBaseURI_3GPP_AS_SESSION_WITH_QOS_v1
	case models.ServiceName_3GPP_PFD_MANAGEMENT:
		suffix = ServiceBaseURI_3GPP_PFD_MANAGEMENT_v1
	case models.ServiceName_3GPP_TRAFFIC_INFLUENCE:
		suffix = ServiceBaseURI_3GPP_TRAFFIC_INFLUENCE_v1
	// case models.ServiceName_NNEF_SMCONTEXT:
	// 	suffix = ServiceBaseURI_NNEF_SMCONTEXT
	// case models.ServiceName_NNEF_EVENTEXPOSURE:
	// 	suffix = ServiceBaseURI_NNEF_EVENTEXPOSURE
	case models.ServiceName_NNEF_OAM:
		suffix = ServiceBaseURI_NNEF_OAM_v1
	case models.ServiceName_NNEF_CMI:
		suffix = ServiceBaseURI_NNEF_CMI_v1
	case models.ServiceName_NPCF_AM_POLICY_CONTROL:
		suffix = ServiceBaseURI_NPCF_AM_POLICY_CONTROL_v1
	case models.ServiceName_NPCF_SMPOLICYCONTROL:
		suffix = ServiceBaseURI_NPCF_SMPOLICYCONTROL_v1
	case models.ServiceName_NPCF_POLICYAUTHORIZATION:
		suffix = ServiceBaseURI_NPCF_POLICYAUTHORIZATION_v1
	case models.ServiceName_NPCF_BDTPOLICYCONTROL:
		suffix = ServiceBaseURI_NPCF_BDTPOLICYCONTROL_v1
	// case models.ServiceName_NPCF_EVENTEXPOSURE:
	// 	suffix = ServiceBaseURI_NPCF_EVENTEXPOSURE
	case models.ServiceName_NPCF_UE_POLICY_CONTROL:
		suffix = ServiceBaseURI_NPCF_UE_POLICY_CONTROL_v1
	case models.ServiceName_NPCF_OAM:
		suffix = ServiceBaseURI_NPCF_OAM_v1
	case models.ServiceName_NPCF_CMI:
		suffix = ServiceBaseURI_NPCF_CMI_v1
	// case models.ServiceName_NSMSF_SMS:
	// 	suffix = ServiceBaseURI_NSMSF_SMS
	case models.ServiceName_NNSSF_NSSELECTION:
		suffix = ServiceBaseURI_NNSSF_NSSELECTION_v2
	case models.ServiceName_NNSSF_NSSAIAVAILABILITY:
		suffix = ServiceBaseURI_NNSSF_NSSAIAVAILABILITY_v1
	case models.ServiceName_NNSSF_OAM:
		suffix = ServiceBaseURI_NNSSF_OAM_v1
	case models.ServiceName_NNSSF_CMI:
		suffix = ServiceBaseURI_NNSSF_CMI_v1
	case models.ServiceName_NUDR_DR:
		suffix = ServiceBaseURI_NUDR_DR_v2
	// case models.ServiceName_NUDR_GROUP_ID_MAP:
	// 	suffix = ServiceBaseURI_NUDR_GROUP_ID_MAP
	case models.ServiceName_NUDR_OAM:
		suffix = ServiceBaseURI_NUDR_OAM_v1
	case models.ServiceName_NUDR_CMI:
		suffix = ServiceBaseURI_NUDR_CMI_v1
	// case models.ServiceName_NLMF_LOC:
	// 	suffix = ServiceBaseURI_NLMF_LOC
	// case models.ServiceName_N5G_EIR_EIC:
	// 	suffix = ServiceBaseURI_N5G_EIR_EIC
	// case models.ServiceName_NBSF_MANAGEMENT:
	// 	suffix = ServiceBaseURI_NBSF_MANAGEMENT
	// case models.ServiceName_NCHF_SPENDINGLIMITCONTROL:
	// 	suffix = ServiceBaseURI_NCHF_SPENDINGLIMITCONTROL
	// case models.ServiceName_NCHF_CONVERGEDCHARGING:
	// 	suffix = ServiceBaseURI_NCHF_CONVERGEDCHARGING
	// case models.ServiceName_NCHF_OFFLINEONLYCHARGING:
	// 	suffix = ServiceBaseURI_NCHF_OFFLINEONLYCHARGING
	// case models.ServiceName_NNWDAF_EVENTSSUBSCRIPTION:
	// 	suffix = ServiceBaseURI_NNWDAF_EVENTSSUBSCRIPTION
	// case models.ServiceName_NNWDAF_ANALYTICSINFO:
	// 	suffix = ServiceBaseURI_NNWDAF_ANALYTICSINFO
	// case models.ServiceName_NGMLC_LOC:
	// 	suffix = ServiceBaseURI_NGMLC_LOC
	// case models.ServiceName_NUCMF_PROVISIONING:
	// 	suffix = ServiceBaseURI_NUCMF_PROVISIONING
	// case models.ServiceName_NUCMF_UECAPABILITYMANAGEMENT:
	// 	suffix = ServiceBaseURI_NUCMF_UECAPABILITYMANAGEMENT
	// case models.ServiceName_NHSS_SDM:
	// 	suffix = ServiceBaseURI_NHSS_SDM
	// case models.ServiceName_NHSS_UECM:
	// 	suffix = ServiceBaseURI_NHSS_UECM
	// case models.ServiceName_NHSS_UEAU:
	// 	suffix = ServiceBaseURI_NHSS_UEAU
	// case models.ServiceName_NHSS_EE:
	// 	suffix = ServiceBaseURI_NHSS_EE
	case models.ServiceName_NHSS_IMS_SDM:
		suffix = ServiceBaseURI_NHSS_IMS_SDM_v1
	case models.ServiceName_NHSS_IMS_UECM:
		suffix = ServiceBaseURI_NHSS_IMS_UECM_v1
	case models.ServiceName_NHSS_IMS_UEAU:
		suffix = ServiceBaseURI_NHSS_IMS_UEAU_v1
	// case models.ServiceName_NSEPP_TELESCOPIC:
	// 	suffix = ServiceBaseURI_NSEPP_TELESCOPIC
	// case models.ServiceName_NSORAF_SOR:
	// 	suffix = ServiceBaseURI_NSORAF_SOR
	// case models.ServiceName_NSPAF_SECURED_PACKET:
	// 	suffix = ServiceBaseURI_NSPAF_SECURED_PACKET
	// case models.ServiceName_NUDSF_DR:
	// 	suffix = ServiceBaseURI_NUDSF_DR
	// case models.ServiceName_NNSSAAF_NSSAA:
	// 	suffix = ServiceBaseURI_NNSSAAF_NSSAA
	case models.ServiceName_NUPF_OAM:
		suffix = ServiceBaseURI_NUPF_OAM_v1
	case models.ServiceName_NUPF_CMI:
		suffix = ServiceBaseURI_NUPF_CMI_v1
	default:
	}
	return apiPrefix + suffix
}

func ApiVersion(name models.ServiceName) string {
	s := ServiceBaseUri(name)
	if s != "" {
		return s[strings.LastIndex(s, "/")+1:]
	}
	return ""
}

func GetServiceNfProfileAndUri(
	nfInstances []models.NrfNfDiscoveryNfProfile,
	srvName models.ServiceName,
) (*models.NrfNfDiscoveryNfProfile, string, error) {
	// select the first ServiceName
	// TODO: select base on other info
	var nfProf *models.NrfNfDiscoveryNfProfile
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
	nfProf *models.NrfNfDiscoveryNfProfile,
	srvName models.ServiceName,
) string {
	if nfProf == nil {
		return ""
	}

	nfUri := ""
	for _, srv := range nfProf.NfServices {
		if srv.ServiceName == srvName &&
			srv.NfServiceStatus == models.NfServiceStatus_REGISTERED {
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
	if s.Sst == t.Sst && strings.EqualFold(s.Sd, t.Sd) {
		return true
	}
	return false
}

func ExtSnssaiEqualFold(s, t models.ExtSnssai) bool {
	if s.Sst == t.Sst && strings.EqualFold(s.Sd, t.Sd) {
		return true
	}
	return false
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
		Sst: int32(sst),
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
