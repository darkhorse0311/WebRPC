// hello-webrpc v1.0.0 e7d03ae5c6db914d3d31a21b762b6c3af1af1a07
// --
// Code generated by webrpc-gen@v0.12.x-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=hello-api.ridl -target=golang -pkg=main -server -out=./server/hello_api.gen.go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
	return "e7d03ae5c6db914d3d31a21b762b6c3af1af1a07"
}

//
// Types
//

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
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
}

type ExampleService interface {
	Ping(ctx context.Context) (bool, error)
	GetUser(ctx context.Context, userID uint64) (*User, error)
}

var WebRPCServices = map[string][]string{
	"ExampleService": {
		"Ping",
		"GetUser",
	},
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type exampleServiceServer struct {
	ExampleService
}

func NewExampleServiceServer(svc ExampleService) WebRPCServer {
	return &exampleServiceServer{
		ExampleService: svc,
	}
}

func (s *exampleServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "ExampleService")

	if r.Method != "POST" {
		err := ErrorWithCause(ErrWebrpcBadMethod, fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/ExampleService/Ping":
		s.servePing(ctx, w, r)
		return
	case "/rpc/ExampleService/GetUser":
		s.serveGetUser(ctx, w, r)
		return
	default:
		err := ErrorWithCause(ErrWebrpcBadRoute, fmt.Errorf("no handler for path %q", r.URL.Path))
		RespondWithError(w, err)
		return
	}
}

func (s *exampleServiceServer) servePing(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

func (s *exampleServiceServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Ping")

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
		ret0, err = s.ExampleService.Ping(ctx)
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

func (s *exampleServiceServer) serveGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetUserJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *exampleServiceServer) serveGetUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetUser")
	reqContent := struct {
		Arg0 uint64 `json:"userID"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.ExampleService.GetUser(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *User `json:"user"`
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
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
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
