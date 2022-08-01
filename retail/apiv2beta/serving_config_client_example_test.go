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

package retail_test

import (
	"context"

	retail "cloud.google.com/go/retail/apiv2beta"
	"google.golang.org/api/iterator"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2beta"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewServingConfigClient() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleServingConfigClient_CreateServingConfig() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.CreateServingConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#CreateServingConfigRequest.
	}
	resp, err := c.CreateServingConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServingConfigClient_DeleteServingConfig() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.DeleteServingConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#DeleteServingConfigRequest.
	}
	err = c.DeleteServingConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleServingConfigClient_UpdateServingConfig() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateServingConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#UpdateServingConfigRequest.
	}
	resp, err := c.UpdateServingConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServingConfigClient_GetServingConfig() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetServingConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#GetServingConfigRequest.
	}
	resp, err := c.GetServingConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServingConfigClient_ListServingConfigs() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ListServingConfigsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#ListServingConfigsRequest.
	}
	it := c.ListServingConfigs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleServingConfigClient_AddControl() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.AddControlRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#AddControlRequest.
	}
	resp, err := c.AddControl(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServingConfigClient_RemoveControl() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.RemoveControlRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2beta#RemoveControlRequest.
	}
	resp, err := c.RemoveControl(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServingConfigClient_GetOperation() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServingConfigClient_ListOperations() {
	ctx := context.Background()
	c, err := retail.NewServingConfigClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
