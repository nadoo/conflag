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
