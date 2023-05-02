// Package shipping provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package shipping

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/icowan/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAccount request
	GetAccount(ctx context.Context) (*http.Response, error)

	// PurchaseShipment request  with any body
	PurchaseShipmentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	PurchaseShipment(ctx context.Context, body PurchaseShipmentJSONRequestBody) (*http.Response, error)

	// GetRates request  with any body
	GetRatesWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	GetRates(ctx context.Context, body GetRatesJSONRequestBody) (*http.Response, error)

	// CreateShipment request  with any body
	CreateShipmentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	CreateShipment(ctx context.Context, body CreateShipmentJSONRequestBody) (*http.Response, error)

	// GetShipment request
	GetShipment(ctx context.Context, shipmentId string) (*http.Response, error)

	// CancelShipment request
	CancelShipment(ctx context.Context, shipmentId string) (*http.Response, error)

	// RetrieveShippingLabel request  with any body
	RetrieveShippingLabelWithBody(ctx context.Context, shipmentId string, trackingId string, contentType string, body io.Reader) (*http.Response, error)

	RetrieveShippingLabel(ctx context.Context, shipmentId string, trackingId string, body RetrieveShippingLabelJSONRequestBody) (*http.Response, error)

	// PurchaseLabels request  with any body
	PurchaseLabelsWithBody(ctx context.Context, shipmentId string, contentType string, body io.Reader) (*http.Response, error)

	PurchaseLabels(ctx context.Context, shipmentId string, body PurchaseLabelsJSONRequestBody) (*http.Response, error)

	// GetTrackingInformation request
	GetTrackingInformation(ctx context.Context, trackingId string) (*http.Response, error)
}

func (c *Client) GetAccount(ctx context.Context) (*http.Response, error) {
	req, err := NewGetAccountRequest(c.Endpoint)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PurchaseShipmentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPurchaseShipmentRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PurchaseShipment(ctx context.Context, body PurchaseShipmentJSONRequestBody) (*http.Response, error) {
	req, err := NewPurchaseShipmentRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetRatesWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGetRatesRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetRates(ctx context.Context, body GetRatesJSONRequestBody) (*http.Response, error) {
	req, err := NewGetRatesRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CreateShipmentWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateShipmentRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateShipment(ctx context.Context, body CreateShipmentJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateShipmentRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetShipment(ctx context.Context, shipmentId string) (*http.Response, error) {
	req, err := NewGetShipmentRequest(c.Endpoint, shipmentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CancelShipment(ctx context.Context, shipmentId string) (*http.Response, error) {
	req, err := NewCancelShipmentRequest(c.Endpoint, shipmentId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) RetrieveShippingLabelWithBody(ctx context.Context, shipmentId string, trackingId string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewRetrieveShippingLabelRequestWithBody(c.Endpoint, shipmentId, trackingId, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) RetrieveShippingLabel(ctx context.Context, shipmentId string, trackingId string, body RetrieveShippingLabelJSONRequestBody) (*http.Response, error) {
	req, err := NewRetrieveShippingLabelRequest(c.Endpoint, shipmentId, trackingId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) PurchaseLabelsWithBody(ctx context.Context, shipmentId string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPurchaseLabelsRequestWithBody(c.Endpoint, shipmentId, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PurchaseLabels(ctx context.Context, shipmentId string, body PurchaseLabelsJSONRequestBody) (*http.Response, error) {
	req, err := NewPurchaseLabelsRequest(c.Endpoint, shipmentId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetTrackingInformation(ctx context.Context, trackingId string) (*http.Response, error) {
	req, err := NewGetTrackingInformationRequest(c.Endpoint, trackingId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetAccountRequest generates requests for GetAccount
func NewGetAccountRequest(endpoint string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/account")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPurchaseShipmentRequest calls the generic PurchaseShipment builder with application/json body
func NewPurchaseShipmentRequest(endpoint string, body PurchaseShipmentJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPurchaseShipmentRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewPurchaseShipmentRequestWithBody generates requests for PurchaseShipment with any type of body
func NewPurchaseShipmentRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/purchaseShipment")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetRatesRequest calls the generic GetRates builder with application/json body
func NewGetRatesRequest(endpoint string, body GetRatesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetRatesRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewGetRatesRequestWithBody generates requests for GetRates with any type of body
func NewGetRatesRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/rates")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewCreateShipmentRequest calls the generic CreateShipment builder with application/json body
func NewCreateShipmentRequest(endpoint string, body CreateShipmentJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateShipmentRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewCreateShipmentRequestWithBody generates requests for CreateShipment with any type of body
func NewCreateShipmentRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/shipments")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetShipmentRequest generates requests for GetShipment
func NewGetShipmentRequest(endpoint string, shipmentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/shipments/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCancelShipmentRequest generates requests for CancelShipment
func NewCancelShipmentRequest(endpoint string, shipmentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/shipments/%s/cancel", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRetrieveShippingLabelRequest calls the generic RetrieveShippingLabel builder with application/json body
func NewRetrieveShippingLabelRequest(endpoint string, shipmentId string, trackingId string, body RetrieveShippingLabelJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRetrieveShippingLabelRequestWithBody(endpoint, shipmentId, trackingId, "application/json", bodyReader)
}

// NewRetrieveShippingLabelRequestWithBody generates requests for RetrieveShippingLabel with any type of body
func NewRetrieveShippingLabelRequestWithBody(endpoint string, shipmentId string, trackingId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "trackingId", trackingId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/shipments/%s/containers/%s/label", pathParam0, pathParam1)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewPurchaseLabelsRequest calls the generic PurchaseLabels builder with application/json body
func NewPurchaseLabelsRequest(endpoint string, shipmentId string, body PurchaseLabelsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPurchaseLabelsRequestWithBody(endpoint, shipmentId, "application/json", bodyReader)
}

// NewPurchaseLabelsRequestWithBody generates requests for PurchaseLabels with any type of body
func NewPurchaseLabelsRequestWithBody(endpoint string, shipmentId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "shipmentId", shipmentId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/shipments/%s/purchaseLabels", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetTrackingInformationRequest generates requests for GetTrackingInformation
func NewGetTrackingInformationRequest(endpoint string, trackingId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "trackingId", trackingId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/shipping/v1/tracking/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAccount request
	GetAccountWithResponse(ctx context.Context) (*GetAccountResp, error)

	// PurchaseShipment request  with any body
	PurchaseShipmentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PurchaseShipmentResp, error)

	PurchaseShipmentWithResponse(ctx context.Context, body PurchaseShipmentJSONRequestBody) (*PurchaseShipmentResp, error)

	// GetRates request  with any body
	GetRatesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetRatesResp, error)

	GetRatesWithResponse(ctx context.Context, body GetRatesJSONRequestBody) (*GetRatesResp, error)

	// CreateShipment request  with any body
	CreateShipmentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateShipmentResp, error)

	CreateShipmentWithResponse(ctx context.Context, body CreateShipmentJSONRequestBody) (*CreateShipmentResp, error)

	// GetShipment request
	GetShipmentWithResponse(ctx context.Context, shipmentId string) (*GetShipmentResp, error)

	// CancelShipment request
	CancelShipmentWithResponse(ctx context.Context, shipmentId string) (*CancelShipmentResp, error)

	// RetrieveShippingLabel request  with any body
	RetrieveShippingLabelWithBodyWithResponse(ctx context.Context, shipmentId string, trackingId string, contentType string, body io.Reader) (*RetrieveShippingLabelResp, error)

	RetrieveShippingLabelWithResponse(ctx context.Context, shipmentId string, trackingId string, body RetrieveShippingLabelJSONRequestBody) (*RetrieveShippingLabelResp, error)

	// PurchaseLabels request  with any body
	PurchaseLabelsWithBodyWithResponse(ctx context.Context, shipmentId string, contentType string, body io.Reader) (*PurchaseLabelsResp, error)

	PurchaseLabelsWithResponse(ctx context.Context, shipmentId string, body PurchaseLabelsJSONRequestBody) (*PurchaseLabelsResp, error)

	// GetTrackingInformation request
	GetTrackingInformationWithResponse(ctx context.Context, trackingId string) (*GetTrackingInformationResp, error)
}

type GetAccountResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetAccountResponse
}

// Status returns HTTPResponse.Status
func (r GetAccountResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAccountResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PurchaseShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *PurchaseShipmentResponse
}

// Status returns HTTPResponse.Status
func (r PurchaseShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PurchaseShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRatesResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetRatesResponse
}

// Status returns HTTPResponse.Status
func (r GetRatesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRatesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CreateShipmentResponse
}

// Status returns HTTPResponse.Status
func (r CreateShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetShipmentResponse
}

// Status returns HTTPResponse.Status
func (r GetShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelShipmentResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *CancelShipmentResponse
}

// Status returns HTTPResponse.Status
func (r CancelShipmentResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelShipmentResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveShippingLabelResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *RetrieveShippingLabelResponse
}

// Status returns HTTPResponse.Status
func (r RetrieveShippingLabelResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveShippingLabelResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PurchaseLabelsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *PurchaseLabelsResponse
}

// Status returns HTTPResponse.Status
func (r PurchaseLabelsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PurchaseLabelsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTrackingInformationResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetTrackingInformationResponse
}

// Status returns HTTPResponse.Status
func (r GetTrackingInformationResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTrackingInformationResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAccountWithResponse request returning *GetAccountResponse
func (c *ClientWithResponses) GetAccountWithResponse(ctx context.Context) (*GetAccountResp, error) {
	rsp, err := c.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetAccountResp(rsp)
}

// PurchaseShipmentWithBodyWithResponse request with arbitrary body returning *PurchaseShipmentResponse
func (c *ClientWithResponses) PurchaseShipmentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PurchaseShipmentResp, error) {
	rsp, err := c.PurchaseShipmentWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePurchaseShipmentResp(rsp)
}

func (c *ClientWithResponses) PurchaseShipmentWithResponse(ctx context.Context, body PurchaseShipmentJSONRequestBody) (*PurchaseShipmentResp, error) {
	rsp, err := c.PurchaseShipment(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParsePurchaseShipmentResp(rsp)
}

// GetRatesWithBodyWithResponse request with arbitrary body returning *GetRatesResponse
func (c *ClientWithResponses) GetRatesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetRatesResp, error) {
	rsp, err := c.GetRatesWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGetRatesResp(rsp)
}

func (c *ClientWithResponses) GetRatesWithResponse(ctx context.Context, body GetRatesJSONRequestBody) (*GetRatesResp, error) {
	rsp, err := c.GetRates(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseGetRatesResp(rsp)
}

// CreateShipmentWithBodyWithResponse request with arbitrary body returning *CreateShipmentResponse
func (c *ClientWithResponses) CreateShipmentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*CreateShipmentResp, error) {
	rsp, err := c.CreateShipmentWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateShipmentResp(rsp)
}

func (c *ClientWithResponses) CreateShipmentWithResponse(ctx context.Context, body CreateShipmentJSONRequestBody) (*CreateShipmentResp, error) {
	rsp, err := c.CreateShipment(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateShipmentResp(rsp)
}

// GetShipmentWithResponse request returning *GetShipmentResponse
func (c *ClientWithResponses) GetShipmentWithResponse(ctx context.Context, shipmentId string) (*GetShipmentResp, error) {
	rsp, err := c.GetShipment(ctx, shipmentId)
	if err != nil {
		return nil, err
	}
	return ParseGetShipmentResp(rsp)
}

// CancelShipmentWithResponse request returning *CancelShipmentResponse
func (c *ClientWithResponses) CancelShipmentWithResponse(ctx context.Context, shipmentId string) (*CancelShipmentResp, error) {
	rsp, err := c.CancelShipment(ctx, shipmentId)
	if err != nil {
		return nil, err
	}
	return ParseCancelShipmentResp(rsp)
}

// RetrieveShippingLabelWithBodyWithResponse request with arbitrary body returning *RetrieveShippingLabelResponse
func (c *ClientWithResponses) RetrieveShippingLabelWithBodyWithResponse(ctx context.Context, shipmentId string, trackingId string, contentType string, body io.Reader) (*RetrieveShippingLabelResp, error) {
	rsp, err := c.RetrieveShippingLabelWithBody(ctx, shipmentId, trackingId, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveShippingLabelResp(rsp)
}

func (c *ClientWithResponses) RetrieveShippingLabelWithResponse(ctx context.Context, shipmentId string, trackingId string, body RetrieveShippingLabelJSONRequestBody) (*RetrieveShippingLabelResp, error) {
	rsp, err := c.RetrieveShippingLabel(ctx, shipmentId, trackingId, body)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveShippingLabelResp(rsp)
}

// PurchaseLabelsWithBodyWithResponse request with arbitrary body returning *PurchaseLabelsResponse
func (c *ClientWithResponses) PurchaseLabelsWithBodyWithResponse(ctx context.Context, shipmentId string, contentType string, body io.Reader) (*PurchaseLabelsResp, error) {
	rsp, err := c.PurchaseLabelsWithBody(ctx, shipmentId, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePurchaseLabelsResp(rsp)
}

func (c *ClientWithResponses) PurchaseLabelsWithResponse(ctx context.Context, shipmentId string, body PurchaseLabelsJSONRequestBody) (*PurchaseLabelsResp, error) {
	rsp, err := c.PurchaseLabels(ctx, shipmentId, body)
	if err != nil {
		return nil, err
	}
	return ParsePurchaseLabelsResp(rsp)
}

// GetTrackingInformationWithResponse request returning *GetTrackingInformationResponse
func (c *ClientWithResponses) GetTrackingInformationWithResponse(ctx context.Context, trackingId string) (*GetTrackingInformationResp, error) {
	rsp, err := c.GetTrackingInformation(ctx, trackingId)
	if err != nil {
		return nil, err
	}
	return ParseGetTrackingInformationResp(rsp)
}

// ParseGetAccountResp parses an HTTP response from a GetAccountWithResponse call
func ParseGetAccountResp(rsp *http.Response) (*GetAccountResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetAccountResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetAccountResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParsePurchaseShipmentResp parses an HTTP response from a PurchaseShipmentWithResponse call
func ParsePurchaseShipmentResp(rsp *http.Response) (*PurchaseShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PurchaseShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest PurchaseShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetRatesResp parses an HTTP response from a GetRatesWithResponse call
func ParseGetRatesResp(rsp *http.Response) (*GetRatesResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetRatesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetRatesResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCreateShipmentResp parses an HTTP response from a CreateShipmentWithResponse call
func ParseCreateShipmentResp(rsp *http.Response) (*CreateShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CreateShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetShipmentResp parses an HTTP response from a GetShipmentWithResponse call
func ParseGetShipmentResp(rsp *http.Response) (*GetShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseCancelShipmentResp parses an HTTP response from a CancelShipmentWithResponse call
func ParseCancelShipmentResp(rsp *http.Response) (*CancelShipmentResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CancelShipmentResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest CancelShipmentResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseRetrieveShippingLabelResp parses an HTTP response from a RetrieveShippingLabelWithResponse call
func ParseRetrieveShippingLabelResp(rsp *http.Response) (*RetrieveShippingLabelResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &RetrieveShippingLabelResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest RetrieveShippingLabelResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParsePurchaseLabelsResp parses an HTTP response from a PurchaseLabelsWithResponse call
func ParsePurchaseLabelsResp(rsp *http.Response) (*PurchaseLabelsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PurchaseLabelsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest PurchaseLabelsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetTrackingInformationResp parses an HTTP response from a GetTrackingInformationWithResponse call
func ParseGetTrackingInformationResp(rsp *http.Response) (*GetTrackingInformationResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetTrackingInformationResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetTrackingInformationResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
