package service

import (
	"context"
	"testing"
	"time"

	"github.com/fleetdm/fleet/server/config"
	"github.com/fleetdm/fleet/server/datastore/inmem"
	"github.com/fleetdm/fleet/server/fleet"
	"github.com/fleetdm/fleet/server/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticate(t *testing.T) {
	ds, err := inmem.New(config.TestConfig())
	require.Nil(t, err)
	svc := newTestService(ds, nil, nil)
	createTestUsers(t, ds)

	var loginTests = []struct {
		name     string
		email    string
		password string
		wantErr  error
	}{
		{
			name:     "admin1",
			email:    testUsers["admin1"].Email,
			password: testUsers["admin1"].PlaintextPassword,
		},
		{
			name:     "user1",
			email:    testUsers["user1"].Email,
			password: testUsers["user1"].PlaintextPassword,
		},
	}

	for _, tt := range loginTests {
		t.Run(tt.email, func(st *testing.T) {
			loggedIn, token, err := svc.Login(test.UserContext(test.UserAdmin), tt.email, tt.password)
			require.Nil(st, err, "login unsuccessful")
			assert.Equal(st, tt.email, loggedIn.Email)
			assert.NotEmpty(st, token)

			sessions, err := svc.GetInfoAboutSessionsForUser(test.UserContext(test.UserAdmin), loggedIn.ID)
			require.Nil(st, err)
			require.Len(st, sessions, 1, "user should have one session")
			session := sessions[0]
			assert.NotZero(st, session.UserID)
			assert.WithinDuration(st, time.Now(), session.AccessedAt, 3*time.Second,
				"access time should be set with current time at session creation")
		})
	}
}

type authViewerService struct {
	fleet.Service
}

func (authViewerService) GetSessionByKey(ctx context.Context, key string) (*fleet.Session, error) {
	return &fleet.Session{}, nil
}

func (authViewerService) UserUnauthorized(ctx context.Context, uid uint) (*fleet.User, error) {
	return &fleet.User{}, nil
}
