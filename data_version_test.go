package lure

import (
	"reflect"
	"testing"
)

func TestNewVersionFromString(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    VersionData
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{ver: ""},
			want:    VersionData{0, 0, 0},
			wantErr: true,
		},
		{
			name:    "normal",
			args:    args{ver: "v3.2.1"},
			want:    VersionData{3, 2, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVersionFromString(tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersionFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersionFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
