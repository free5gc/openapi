package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

type TrafficDescriptorComponent struct {
	Identifier TrafDesComTypeId `json:"identifier" bson:"identifier"`
	Value      []uint8          `json:"value" bson:"value"`
}

func (t *TrafficDescriptorComponent) GetByteLen() (uint8, error) {
	if t.Identifier == 0x00 || (len(t.Value) == 0 && t.Identifier != Traf_Match_all_type) {
		return 0, fmt.Errorf(" TrafficDescriptorComponent is nil!! ")
	}
	return 1 + uint8(len(t.Value)), nil
}

func (t *TrafficDescriptorComponent) EncodeTrafficDescriptorComponent(buf *bytes.Buffer) error {
	err := binary.Write(buf, binary.BigEndian, t.Identifier)
	if err != nil {
		return err
	}
	err = binary.Write(buf, binary.BigEndian, t.Value)
	if err != nil {
		return err
	}

	return nil
}

func (t *TrafficDescriptorComponent) SetValue(value []byte) error {
	if t.Identifier == 0x00 {
		return fmt.Errorf("please set the Identifier first")
	}

	if t.Identifier == Traf_Match_all_type {
		if value != nil {
			return fmt.Errorf(" Match_all_type should not include any value field!! ")
		}
		return nil
	} else if t.Identifier == Traf_IPv4_remote_addr_type {
		if len(value) != 8 {
			return fmt.Errorf(" IP value should contain 4 Octet IP and 4 Octet mask!! ")
		}
		t.Value = value
		return nil
	} else if t.Identifier == Traf_Dst_MAC_addr_type {
		if len(value) != 6 {
			return fmt.Errorf(" Dst MAC value should contain 6 Octets!! ")
		}
		t.Value = value
	} else {
		// TODO: add other format error checking
		return fmt.Errorf("[TrafficDescriptorComponent]- The %v type is not yet supported!! ", t.Identifier)
	}
	return nil
}

func (t *TrafficDescriptorComponent) IPv4remote_MakeByte(ip, mask string) ([]byte, error) {
	ret := []byte{}

	// IP field
	ipArr := strings.Split(ip, ".")
	if len(ipArr) != 4 {
		return nil, fmt.Errorf("the IP address should contain 4 sections")
	}
	for _, ipSec := range ipArr {
		ipSecInt, err := strconv.Atoi(ipSec)
		if err != nil {
			return nil, fmt.Errorf("the Transfer of IP from string to int occurs error")
		}
		ret = append(ret, byte(ipSecInt))
	}

	// Mask field
	maskArr := strings.Split(mask, ".")
	if len(maskArr) != 4 {
		return nil, fmt.Errorf("the IP mask address should contain 4 sections")
	}
	for _, maskSec := range maskArr {
		maskSecInt, err := strconv.Atoi(maskSec)
		if err != nil {
			return nil, fmt.Errorf("the Transfer of IP mask from string to int occurs error")
		}
		ret = append(ret, byte(maskSecInt))
	}
	return ret, nil
}

func (t *TrafficDescriptorComponent) DecodeTrafDescComp(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &t.Identifier); err != nil {
		return err
	}

	switch t.Identifier {
	case Traf_Match_all_type:
		t.Value = make([]byte, 0)
	case Traf_IPv4_remote_addr_type:
		t.Value = make([]byte, 8)
		if err := binary.Read(buf, binary.BigEndian, t.Value[:len(t.Value)]); err != nil {
			return err
		}
	case Traf_Dst_MAC_addr_type:
		t.Value = make([]byte, 6)
		if err := binary.Read(buf, binary.BigEndian, t.Value[:len(t.Value)]); err != nil {
			return err
		}
	default:
		// TODO: add other foramt byte slice read
		return fmt.Errorf("[TrafficDescriptorComponent] the format of Traffic descriptor is not supported yet: %b", t.Identifier)
	}
	return nil
}
