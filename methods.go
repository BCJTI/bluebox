package bluebox

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseUrl = "https://api.bluebox-systems-cargo.com"

// Params serves to map data for post in the request
type Params map[string]interface{}

// Headers serves to add extra request headers
type Headers map[string]string

// Make request and return the response
func (c *Client) execute(method string, path string, params interface{}, headers Headers, model interface{}) error {

	var request *http.Request

	// init vars
	endpoint := baseUrl + path

	// check for params
	if params != nil {

		// marshal params
		b, err := json.Marshal(params)
		if err != nil {
			return err
		}

		fmt.Println(string(b))

		// send as body
		if method != http.MethodGet {

			// set payload with params
			payload := strings.NewReader(string(b))

			// set request with payload
			request, _ = http.NewRequest(method, endpoint, payload)

		} else {

			var values Params

			// convert any type to params
			if err = json.Unmarshal(b, &values); err != nil {
				return err
			}

			// init request
			request, _ = http.NewRequest(method, endpoint, nil)

			// init query string
			query := request.URL.Query()

			// add params
			for key, value := range values {
				query.Add(key, AnyToString(value))
			}

			// set query string
			request.URL.RawQuery = query.Encode()

		}

	} else {

		// set request without payload
		request, _ = http.NewRequest(method, endpoint, nil)

	}

	// set basic auth
	// request.SetBasicAuth(c.username, c.password)

	// define json content type
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Basic YXBpLXdpbmRsb2c6V2luZGxvZzEyMzQ1Njc4")

	// add extra headers
	if headers != nil {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{},
	}}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read response
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// init error response
	msg := &ErrMessage{}

	// check for error message
	if err = json.Unmarshal(data, msg); err == nil && msg.ErrorMessage != "" {
		return msg
	}

	// verify status code
	if NotIn(response.StatusCode, http.StatusOK, http.StatusCreated, http.StatusAccepted) {

		// return body as error message
		if len(data) > 0 {
			return errors.New(string(data))
		}

		// return http status as error
		return errors.New(response.Status)

	}

	// parse data
	return json.Unmarshal(data, model)

}

// Get executes GET requests
func (c *Client) Get(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodGet, path, params, headers, model)
}

// Post executes POST requests
func (c *Client) Post(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodPost, path, params, headers, model)
}

// Put executes PUT requests
func (c *Client) Put(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodPut, path, params, headers, model)
}

// Patch executes PATCH requests
func (c *Client) Patch(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodPatch, path, params, headers, model)
}

// Delete executes DELETE requests
func (c *Client) Delete(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodDelete, path, params, headers, model)
}
