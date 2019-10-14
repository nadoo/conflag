package main

import (
	"fmt"

	"github.com/nadoo/conflag"
)

var conf struct {
	Name  string
	Age   int
	Male  bool
	Part1 string
	Part2 string
}

func main() {
	// get a new conflag instance
	flag := conflag.New()

	// setup flags as the standard flag package
	flag.StringVar(&conf.Name, "name", "", "your name")
	flag.IntVar(&conf.Age, "age", 0, "your age")
	flag.BoolVar(&conf.Male, "male", false, "your sex")
	flag.StringVar(&conf.Part1, "part1", "", "part1")
	flag.StringVar(&conf.Part2, "part2", "", "part2")

	// parse before access flags
	flag.Parse()

	// now you're able to get the parsed flag values
	fmt.Printf("  Name: %s\n", conf.Name)
	fmt.Printf("  Age: %d\n", conf.Age)
	fmt.Printf("  Male: %v\n", conf.Male)

	// flag values from include file
	fmt.Printf("  Part1: %s\n", conf.Part1)
	fmt.Printf("  Part2: %s\n", conf.Part2)
}
