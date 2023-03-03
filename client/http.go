package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"context"

	"github.com/Fred78290/vmrest-go-client/client/api"
	ctx "github.com/Fred78290/vmrest-go-client/client/context"
	"github.com/Fred78290/vmrest-go-client/client/model"
)

const VMREST_CONTENT_TYPE = "application/vnd.vmware.vmw.rest-v1+json"

// Logger is the interface that should be implemented for loggers that wish to
// log HTTP requests and HTTP responses.
type Logger interface {
	// LogRequest logs an HTTP request.
	LogRequest(*http.Request)

	// LogResponse logs an HTTP response.
	LogResponse(*http.Response)
}

type clientWrapper struct {
	// Client is the underlying HTTP client used to run the requests. It may be overloaded but a default one is instanciated in ``NewClient`` by default.
	APIEndPoint  string
	APIUserAgent string
	APIPassword  string
	APIUserName  string
	Client       *http.Client
	Timeout      time.Duration
	// Logger is used to log HTTP requests and responses.
	Logger Logger
}

/*
NewHttpClient Changes the VM power state
  - @param endpoint the endpoint to joint vmrest
  - @param userAgent
  - @param userName
  - @param password
  - @param timeout operation timeout in seconds
  - @param unsecureTLS unsecure tls

@return api.Client
*/
func NewHttpClient(endpoint, userAgent, userName, password string, timeout time.Duration, unsecureTLS bool) (api.Client, error) {

	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1:8697"
	}

	client := clientWrapper{
		APIEndPoint:  endpoint,
		APIUserAgent: userAgent,
		APIPassword:  password,
		APIUserName:  userName,
		Timeout:      timeout,
		Client: &http.Client{
			Timeout: timeout * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: unsecureTLS,
				},
			},
		},
	}

	// Get and check the configuration
	if err := client.validate(); err != nil {
		return nil, err
	}

	return &client, nil
}

//
// Common request wrappers
//

// Get is a wrapper for the GET method
func (c *clientWrapper) Get(url string, resType interface{}) error {
	return c.CallAPI("GET", url, nil, resType, true)
}

// Patch is a wrapper for the POST method
func (c *clientWrapper) Patch(url string, reqBody, resType interface{}) error {
	return c.CallAPI("PATCH", url, reqBody, resType, true)
}

// Post is a wrapper for the POST method
func (c *clientWrapper) Post(url string, reqBody, resType interface{}) error {
	return c.CallAPI("POST", url, reqBody, resType, true)
}

// Put is a wrapper for the PUT method
func (c *clientWrapper) Put(url string, reqBody, resType interface{}) error {
	return c.CallAPI("PUT", url, reqBody, resType, true)
}

// Delete is a wrapper for the DELETE method
func (c *clientWrapper) Delete(url string, resType interface{}) error {
	return c.CallAPI("DELETE", url, nil, resType, true)
}

// GetWithContext is a wrapper for the GET method
func (c *clientWrapper) GetWithContext(ctx context.Context, url string, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "GET", url, nil, resType, true)
}

// PatchWithContext is a wrapper for the POST method
func (c *clientWrapper) PatchWithContext(ctx context.Context, url string, reqBody, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "PATCH", url, reqBody, resType, true)
}

// PostWithContext is a wrapper for the POST method
func (c *clientWrapper) PostWithContext(ctx context.Context, url string, reqBody, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "POST", url, reqBody, resType, true)
}

// PutWithContext is a wrapper for the PUT method
func (c *clientWrapper) PutWithContext(ctx context.Context, url string, reqBody, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "PUT", url, reqBody, resType, true)
}

// DeleteWithContext is a wrapper for the DELETE method
func (c *clientWrapper) DeleteWithContext(ctx context.Context, url string, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "DELETE", url, nil, resType, true)
}

// NewRequest returns a new HTTP request
func (c *clientWrapper) NewRequest(method, path string, reqBody interface{}, needAuth bool) (*http.Request, error) {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
	}

	target := fmt.Sprintf("%s%s", c.APIEndPoint, path)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Inject headers
	if body != nil {
		req.Header.Set("Content-Type", VMREST_CONTENT_TYPE)
	}

	if len(c.APIUserName) > 0 && len(c.APIPassword) > 0 {
		req.SetBasicAuth(c.APIUserName, c.APIPassword)
	}

	req.Header.Set("Accept", VMREST_CONTENT_TYPE)
	req.Header.Set("User-Agent", c.APIUserAgent)

	// Send the request with requested timeout
	c.Client.Timeout = c.Timeout

	return req, nil
}

// Do sends an HTTP request and returns an HTTP response
func (c *clientWrapper) Do(req *http.Request) (*http.Response, error) {
	if c.Logger != nil {
		c.Logger.LogRequest(req)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.Logger != nil {
		c.Logger.LogResponse(resp)
	}
	return resp, nil
}

// CallAPI is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everything went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *clientWrapper) CallAPI(method, path string, reqBody, resType interface{}, needAuth bool) error {
	return c.CallAPIWithContext(ctx.NewContext(c.Timeout), method, path, reqBody, resType, needAuth)
}

// CallAPIWithContext is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// # Context is used by http.Client to handle context cancelation
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everything went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *clientWrapper) CallAPIWithContext(ctx context.Context, method, path string, reqBody, resType interface{}, needAuth bool) error {
	req, err := c.NewRequest(method, path, reqBody, needAuth)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	response, err := c.Do(req)
	if err != nil {
		return err
	}
	return c.UnmarshalResponse(response, resType)
}

// UnmarshalResponse checks the response and unmarshals it into the response
// type if needed Helper function, called from CallAPI
func (c *clientWrapper) UnmarshalResponse(response *http.Response, resType interface{}) error {
	// Read all the response body
	defer response.Body.Close()

	if body, err := io.ReadAll(response.Body); err != nil {
		return err
	} else if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		if resType != nil {
			return c.decode(resType, body, response.Header.Get("Content-Type"))
		}
	} else {
		var v model.ErrorModel

		newErr := GenericSwaggerError{
			body:  body,
			error: response.Status,
		}

		if err = c.decode(&v, body, response.Header.Get("Content-Type")); err != nil {
			newErr.error = err.Error()
		} else {
			newErr.model = v
		}

		return newErr
	}

	return nil
}

func (c *clientWrapper) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

func (c *clientWrapper) validate() error {
	var response interface{}

	if err := c.Get("/", response); err != nil {
		return err
	}

	return nil
}
