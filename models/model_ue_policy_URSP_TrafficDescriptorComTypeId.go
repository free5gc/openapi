package models

// 3GPP TS 24.526 V17.8.0, Table 5.2.1: UE policy part contents including a URSP rule
type TrafDesComTypeId uint8

// All other values are spare. If received they shall be interpreted as unknown.
const (
	Traf_Match_all_type                   TrafDesComTypeId = 0x01
	Traf_OSId_plus_OSAppId_type           TrafDesComTypeId = 0x08
	Traf_IPv4_remote_addr_type            TrafDesComTypeId = 0x10
	Traf_IPv6_remote_addr_Prefix_len_type TrafDesComTypeId = 0x21
	Traf_ProtocolId_Next_header_type      TrafDesComTypeId = 0x30
	Traf_Single_remote_port_type0         TrafDesComTypeId = 0x50
	Traf_Remote_port_range_type           TrafDesComTypeId = 0x51
	Traf_IP_3_tuple_type                  TrafDesComTypeId = 0x52
	Traf_Secur_para_idx_type              TrafDesComTypeId = 0x60
	Traf_TypeOf_service_Traf_class_type   TrafDesComTypeId = 0x70
	Traf_Flow_label_type                  TrafDesComTypeId = 0x80
	Traf_Dst_MAC_addr_type                TrafDesComTypeId = 0x81
	Traf_802_1Q_C_TAG_VID_type            TrafDesComTypeId = 0x83
	Traf_802_1Q_S_TAG_VID_type            TrafDesComTypeId = 0x84
	Traf__802_1Q_C_TAG_PCP_DEI_type       TrafDesComTypeId = 0x85
	Traf_802_1Q_S_TAG_PCP_DEI_type        TrafDesComTypeId = 0x86
	Traf_Ethertype_type                   TrafDesComTypeId = 0x87
	Traf_DNN_type                         TrafDesComTypeId = 0x88
	Traf_Conn_capabilities_type           TrafDesComTypeId = 0x90
	Traf_Dst_FQDN                         TrafDesComTypeId = 0x91
	Traf_Regular_expression               TrafDesComTypeId = 0x92
	Traf_OS_App_Id_type                   TrafDesComTypeId = 0xA0
	Traf_Dst_MAC_addr_range_type          TrafDesComTypeId = 0xA1
)

// func (t *TrafDesComTypeId) SetType(typeId uint8) {
// 	t = typeId
// }

// func (t *TrafDesComTypeId) GetType() TrafDesComTypeId {
// 	return t
// }
