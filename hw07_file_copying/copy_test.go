package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name        string
		args        args
		compareFile string
	}{
		{name: "limit 0", args: args{fromPath: "testdata/input.txt", toPath: "", offset: 0, limit: 0}, compareFile: "testdata/out_offset0_limit0.txt"},
		{name: "limit 10", args: args{fromPath: "testdata/input.txt", toPath: "", offset: 0, limit: 10}, compareFile: "testdata/out_offset0_limit10.txt"},
		{name: "limit 1000", args: args{fromPath: "testdata/input.txt", toPath: "", offset: 0, limit: 1000}, compareFile: "testdata/out_offset0_limit1000.txt"},
		{name: "limit 10000", args: args{fromPath: "testdata/input.txt", toPath: "", offset: 0, limit: 10000}, compareFile: "testdata/out_offset0_limit10000.txt"},
		{name: "offset 100 and limit 1000", args: args{fromPath: "testdata/input.txt", toPath: "", offset: 100, limit: 1000}, compareFile: "testdata/out_offset100_limit1000.txt"},
		{name: "offset 6000 and limit 1000", args: args{fromPath: "testdata/input.txt", toPath: "", offset: 6000, limit: 1000}, compareFile: "testdata/out_offset6000_limit1000.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dstFile *os.File
			var err error
			if len(tt.args.toPath) == 0 {
				dstFile, err = os.CreateTemp("", "")
				assert.NoError(t, err)
			} else {
				dstFile, err = os.OpenFile(tt.args.toPath, os.O_RDWR, 0755)
				assert.NoError(t, err)
			}
			err = Copy(tt.args.fromPath, dstFile.Name(), tt.args.offset, tt.args.limit)
			if assert.NoError(t, err) {
				dstStat, err := dstFile.Stat()
				assert.NoError(t, err)
				cmpFile, err := os.OpenFile(tt.compareFile, os.O_RDWR, 0755)
				assert.NoError(t, err)
				cmpStat, err := cmpFile.Stat()
				assert.NoError(t, err)
				assert.Equal(t, dstStat.Size(), cmpStat.Size())
			}
		})
	}
}
