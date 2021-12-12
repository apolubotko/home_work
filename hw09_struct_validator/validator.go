package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	tagName         = "validate"
	tagStringLen    = "len"
	tagStringRegexp = "regexp"
	tagSeparator    = ":"
	tagIntMax       = "max"
	tagIntMin       = "min"
	tagInValues     = "in"
	tagMinSize      = 2
)

var (
	errStringTooLong       = errors.New("the string is too long")
	errStringRegexpNotMuch = errors.New("the string does't contains any match of the regular expression")
	errUnknowTagName       = errors.New("unknown tag name")
	errIntMin              = errors.New("the number is too low")
	errIntMax              = errors.New("the number is too high")
	errValueNotInList      = errors.New("the value is not allowed")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var errStr strings.Builder

	for _, ve := range v {
		if ve.Err != nil {
			// errStr.WriteString(fmt.Sprintf("field %v has error: %v\n", ve.Field, ve.Err.Error()))
			errStr.WriteString(ve.Err.Error())
		}
	}
	return errStr.String()
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors

	iv := reflect.ValueOf(v)
	if iv.Kind() != reflect.Ptr {
		validationErrors = append(validationErrors, ValidationError{
			Field: iv.String(),
			Err:   fmt.Errorf("%T is not a pointer", iv),
		})
		return validationErrors
	}

	iv = iv.Elem()
	if iv.Kind() != reflect.Struct {
		validationErrors = append(validationErrors, ValidationError{
			Field: iv.String(),
			Err:   fmt.Errorf("%T is not a pointer to struct", iv),
		})
		return validationErrors
	}

	t := iv.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if len(tag) > 0 {
			switch field.Type.Kind() {
			case reflect.String:
				value := iv.Field(i).String()
				err := validateStr(value, tag)
				if err != nil {
					validationErrors = append(validationErrors, *err)
				}
			case reflect.Int:
				value := iv.Field(i).Int()
				err := validateInt(value, tag)
				if err != nil {
					validationErrors = append(validationErrors, *err)
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateStr(field, tag string) *ValidationError {
	validationError := &ValidationError{
		Field: field,
	}
	tokens := strings.Split(tag, "|")
	for _, token := range tokens {
		values := strings.Split(token, tagSeparator)
		if len(values) < tagMinSize {
			continue
		}
		key := values[0]
		val := values[1]
		switch key {
		case tagStringLen:
			valInt, err := strconv.Atoi(val)
			if err != nil {
				validationError.Err = err
			}
			if len(field) > valInt {
				validationError.Err = errStringTooLong
			}
		case tagStringRegexp:
			reg, err := regexp.Compile(val)
			if err != nil {
				validationError.Err = err
			}
			if !reg.MatchString(field) {
				validationError.Err = errStringRegexpNotMuch
			}
		case tagInValues:
			if err := processInValue(field, val); err != nil {
				validationError.Err = err
			}
		default:
			validationError.Err = errUnknowTagName
		}
	}

	if validationError.Err != nil {
		return validationError
	}

	return nil
}

func validateInt(field int64, tag string) *ValidationError {
	validationError := &ValidationError{
		Field: strconv.Itoa(int(field)),
	}
	tokens := strings.Split(tag, "|")
	for _, token := range tokens {
		values := strings.Split(token, tagSeparator)
		if len(values) < tagMinSize {
			continue
		}

		key := values[0]
		val := values[1]

		switch key {
		case tagIntMax:
			valInt, err := strconv.Atoi(val)
			if err != nil {
				validationError.Err = err
			}
			if valInt < int(field) {
				validationError.Err = errIntMax
			}

		case tagIntMin:
			valInt, err := strconv.Atoi(val)
			if err != nil {
				validationError.Err = err
			}
			if int64(valInt) > field {
				validationError.Err = errIntMin
			}

		case tagInValues:
			if err := processInValue(field, val); err != nil {
				validationError.Err = err
			}
		default:
			validationError.Err = errUnknowTagName
			return validationError
		}
	}

	if validationError.Err != nil {
		return validationError
	}

	return nil
}

func processInValue(field interface{}, str string) error {
	values := strings.Split(str, ",")

	switch val := field.(type) {
	case string:
		for _, v := range values {
			if v == val {
				return nil
			}
		}
	case int64:
		for _, v := range values {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			if int64(vInt) == val {
				return nil
			}
		}
	}

	return errValueNotInList
}
