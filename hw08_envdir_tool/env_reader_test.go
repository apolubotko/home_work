package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
		key     string
		wantErr bool
	}{
		{
			name: "Basic tests",
			args: args{
				dir: "./testdata/env",
			},
			want: Environment{
				"BAR": EnvValue{
					Value:      "bar",
					NeedRemove: false,
				},
				"FOO": EnvValue{
					Value:      "   foo\nwith new line",
					NeedRemove: true,
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDir(tt.args.dir)
			require.Equal(t, tt.want["BAR"].Value, got["BAR"].Value)
			require.Equal(t, tt.want["FOO"].Value, got["FOO"].Value)
			require.Len(t, got["EMPTY"].Value, 0)
			require.NoError(t, err)
		})
	}
}
