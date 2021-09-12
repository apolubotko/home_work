package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

const (
	bufferSize = 16
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var limitCounter, size int64
	fileSrc, err := os.OpenFile(fromPath, os.O_RDWR, 0755)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fileSrc.Close()
	fileDst, err := os.OpenFile(toPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fileDst.Close()

	sourceStat, err := fileSrc.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	if sourceStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit > 0 && limit <= sourceStat.Size() {
		size = limit
	} else {
		size = sourceStat.Size() - offset
	}

	bar := pb.New(int(size)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 20)
	bar.ShowPercent = true

	data := make([]byte, bufferSize)

	_, err = fileSrc.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}
	bar.Start()
	for {
		n, err := fileSrc.Read(data)

		if errors.Is(err, io.EOF) {
			break
		}
		limitCounter += int64(n)
		if limit > 0 && limitCounter > limit {
			l := limit % bufferSize
			fileDst.Write(data[:l])
			bar.Add(int(l))
			break
		}
		bar.Add(n)
		if _, err := fileDst.Write(data[:n]); err != nil {
			fmt.Printf("Write error: %v", err)
			break
		}
		time.Sleep(time.Millisecond * 10)
	}

	bar.FinishPrint("Done!")

	return nil
}
