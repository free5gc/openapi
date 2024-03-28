package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSupportedFeature(t *testing.T) {
	suppFeat, err := NewSupportedFeature("03")
	assert.Nil(t, err, "")
	assert.Equal(t, SupportedFeature{0x03}, suppFeat)

	suppFeat, err = NewSupportedFeature("03FF")
	assert.Nil(t, err, "")
	assert.Equal(t, SupportedFeature{0x03, 0xFF}, suppFeat)

	suppFeat, err = NewSupportedFeature("0324")
	assert.Nil(t, err, "")
	assert.Equal(t, SupportedFeature{0x03, 0x24}, suppFeat)

	// error case
	suppFeat, err = NewSupportedFeature("ZXCD")
	assert.NotNil(t, err, "should retrun error")
	assert.Equal(t, SupportedFeature{}, suppFeat)
}
func TestGetFeatureOfSupportedFeature(t *testing.T) {
	suppFeat, err := NewSupportedFeature("1324")
	assert.Nil(t, err)

	assert.False(t, suppFeat.GetFeature(1))
	assert.False(t, suppFeat.GetFeature(2))
	assert.True(t, suppFeat.GetFeature(3))
	assert.False(t, suppFeat.GetFeature(4))

	assert.False(t, suppFeat.GetFeature(5))
	assert.True(t, suppFeat.GetFeature(6))
	assert.False(t, suppFeat.GetFeature(7))
	assert.False(t, suppFeat.GetFeature(8))

	assert.True(t, suppFeat.GetFeature(9))
	assert.True(t, suppFeat.GetFeature(10))
	assert.False(t, suppFeat.GetFeature(11))
	assert.False(t, suppFeat.GetFeature(12))

	assert.True(t, suppFeat.GetFeature(13))
	assert.False(t, suppFeat.GetFeature(14))
	assert.False(t, suppFeat.GetFeature(15))
	assert.False(t, suppFeat.GetFeature(16))
}

func TestStringOfSupportedFeature(t *testing.T) {
	suppFeat, err := NewSupportedFeature("1324")
	assert.Nil(t, err)
	assert.Equal(t, "1324", suppFeat.String())

	// testing padding
	suppFeat, err = NewSupportedFeature("1")
	assert.Nil(t, err)
	assert.Equal(t, "01", suppFeat.String())

	suppFeat, err = NewSupportedFeature("ABCDE")
	assert.Nil(t, err)
	assert.Equal(t, "0abcde", suppFeat.String())
}

func TestNegotiateWithOfSupportedFeature(t *testing.T) {
	var suppFeatA, suppFeatB, negotiatedFeat SupportedFeature
	suppFeatA, err := NewSupportedFeature("0FFF")
	assert.Nil(t, err)
	suppFeatB, err = NewSupportedFeature("1324")
	assert.Nil(t, err)
	negotiatedFeat = suppFeatA.NegotiateWith(suppFeatB)
	assert.Equal(t, SupportedFeature{0x03, 0x24}, negotiatedFeat)

	suppFeatA, err = NewSupportedFeature("0234")
	assert.Nil(t, err)
	suppFeatB, err = NewSupportedFeature("0001")
	assert.Nil(t, err)
	negotiatedFeat = suppFeatA.NegotiateWith(suppFeatB)
	assert.Equal(t, SupportedFeature{0x00, 0x00}, negotiatedFeat)

	suppFeatA, err = NewSupportedFeature("FFFF")
	assert.Nil(t, err)
	suppFeatB, err = NewSupportedFeature("F")
	assert.Nil(t, err)
	negotiatedFeat = suppFeatA.NegotiateWith(suppFeatB)
	assert.Equal(t, SupportedFeature{0x00, 0x0F}, negotiatedFeat)

	suppFeatA, err = NewSupportedFeature("3000")
	assert.Nil(t, err)
	suppFeatB, err = NewSupportedFeature("3")
	assert.Nil(t, err)
	negotiatedFeat = suppFeatA.NegotiateWith(suppFeatB)
	assert.Equal(t, SupportedFeature{0x00, 0x00}, negotiatedFeat)

	suppFeatA, err = NewSupportedFeature("23E3")
	assert.Nil(t, err)
	suppFeatB, err = NewSupportedFeature("1")
	assert.Nil(t, err)
	negotiatedFeat = suppFeatA.NegotiateWith(suppFeatB)
	assert.Equal(t, SupportedFeature{0x00, 0x01}, negotiatedFeat)
}
