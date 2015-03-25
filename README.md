gotemp
=========

A simple package to return current temperature in Fahrenheit from Yahoo API.

Install
-------
```Shell

$ go get github.com/jochasinga/gotemp
$ go install github.com/jochasinga/gotemp

```

Examples
--------

Get a current temperature from New York City

```Go

package main

import (
        fmt
        temp "github.com/jochasinga/gotemp"
)

func main() {
		mytemp := temp.Now("New York", "US")
}

```

