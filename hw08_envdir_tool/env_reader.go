package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var (
	errUnsetVar = errors.New("unset variable")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	var needRemove bool
	environment := make(Environment)
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		envValue, err := getEnvValue(dir, entry.Name())
		if err != nil {
			if errors.Is(err, errUnsetVar) {
				needRemove = true
			} else {
				return environment, err
			}
		}
		environment[entry.Name()] = EnvValue{Value: envValue, NeedRemove: needRemove}
		needRemove = false
	}

	return environment, nil
}

func getEnvValue(dir, file string) (string, error) {
	var buf bytes.Buffer

	filePath := dir + string(os.PathSeparator) + file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// Check the empty files if so then unset the VAR
	if len(data) < 1 {
		return "", errUnsetVar
	}

	for _, b := range data {
		if b == '\n' {
			break
		}
		if b == 0x00 {
			b = 0x0A
		}
		buf.WriteByte(b)
	}
	str := strings.TrimRight(buf.String(), " ")

	return str, nil
}
