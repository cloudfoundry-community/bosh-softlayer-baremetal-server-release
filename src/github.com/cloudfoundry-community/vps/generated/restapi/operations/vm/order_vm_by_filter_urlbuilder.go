package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
)

// OrderVMByFilterURL generates an URL for the order Vm by filter operation
type OrderVMByFilterURL struct {
}

// Build a url path and query string
func (o *OrderVMByFilterURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/vms/order"

	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *OrderVMByFilterURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *OrderVMByFilterURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *OrderVMByFilterURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on OrderVMByFilterURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on OrderVMByFilterURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *OrderVMByFilterURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
