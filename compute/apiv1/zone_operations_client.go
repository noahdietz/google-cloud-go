// Copyright 2021 Google LLC
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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newZoneOperationsClientHook clientHook

// ZoneOperationsCallOptions contains the retry settings for each method of ZoneOperationsClient.
type ZoneOperationsCallOptions struct {
	Delete []gax.CallOption
	Get    []gax.CallOption
	List   []gax.CallOption
	Wait   []gax.CallOption
}

// internalZoneOperationsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalZoneOperationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteZoneOperationRequest, ...gax.CallOption) (*computepb.DeleteZoneOperationResponse, error)
	Get(context.Context, *computepb.GetZoneOperationRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListZoneOperationsRequest, ...gax.CallOption) *OperationIterator
	Wait(context.Context, *computepb.WaitZoneOperationRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// ZoneOperationsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The ZoneOperations API.
type ZoneOperationsClient struct {
	// The internal transport-dependent client.
	internalClient internalZoneOperationsClient

	// The call options for this service.
	CallOptions *ZoneOperationsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ZoneOperationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ZoneOperationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ZoneOperationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified zone-specific Operations resource.
func (c *ZoneOperationsClient) Delete(ctx context.Context, req *computepb.DeleteZoneOperationRequest, opts ...gax.CallOption) (*computepb.DeleteZoneOperationResponse, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get retrieves the specified zone-specific Operations resource.
func (c *ZoneOperationsClient) Get(ctx context.Context, req *computepb.GetZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves a list of Operation resources contained within the specified zone.
func (c *ZoneOperationsClient) List(ctx context.Context, req *computepb.ListZoneOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Wait waits for the specified Operation resource to return as DONE or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method differs from the GET method in that it waits for no more than the default deadline (2 minutes) and then returns the current state of the operation, which might be DONE or still in progress.
//
// This method is called on a best-effort basis. Specifically:
//
//   In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds.
//
//   If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not DONE.
func (c *ZoneOperationsClient) Wait(ctx context.Context, req *computepb.WaitZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Wait(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type zoneOperationsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewZoneOperationsRESTClient creates a new zone operations rest client.
//
// The ZoneOperations API.
func NewZoneOperationsRESTClient(ctx context.Context, opts ...option.ClientOption) (*ZoneOperationsClient, error) {
	clientOpts := append(defaultZoneOperationsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &zoneOperationsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &ZoneOperationsClient{internalClient: c, CallOptions: &ZoneOperationsCallOptions{}}, nil
}

func defaultZoneOperationsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *zoneOperationsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *zoneOperationsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *zoneOperationsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified zone-specific Operations resource.
func (c *zoneOperationsRESTClient) Delete(ctx context.Context, req *computepb.DeleteZoneOperationRequest, opts ...gax.CallOption) (*computepb.DeleteZoneOperationResponse, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations/%v", req.GetProject(), req.GetZone(), req.GetOperation())

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.DeleteZoneOperationResponse{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get retrieves the specified zone-specific Operations resource.
func (c *zoneOperationsRESTClient) Get(ctx context.Context, req *computepb.GetZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations/%v", req.GetProject(), req.GetZone(), req.GetOperation())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// List retrieves a list of Operation resources contained within the specified zone.
func (c *zoneOperationsRESTClient) List(ctx context.Context, req *computepb.ListZoneOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*computepb.ListZoneOperationsRequest)
	m := protojson.MarshalOptions{AllowPartial: true, UseProtoNames: false}
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Operation, string, error) {
		resp := &computepb.OperationList{}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		req.PageToken = proto.String(pageToken)

		jsonReq, err := m.Marshal(req)
		if err != nil {
			return nil, "", err
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations", req.GetProject(), req.GetZone())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if httpRsp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf(httpRsp.Status)
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Wait waits for the specified Operation resource to return as DONE or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method differs from the GET method in that it waits for no more than the default deadline (2 minutes) and then returns the current state of the operation, which might be DONE or still in progress.
//
// This method is called on a best-effort basis. Specifically:
//
//   In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds.
//
//   If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not DONE.
func (c *zoneOperationsRESTClient) Wait(ctx context.Context, req *computepb.WaitZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations/%v/wait", req.GetProject(), req.GetZone(), req.GetOperation())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}
