package hasher

import (
	"io"
	"strings"
	"testing"
)

var testString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func TestClient_Md5(t *testing.T) {
	type fields struct {
		Progress io.Writer
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct MD5 checksum is calculated",
			args: args{
				r: strings.NewReader(testString),
			},
			want: "818c6e601a24f72750da0f6c9b8ebe28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Progress: tt.fields.Progress,
			}
			got, err := c.Md5(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Md5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Sha1(t *testing.T) {
	type fields struct {
		Progress io.Writer
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct SHA1 checksum is calculated",
			args: args{
				r: strings.NewReader(testString),
			},
			want: "cca0871ecbe200379f0a1e4b46de177e2d62e655",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Progress: tt.fields.Progress,
			}
			got, err := c.Sha1(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Sha224(t *testing.T) {
	type fields struct {
		Progress io.Writer
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct SHA224 checksum is calculated",
			args: args{
				r: strings.NewReader(testString),
			},
			want: "730fbab30ac9c9fce3b7e3af8e0f637ad9705191b1424e0ab798896d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Progress: tt.fields.Progress,
			}
			got, err := c.Sha224(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha224() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha224() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Sha256(t *testing.T) {
	type fields struct {
		Progress io.Writer
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct SHA256 checksum is calculated",
			args: args{
				r: strings.NewReader(testString),
			},
			want: "973153f86ec2da1748e63f0cf85b89835b42f8ee8018c549868a1308a19f6ca3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Progress: tt.fields.Progress,
			}
			got, err := c.Sha256(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha256() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Sha384(t *testing.T) {
	type fields struct {
		Progress io.Writer
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct SHA384 checksum is calculated",
			args: args{
				r: strings.NewReader(testString),
			},
			want: "8ac70afc3c5cb8daf3c332fbe88b7b9c2a5104217693b065a42db8de1bf40aca0bdaad7dbc6c405e97cf0809d1afcf52",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Progress: tt.fields.Progress,
			}
			got, err := c.Sha384(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha384() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha384() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Sha512(t *testing.T) {
	type fields struct {
		Progress io.Writer
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct SHA512 checksum is calculated",
			args: args{
				r: strings.NewReader(testString),
			},
			want: "83cd8866be238eda447cb0ee94a6bfa6248109346b1ce3c75f8a67d35f3d8ab1697b46703065c094fcc7d3a61acc1e8ee85a4f306f13cc1a7aea7651781199b3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Progress: tt.fields.Progress,
			}
			got, err := c.Sha512(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha512() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha512() got = %v, want %v", got, tt.want)
			}
		})
	}
}
