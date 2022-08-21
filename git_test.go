package git

import (
	"testing"
)

func TestGIT_PlainCloneSSH(t *testing.T) {
	type args struct {
		directory      string
		url            string
		privateKeyFile string
		skiptls        bool
		depth          int
	}
	tests := []struct {
		name    string
		g       *GIT
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GIT{}
			if err := g.PlainCloneSSH(tt.args.directory, tt.args.url, tt.args.privateKeyFile, tt.args.skiptls, tt.args.depth); (err != nil) != tt.wantErr {
				t.Errorf("GIT.PlainCloneSSH() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGIT_PlainCloneHTTP(t *testing.T) {
	type args struct {
		directory string
		url       string
		token     string
		skiptls   bool
		depth     int
	}
	tests := []struct {
		name    string
		g       *GIT
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GIT{}
			if err := g.PlainCloneHTTP(tt.args.directory, tt.args.url, tt.args.token, tt.args.skiptls, tt.args.depth); (err != nil) != tt.wantErr {
				t.Errorf("GIT.PlainCloneHTTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
