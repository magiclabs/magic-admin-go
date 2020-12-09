package magic

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// APIVersion is the version of the library.
	APIVersion = "v1.0.0"

	// APIURL is the URL of the API service backend.
	APIURL = "https://api.magic.link"

	// APISecretHeader holds the header name for api authorization.
	APISecretHeader = "X-Magic-Secret-Key"
)

var (
	ErrRespQuotaExceeded = errors.New("quota exceeded")
)

// Response default response data structure of magic backend server.
type Response struct {
	Data      interface{} `json:"data, required"`
	ErrorCode string      `json:"error_code, string, required"`
	Message   string      `json:"message, required"`
	Status    string      `json:"status, required"`
}

// Error implements error interface in case of failed request if the status is not equal to "ok".
func (r *Response) Error() string {
	if r.Status == "ok" {
		return r.Status
	}

	return fmt.Sprintf("request is failed, error code: %s, with message: %s",
		r.ErrorCode, r.Message)
}

// NewDefaultClient creates backend client with default configuration of retries.
func NewDefaultClient() *resty.Client {
	return NewClientWithRetry(3, time.Second, 10 * time.Second)
}

// NewClient creates new backend client with default api url.
func NewClient() *resty.Client {
	return resty.New().SetHostURL(APIURL).SetError(new(Response))
}

// NewClientWithRetry creates backend client with backoff retry configuration.
func NewClientWithRetry(retries int, retryWait, timeout time.Duration) *resty.Client {
	client := NewClient()

	// Set retry count to non zero to enable retries
	client.SetRetryCount(retries).
		// Retry wait time till the next request is risen.
		SetRetryWaitTime(retryWait).
		// Wait time of backend response.
		SetRetryMaxWaitTime(timeout).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter.
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, ErrRespQuotaExceeded
		}).
		AddRetryCondition(
			// RetryConditionFunc identify which responses could be considered as retryable.
			func(r *resty.Response, err error) bool {
				return r.StatusCode() == http.StatusTooManyRequests
			},
		)

	return client
}
