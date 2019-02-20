// sample v0.0.1
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/golang
// Do not edit by hand. Update your webrpc schema and re-generate.
package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/webrpc/webrpc/lib/webrpc-go"
)

type Kind uint32

const (
	Kind_USER  Kind = 1
	Kind_ADMIN Kind = 2
)

var Kind_name = map[uint32]string{
	1: "USER",
	2: "ADMIN",
}

var Kind_value = map[string]uint32{
	"USER":  1,
	"ADMIN": 2,
}

func (x Kind) String() string {
	return Kind_name[uint32(x)]
}

func (x Kind) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(Kind_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *Kind) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = Kind(Kind_value[j])
	return nil
}

type Empty struct {
}

type User struct {
	ID        uint64     `json:"id" db:"id"`
	Username  string     `json:"USERNAME" db:"username"`
	Role      string     `json:"role" db:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
}

type ComplexType struct {
	Meta              map[string]interface{}       `json:"meta"`
	MetaNestedExample map[string]map[string]uint32 `json:"metaNestedExample"`
	NamesList         []string                     `json:"namesList"`
	NumsList          []int64                      `json:"numsList"`
	DoubleArray       [][]string                   `json:"doubleArray"`
	ListOfMaps        []map[string]uint32          `json:"listOfMaps"`
	ListOfUsers       []*User                      `json:"listOfUsers"`
	MapOfUsers        map[string]*User             `json:"mapOfUsers"`
	User              *User                        `json:"user"`
}

type ExampleRPC interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	GetUser(ctx context.Context, header map[string]string, userID uint64) (uint32, *User, error)
}

type AnotherRPC interface {
	Owner(ctx context.Context, q *string, id *uint64, desc string) (uint32, *User, []string, error)
}

var Services = map[string][]string{
	"ExampleRPC": {
		"Ping",
		"Status",
		"GetUser",
	},
	"AnotherRPC": {
		"Owner",
	},
}

// Client

const ExampleRPCPathPrefix = "/rpc/ExampleRPC/"

const AnotherRPCPathPrefix = "/rpc/AnotherRPC/"

type exampleRPCClient struct {
	client HTTPClient
	urls   [3]string
}

func NewExampleRPCClient(addr string, client HTTPClient) ExampleRPC {
	prefix := urlBase(addr) + ExampleRPCPathPrefix
	urls := [3]string{
		prefix + "Ping",
		prefix + "Status",
		prefix + "GetUser",
	}
	return &exampleRPCClient{
		client: client,
		urls:   urls,
	}
}

func (c *exampleRPCClient) Ping(ctx context.Context) error {

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, nil)
	return err
}

func (c *exampleRPCClient) Status(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *exampleRPCClient) GetUser(ctx context.Context, header map[string]string, userID uint64) (uint32, *User, error) {
	in := struct {
		Arg0 map[string]string `json:"header"`
		Arg1 uint64            `json:"userID"`
	}{header, userID}
	out := struct {
		Ret0 uint32 `json:"code"`
		Ret1 *User  `json:"user"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], in, &out)
	return out.Ret0, out.Ret1, err
}

type anotherRPCClient struct {
	client HTTPClient
	urls   [1]string
}

func NewAnotherRPCClient(addr string, client HTTPClient) AnotherRPC {
	prefix := urlBase(addr) + AnotherRPCPathPrefix
	urls := [1]string{
		prefix + "Owner",
	}
	return &anotherRPCClient{
		client: client,
		urls:   urls,
	}
}

func (c *anotherRPCClient) Owner(ctx context.Context, q *string, id *uint64, desc string) (uint32, *User, []string, error) {
	in := struct {
		Arg0 *string `json:"q"`
		Arg1 *uint64 `json:"id"`
		Arg2 string  `json:"desc"`
	}{q, id, desc}
	out := struct {
		Ret0 uint32   `json:"code"`
		Ret1 *User    `json:"user"`
		Ret2 []string `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], in, &out)
	return out.Ret0, out.Ret1, out.Ret2, err
}

// Server

type exampleRPCServer struct {
	ExampleRPC
}

func NewExampleRPCServer(svc ExampleRPC) WebRPCServer {
	return &exampleRPCServer{
		ExampleRPC: svc,
	}
}

func (s *exampleRPCServer) WebRPCVersion() string {
	return "v0.0.1"
}

func (s *exampleRPCServer) ServiceVersion() string {
	return "v0.1.0"
}

func (s *exampleRPCServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = webrpc.WithResponseWriter(ctx, w)
	ctx = webrpc.WithServiceName(ctx, "ExampleRPC")

	if r.Method != "POST" {
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unsupported method %q (only POST is allowed)", r.Method)
		writeJSONError(ctx, w, r, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/ExampleRPC/Ping":
		s.servePing(ctx, w, r)
		return
	case "/rpc/ExampleRPC/Status":
		s.serveStatus(ctx, w, r)
		return
	case "/rpc/ExampleRPC/GetUser":
		s.serveGetUser(ctx, w, r)
		return
	default:
		err := webrpc.Errorf(webrpc.ErrBadRoute, "no handler for path %q", r.URL.Path)
		writeJSONError(ctx, w, r, err)
		return
	}
}

func (s *exampleRPCServer) servePing(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePingJSON(ctx, w, r)
	default:
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		writeJSONError(ctx, w, r, err)
	}
}

func (s *exampleRPCServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "Ping")

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				writeJSONError(ctx, w, r, webrpc.ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		err = s.ExampleRPC.Ping(ctx)
	}()

	if err != nil {
		writeJSONError(ctx, w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *exampleRPCServer) serveStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveStatusJSON(ctx, w, r)
	default:
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		writeJSONError(ctx, w, r, err)
	}
}

func (s *exampleRPCServer) serveStatusJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "Status")

	// Call service method
	var ret0 bool
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				writeJSONError(ctx, w, r, webrpc.ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.ExampleRPC.Status(ctx)
	}()
	respContent := struct {
		Ret0 bool `json:"status"`
	}{ret0}

	if err != nil {
		writeJSONError(ctx, w, r, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to marshal json response")
		writeJSONError(ctx, w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleRPCServer) serveGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetUserJSON(ctx, w, r)
	default:
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		writeJSONError(ctx, w, r, err)
	}
}

func (s *exampleRPCServer) serveGetUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "GetUser")
	reqContent := struct {
		Arg0 map[string]string `json:"header"`
		Arg1 uint64            `json:"userID"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to read request data")
		writeJSONError(ctx, w, r, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInvalidArgument, err, "failed to unmarshal request data")
		writeJSONError(ctx, w, r, err)
		return
	}

	// Call service method
	var ret0 uint32
	var ret1 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				writeJSONError(ctx, w, r, webrpc.ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, ret1, err = s.ExampleRPC.GetUser(ctx, reqContent.Arg0, reqContent.Arg1)
	}()
	respContent := struct {
		Ret0 uint32 `json:"code"`
		Ret1 *User  `json:"user"`
	}{ret0, ret1}

	if err != nil {
		writeJSONError(ctx, w, r, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to marshal json response")
		writeJSONError(ctx, w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

type anotherRPCServer struct {
	AnotherRPC
}

func NewAnotherRPCServer(svc AnotherRPC) WebRPCServer {
	return &anotherRPCServer{
		AnotherRPC: svc,
	}
}

func (s *anotherRPCServer) WebRPCVersion() string {
	return "v0.0.1"
}

func (s *anotherRPCServer) ServiceVersion() string {
	return "v0.1.0"
}

func (s *anotherRPCServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = webrpc.WithResponseWriter(ctx, w)
	ctx = webrpc.WithServiceName(ctx, "AnotherRPC")

	if r.Method != "POST" {
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unsupported method %q (only POST is allowed)", r.Method)
		writeJSONError(ctx, w, r, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/AnotherRPC/Owner":
		s.serveOwner(ctx, w, r)
		return
	default:
		err := webrpc.Errorf(webrpc.ErrBadRoute, "no handler for path %q", r.URL.Path)
		writeJSONError(ctx, w, r, err)
		return
	}
}

func (s *anotherRPCServer) serveOwner(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveOwnerJSON(ctx, w, r)
	default:
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		writeJSONError(ctx, w, r, err)
	}
}

func (s *anotherRPCServer) serveOwnerJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "Owner")
	reqContent := struct {
		Arg0 *string `json:"q"`
		Arg1 *uint64 `json:"id"`
		Arg2 string  `json:"desc"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to read request data")
		writeJSONError(ctx, w, r, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInvalidArgument, err, "failed to unmarshal request data")
		writeJSONError(ctx, w, r, err)
		return
	}

	// Call service method
	var ret0 uint32
	var ret1 *User
	var ret2 []string
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				writeJSONError(ctx, w, r, webrpc.ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, ret1, ret2, err = s.AnotherRPC.Owner(ctx, reqContent.Arg0, reqContent.Arg1, reqContent.Arg2)
	}()
	respContent := struct {
		Ret0 uint32   `json:"code"`
		Ret1 *User    `json:"user"`
		Ret2 []string `json:"status"`
	}{ret0, ret1, ret2}

	if err != nil {
		writeJSONError(ctx, w, r, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to marshal json response")
		writeJSONError(ctx, w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

//
// Helpers
//

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WebRPCServer interface {
	http.Handler
	WebRPCVersion() string
	ServiceVersion() string
}

type errResponse struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Cause  string `json:"cause,omitempty"`
}

func writeJSONError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	rpcErr, ok := err.(webrpc.Error)
	if !ok {
		rpcErr = webrpc.WrapError(webrpc.ErrInternal, err, "webrpc error")
	}

	statusCode := webrpc.HTTPStatusFromErrorCode(rpcErr.Code())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResp := errResponse{
		Status: statusCode,
		Code:   string(rpcErr.Code()),
		Msg:    rpcErr.Error(),
	}
	respBody, _ := json.Marshal(errResp)
	w.Write(respBody)
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
	return req, nil
}

// doJSONRequest is common code to make a request to the remote service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("request failed", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	if out != nil {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return clientError("failed to read response body", err)
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return clientError("failed to unmarshal json response body", err)
		}
		if err = ctx.Err(); err != nil {
			return clientError("aborted because context was done", err)
		}
	}

	return nil
}

// errorFromResponse builds a webrpc.Error from a non-200 HTTP response.
func errorFromResponse(resp *http.Response) webrpc.Error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}

	var respErr errResponse
	if err := json.Unmarshal(respBody, &respErr); err != nil {
		return clientError("failed unmarshal error response", err)
	}

	errCode := webrpc.ErrorCode(respErr.Code)

	if webrpc.HTTPStatusFromErrorCode(errCode) == 0 {
		return webrpc.ErrorInternal("invalid code returned from server error response: %s", respErr.Code)
	}

	return webrpc.Errorf(errCode, respErr.Msg)
}

func clientError(desc string, err error) webrpc.Error {
	return webrpc.WrapError(webrpc.ErrInternal, err, desc)
}
