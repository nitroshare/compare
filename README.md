## compare

[![Build Status](https://github.com/nitroshare/compare/actions/workflows/test.yml/badge.svg)](https://github.com/nitroshare/compare/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/nitroshare/compare/badge.svg?branch=main)](https://coveralls.io/github/nitroshare/compare?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/nitroshare/compare.svg)](https://pkg.go.dev/github.com/nitroshare/compare)
[![MIT License](https://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](https://opensource.org/licenses/MIT)

This package provides a simple utility function for comparing values in a test. The goal is extreme simplicity:

```golang
import (
    "testing"

    "github.com/nitroshare/compare"
)

func TestSomething(t *testing.T) {

    // Second and third parameter are compared and the fourth parameter
    // determines if they should match; in this case, they should...
    compare.Compare(t, 1, 1, true)

    // ...and in this case they should not
    compare.Compare(t, 1, 2, false)

    // Functions are a special exception - they cannot be directly compared
    // with Compare, so a separate function is provided that will work in
    // most circumstances:
    compare.CompareFn(t, testing.Init, testing.Init, true)
}
```
