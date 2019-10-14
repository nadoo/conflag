/*
Package conflag is a drop-in replacement for Go's standard flag package with config file support.

Usage

Create a conflag instance and then you can use it like the standard flag package.

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

Config File

Just use the command line flag name as key name:

	## config file
	# comment line starts with "#"

	# format:
	#KEY=VALUE,
	# just use the command line flag name as the key name

	# your name
	name=Jason

	# your age
	age=20

	# are you male?
	male=true

	# use include to include more config files
	include=part1.inc.conf
	include=part2.inc.conf
*/
package conflag
