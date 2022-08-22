package files

import "net/url"

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
	return sav.err
}

func (sav StringArrayValue) AllErrors() []error {
	return sav.errChain
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
	return uv.err
}

func (uv UrlValue) AllErrors() []error {
	return uv.errChain
}

// JsonObjectValue
// Represents value designated as JSON Object
// *should start with {}
type JsonObjectValue struct {
	product  map[interface{}]interface{}
	err      error
	errChain []error
}

func (jov JsonObjectValue) Value() map[interface{}]interface{} {
	return jov.product
}

func (jov JsonObjectValue) Error() error {
	return jov.err
}

func (jov JsonObjectValue) AllErrors() []error {
	return jov.errChain
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
	return bv.err
}

func (bv BoolValue) AllErrors() []error {
	return bv.errChain
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
	return iv.err
}

func (iv IntValue) AllErrors() []error {
	return iv.errChain
}
