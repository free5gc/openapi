package openapi

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Regression test: all HTTP clients must wrap their transport with otelhttp.NewTransport
// so that trace IDs are propagated in outgoing requests. If this breaks,
// downstream services will not receive the traceparent header and cross-service tracing fails.
func TestTransport_WrapsWithOtelTransport(t *testing.T) {
	t.Run("HTTP2CleartextClient.SetTransport", func(t *testing.T) {
		h2c := &HTTP2CleartextClient{client: &http.Client{}}
		h2c.SetTransport(1*time.Second, 2*time.Second, nil)

		_, ok := h2c.client.Transport.(*otelhttp.Transport)
		require.True(t, ok,
			"Transport must be *otelhttp.Transport to propagate trace IDs; "+
				"raw http2.Transport was set instead")
	})

	t.Run("NewHttp2Client", func(t *testing.T) {
		c := NewHttp2Client(false, 5*time.Second, 1*time.Second, 2*time.Second)

		_, ok := c.Transport.(*otelhttp.Transport)
		require.True(t, ok,
			"Transport must be *otelhttp.Transport to propagate trace IDs; "+
				"raw http2.Transport was set instead")
	})
}

func TestParameterToString(t *testing.T) {
	var testCases = []struct {
		name   string
		obj    interface{}
		format string
		para   string
	}{
		{
			name:   "strings csv",
			obj:    []string{"red", "blue", "yellow"},
			format: "csv",
			para:   "red,blue,yellow",
		},
		{
			name:   "strings pipes",
			obj:    []string{"red", "blue", "yellow"},
			format: "pipes",
			para:   "red|blue|yellow",
		},
		{
			name:   "strings ssv",
			obj:    []string{"red", "blue", "yellow"},
			format: "ssv",
			para:   "red blue yellow",
		},
		{
			name:   "strings json",
			obj:    []string{"red", "blue", "yellow"},
			format: "application/json",
			para:   "[\"red\",\"blue\",\"yellow\"]",
		},
		{
			name: "obj json",
			obj: struct {
				Name   string `json:"name"`
				Age    int    `json:"age"`
				Active bool   `json:"active"`
			}{
				Name:   "free5GC",
				Age:    3,
				Active: true,
			},
			format: "application/json",
			para:   "{\"name\":\"free5GC\",\"age\":3,\"active\":true}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			para := ParameterToString(tc.obj, tc.format)
			require.Equal(t, tc.para, para)
		})
	}
}

func getStrPtr(s string) *string {
	return &s
}

func TestAddQueryParams(t *testing.T) {
	var testCases = []struct {
		name             string
		obj              any
		format           string
		queryParamString string
	}{
		{
			name:             "string multi",
			obj:              "red",
			format:           "multi",
			queryParamString: "color=red",
		},
		{
			name:             "string ptr multi",
			obj:              getStrPtr("red"),
			format:           "multi",
			queryParamString: "color=red",
		},
		{
			name:             "strings multi",
			obj:              []string{"red", "blue", "yellow"},
			format:           "multi",
			queryParamString: "color=red&color=blue&color=yellow",
		},
		{
			name:             "strings csv",
			obj:              []string{"red", "blue", "yellow"},
			format:           "csv",
			queryParamString: "color=" + url.QueryEscape("red,blue,yellow"),
		},
		{
			name:             "strings pipes",
			obj:              []string{"red", "blue", "yellow"},
			format:           "pipes",
			queryParamString: "color=" + url.QueryEscape("red|blue|yellow"),
		},
		{
			name:             "strings ssv",
			obj:              []string{"red", "blue", "yellow"},
			format:           "ssv",
			queryParamString: "color=" + url.QueryEscape("red blue yellow"),
		},
		{
			name:             "strings json",
			obj:              []string{"red", "blue", "yellow"},
			format:           "application/json",
			queryParamString: "color=" + url.QueryEscape(`["red","blue","yellow"]`),
		},
		{
			name: "obj json",
			obj: struct {
				Red    int `json:"red"`
				Blue   int `json:"blue"`
				Yellow int `json:"yellow"`
			}{
				Red:    255,
				Blue:   255,
				Yellow: 255,
			},
			format:           "application/json",
			queryParamString: "color=" + url.QueryEscape(`{"red":255,"blue":255,"yellow":255}`),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			queryParams := url.Values{}
			err := AddQueryParams(&queryParams, "color", tc.obj, tc.format)
			require.NoError(t, err)
			require.Equal(t, tc.queryParamString, queryParams.Encode())
		})
	}
}

func TestMultipartDeserialize_LargerThanBuffer(t *testing.T) {
	reqJsonBody := `{
		"n1MessageContainer":{
			"n1MessageClass":"5GMM",
			"n1MessageContent":{
				"contentId":"n1Msg"
			}
		},
		"registrationCtxtContainer":{
			"ueContext":{
				"supi":"imsi-2089300007487",
				"pei":"imeisv-1110000000000000",
				"seafData":{
					"ngKsi":{
						"tsc":"NATIVE",
						"ksi":0
					},
					"keyAmf":{
						"keyType":"",
						"keyVal":"d04d0dfb241eee4e7fc070866ee9be5724e74a362efade8b6d3e84defd03195c"
					},
					"nh":"f6cdd59292429570efc9558a548f244a8bb542e75b4b48acc59080755432fa34",
					"ncc":1
				}
			},
			"anType":"3GPP_ACCESS",
			"anN2ApId":1,
			"ranNodeId":{
				"plmnId":{
					"mcc":"208",
					"mnc":"93"
				},
				"gNbId":{
					"bitLength":22,
					"gNBValue":"000001"
				}
			},
			"initialAmfName":"amf1",
			"userLocation":{
				"nrLocation":{
					"tai":{
						"plmnId":{
							"mcc":"208",
							"mnc":"93"
						},
						"tac":"000001"
					},
					"ncgi":{
						"plmnId":{
							"mcc":"208",
							"mnc":"93"
						},
						"nrCellId":"000004001"
					},
					"ueLocationTimestamp":"2024-04-12T09:54:54.152026945Z"
				}
			},
			"rrcEstCause":"2",
			"anN2IPv4Addr":"172.16.21.2:39249",
			"allowedNssai":{
				"allowedSnssaiList":[
					{
						"allowedSnssai":{
							"sst":1,
							"sd":"111111"
						},
						"nsiInformationList":[
							{
								"nrfId":"http://10.100.200.10:8000/nnrf-nfm/v1/nf-instances",
								"nsiId":"2"
							}
						]
					}
				],
				"accessType":"3GPP_ACCESS"
			}
		}
	}`
	reqBody := "--D4/?XquckV=bx9{dC}EG[TUQEzEb.'(CG(k8,B:D++6w3M5qJ?P]B]-Xk" +
		"\r\n" +
		"Content-Type: application/json" +
		"\r\n\r\n" +
		reqJsonBody +
		"\r\n" +
		"--D4/?XquckV=bx9{dC}EG[TUQEzEb.'(CG(k8,B:D++6w3M5qJ?P]B]-Xk--" +
		"\r\n"
	fmt.Println(reqBody)
	notify := models.N1MessageNotifyRequestBody{}
	boundary := "D4/?XquckV=bx9{dC}EG[TUQEzEb.'(CG(k8,B:D++6w3M5qJ?P]B]-Xk"
	err := multipart.Unmarshal(boundary, []byte(reqBody), &notify)
	require.NoError(t, err)
	require.NotNil(t, notify.JsonData)
	require.NotNil(t, notify.JsonData.RegistrationCtxtContainer)
	require.NotNil(t, notify.JsonData.N1MessageContainer)
}
