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

package storagetransfer

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	storagetransferpb "google.golang.org/genproto/googleapis/storagetransfer/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	GetGoogleServiceAccount []gax.CallOption
	CreateTransferJob       []gax.CallOption
	UpdateTransferJob       []gax.CallOption
	GetTransferJob          []gax.CallOption
	ListTransferJobs        []gax.CallOption
	PauseTransferOperation  []gax.CallOption
	ResumeTransferOperation []gax.CallOption
	RunTransferJob          []gax.CallOption
	CreateAgentPool         []gax.CallOption
	UpdateAgentPool         []gax.CallOption
	GetAgentPool            []gax.CallOption
	ListAgentPools          []gax.CallOption
	DeleteAgentPool         []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("storagetransfer.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("storagetransfer.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://storagetransfer.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetGoogleServiceAccount: []gax.CallOption{},
		CreateTransferJob:       []gax.CallOption{},
		UpdateTransferJob:       []gax.CallOption{},
		GetTransferJob:          []gax.CallOption{},
		ListTransferJobs:        []gax.CallOption{},
		PauseTransferOperation:  []gax.CallOption{},
		ResumeTransferOperation: []gax.CallOption{},
		RunTransferJob:          []gax.CallOption{},
		CreateAgentPool:         []gax.CallOption{},
		UpdateAgentPool:         []gax.CallOption{},
		GetAgentPool:            []gax.CallOption{},
		ListAgentPools:          []gax.CallOption{},
		DeleteAgentPool:         []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Storage Transfer API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetGoogleServiceAccount(context.Context, *storagetransferpb.GetGoogleServiceAccountRequest, ...gax.CallOption) (*storagetransferpb.GoogleServiceAccount, error)
	CreateTransferJob(context.Context, *storagetransferpb.CreateTransferJobRequest, ...gax.CallOption) (*storagetransferpb.TransferJob, error)
	UpdateTransferJob(context.Context, *storagetransferpb.UpdateTransferJobRequest, ...gax.CallOption) (*storagetransferpb.TransferJob, error)
	GetTransferJob(context.Context, *storagetransferpb.GetTransferJobRequest, ...gax.CallOption) (*storagetransferpb.TransferJob, error)
	ListTransferJobs(context.Context, *storagetransferpb.ListTransferJobsRequest, ...gax.CallOption) *TransferJobIterator
	PauseTransferOperation(context.Context, *storagetransferpb.PauseTransferOperationRequest, ...gax.CallOption) error
	ResumeTransferOperation(context.Context, *storagetransferpb.ResumeTransferOperationRequest, ...gax.CallOption) error
	RunTransferJob(context.Context, *storagetransferpb.RunTransferJobRequest, ...gax.CallOption) (*RunTransferJobOperation, error)
	RunTransferJobOperation(name string) *RunTransferJobOperation
	CreateAgentPool(context.Context, *storagetransferpb.CreateAgentPoolRequest, ...gax.CallOption) (*storagetransferpb.AgentPool, error)
	UpdateAgentPool(context.Context, *storagetransferpb.UpdateAgentPoolRequest, ...gax.CallOption) (*storagetransferpb.AgentPool, error)
	GetAgentPool(context.Context, *storagetransferpb.GetAgentPoolRequest, ...gax.CallOption) (*storagetransferpb.AgentPool, error)
	ListAgentPools(context.Context, *storagetransferpb.ListAgentPoolsRequest, ...gax.CallOption) *AgentPoolIterator
	DeleteAgentPool(context.Context, *storagetransferpb.DeleteAgentPoolRequest, ...gax.CallOption) error
}

// Client is a client for interacting with Storage Transfer API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Storage Transfer Service and its protos.
// Transfers data between between Google Cloud Storage buckets or from a data
// source external to Google to a Cloud Storage bucket.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetGoogleServiceAccount returns the Google service account that is used by Storage Transfer
// Service to access buckets in the project where transfers
// run or in other projects. Each Google service account is associated
// with one Google Cloud project. Users
// should add this service account to the Google Cloud Storage bucket
// ACLs to grant access to Storage Transfer Service. This service
// account is created and owned by Storage Transfer Service and can
// only be used by Storage Transfer Service.
func (c *Client) GetGoogleServiceAccount(ctx context.Context, req *storagetransferpb.GetGoogleServiceAccountRequest, opts ...gax.CallOption) (*storagetransferpb.GoogleServiceAccount, error) {
	return c.internalClient.GetGoogleServiceAccount(ctx, req, opts...)
}

// CreateTransferJob creates a transfer job that runs periodically.
func (c *Client) CreateTransferJob(ctx context.Context, req *storagetransferpb.CreateTransferJobRequest, opts ...gax.CallOption) (*storagetransferpb.TransferJob, error) {
	return c.internalClient.CreateTransferJob(ctx, req, opts...)
}

// UpdateTransferJob updates a transfer job. Updating a job’s transfer spec does not affect
// transfer operations that are running already.
//
// Note: The job’s status field can be modified
// using this RPC (for example, to set a job’s status to
// DELETED,
// DISABLED, or
// ENABLED).
func (c *Client) UpdateTransferJob(ctx context.Context, req *storagetransferpb.UpdateTransferJobRequest, opts ...gax.CallOption) (*storagetransferpb.TransferJob, error) {
	return c.internalClient.UpdateTransferJob(ctx, req, opts...)
}

// GetTransferJob gets a transfer job.
func (c *Client) GetTransferJob(ctx context.Context, req *storagetransferpb.GetTransferJobRequest, opts ...gax.CallOption) (*storagetransferpb.TransferJob, error) {
	return c.internalClient.GetTransferJob(ctx, req, opts...)
}

// ListTransferJobs lists transfer jobs.
func (c *Client) ListTransferJobs(ctx context.Context, req *storagetransferpb.ListTransferJobsRequest, opts ...gax.CallOption) *TransferJobIterator {
	return c.internalClient.ListTransferJobs(ctx, req, opts...)
}

// PauseTransferOperation pauses a transfer operation.
func (c *Client) PauseTransferOperation(ctx context.Context, req *storagetransferpb.PauseTransferOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.PauseTransferOperation(ctx, req, opts...)
}

// ResumeTransferOperation resumes a transfer operation that is paused.
func (c *Client) ResumeTransferOperation(ctx context.Context, req *storagetransferpb.ResumeTransferOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.ResumeTransferOperation(ctx, req, opts...)
}

// RunTransferJob attempts to start a new TransferOperation for the current TransferJob. A
// TransferJob has a maximum of one active TransferOperation. If this method
// is called while a TransferOperation is active, an error will be returned.
func (c *Client) RunTransferJob(ctx context.Context, req *storagetransferpb.RunTransferJobRequest, opts ...gax.CallOption) (*RunTransferJobOperation, error) {
	return c.internalClient.RunTransferJob(ctx, req, opts...)
}

// RunTransferJobOperation returns a new RunTransferJobOperation from a given name.
// The name must be that of a previously created RunTransferJobOperation, possibly from a different process.
func (c *Client) RunTransferJobOperation(name string) *RunTransferJobOperation {
	return c.internalClient.RunTransferJobOperation(name)
}

// CreateAgentPool creates an agent pool resource.
func (c *Client) CreateAgentPool(ctx context.Context, req *storagetransferpb.CreateAgentPoolRequest, opts ...gax.CallOption) (*storagetransferpb.AgentPool, error) {
	return c.internalClient.CreateAgentPool(ctx, req, opts...)
}

// UpdateAgentPool updates an existing agent pool resource.
func (c *Client) UpdateAgentPool(ctx context.Context, req *storagetransferpb.UpdateAgentPoolRequest, opts ...gax.CallOption) (*storagetransferpb.AgentPool, error) {
	return c.internalClient.UpdateAgentPool(ctx, req, opts...)
}

// GetAgentPool gets an agent pool.
func (c *Client) GetAgentPool(ctx context.Context, req *storagetransferpb.GetAgentPoolRequest, opts ...gax.CallOption) (*storagetransferpb.AgentPool, error) {
	return c.internalClient.GetAgentPool(ctx, req, opts...)
}

// ListAgentPools lists agent pools.
func (c *Client) ListAgentPools(ctx context.Context, req *storagetransferpb.ListAgentPoolsRequest, opts ...gax.CallOption) *AgentPoolIterator {
	return c.internalClient.ListAgentPools(ctx, req, opts...)
}

// DeleteAgentPool deletes an agent pool.
func (c *Client) DeleteAgentPool(ctx context.Context, req *storagetransferpb.DeleteAgentPoolRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAgentPool(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Storage Transfer API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client storagetransferpb.StorageTransferServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new storage transfer service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Storage Transfer Service and its protos.
// Transfers data between between Google Cloud Storage buckets or from a data
// source external to Google to a Cloud Storage bucket.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           storagetransferpb.NewStorageTransferServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) GetGoogleServiceAccount(ctx context.Context, req *storagetransferpb.GetGoogleServiceAccountRequest, opts ...gax.CallOption) (*storagetransferpb.GoogleServiceAccount, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetGoogleServiceAccount[0:len((*c.CallOptions).GetGoogleServiceAccount):len((*c.CallOptions).GetGoogleServiceAccount)], opts...)
	var resp *storagetransferpb.GoogleServiceAccount
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetGoogleServiceAccount(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateTransferJob(ctx context.Context, req *storagetransferpb.CreateTransferJobRequest, opts ...gax.CallOption) (*storagetransferpb.TransferJob, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateTransferJob[0:len((*c.CallOptions).CreateTransferJob):len((*c.CallOptions).CreateTransferJob)], opts...)
	var resp *storagetransferpb.TransferJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateTransferJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateTransferJob(ctx context.Context, req *storagetransferpb.UpdateTransferJobRequest, opts ...gax.CallOption) (*storagetransferpb.TransferJob, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "job_name", url.QueryEscape(req.GetJobName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateTransferJob[0:len((*c.CallOptions).UpdateTransferJob):len((*c.CallOptions).UpdateTransferJob)], opts...)
	var resp *storagetransferpb.TransferJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateTransferJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetTransferJob(ctx context.Context, req *storagetransferpb.GetTransferJobRequest, opts ...gax.CallOption) (*storagetransferpb.TransferJob, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "job_name", url.QueryEscape(req.GetJobName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTransferJob[0:len((*c.CallOptions).GetTransferJob):len((*c.CallOptions).GetTransferJob)], opts...)
	var resp *storagetransferpb.TransferJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetTransferJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListTransferJobs(ctx context.Context, req *storagetransferpb.ListTransferJobsRequest, opts ...gax.CallOption) *TransferJobIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListTransferJobs[0:len((*c.CallOptions).ListTransferJobs):len((*c.CallOptions).ListTransferJobs)], opts...)
	it := &TransferJobIterator{}
	req = proto.Clone(req).(*storagetransferpb.ListTransferJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*storagetransferpb.TransferJob, string, error) {
		resp := &storagetransferpb.ListTransferJobsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListTransferJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTransferJobs(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) PauseTransferOperation(ctx context.Context, req *storagetransferpb.PauseTransferOperationRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PauseTransferOperation[0:len((*c.CallOptions).PauseTransferOperation):len((*c.CallOptions).PauseTransferOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.PauseTransferOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ResumeTransferOperation(ctx context.Context, req *storagetransferpb.ResumeTransferOperationRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ResumeTransferOperation[0:len((*c.CallOptions).ResumeTransferOperation):len((*c.CallOptions).ResumeTransferOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.ResumeTransferOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) RunTransferJob(ctx context.Context, req *storagetransferpb.RunTransferJobRequest, opts ...gax.CallOption) (*RunTransferJobOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "job_name", url.QueryEscape(req.GetJobName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RunTransferJob[0:len((*c.CallOptions).RunTransferJob):len((*c.CallOptions).RunTransferJob)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RunTransferJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunTransferJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) CreateAgentPool(ctx context.Context, req *storagetransferpb.CreateAgentPoolRequest, opts ...gax.CallOption) (*storagetransferpb.AgentPool, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAgentPool[0:len((*c.CallOptions).CreateAgentPool):len((*c.CallOptions).CreateAgentPool)], opts...)
	var resp *storagetransferpb.AgentPool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAgentPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateAgentPool(ctx context.Context, req *storagetransferpb.UpdateAgentPoolRequest, opts ...gax.CallOption) (*storagetransferpb.AgentPool, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "agent_pool.name", url.QueryEscape(req.GetAgentPool().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateAgentPool[0:len((*c.CallOptions).UpdateAgentPool):len((*c.CallOptions).UpdateAgentPool)], opts...)
	var resp *storagetransferpb.AgentPool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateAgentPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetAgentPool(ctx context.Context, req *storagetransferpb.GetAgentPoolRequest, opts ...gax.CallOption) (*storagetransferpb.AgentPool, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAgentPool[0:len((*c.CallOptions).GetAgentPool):len((*c.CallOptions).GetAgentPool)], opts...)
	var resp *storagetransferpb.AgentPool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAgentPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListAgentPools(ctx context.Context, req *storagetransferpb.ListAgentPoolsRequest, opts ...gax.CallOption) *AgentPoolIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListAgentPools[0:len((*c.CallOptions).ListAgentPools):len((*c.CallOptions).ListAgentPools)], opts...)
	it := &AgentPoolIterator{}
	req = proto.Clone(req).(*storagetransferpb.ListAgentPoolsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*storagetransferpb.AgentPool, string, error) {
		resp := &storagetransferpb.ListAgentPoolsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListAgentPools(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAgentPools(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) DeleteAgentPool(ctx context.Context, req *storagetransferpb.DeleteAgentPoolRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteAgentPool[0:len((*c.CallOptions).DeleteAgentPool):len((*c.CallOptions).DeleteAgentPool)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteAgentPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// RunTransferJobOperation manages a long-running operation from RunTransferJob.
type RunTransferJobOperation struct {
	lro *longrunning.Operation
}

// RunTransferJobOperation returns a new RunTransferJobOperation from a given name.
// The name must be that of a previously created RunTransferJobOperation, possibly from a different process.
func (c *gRPCClient) RunTransferJobOperation(name string) *RunTransferJobOperation {
	return &RunTransferJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunTransferJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RunTransferJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RunTransferJobOperation) Metadata() (*storagetransferpb.TransferOperation, error) {
	var meta storagetransferpb.TransferOperation
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunTransferJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunTransferJobOperation) Name() string {
	return op.lro.Name()
}

// AgentPoolIterator manages a stream of *storagetransferpb.AgentPool.
type AgentPoolIterator struct {
	items    []*storagetransferpb.AgentPool
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*storagetransferpb.AgentPool, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AgentPoolIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AgentPoolIterator) Next() (*storagetransferpb.AgentPool, error) {
	var item *storagetransferpb.AgentPool
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AgentPoolIterator) bufLen() int {
	return len(it.items)
}

func (it *AgentPoolIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TransferJobIterator manages a stream of *storagetransferpb.TransferJob.
type TransferJobIterator struct {
	items    []*storagetransferpb.TransferJob
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*storagetransferpb.TransferJob, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TransferJobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TransferJobIterator) Next() (*storagetransferpb.TransferJob, error) {
	var item *storagetransferpb.TransferJob
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TransferJobIterator) bufLen() int {
	return len(it.items)
}

func (it *TransferJobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
