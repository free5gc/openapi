package models

type RouteSelComTypeId uint8

const (
	Route_SSC_mode_type                                  RouteSelComTypeId = 0x01
	Route_S_NSSAI_type                                   RouteSelComTypeId = 0x02
	Route_DNN_type                                       RouteSelComTypeId = 0x04
	Route_PDU_session_type_type                          RouteSelComTypeId = 0x08
	Route_Preferred_access_type_type                     RouteSelComTypeId = 0x10
	Route_Multi_access_preference_type                   RouteSelComTypeId = 0x11
	Route_Nonseamless_Non3GPP_off_indi_type              RouteSelComTypeId = 0x20
	Route_Location_criteria_type                         RouteSelComTypeId = 0x40
	Route_Time_window_type                               RouteSelComTypeId = 0x80
	Route_NR5G_ProSe_lay3_UE2network_relay_off_indi_type RouteSelComTypeId = 0x81
	Route_PDU_session_pair_ID_type                       RouteSelComTypeId = 0x82
	Route_RSN_type                                       RouteSelComTypeId = 0x83
)

func (r *RouteSelComTypeId) SetTypeId(typeId RouteSelComTypeId) {
	*r = typeId
}
