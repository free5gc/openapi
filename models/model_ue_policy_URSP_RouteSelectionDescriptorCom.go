package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type RouteSelectionComponent struct {
	Identifier RouteSelComTypeId
	Value      []uint8
}

func (r *RouteSelectionComponent) GetByteLen() (uint8, error) {
	if r.Identifier == 0x00 || len(r.Value) == 0 {
		return 0, fmt.Errorf(" RouteSelectionComponent is nil, type is %v", r.Identifier)
	}
	return uint8(1 + len(r.Value)), nil
}

func (r *RouteSelectionComponent) EncodeRouteSelDescComopent(buf *bytes.Buffer) error {
	err := binary.Write(buf, binary.BigEndian, r.Identifier)
	if err != nil {
		return err
	}
	err = binary.Write(buf, binary.BigEndian, r.Value)
	if err != nil {
		return err
	}
	return nil
}

func (r *RouteSelectionComponent) SetValue(value []byte) error {
	newValue := []byte{}
	if r.Identifier == 0x00 {
		return fmt.Errorf(" Please set the Identifier first!! ")
	}

	if r.Identifier == Route_S_NSSAI_type {
		if len(value) != 6 && len(value) != 3 {
			// ref to TS24.501 v17, 9.11.2.8 S-NSSAI
			// only support SST and (SST and SD) type here
			return fmt.Errorf(" Route_S_NSSAI_type should have SST, or with SD!! ")
		}
		newValue = append(newValue, byte(len(value)))
	} else if r.Identifier == Route_DNN_type {
		// The DNN value contains an APN as defined in 3GPP TS 23.003 [4], APN(4G)=DNN(5G)
		// The DNN is compose of "APN Network Id(Mand)" and "APN Operator Identifier(Opt)"
		if len(value) == 0 {
			return fmt.Errorf(" Route_DNN_type should not be empty!! ")
		}
		newValue = append(newValue, byte(len(value)))
	} else if r.Identifier == Route_PDU_session_type_type {
		// be encoded as a one octet PDU session type field. The bits 8 through 4 of the octet
		// shall be spare, and the bits 3 through 1 shall be encoded as the value part of the PDU
		// session type information element defined in clause 9.11.4.11 of 3GPP TS 24.501 [11].
		if len(value) != 1 {
			return fmt.Errorf(" Route_PDU_session_type_type should contain only 1 byte")
		}
	} else {
		return fmt.Errorf("[RouteSelectionComponent]-The %v type is not yet supported!! ", r.Identifier)
	}
	newValue = append(newValue, value...)
	r.Value = newValue
	return nil
}

func (r *RouteSelectionComponent) MakeByte_SNSSAI(sst uint8, sd []uint8) []byte {
	// ref to: TS24.501 v17, 9.11.2.8 S-NSSAI
	ret := []byte{}

	// "S-NSSAI IEI", 1 Octet,
	// TODO: I dont know how to assign this value
	ret = append(ret, 0x00)

	// "Length of S-NSSAI contents", 1 Octet
	if sd == nil {
		ret = append(ret, 0x01)
	} else {
		ret = append(ret, 0x04)
	}

	// "SST", 1 Octet
	ret = append(ret, sst)

	// "SD", 3 Octet
	if sd != nil {
		ret = append(ret, sd...)
	}
	return ret
}

func (r *RouteSelectionComponent) MakeByte_DNN(dnn string) []byte {
	// ref to: TS 23003 R17, 9A Definition of Data Network Name
	// TODO: Let the format of DNN be the same as in the spec
	return []byte(dnn)
}

func (r *RouteSelectionComponent) MakeByte_PDUsessType(pDUtpye string) ([]byte, error) {
	lowType := strings.ToLower(pDUtpye)

	if lowType == "ipv4" {
		return []byte{0x01}, nil
	} else if lowType == "ipv6" {
		return []byte{0x02}, nil
	} else if lowType == "ipv4v6" {
		return []byte{0x03}, nil
	} else if lowType == "unstructured" {
		return []byte{0x04}, nil
	} else if lowType == "ethernet" {
		return []byte{0x05}, nil
	} else if lowType == "reserved" {
		return []byte{0x07}, nil
	}
	// All other values are unused and shall be interpreted as "IPv4v6", if received by the UE or the network.
	return []byte{0x03}, nil

	// return nil, fmt.Errorf("pdu session type doesn't follow the format!! ")
}

func (r *RouteSelectionComponent) DecodeRouSelDescComp(buf *bytes.Buffer) error {
	// read component Identifier
	if err := binary.Read(buf, binary.BigEndian, &r.Identifier); err != nil {
		return err
	}
	// read component value, case by case
	switch r.Identifier {
	case Route_S_NSSAI_type:
		// read the SNSSAI length first
		var snssaiLen uint8
		if err := binary.Read(buf, binary.BigEndian, &snssaiLen); err != nil {
			return err
		}
		r.Value = make([]byte, snssaiLen+1)
		r.Value[0] = snssaiLen
		// read the SNSSAI value
		if err := binary.Read(buf, binary.BigEndian, r.Value[1:snssaiLen+1]); err != nil {
			return err
		}
	case Route_DNN_type:
		var dnnLen uint8
		if err := binary.Read(buf, binary.BigEndian, &dnnLen); err != nil {
			return err
		}
		r.Value = make([]byte, dnnLen+1)
		r.Value[0] = dnnLen
		// read the dnn value
		if err := binary.Read(buf, binary.BigEndian, r.Value[1:dnnLen+1]); err != nil {
			return err
		}
	case Route_PDU_session_type_type:
		r.Value = make([]byte, 1)
		if err := binary.Read(buf, binary.BigEndian, r.Value[:1]); err != nil {
			return err
		}
	default:
		return fmt.Errorf("[DecodeRouSelDescComp] The %v type is not yet supported!! ", r.Identifier)
	}
	// fmt.Printf("[DecodeRouSelDescComp] decode success:%+v\n", *r)
	return nil
}
