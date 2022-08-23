package files

import (
	"fmt"
	"net/url"
	"strings"
)

// StringArrayValue
// Represents value of array of strings
//
type StringArrayValue struct {
	product  []string
	err      error
	errChain []error
}

func (sav StringArrayValue) Value() []string {
	return sav.product
}

func (sav StringArrayValue) Error() error {
	if len(sav.errChain) > 0 || sav.err != nil {
		errStr := ""
		if sav.err != nil {
			sav.errChain = append(sav.errChain, sav.err)
		}
		for index, err := range sav.errChain {
			errStr += fmt.Errorf("[%d] %w ", index, err).Error()
		}

		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (sav StringArrayValue) AsError(errStr string) error {
	if len(sav.errChain) > 0 || sav.err != nil {
		if sav.err != nil {
			sav.errChain = append(sav.errChain, sav.err)
		}
		for _, err := range sav.errChain {
			if strings.Contains(err.Error(), errStr) {
				return err
			}
		}
	}
	return nil
}

// UrlValue
// Represents value designated as URL
//
type UrlValue struct {
	product  *url.URL
	err      error
	errChain []error
}

func (uv UrlValue) Value() *url.URL {
	return uv.product
}

func (uv UrlValue) Error() error {
	if len(uv.errChain) > 0 || uv.err != nil {
		errStr := ""
		if uv.err != nil {
			uv.errChain = append(uv.errChain, uv.err)
		}
		for index, err := range uv.errChain {
			errStr += fmt.Errorf("[%d] %w ", index, err).Error()
		}

		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (uv UrlValue) AsError(errStr string) error {
	if len(uv.errChain) > 0 || uv.err != nil {
		if uv.err != nil {
			uv.errChain = append(uv.errChain, uv.err)
		}
		for _, err := range uv.errChain {
			if strings.Contains(err.Error(), errStr) {
				return err
			}
		}
	}
	return nil
}

// JsonObjectValue
// Represents value designated as JSON Object
// *should start with {}
type JsonObjectValue struct {
	product  map[string]interface{}
	err      error
	errChain []error
}

func (jov JsonObjectValue) Value() map[string]interface{} {
	return jov.product
}

func (jov JsonObjectValue) Error() error {
	if len(jov.errChain) > 0 || jov.err != nil {
		errStr := ""
		if jov.err != nil {
			jov.errChain = append(jov.errChain, jov.err)
		}
		for index, err := range jov.errChain {
			errStr += fmt.Errorf("[%d] %w ", index, err).Error()
		}

		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (jov JsonObjectValue) AsError(errStr string) error {
	if len(jov.errChain) > 0 || jov.err != nil {
		if jov.err != nil {
			jov.errChain = append(jov.errChain, jov.err)
		}
		for _, err := range jov.errChain {
			if strings.Contains(err.Error(), errStr) {
				return err
			}
		}
	}
	return nil
}

// BoolValue
// Represents boolean value
//
type BoolValue struct {
	product  bool
	err      error
	errChain []error
}

func (bv BoolValue) Value() bool {
	return bv.product
}

func (bv BoolValue) Error() error {
	if len(bv.errChain) > 0 || bv.err != nil {
		errStr := ""
		if bv.err != nil {
			bv.errChain = append(bv.errChain, bv.err)
		}
		for index, err := range bv.errChain {
			errStr += fmt.Errorf("[%d] %w ", index, err).Error()
		}

		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (bv BoolValue) AsError(errStr string) error {
	if len(bv.errChain) > 0 || bv.err != nil {
		if bv.err != nil {
			bv.errChain = append(bv.errChain, bv.err)
		}
		for _, err := range bv.errChain {
			if strings.Contains(err.Error(), errStr) {
				return err
			}
		}
	}
	return nil
}

// IntValue
// Represents int64 value
//
type IntValue struct {
	product  int64
	err      error
	errChain []error
}

func (iv IntValue) Value() int64 {
	return iv.product
}

func (iv IntValue) Error() error {
	if len(iv.errChain) > 0 || iv.err != nil {
		errStr := ""
		if iv.err != nil {
			iv.errChain = append(iv.errChain, iv.err)
		}
		for index, err := range iv.errChain {
			errStr += fmt.Errorf("[%d] %w ", index, err).Error()
		}

		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (iv IntValue) AsError(errStr string) error {
	if len(iv.errChain) > 0 || iv.err != nil {
		if iv.err != nil {
			iv.errChain = append(iv.errChain, iv.err)
		}
		for _, err := range iv.errChain {
			if strings.Contains(err.Error(), errStr) {
				return err
			}
		}
	}
	return nil
}
