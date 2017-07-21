# conflag
conflag is a config file and command line parser based on Go's standard flag package.

## Usage

### Your code:
```Go
package main

import (
	"fmt"

	"github.com/nadoo/conflag"
)

var conf struct {
	Name  string
	Age   int
	IsMan bool
}

func main() {

	// get a new conflag instance
	flag := conflag.New()

	// setup flags as the standard flag package
	flag.StringVar(&conf.Name, "name", "", "your name")
	flag.IntVar(&conf.Age, "age", 0, "your age")
	flag.BoolVar(&conf.IsMan, "isman", false, "are you man")

	// parse before access flags
	flag.Parse()

	// now you're able to get the parsed flag values
	fmt.Printf("  Name: %s\n", conf.Name)
	fmt.Printf("  Age: %d\n", conf.Age)
	fmt.Printf("  IsMan: %v\n", conf.IsMan)

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
  IsMan: false
```

### Run with config file(-config):
sample.conf:
```bash
name=Jason
age=20
isman
```
command: **use "-config" flag to specify the config file path.**
```bash
sample -config sample.conf
```
output:
```bash
  Name: Jason
  Age: 20
  IsMan: true
```

### Run with config file and OVERRIDE a flag value using commandline:
sample.conf:
```bash
name=Jason
age=20
isman
```
command:
```bash
sample -config sample.conf -name Michael
```
output:
```bash
  Name: Michael
  Age: 20
  IsMan: true
```