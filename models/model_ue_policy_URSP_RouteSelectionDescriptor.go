package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type RouteSelectionDescriptor struct {
	Length                uint8                     `json:"length" bson:"length"`
	PrecedenceValue       uint8                     `json:"precedenceValue" bson:"precedenceValue"`
	LengthContent         uint8                     `json:"LengthContent" bson:"LengthContent"`
	RouteSelectionContent []RouteSelectionComponent `json:"routeSelectionContent" bson:"routeSelectionContent"`
}

func (r *RouteSelectionDescriptor) GetByteLen() (uint8, error) {
	r.LengthContent = 0
	for _, routeSelCom := range r.RouteSelectionContent {
		len, err := routeSelCom.GetByteLen()
		if err != nil {
			return 0, err
		}
		r.LengthContent += len
	}
	// byte length of content and precedence and lengthcontent
	r.Length = 2 + r.LengthContent

	// because this func would be used to count total byte length, plus 1 means IE:"Length"
	return r.Length + 1, nil
}

func (r *RouteSelectionDescriptor) EncodeRouteSelectionDescriptor(buf *bytes.Buffer) error {
	// get the total byte length of this RouteSelect Descriptor first
	if _, err := r.GetByteLen(); err != nil {
		return err
	}

	if err := binary.Write(buf, binary.BigEndian, r.Length); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, r.PrecedenceValue); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, r.LengthContent); err != nil {
		return err
	}

	for _, rouSelDescCom := range r.RouteSelectionContent {
		if err := rouSelDescCom.EncodeRouteSelDescComopent(buf); err != nil {
			return err
		}
	}

	return nil
}

func (r *RouteSelectionDescriptor) DecodeRouteSelDesc(buf *bytes.Buffer) error {
	// read length of this route select Descriptor
	if err := binary.Read(buf, binary.BigEndian, &r.Length); err != nil {
		return err
	}
	// create a new buffer to decode advacing IE
	byRouSelDesc := buf.Next(int(r.Length))
	if len(byRouSelDesc) != int(r.Length) {
		return fmt.Errorf("length of buffer is less than length of RouteSelectionDescriptor")
	}
	newbuf := bytes.NewBuffer(byRouSelDesc)

	// read precedence
	if err := binary.Read(newbuf, binary.BigEndian, &r.PrecedenceValue); err != nil {
		return err
	}
	// read Length of route select descriptor Content
	if err := binary.Read(newbuf, binary.BigEndian, &r.LengthContent); err != nil {
		return err
	}
	// read route select descriptor Content
	for {
		var rouSelDescCom RouteSelectionComponent
		if err := rouSelDescCom.DecodeRouSelDescComp(newbuf); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		r.RouteSelectionContent = append(r.RouteSelectionContent, rouSelDescCom)
	}
}
