// example-api-service v1.0.0 f1b366018b1650b6a0a4f09ff793089ec62830a6
// --
// Code generated by webrpc-gen@v0.11.x-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=./proto/api.ridl -target=golang -pkg=main -server -client -out=./api.gen.go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v1.0.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "f1b366018b1650b6a0a4f09ff793089ec62830a6"
}

//
// Types
//

type User struct {
	Username string `json:"username"`
	Age      uint32 `json:"age"`
}

type Location uint32

const (
	Location_TORONTO  Location = 0
	Location_NEW_YORK Location = 1
)

var Location_name = map[uint32]string{
	0: "TORONTO",
	1: "NEW_YORK",
}

var Location_value = map[string]uint32{
	"TORONTO":  0,
	"NEW_YORK": 1,
}

func (x Location) String() string {
	return Location_name[uint32(x)]
}

func (x Location) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(Location_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *Location) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = Location(Location_value[j])
	return nil
}

type ExampleAPI interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	GetUsers(ctx context.Context) ([]*User, *Location, error)
}

var WebRPCServices = map[string][]string{
	"ExampleAPI": {
		"Ping",
		"Status",
		"GetUsers",
	},
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type exampleAPIServer struct {
	ExampleAPI
}

func NewExampleAPIServer(svc ExampleAPI) WebRPCServer {
	return &exampleAPIServer{
		ExampleAPI: svc,
	}
}

func (s *exampleAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "ExampleAPI")

	if r.Method != "POST" {
		err := ErrorWithCause(ErrWebrpcBadMethod, fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/ExampleAPI/Ping":
		s.servePing(ctx, w, r)
		return
	case "/rpc/ExampleAPI/Status":
		s.serveStatus(ctx, w, r)
		return
	case "/rpc/ExampleAPI/GetUsers":
		s.serveGetUsers(ctx, w, r)
		return
	default:
		err := ErrorWithCause(ErrWebrpcBadRoute, fmt.Errorf("no handler for path %q", r.URL.Path))
		RespondWithError(w, err)
		return
	}
}

func (s *exampleAPIServer) servePing(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePingJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleAPIServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Ping")

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		err = s.ExampleAPI.Ping(ctx)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *exampleAPIServer) serveStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveStatusJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleAPIServer) serveStatusJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Status")

	// Call service method
	var ret0 bool
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.ExampleAPI.Status(ctx)
	}()
	respContent := struct {
		Ret0 bool `json:"status"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleAPIServer) serveGetUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetUsersJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleAPIServer) serveGetUsersJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetUsers")

	// Call service method
	var ret0 []*User
	var ret1 *Location
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, ret1, err = s.ExampleAPI.GetUsers(ctx)
	}()
	respContent := struct {
		Ret0 []*User   `json:"users"`
		Ret1 *Location `json:"location"`
	}{ret0, ret1}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrorWithCause(ErrWebrpcEndpoint, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

//
// Client
//

const ExampleAPIPathPrefix = "/rpc/ExampleAPI/"

type exampleAPIClient struct {
	client HTTPClient
	urls   [3]string
}

func NewExampleAPIClient(addr string, client HTTPClient) ExampleAPI {
	prefix := urlBase(addr) + ExampleAPIPathPrefix
	urls := [3]string{
		prefix + "Ping",
		prefix + "Status",
		prefix + "GetUsers",
	}
	return &exampleAPIClient{
		client: client,
		urls:   urls,
	}
}

func (c *exampleAPIClient) Ping(ctx context.Context) error {

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, nil)
	return err
}

func (c *exampleAPIClient) Status(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *exampleAPIClient) GetUsers(ctx context.Context) ([]*User, *Location, error) {
	out := struct {
		Ret0 []*User   `json:"users"`
		Ret1 *Location `json:"location"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, out.Ret1, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doJSONRequest is common code to make a request to the remote service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("could not build request: %w", err))
	}
	resp, err := client.Do(req)
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("failed to close response body: %w", cerr))
		}
	}()

	if err = ctx.Err(); err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("aborted because context was done: %w", err))
	}

	if resp.StatusCode != 200 {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return rpcErr
	}

	if out != nil {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	return fmt.Sprintf("Error %d %s: %s", e.Code, e.Name, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	err := rpcErr
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Webrpc errors
var (
	ErrWebrpcEndpoint      = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 0}
	ErrWebrpcBadRoute      = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod     = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest    = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse   = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic   = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
)
