# conflag

[![Go Report Card](https://goreportcard.com/badge/github.com/nadoo/conflag?style=flat-square)](https://goreportcard.com/report/github.com/nadoo/conflag)
[![GitHub tag](https://img.shields.io/github/v/tag/nadoo/conflag.svg?sort=semver&style=flat-square)](https://github.com/nadoo/conflag/releases)
[![Go Document](https://img.shields.io/badge/go-document-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/nadoo/conflag )

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
example -name Jay -age 30
```
output:
```bash
  Name: Jay
  Age: 30
  Male: false
```

### Run with config file and environment variable(-config):
example.conf:
```bash
name={$NAME}
age=20
male
```
command: **use "-config" flag to specify the config file path.**
```bash
NAME=Jason example -config example.conf
```
output:
```bash
  Name: Jason
  Age: 20
  Male: true
```

### Run with config file and OVERRIDE a flag value using commandline:
example.conf:
```bash
name=Jason
age=20
male
```
command:
```bash
example -config example.conf -name Michael
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
# use {$ENV_VAR_NAME} in VALUE to get the Environment Variable value

# your name
name={$NAME}

# your age
age=20

# are you male?
male=true

# use include to include more config files
include=part1.inc.conf
include=part2.inc.conf
```
See [example.conf](example/example.conf)