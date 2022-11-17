package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/wallet"
)

const testDIDToken = "WyIweGFhNTBiZTcwNzI5Y2E3MDViYTdjOGQwMDE4NWM2ZjJkYTQ3OWQwZm" +
	"NkZTUzMTFjYTRjZTViMWJhNzE1YzhhNzIxYzVmMTk0ODQzNGY5NmZmNTc3ZDdiMmI2YWQ4MmQ" +
	"zZGQ1YTI0NTdmZTY5OThiMTM3ZWQ5YmMwOGQzNmU1NDljMWIiLCJ7XCJpYXRcIjoxNTg2NzY0" +
	"MjcwLFwiZXh0XCI6MTExNzM1Mjg1MDAsXCJpc3NcIjpcImRpZDpldGhyOjB4NEI3M0M1ODM3M" +
	"EFFZmNFZjg2QTYwMjFhZkNEZTU2NzM1MTEzNzZCMlwiLFwic3ViXCI6XCJOanJBNTNTY1E4SV" +
	"Y4ME5Kbng0dDNTaGk5LWtGZkY1cWF2RDJWcjBkMWRjPVwiLFwiYXVkXCI6XCJkaWQ6bWFnaWM" +
	"6NzMxODQ4Y2MtMDg0ZS00MWZmLWJiZGYtN2YxMDM4MTdlYTZiXCIsXCJuYmZcIjoxNTg2NzY0" +
	"MjcwLFwidGlkXCI6XCJlYmNjODgwYS1mZmM5LTQzNzUtODRhZS0xNTRjY2Q1Yzc0NmRcIixcI" +
	"mFkZFwiOlwiMHg4NGQ2ODM5MjY4YTFhZjkxMTFmZGVjY2QzOTZmMzAzODA1ZGNhMmJjMDM0NT" +
	"BiN2ViMTE2ZTJmNWZjOGM1YTcyMmQxZmI5YWYyMzNhYTczYzVjMTcwODM5Y2U1YWQ4MTQxYjl" +
	"iNDY0MzM4MDk4MmRhNGJmYmIwYjExMjg0OTg4ZjFiXCJ9Il0="

const testSecret = "sk_test_E123E4567E8901D2"

func TestUserGetMetadata(t *testing.T) {
	srv := createServerSuccess(t)
	defer srv.Close()

	// Replace host url to test one.
	client := magic.NewDefaultClient()
	client.SetHostURL(srv.URL)

	uClient := NewUserClient(testSecret, client)

	meta, err := uClient.GetMetadataByToken(testDIDToken)
	require.NoError(t, err, "can't create new token")

	assert.Equal(t, "user@email.com", meta.Email)
	assert.Equal(t, "did:ethr:0x4B73C58370AEfcEf86A6021afCDe5673511376B2", meta.Issuer)
	assert.Equal(t, "0x4B73C58370AEfcEf86A6021afCDe5673511376B2", meta.PublicAddress)
}

func TestUserGetMetadataWithWallet(t *testing.T) {
	srv := createServerSuccess(t)
	defer srv.Close()

	// Replace host url to test one.
	client := magic.NewDefaultClient()
	client.SetHostURL(srv.URL)

	uClient := NewUserClient(testSecret, client)

	meta, err := uClient.GetMetadataByTokenAndWallet(testDIDToken, wallet.ETH)
	require.NoError(t, err, "can't create new token")

	assert.Equal(t, "user@email.com", meta.Email)
	assert.Equal(t, "did:ethr:0x4B73C58370AEfcEf86A6021afCDe5673511376B2", meta.Issuer)
	assert.Equal(t, "0x4B73C58370AEfcEf86A6021afCDe5673511376B2", meta.PublicAddress)
	assert.NotNil(t, meta.Wallets)
}

func TestUserGetMetadataWrongSecret(t *testing.T) {
	srv := createServerSuccess(t)
	defer srv.Close()

	// Replace host url to test one.
	client := magic.NewDefaultClient()
	client.SetHostURL(srv.URL)

	uClient := NewUserClient("wrong_secret", client)

	_, err := uClient.GetMetadataByToken(testDIDToken)
	require.Error(t, err, "server must return error")
	_, ok := err.(*magic.AuthenticationError)
	require.True(t, ok, "Error type must be AuthenticationError")
}

func TestUserGetMetadataBackendFailure(t *testing.T) {
	srv := createServerFail(t)
	defer srv.Close()

	// Replace host url to test one.
	client := magic.NewDefaultClient()
	client.SetHostURL(srv.URL)

	uClient := NewUserClient("wrong_secret", client)

	_, err := uClient.GetMetadataByToken(testDIDToken)
	require.Error(t, err, "server must return error")
	_, ok := err.(*magic.APIError)
	require.True(t, ok, "Error type must be APIError")
}

func TestUserLogout(t *testing.T) {
	srv := createServerSuccess(t)
	defer srv.Close()

	// Replace host url to test one.
	client := magic.NewDefaultClient()
	client.SetHostURL(srv.URL)

	uClient := NewUserClient(testSecret, client)

	err := uClient.LogoutByToken(testDIDToken)
	require.NoError(t, err, "can't logout user by DID token")
}

func TestUserLogoutBackendFailure(t *testing.T) {
	srv := createServerFail(t)
	defer srv.Close()

	// Replace host url to test one.
	client := magic.NewDefaultClient()
	client.SetHostURL(srv.URL)

	uClient := NewUserClient("wrong_secret", client)

	err := uClient.LogoutByToken(testDIDToken)
	require.Error(t, err, "server must return error")
	_, ok := err.(*magic.APIError)
	require.True(t, ok, "Error type must be APIError")
}

// createServerSuccess creates internal server which simulates positive case for backend api requests.
func createServerSuccess(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Method: %v", r.Method)
		t.Logf("Path: %v", r.URL.Path)

		w.Header().Add("Content-Type", "application/json")

		secret := r.Header.Get(magic.APISecretHeader)
		if secret != testSecret {
			w.WriteHeader(http.StatusUnauthorized)
			resp := magic.Response{
				ErrorCode: "err_code_unauthorized",
				Message:   "unauthorized",
				Status:    "fail",
			}
			data, err := json.Marshal(resp)
			require.NoError(t, err, "can't marshal test data")
			_, _ = w.Write(data)
			t.Log("401 - Unauthorized")
			return
		}

		switch r.Method {
		case http.MethodGet:
			switch r.URL.Path {
			case userInfoV1:
				resp := magic.Response{
					Data: &magic.UserInfo{
						Email:         "user@email.com",
						Issuer:        "did:ethr:0x4B73C58370AEfcEf86A6021afCDe5673511376B2",
						PublicAddress: "0x4B73C58370AEfcEf86A6021afCDe5673511376B2",
					},
					Status: "ok",
				}
				data, err := json.Marshal(resp)
				require.NoError(t, err, "can't marshal test data")
				_, _ = w.Write(data)
			}

		case http.MethodPost:
			switch r.URL.Path {
			case userLogoutV2:
				resp := magic.Response{
					Status: "ok",
				}
				data, err := json.Marshal(resp)
				require.NoError(t, err, "can't marshal test data")
				_, _ = w.Write(data)
			}
		}
	}))
}

// createServerFail creates internal server which simulates negative case for backend api requests.
func createServerFail(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Method: %v", r.Method)
		t.Logf("Path: %v", r.URL.Path)

		w.Header().Add("Content-Type", "application/json")

		switch r.URL.Path {
		case userInfoV1, userLogoutV2:
			w.WriteHeader(http.StatusInternalServerError)
			resp := magic.Response{
				ErrorCode: "err_code_internal_server_error",
				Message:   "internal server error",
				Status:    "fail",
			}
			data, err := json.Marshal(resp)
			require.NoError(t, err, "can't marshal test data")
			_, _ = w.Write(data)
			t.Log("500 - Internal server error")
			return
		}
	}))
}
