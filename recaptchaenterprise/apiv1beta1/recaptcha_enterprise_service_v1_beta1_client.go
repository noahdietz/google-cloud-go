// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recaptchaenterprise

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newRecaptchaEnterpriseServiceV1Beta1ClientHook clientHook

// RecaptchaEnterpriseServiceV1Beta1CallOptions contains the retry settings for each method of RecaptchaEnterpriseServiceV1Beta1Client.
type RecaptchaEnterpriseServiceV1Beta1CallOptions struct {
	CreateAssessment   []gax.CallOption
	AnnotateAssessment []gax.CallOption
}

func defaultRecaptchaEnterpriseServiceV1Beta1GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recaptchaenterprise.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("recaptchaenterprise.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://recaptchaenterprise.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultRecaptchaEnterpriseServiceV1Beta1CallOptions() *RecaptchaEnterpriseServiceV1Beta1CallOptions {
	return &RecaptchaEnterpriseServiceV1Beta1CallOptions{
		CreateAssessment:   []gax.CallOption{},
		AnnotateAssessment: []gax.CallOption{},
	}
}

func defaultRecaptchaEnterpriseServiceV1Beta1RESTCallOptions() *RecaptchaEnterpriseServiceV1Beta1CallOptions {
	return &RecaptchaEnterpriseServiceV1Beta1CallOptions{
		CreateAssessment:   []gax.CallOption{},
		AnnotateAssessment: []gax.CallOption{},
	}
}

// internalRecaptchaEnterpriseServiceV1Beta1Client is an interface that defines the methods available from reCAPTCHA Enterprise API.
type internalRecaptchaEnterpriseServiceV1Beta1Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateAssessment(context.Context, *recaptchaenterprisepb.CreateAssessmentRequest, ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error)
	AnnotateAssessment(context.Context, *recaptchaenterprisepb.AnnotateAssessmentRequest, ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error)
}

// RecaptchaEnterpriseServiceV1Beta1Client is a client for interacting with reCAPTCHA Enterprise API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to determine the likelihood an event is legitimate.
type RecaptchaEnterpriseServiceV1Beta1Client struct {
	// The internal transport-dependent client.
	internalClient internalRecaptchaEnterpriseServiceV1Beta1Client

	// The call options for this service.
	CallOptions *RecaptchaEnterpriseServiceV1Beta1CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RecaptchaEnterpriseServiceV1Beta1Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RecaptchaEnterpriseServiceV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *RecaptchaEnterpriseServiceV1Beta1Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateAssessment creates an Assessment of the likelihood an event is legitimate.
func (c *RecaptchaEnterpriseServiceV1Beta1Client) CreateAssessment(ctx context.Context, req *recaptchaenterprisepb.CreateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error) {
	return c.internalClient.CreateAssessment(ctx, req, opts...)
}

// AnnotateAssessment annotates a previously created Assessment to provide additional information
// on whether the event turned out to be authentic or fradulent.
func (c *RecaptchaEnterpriseServiceV1Beta1Client) AnnotateAssessment(ctx context.Context, req *recaptchaenterprisepb.AnnotateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error) {
	return c.internalClient.AnnotateAssessment(ctx, req, opts...)
}

// recaptchaEnterpriseServiceV1Beta1GRPCClient is a client for interacting with reCAPTCHA Enterprise API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type recaptchaEnterpriseServiceV1Beta1GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing RecaptchaEnterpriseServiceV1Beta1Client
	CallOptions **RecaptchaEnterpriseServiceV1Beta1CallOptions

	// The gRPC API client.
	recaptchaEnterpriseServiceV1Beta1Client recaptchaenterprisepb.RecaptchaEnterpriseServiceV1Beta1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewRecaptchaEnterpriseServiceV1Beta1Client creates a new recaptcha enterprise service v1 beta1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to determine the likelihood an event is legitimate.
func NewRecaptchaEnterpriseServiceV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*RecaptchaEnterpriseServiceV1Beta1Client, error) {
	clientOpts := defaultRecaptchaEnterpriseServiceV1Beta1GRPCClientOptions()
	if newRecaptchaEnterpriseServiceV1Beta1ClientHook != nil {
		hookOpts, err := newRecaptchaEnterpriseServiceV1Beta1ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RecaptchaEnterpriseServiceV1Beta1Client{CallOptions: defaultRecaptchaEnterpriseServiceV1Beta1CallOptions()}

	c := &recaptchaEnterpriseServiceV1Beta1GRPCClient{
		connPool:                                connPool,
		disableDeadlines:                        disableDeadlines,
		recaptchaEnterpriseServiceV1Beta1Client: recaptchaenterprisepb.NewRecaptchaEnterpriseServiceV1Beta1Client(connPool),
		CallOptions:                             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *recaptchaEnterpriseServiceV1Beta1GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *recaptchaEnterpriseServiceV1Beta1GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *recaptchaEnterpriseServiceV1Beta1GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type recaptchaEnterpriseServiceV1Beta1RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing RecaptchaEnterpriseServiceV1Beta1Client
	CallOptions **RecaptchaEnterpriseServiceV1Beta1CallOptions
}

// NewRecaptchaEnterpriseServiceV1Beta1RESTClient creates a new recaptcha enterprise service v1 beta1 rest client.
//
// Service to determine the likelihood an event is legitimate.
func NewRecaptchaEnterpriseServiceV1Beta1RESTClient(ctx context.Context, opts ...option.ClientOption) (*RecaptchaEnterpriseServiceV1Beta1Client, error) {
	clientOpts := append(defaultRecaptchaEnterpriseServiceV1Beta1RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRecaptchaEnterpriseServiceV1Beta1RESTCallOptions()
	c := &recaptchaEnterpriseServiceV1Beta1RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &RecaptchaEnterpriseServiceV1Beta1Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRecaptchaEnterpriseServiceV1Beta1RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://recaptchaenterprise.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://recaptchaenterprise.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://recaptchaenterprise.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *recaptchaEnterpriseServiceV1Beta1RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *recaptchaEnterpriseServiceV1Beta1RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *recaptchaEnterpriseServiceV1Beta1RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *recaptchaEnterpriseServiceV1Beta1GRPCClient) CreateAssessment(ctx context.Context, req *recaptchaenterprisepb.CreateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAssessment[0:len((*c.CallOptions).CreateAssessment):len((*c.CallOptions).CreateAssessment)], opts...)
	var resp *recaptchaenterprisepb.Assessment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recaptchaEnterpriseServiceV1Beta1Client.CreateAssessment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *recaptchaEnterpriseServiceV1Beta1GRPCClient) AnnotateAssessment(ctx context.Context, req *recaptchaenterprisepb.AnnotateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AnnotateAssessment[0:len((*c.CallOptions).AnnotateAssessment):len((*c.CallOptions).AnnotateAssessment)], opts...)
	var resp *recaptchaenterprisepb.AnnotateAssessmentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recaptchaEnterpriseServiceV1Beta1Client.AnnotateAssessment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateAssessment creates an Assessment of the likelihood an event is legitimate.
func (c *recaptchaEnterpriseServiceV1Beta1RESTClient) CreateAssessment(ctx context.Context, req *recaptchaenterprisepb.CreateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetAssessment()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/assessments", req.GetParent())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateAssessment[0:len((*c.CallOptions).CreateAssessment):len((*c.CallOptions).CreateAssessment)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &recaptchaenterprisepb.Assessment{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// AnnotateAssessment annotates a previously created Assessment to provide additional information
// on whether the event turned out to be authentic or fradulent.
func (c *recaptchaEnterpriseServiceV1Beta1RESTClient) AnnotateAssessment(ctx context.Context, req *recaptchaenterprisepb.AnnotateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:annotate", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).AnnotateAssessment[0:len((*c.CallOptions).AnnotateAssessment):len((*c.CallOptions).AnnotateAssessment)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &recaptchaenterprisepb.AnnotateAssessmentResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
