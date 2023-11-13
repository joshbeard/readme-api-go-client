package readme_test

import "net/http"

// mockPaginatedRequestHeader represents common and valid HTTP headers for paginated requests.
var mockPaginatedRequestHeader = http.Header{
	"Link":          {`<>; rel="next", <>; rel="prev", <>; rel="last"`},
	"X-Total-Count": {"20"},
}