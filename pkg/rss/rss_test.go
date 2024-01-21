package rss

import (
	"GoNews/pkg/storage"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestParse_Success",
			args: args{
				url: "https://example.com/rss",
			},
			want: []storage.Post{
				{
					Title:   "Post 1",
					Content: "Content 1",
					Link:    "https://example.com/post1",
					PubTime: 1631234567,
				},
				{
					Title:   "Post 2",
					Content: "Content 2",
					Link:    "https://example.com/post2",
					PubTime: 1631234568,
				},
			},
			wantErr: false,
		},
		{
			name: "TestParse_Error",
			args: args{
				url: "https://example.com/invalid",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}