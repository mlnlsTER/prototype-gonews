package api

import (
	"net/http"
	"testing"
)

func TestAPI_postsHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		api  *API
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.api.postsHandler(tt.args.w, tt.args.r)
		})
	}
}
