package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

//ref 3GPP TS 24.526 V17.8.0 (2022-12)--section 5.2 Encoding of UE policy part type URSP
//it's a encoding strcut for URSP to transmit to UE

type UePolicyURSP struct {
	URSPruleSet []URSPrule `json:"uRSPruleSet" bson:"uRSPruleSet"`
}

func (u *UePolicyURSP) EncodeURSP() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	for i, URSPrule := range u.URSPruleSet {
		err := URSPrule.EncodeURSPrule(buffer)
		if err != nil {
			return nil, err
		}
		u.URSPruleSet[i] = URSPrule
	}
	return buffer.Bytes(), nil
}

func (u *UePolicyURSP) DecodeURSP(b []byte) error {
	buf := bytes.NewBuffer(b)

	for {
		byUrspRule, ruleLen, err := parseURSPrule(buf)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		var urspRule URSPrule
		err = urspRule.DecodeURSPrule(byUrspRule, ruleLen)
		if err != nil {
			return err
		}
		u.URSPruleSet = append(u.URSPruleSet, urspRule)
	}
}

func parseURSPrule(buf *bytes.Buffer) ([]byte, uint8, error) {
	var urspLen uint8

	if err := binary.Read(buf, binary.BigEndian, &urspLen); err != nil {
		return nil, 0, err
	}

	// return byteslice of a single URSP rule
	readByte := buf.Next(int(urspLen))
	// if buffer length is less than  URSP rule length
	if len(readByte) != int(urspLen) {
		return nil, 0, fmt.Errorf("parsing a single URSPrule error, buffer length is less than  URSP rule length")
	}
	return readByte, urspLen, nil
}
