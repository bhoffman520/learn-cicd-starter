package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			want:    "1234567890",
			wantErr: nil,
		},
		{
			name:    "no api key",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetAPIKey(test.headers)
			if got != test.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, test.want)
			}
			if err != test.wantErr {
				t.Errorf("GetAPIKey() error = %v, want %v", err, test.wantErr)
			}
		})
	}
}
