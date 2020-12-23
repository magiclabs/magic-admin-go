package magic

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// APIVersion is the version of the library.
	APIVersion = "v0.1.0"

	// APIURL is the URL of the API service backend.
	APIURL = "https://api.magic.link"

	// APISecretHeader holds the header name for api authorization.
	APISecretHeader = "X-Magic-Secret-Key"
)

var (
	ErrRespQuotaExceeded = errors.New("quota exceeded")
)

// Default response data structure of magic backend server.
type Response struct {
	Data      interface{} `json:"data"`
	ErrorCode ErrorCode   `json:"error_code"`
	Message   string      `json:"message"`
	Status    string      `json:"status"`
}

// NewDefaultClient creates backend client with default configuration of retries.
func NewDefaultClient() *resty.Client {
	return NewClientWithRetry(3, time.Second, 10 * time.Second)
}

// NewClient creates new backend client with default api url.
func NewClient() *resty.Client {
	return resty.New().SetHostURL(APIURL).SetError(new(Error))
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
