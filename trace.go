package sg

import (
	"net/http"
	"net/http/httputil"
)

// Tracer is the implementation of a tracer that can print debug information of
// a Client SendGrid API requests and responses. Setting a tracer on the client
// can help us debug errors.
//
// The standard Logger implements this interface.
type Tracer interface {
	Printf(string, ...interface{})
}

var dumpRequest = func(t Tracer, request *http.Request) {
	if t == nil {
		return
	}

	if dump, err := httputil.DumpRequest(request, true); err == nil {
		t.Printf("\nRequest\n%s\n", dump)
	} else {
		t.Printf("\nRequest\n%s\n", request)
	}
}

var dumpResponse = func(t Tracer, response *http.Response) {
	if t == nil {
		return
	}

	if dump, err := httputil.DumpResponse(response, true); err == nil {
		t.Printf("\nReponse\n%s\n", dump)
	} else {
		t.Printf("\nResponse\n%s\n", response)
	}
}
