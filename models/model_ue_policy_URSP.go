package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

//ref 3GPP TS 24.526 V17.8.0 (2022-12)--section 5.2 Encoding of UE policy part type URSP
//it's a encoding strcut for URSP to transmit to UE

type URSPrule struct {
	Length                             uint8                        `json:"length" bson:"length"`
	PrecedenceValue                    uint8                        `json:"precedenceValue" bson:"precedenceValue"`
	LengthTrafficDescriptor            uint8                        `json:"lengthTrafficDescriptor" bson:"lengthTrafficDescriptor"`
	TrafficDescriptor                  []TrafficDescriptorComponent `json:"trafficDescriptor" bson:"trafficDescriptor"`
	LengthRouteSelectionDescriptorList uint8                        `json:"lengthRouteSelectionDescriptorList" bson:"lengthRouteSelectionDescriptorList"`
	RouteSelectionDescriptorList       []RouteSelectionDescriptor   `json:"routeSelectionDescriptorList" bson:"routeSelectionDescriptorList"`
}

func (u *URSPrule) CntURSPlen() error {
	u.LengthTrafficDescriptor = 0
	for _, trafDesCom := range u.TrafficDescriptor {
		len, err := trafDesCom.GetByteLen()
		if err != nil {
			return err
		}
		u.LengthTrafficDescriptor += len
	}
	u.LengthRouteSelectionDescriptorList = 0
	for _, routeSelDesList := range u.RouteSelectionDescriptorList {
		len, err := routeSelDesList.GetByteLen()
		if err != nil {
			return err
		}
		u.LengthRouteSelectionDescriptorList += len
	}

	u.Length = 3 + u.LengthTrafficDescriptor + u.LengthRouteSelectionDescriptorList
	return nil
}

func (u *URSPrule) EncodeURSPrule(buf *bytes.Buffer) error {
	// count total byte length of this URSP rule first then encoding in next step
	if err := u.CntURSPlen(); err != nil {
		return err
	}

	if err := binary.Write(buf, binary.BigEndian, u.Length); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, u.PrecedenceValue); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, u.LengthTrafficDescriptor); err != nil {
		return err
	}
	for _, trafDesCom := range u.TrafficDescriptor {
		err := trafDesCom.EncodeTrafficDescriptorComponent(buf)
		if err != nil {
			return err
		}
	}
	if err := binary.Write(buf, binary.BigEndian, u.LengthRouteSelectionDescriptorList); err != nil {
		return err
	}
	for i, routeSelDesList := range u.RouteSelectionDescriptorList {
		err := routeSelDesList.EncodeRouteSelectionDescriptor(buf)
		if err != nil {
			return err
		}
		u.RouteSelectionDescriptorList[i] = routeSelDesList
	}

	return nil
}

func (u *URSPrule) DecodeURSPrule(b []byte, ruleLen uint8) error {
	buf := bytes.NewBuffer(b)
	// this IE have been deocde in prevoius caller func
	u.Length = ruleLen

	if err := binary.Read(buf, binary.BigEndian, &u.PrecedenceValue); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &u.LengthTrafficDescriptor); err != nil {
		return err
	}
	// decode Traffic Desc to multi component
	if err := u.decodeTrafDesc(buf); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &u.LengthRouteSelectionDescriptorList); err != nil {
		return err
	}
	// decode Route Desc to multi component
	if err := u.decodeRouteSelDescLs(buf); err != nil {
		return err
	}
	return nil
}

func (u *URSPrule) decodeTrafDesc(buf *bytes.Buffer) error {
	byTrafDesc := buf.Next(int(u.LengthTrafficDescriptor))
	if len(byTrafDesc) != int(u.LengthTrafficDescriptor) {
		return fmt.Errorf("the buffer len is less than LengthTrafficDescriptor")
	}

	newbuf := bytes.NewBuffer(byTrafDesc)
	for {
		var trafDescComp TrafficDescriptorComponent
		if err := trafDescComp.DecodeTrafDescComp(newbuf); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		u.TrafficDescriptor = append(u.TrafficDescriptor, trafDescComp)
	}
}

func (u *URSPrule) decodeRouteSelDescLs(buf *bytes.Buffer) error {
	byRouteSelDescLs := buf.Next(int(u.LengthRouteSelectionDescriptorList))
	if len(byRouteSelDescLs) != int(u.LengthRouteSelectionDescriptorList) {
		return fmt.Errorf("the buffer len is less than LengthRouteSelectionDescriptorList")
	}

	newbuf := bytes.NewBuffer(byRouteSelDescLs)
	for {
		var routeSelDesc RouteSelectionDescriptor
		if err := routeSelDesc.DecodeRouteSelDesc(newbuf); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		u.RouteSelectionDescriptorList = append(u.RouteSelectionDescriptorList, routeSelDesc)
	}
}
