package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input_headers http.Header
		want_key      string
		want_error    error
	}{
		"simple": {
			input_headers: http.Header{"Authorization": []string{"ApiKey deadbeef"}},
			want_key:      "deadbeef",
			want_error:    nil,
		},
		"no-auth-header": {
			input_headers: http.Header{"Hello": []string{"World"}},
			want_key:      "",
			want_error:    errors.New("no authorization header included"),
		},
		"malformed-header": {
			input_headers: http.Header{"Authorization": []string{"ApiKey"}},
			want_key:      "",
			want_error:    errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got_key, got_error := GetAPIKey(tc.input_headers)

			if !reflect.DeepEqual(got_key, tc.want_key) {
				t.Fatalf("expected: %#v, got: %#v", tc.want_key, got_key)
			}

			if got_error != nil && tc.want_error == nil {
				t.Fatalf("Unexpected error: %v", got_error)
			} else if got_error == nil && tc.want_error != nil {
				t.Fatalf("Expected error: %v, got nil", tc.want_error)
			}
		})
	}
}
