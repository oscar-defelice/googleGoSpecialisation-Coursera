/*

Package errors to implement user defined errors.

Created by Oscar de Felice on 04/09/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/


// Package errors implements functions to manipulate errors.
package errors

// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
