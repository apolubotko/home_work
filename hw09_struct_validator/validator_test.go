package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: &User{
				ID:     "cesh7Luu4GeePie3okae",
				Name:   "Andrey",
				Age:    40,
				Email:  "testg@gmail.com",
				Role:   "admin",
				Phones: []string{"79262000011", "79262000012"},
				meta:   []byte(`{"name":"Andrey","gender":"male"}`),
			},
			expectedErr: nil,
		},
		{
			in: &User{
				ID:     "cesh7Luu4GeePie3okaecesh7Luu4GeePie3okae",
				Name:   "Andrey",
				Age:    40,
				Email:  "test@gmail.com",
				Role:   "admin",
				Phones: []string{"79262000011", "79262000012"},
				meta:   []byte(`{"name":"Andrey","gender":"male"}`),
			},
			expectedErr: ValidationErrors{{Err: errStringTooLong}},
		},
		{
			in: &User{
				ID:     "cesh7Luu4GeePie3okae",
				Name:   "Andrey",
				Age:    16,
				Email:  "test@gmail.com",
				Role:   "admin",
				Phones: []string{"79262000011", "79262000012"},
				meta:   []byte(`{"name":"Andrey","gender":"male"}`),
			},
			expectedErr: ValidationErrors{{Err: errIntMin}},
		},
		{
			in: &User{
				ID:     "cesh7Luu4GeePie3okae",
				Name:   "Andrey",
				Age:    60,
				Email:  "test@gmail.com",
				Role:   "admin",
				Phones: []string{"79262000011", "79262000012"},
				meta:   []byte(`{"name":"Andrey","gender":"male"}`),
			},
			expectedErr: ValidationErrors{{Err: errIntMax}},
		},
		{
			in: &User{
				ID:     "cesh7Luu4GeePie3okae",
				Name:   "Andrey",
				Age:    40,
				Email:  "test@gmail",
				Role:   "admin",
				Phones: []string{"79262000011", "79262000012"},
				meta:   []byte(`{"name":"Andrey","gender":"male"}`),
			},
			expectedErr: ValidationErrors{{Err: errStringRegexpNotMuch}},
		},
		{
			in: &User{
				ID:     "cesh7Luu4GeePie3okae",
				Name:   "Andrey",
				Age:    40,
				Email:  "test@gmail.com",
				Role:   "Engineer",
				Phones: []string{"79262000011", "79262000012"},
				meta:   []byte(`{"name":"Andrey","gender":"male"}`),
			},
			expectedErr: ValidationErrors{{Err: errValueNotInList}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			_ = tt
			t.Parallel()
			err := Validate(tt.in)
			if tt.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
