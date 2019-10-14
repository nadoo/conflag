# conflag

[![Go Report Card](https://goreportcard.com/badge/github.com/nadoo/conflag?style=flat-square)](https://goreportcard.com/report/github.com/nadoo/conflag)
[![GitHub tag](https://img.shields.io/github/tag/nadoo/conflag.svg?style=flat-square)](https://github.com/nadoo/conflag/releases)
[![GoDoc](https://img.shields.io/badge/go-doc-blue.svg?style=flat-square)](http://godoc.org/github.com/nadoo/conflag)

conflag is a drop-in replacement for Go's standard flag package with config file support.

## Usage

### Your code:
```Go
package main

import (
	"fmt"

	"github.com/nadoo/conflag"
)

var conf struct {
	Name string
	Age  int
	Male bool
}

func main() {
	// get a new conflag instance
	flag := conflag.New()

	// setup flags as the standard flag package
	flag.StringVar(&conf.Name, "name", "", "your name")
	flag.IntVar(&conf.Age, "age", 0, "your age")
	flag.BoolVar(&conf.Male, "male", false, "your sex")

	// parse before access flags
	flag.Parse()

	// now you're able to get the parsed flag values
	fmt.Printf("  Name: %s\n", conf.Name)
	fmt.Printf("  Age: %d\n", conf.Age)
	fmt.Printf("  Male: %v\n", conf.Male)
}
```

### Run without config file:
command:
```bash
sample -name Jay -age 30
```
output:
```bash
  Name: Jay
  Age: 30
  Male: false
```

### Run with config file(-config):
sample.conf:
```bash
name=Jason
age=20
male
```
command: **use "-config" flag to specify the config file path.**
```bash
sample -config sample.conf
```
output:
```bash
  Name: Jason
  Age: 20
  Male: true
```

### Run with config file and OVERRIDE a flag value using commandline:
sample.conf:
```bash
name=Jason
age=20
male
```
command:
```bash
sample -config sample.conf -name Michael
```
output:
```bash
  Name: Michael
  Age: 20
  Male: true
```

## Config File
- format: KEY=VALUE

**just use the command line flag name as key name**:

```bash
## config file
# comment line starts with "#"

# format:
#KEY=VALUE, 
# just use the command line flag name as key name

# your name
name=Jason

# your age
age=20

# are you male?
male
```
See [simple.conf](examples/simple/simple.conf)