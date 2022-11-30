package shipdis

import (
	"bytes"
	"io"
	"net/http"
)

func (rp *RequestPayload) MakeHttpRequest() (resp []byte, err error) {
	client := &http.Client{}

	// Create request - Add payload
	request, err := http.NewRequest(rp.Method, rp.URL, bytes.NewBuffer(rp.Payload))
	if rp.Payload == nil {
		request, err = http.NewRequest(rp.Method, rp.URL, nil)
	}

	if err != nil {
		return
	}

	// Headers
	request.Header = rp.Headers

	// Query Params
	request.URL.RawQuery = rp.QueryParams.Encode()

	//Request
	response, err := client.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()
	resp, err = io.ReadAll(response.Body)
	return
}
