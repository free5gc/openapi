package openapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

type MultipartRelatedBinding struct{}

func (MultipartRelatedBinding) Name() string {
	return "json"
}

func (MultipartRelatedBinding) Bind(req *http.Request, obj interface{}) error {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return Deserialize(obj, b, req.Header.Get("Content-Type"))
}

func (MultipartRelatedBinding) BindBody(body []byte, obj interface{}) error {
	re, err := regexp.Compile(`--([a-zA-Z0-9+\-_]+)--`)
	if err != nil {
		return err
	}
	submatch := re.FindSubmatch(body)
	if len(submatch) < 1 {
		return errors.New("cannot parse multipart boundary")
	}
	return Deserialize(obj, body, "multipart/related; boundary="+string(submatch[1]))
}

type MultipartRelatedRender struct {
	Data        interface{}
	contentType string
}

func (r MultipartRelatedRender) Render(w http.ResponseWriter) (err error) {
	payloadBuf := &bytes.Buffer{}
	ct, err := MultipartEncode(r.Data, payloadBuf)
	if err != nil {
		return errors.Wrap(err, "multipart rending fail")
	}
	r.contentType = ct
	w.Header().Set("Content-Type", r.contentType)
	_, err = payloadBuf.WriteTo(w)
	if err != nil {
		return errors.Wrap(err, "multipart rending fail")
	}
	return
}

func (r MultipartRelatedRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.contentType)
}
