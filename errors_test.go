package errors

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		msg string
		exp error
	}{
		{"", errors.New("")},
		{"msg", errors.New("msg")},
		{"%s", errors.New("%s")},
	}

	for i, test := range tests {
		got := New(test.msg)
		if got.Error() != test.exp.Error() {
			t.Errorf("case %d: expected %q, got %q", i, test.exp, got)
		}
	}
}

func TestErrorf(t *testing.T) {
	tests := []struct {
		format string
		args   []interface{}
		exp    error
	}{
		{"msg", nil, fmt.Errorf("msg")},
		{"some %s with code %d", []interface{}{"error", 200}, fmt.Errorf("some %s with code %d", "error", 200)},
	}

	for i, test := range tests {
		got := Errorf(test.format, test.args...)
		if got.Error() != test.exp.Error() {
			t.Errorf("case %d: expected %q, got %q", i, test.exp, got)
		}
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		err error
		msg string
		exp string
	}{
		{io.EOF, "read error", "read error: EOF"},
		{Wrap(io.EOF, "read error"), "client error", "client error: read error: EOF"},
	}

	for i, test := range tests {
		got := Wrap(test.err, test.msg)
		if got.Error() != test.exp {
			t.Errorf("case %d: expected %q, got %q", i, test.exp, got)
		}
	}
}

func TestWrapf(t *testing.T) {
	tests := []struct {
		err    error
		format string
		args   []interface{}
		exp    string
	}{
		{io.EOF, "read error", nil, "read error: EOF"},
		{io.EOF, "%s error", []interface{}{"read"}, "read error: EOF"},
		{Wrapf(io.EOF, "read error with %d", 1), "%s error", []interface{}{"client"}, "client error: read error with 1: EOF"},
	}

	for i, test := range tests {
		got := Wrapf(test.err, test.format, test.args...)
		if got.Error() != test.exp {
			t.Errorf("case %d: expected %q, got %q", i, test.exp, got)
		}
	}
}

func TestWrapNil(t *testing.T) {
	got := Wrap(nil, "no error")
	if got != nil {
		t.Errorf("expected nil, got %#v", got)
	}
}

func TestWrapfNil(t *testing.T) {
	got := Wrapf(nil, "no %s", "error")
	if got != nil {
		t.Errorf("expected nil, got %#v", got)
	}
}

func TestCause(t *testing.T) {
	err := New("error")
	tests := []struct {
		err error
		exp error
	}{
		{nil, nil},
		{io.EOF, io.EOF},
		{err, err},
		{Wrap(nil, "no error"), nil},
		{Wrap(io.EOF, "wrapper"), io.EOF},
	}

	for i, test := range tests {
		got := Cause(test.err)
		if !reflect.DeepEqual(got, test.exp) {
			t.Errorf("case %d: expected %#v, got %#v", i, test.exp, got)
		}
	}
}
