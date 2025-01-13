// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/fleetdm/fleet/v4/server/fleet"
)

var _ fleet.MDMBootstrapPackageStore = (*MDMBootstrapPackageStore)(nil)

type GetFunc func(ctx context.Context, packageID string) (io.ReadCloser, int64, error)

type PutFunc func(ctx context.Context, packageID string, content io.ReadSeeker) error

type ExistsFunc func(ctx context.Context, packageID string) (bool, error)

type CleanupFunc func(ctx context.Context, usedPackageIDs []string, removeCreatedBefore time.Time) (int, error)

type SignFunc func(ctx context.Context, fileID string) (string, error)

type MDMBootstrapPackageStore struct {
	GetFunc        GetFunc
	GetFuncInvoked bool

	PutFunc        PutFunc
	PutFuncInvoked bool

	ExistsFunc        ExistsFunc
	ExistsFuncInvoked bool

	CleanupFunc        CleanupFunc
	CleanupFuncInvoked bool

	SignFunc        SignFunc
	SignFuncInvoked bool

	mu sync.Mutex
}

func (fs *MDMBootstrapPackageStore) Get(ctx context.Context, packageID string) (io.ReadCloser, int64, error) {
	fs.mu.Lock()
	fs.GetFuncInvoked = true
	fs.mu.Unlock()
	return fs.GetFunc(ctx, packageID)
}

func (fs *MDMBootstrapPackageStore) Put(ctx context.Context, packageID string, content io.ReadSeeker) error {
	fs.mu.Lock()
	fs.PutFuncInvoked = true
	fs.mu.Unlock()
	return fs.PutFunc(ctx, packageID, content)
}

func (fs *MDMBootstrapPackageStore) Exists(ctx context.Context, packageID string) (bool, error) {
	fs.mu.Lock()
	fs.ExistsFuncInvoked = true
	fs.mu.Unlock()
	return fs.ExistsFunc(ctx, packageID)
}

func (fs *MDMBootstrapPackageStore) Cleanup(ctx context.Context, usedPackageIDs []string, removeCreatedBefore time.Time) (int, error) {
	fs.mu.Lock()
	fs.CleanupFuncInvoked = true
	fs.mu.Unlock()
	return fs.CleanupFunc(ctx, usedPackageIDs, removeCreatedBefore)
}

func (fs *MDMBootstrapPackageStore) Sign(ctx context.Context, fileID string) (string, error) {
	fs.mu.Lock()
	fs.SignFuncInvoked = true
	fs.mu.Unlock()
	return fs.SignFunc(ctx, fileID)
}