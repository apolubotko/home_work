package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	type args struct {
		cmd []string
		env Environment
	}
	tests := []struct {
		name           string
		args           args
		wantReturnCode int
	}{
		{
			name: "Failure exit code",
			args: args{
				cmd: []string{"/test/sh", "arg1"},
			},
			wantReturnCode: -1,
		},
		{
			name: "Success exit code",
			args: args{
				cmd: []string{"/bin/bash", "testdata/echo.sh", "hello"},
			},
			wantReturnCode: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := RunCmd(tt.args.cmd, tt.args.env)
			assert.Equal(t, rc, tt.wantReturnCode)
		})
	}
}
