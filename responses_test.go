package responses

import (
	"net/http"
	"testing"
)

type TestResponseWriter struct{}

func (t TestResponseWriter) Header() http.Header {
	return http.Header{}
}

func (t TestResponseWriter) Write(data []byte) (int, error) {
	return 0, nil
}

func (t TestResponseWriter) WriteHeader(statusCode int) {}

func TestWrite(t *testing.T) {
	w := TestResponseWriter{}
	err := Write(w, "daldalerm")
	if err != nil {
		t.Errorf("%v", err)
	}
	err = Write(w, 1234)
	if err != nil {
		t.Errorf("%v", err)
	}
	err = Write(w, []byte("test"))
	if err != nil {
		t.Errorf("%v", err)
	}
	err = Write(w, 1234.4)
	if err != nil {
		t.Errorf("%v", err)

	}
}
