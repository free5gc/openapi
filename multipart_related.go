package openapi

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"regexp"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/pkg/errors"
)

type MultipartRelatedBinding struct{}

func (MultipartRelatedBinding) Name() string {
	return "json"
}

func (MultipartRelatedBinding) Bind(req *http.Request, obj interface{}) error {
	b, err := io.ReadAll(req.Body)
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
	boundary, reqBodyBytes, err := multipart.Marshal(r.Data)
	if err != nil {
		return errors.Wrap(err, "multipart rending fail")
	}
	contentType := mime.FormatMediaType("multipart/related", map[string]string{"boundary": boundary})
	r.contentType = contentType
	w.Header().Set("Content-Type", r.contentType)
	_, err = bytes.NewBuffer(reqBodyBytes).WriteTo(w)
	if err != nil {
		return errors.Wrap(err, "multipart rending fail")
	}
	return
}

func (r MultipartRelatedRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.contentType)
}
