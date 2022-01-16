package api

import (
	"net/http"
	"testing"
)

func TestDoc(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Doc(tt.args.w, tt.args.r)
		})
	}
}

func TestAllPages(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AllPages(tt.args.w, tt.args.r)
		})
	}
}

func TestGetPage(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPage(tt.args.w, tt.args.r)
		})
	}
}

func TestCreatePage(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePage(tt.args.w, tt.args.r)
		})
	}
}

func Test_writeJSON(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeJSON(tt.args.w, tt.args.data)
		})
	}
}

func Test_errJSON(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		err    string
		status int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errJSON(tt.args.w, tt.args.err, tt.args.status)
		})
	}
}
