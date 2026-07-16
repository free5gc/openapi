package multimodels

import (
	"errors"
	"reflect"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

func readN2InformationFromRelatedContent(n2InfoContainer *models.Amf_Comm_N2InfoContainer, from []reflect.Value) ([]byte, error) {
	if n2InfoContainer == nil {
		return nil, errors.New("n2InfoContainer is nil")
	}
	if n2InfoContainer.N2InformationClass == "" {
		return nil, errors.New("n2InfoContainer.N2InformationClass is empty")
	}

	var n2Information []byte

	switch n2InfoContainer.N2InformationClass {
	case models.Amf_Comm_N2InformationClass_SM:
		if n2InfoContainer.SmInfo != nil {
			contentId := n2InfoContainer.SmInfo.N2InfoContent.NgapData.ContentId
			if contentId == "" {
				return nil, errors.New("n2InfoContainer.SmInfo.N2InfoContent.NgapData.ContentId is empty")
			}
			content := multipart.MatchRelatedContent(contentId, from)
			if content != nil {
				n2Information = content
			}
		}
	case models.Amf_Comm_N2InformationClass_RAN:
		if n2InfoContainer.RanInfo != nil {
			contentId := n2InfoContainer.RanInfo.N2InfoContent.NgapData.ContentId
			if contentId == "" {
				return nil, errors.New("n2InfoContainer.RanInfo.N2InfoContent.NgapData.ContentId is empty")
			}
			content := multipart.MatchRelatedContent(contentId, from)
			if content != nil {
				n2Information = content
			}
		}
	case models.Amf_Comm_N2InformationClass_V2_X:
		if n2InfoContainer.V2xInfo != nil {
			contentId := n2InfoContainer.V2xInfo.N2Pc5Pol.NgapData.ContentId
			if contentId == "" {
				return nil, errors.New("n2InfoContainer.V2xInfo.N2Pc5Pol.NgapData.ContentId is empty")
			}
			content := multipart.MatchRelatedContent(contentId, from)
			if content != nil {
				n2Information = content
			}
		}
	case models.Amf_Comm_N2InformationClass_PROSE:
		if n2InfoContainer.ProseInfo != nil {
			contentId := n2InfoContainer.ProseInfo.N2Pc5ProSePol.NgapData.ContentId
			if contentId == "" {
				return nil, errors.New("n2InfoContainer.ProseInfo.N2Pc5ProSePol.NgapData.ContentId is empty")
			}
			content := multipart.MatchRelatedContent(contentId, from)
			if content != nil {
				n2Information = content
			}
		}
	case models.Amf_Comm_N2InformationClass_NRP_PA:
		if n2InfoContainer.NrppaInfo != nil {
			contentId := n2InfoContainer.NrppaInfo.NrppaPdu.NgapData.ContentId
			if contentId == "" {
				return nil, errors.New("n2InfoContainer.NrppaInfo.NrppaPdu.NgapData.ContentId is empty")
			}
			content := multipart.MatchRelatedContent(contentId, from)
			if content != nil {
				n2Information = content
			}
		}
	case models.Amf_Comm_N2InformationClass_PWS_BCAL:
		fallthrough
	case models.Amf_Comm_N2InformationClass_PWS_RF:
		fallthrough
	case models.Amf_Comm_N2InformationClass_PWS:
		if n2InfoContainer.PwsInfo != nil {
			contentId := n2InfoContainer.PwsInfo.PwsContainer.NgapData.ContentId
			if contentId == "" {
				return nil, errors.New("n2InfoContainer.PwsInfo.PwsContainer.NgapData.ContentId is empty")
			}
			content := multipart.MatchRelatedContent(contentId, from)
			if content != nil {
				n2Information = content
			}
		}
	}
	return n2Information, nil
}

func putN2InformationToRelatedContent(n2InfoContainer *models.Amf_Comm_N2InfoContainer, n2Information []byte, to reflect.Value) error {
	if n2InfoContainer == nil {
		return errors.New("n2InfoContainer is nil")
	}
	if n2InfoContainer.N2InformationClass == "" {
		return errors.New("n2InfoContainer.N2InformationClass is empty")
	}
	switch n2InfoContainer.N2InformationClass {
	case models.Amf_Comm_N2InformationClass_SM:
		if n2InfoContainer.SmInfo != nil {
			contentId := n2InfoContainer.SmInfo.N2InfoContent.NgapData.ContentId
			if contentId == "" {
				return errors.New("n2InfoContainer.SmInfo.N2InfoContent.NgapData.ContentId is empty")
			}
			to.Set(reflect.ValueOf(&multipart.RelatedContent{
				ContentID: contentId,
				Content:   n2Information,
			}))
		}
	case models.Amf_Comm_N2InformationClass_RAN:
		if n2InfoContainer.RanInfo != nil {
			contentId := n2InfoContainer.RanInfo.N2InfoContent.NgapData.ContentId
			if contentId == "" {
				return errors.New("n2InfoContainer.RanInfo.N2InfoContent.NgapData.ContentId is empty")
			}
			to.Set(reflect.ValueOf(&multipart.RelatedContent{
				ContentID: contentId,
				Content:   n2Information,
			}))
		}
	case models.Amf_Comm_N2InformationClass_V2_X:
		if n2InfoContainer.V2xInfo != nil {
			contentId := n2InfoContainer.V2xInfo.N2Pc5Pol.NgapData.ContentId
			if contentId == "" {
				return errors.New("n2InfoContainer.V2xInfo.N2Pc5Pol.NgapData.ContentId is empty")
			}
			to.Set(reflect.ValueOf(&multipart.RelatedContent{
				ContentID: contentId,
				Content:   n2Information,
			}))
		}
	case models.Amf_Comm_N2InformationClass_PROSE:
		if n2InfoContainer.ProseInfo != nil {
			contentId := n2InfoContainer.ProseInfo.N2Pc5ProSePol.NgapData.ContentId
			if contentId == "" {
				return errors.New("n2InfoContainer.ProseInfo.N2Pc5ProSePol.NgapData.ContentId is empty")
			}
			to.Set(reflect.ValueOf(&multipart.RelatedContent{
				ContentID: contentId,
				Content:   n2Information,
			}))
		}
	case models.Amf_Comm_N2InformationClass_NRP_PA:
		if n2InfoContainer.NrppaInfo != nil {
			contentId := n2InfoContainer.NrppaInfo.NrppaPdu.NgapData.ContentId
			if contentId == "" {
				return errors.New("n2InfoContainer.NrppaInfo.NrppaPdu.NgapData.ContentId is empty")
			}
			to.Set(reflect.ValueOf(&multipart.RelatedContent{
				ContentID: contentId,
				Content:   n2Information,
			}))
		}
	case models.Amf_Comm_N2InformationClass_PWS_BCAL:
		fallthrough
	case models.Amf_Comm_N2InformationClass_PWS_RF:
		fallthrough
	case models.Amf_Comm_N2InformationClass_PWS:
		if n2InfoContainer.PwsInfo != nil {
			contentId := n2InfoContainer.PwsInfo.PwsContainer.NgapData.ContentId
			if contentId == "" {
				return errors.New("n2InfoContainer.PwsInfo.PwsContainer.NgapData.ContentId is empty")
			}
			to.Set(reflect.ValueOf(&multipart.RelatedContent{
				ContentID: contentId,
				Content:   n2Information,
			}))
		}
	}
	return nil
}
