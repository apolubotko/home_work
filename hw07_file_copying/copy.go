package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileSrc, err := os.OpenFile(fromPath, os.O_RDWR, 0755)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fileSrc.Close()

	sourceStat, err := fileSrc.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	if sourceStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	size := limit
	if limit <= 0 || limit > sourceStat.Size() {
		size = sourceStat.Size() - offset
	}

	fileDst, err := os.OpenFile(toPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fileDst.Close()

	_, err = fileSrc.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	reader := io.LimitReader(fileSrc, size)
	writer := fileDst

	// start new bar
	bar := pb.Full.Start64(size).SetRefreshRate(time.Millisecond * 20)

	// create proxy reader
	barReader := bar.NewProxyReader(reader)

	// copy from proxy reader
	io.Copy(writer, barReader)

	// finish bar
	bar.Finish()

	return nil
}
