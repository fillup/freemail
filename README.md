# Free email service domain list
On more than one occasion, today being one of them, I've needed the ability to test if 
a person's email address belongs to a free/public email service, ex: `gmail.com`. So 
rather than compiling yet another list and adding this exact same logic into each app 
I need it in, here's a small and simple standalone package I can use. 

## Build / test status
[![Go](https://github.com/fillup/freemail/actions/workflows/go.yml/badge.svg)](https://github.com/fillup/freemail/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/fillup/freemail)](https://goreportcard.com/report/github.com/fillup/freemail)

## Usage

```go
package main

import (
	"log"

	"github.com/fillup/freemail"
)

func main() {

	// Check by email address
	email := "joe.schmoe@gmail.com"
	isFree, err := freemail.IsFreeEmail(email)
	if err != nil {
		log.Fatal(err)
	} else if isFree {
		log.Fatal("Sorry, public email addresses are not allowed")
    }
	
	// Check by domain, including additional domains freemail is not aware of
	domain := "1234567.io"
	if freemail.IsFreeDomain(domain, "1234567.10") {
		log.Fatal("Sorry, public email addresses are not allowed")
    }
	
}

```

## Extensibility
The list of free/public email providers grows probably every day. If you find or know of any 
not already included in this package's default list, please open a PR to add it, or at least
an issue. I'll try to get things updated quickly, but in the meantime both the `IsFreeEmail`
and `IsFreeDomain` methods support providing additional domains for consideration as well.  

## License
MIT License

Copyright (c) 2024 Phillip Shipley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.