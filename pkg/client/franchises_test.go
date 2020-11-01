package client_test

import (
	"net/http"
	"testing"

	"github.com/gira-games/client/internal/fixtures"

	"github.com/gira-games/client/pkg/client"

	"github.com/stretchr/testify/require"
)

var (
	franchise = &client.Franchise{
		ID:   "1",
		Name: "Batman",
	}
	franchises         = []*client.Franchise{franchise}
	franchisesResponse = &client.GetFranchisesResponse{
		Franchises: franchises,
	}
)

func TestFranchisesGet(t *testing.T) {
	ts := fixtures.NewTestServer(t).
		Path("/franchises").
		Token(token).
		Method(http.MethodGet).
		Data(franchisesResponse).
		Build()
	defer ts.Close()

	cl := newClient(t, ts.URL)

	resp, err := cl.GetFranchises(&client.GetFranchisesRequest{Token: token})
	require.NoError(t, err)
	require.Equal(t, franchises, resp.Franchises)
}

func TestFranchisesCreate(t *testing.T) {
	ts := fixtures.NewTestServer(t).
		Path("/franchises").
		Token(token).
		Method(http.MethodPost).
		Data(franchise).
		Build()
	defer ts.Close()

	cl := newClient(t, ts.URL)

	frResponse, err := cl.CreateFranchise(&client.CreateFranchiseRequest{
		Name:  "AC",
		Token: token,
	})
	require.NoError(t, err)
	require.Equal(t, franchise, frResponse.Franchise)
}
