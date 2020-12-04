// Copyright 2020 Google LLC
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

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newSessionEntityTypesClientHook clientHook

// SessionEntityTypesCallOptions contains the retry settings for each method of SessionEntityTypesClient.
type SessionEntityTypesCallOptions struct {
	ListSessionEntityTypes  []gax.CallOption
	GetSessionEntityType    []gax.CallOption
	CreateSessionEntityType []gax.CallOption
	UpdateSessionEntityType []gax.CallOption
	DeleteSessionEntityType []gax.CallOption
}

func defaultSessionEntityTypesClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSessionEntityTypesCallOptions() *SessionEntityTypesCallOptions {
	return &SessionEntityTypesCallOptions{
		ListSessionEntityTypes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// SessionEntityTypesClient is a client for interacting with Dialogflow API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type SessionEntityTypesClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	sessionEntityTypesClient cxpb.SessionEntityTypesClient

	// The call options for this service.
	CallOptions *SessionEntityTypesCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSessionEntityTypesClient creates a new session entity types client.
//
// Service for managing SessionEntityTypes.
func NewSessionEntityTypesClient(ctx context.Context, opts ...option.ClientOption) (*SessionEntityTypesClient, error) {
	clientOpts := defaultSessionEntityTypesClientOptions()

	if newSessionEntityTypesClientHook != nil {
		hookOpts, err := newSessionEntityTypesClientHook(ctx, clientHookParams{})
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
	c := &SessionEntityTypesClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultSessionEntityTypesCallOptions(),

		sessionEntityTypesClient: cxpb.NewSessionEntityTypesClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SessionEntityTypesClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SessionEntityTypesClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SessionEntityTypesClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListSessionEntityTypes returns the list of all session entity types in the specified session.
func (c *SessionEntityTypesClient) ListSessionEntityTypes(ctx context.Context, req *cxpb.ListSessionEntityTypesRequest, opts ...gax.CallOption) *SessionEntityTypeIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListSessionEntityTypes[0:len(c.CallOptions.ListSessionEntityTypes):len(c.CallOptions.ListSessionEntityTypes)], opts...)
	it := &SessionEntityTypeIterator{}
	req = proto.Clone(req).(*cxpb.ListSessionEntityTypesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.SessionEntityType, string, error) {
		var resp *cxpb.ListSessionEntityTypesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.sessionEntityTypesClient.ListSessionEntityTypes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSessionEntityTypes(), resp.GetNextPageToken(), nil
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

// GetSessionEntityType retrieves the specified session entity type.
func (c *SessionEntityTypesClient) GetSessionEntityType(ctx context.Context, req *cxpb.GetSessionEntityTypeRequest, opts ...gax.CallOption) (*cxpb.SessionEntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetSessionEntityType[0:len(c.CallOptions.GetSessionEntityType):len(c.CallOptions.GetSessionEntityType)], opts...)
	var resp *cxpb.SessionEntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionEntityTypesClient.GetSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateSessionEntityType creates a session entity type.
//
// If the specified session entity type already exists, overrides the
// session entity type.
func (c *SessionEntityTypesClient) CreateSessionEntityType(ctx context.Context, req *cxpb.CreateSessionEntityTypeRequest, opts ...gax.CallOption) (*cxpb.SessionEntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateSessionEntityType[0:len(c.CallOptions.CreateSessionEntityType):len(c.CallOptions.CreateSessionEntityType)], opts...)
	var resp *cxpb.SessionEntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionEntityTypesClient.CreateSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateSessionEntityType updates the specified session entity type.
func (c *SessionEntityTypesClient) UpdateSessionEntityType(ctx context.Context, req *cxpb.UpdateSessionEntityTypeRequest, opts ...gax.CallOption) (*cxpb.SessionEntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session_entity_type.name", url.QueryEscape(req.GetSessionEntityType().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateSessionEntityType[0:len(c.CallOptions.UpdateSessionEntityType):len(c.CallOptions.UpdateSessionEntityType)], opts...)
	var resp *cxpb.SessionEntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionEntityTypesClient.UpdateSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteSessionEntityType deletes the specified session entity type.
func (c *SessionEntityTypesClient) DeleteSessionEntityType(ctx context.Context, req *cxpb.DeleteSessionEntityTypeRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteSessionEntityType[0:len(c.CallOptions.DeleteSessionEntityType):len(c.CallOptions.DeleteSessionEntityType)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.sessionEntityTypesClient.DeleteSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SessionEntityTypeIterator manages a stream of *cxpb.SessionEntityType.
type SessionEntityTypeIterator struct {
	items    []*cxpb.SessionEntityType
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.SessionEntityType, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SessionEntityTypeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SessionEntityTypeIterator) Next() (*cxpb.SessionEntityType, error) {
	var item *cxpb.SessionEntityType
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SessionEntityTypeIterator) bufLen() int {
	return len(it.items)
}

func (it *SessionEntityTypeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
