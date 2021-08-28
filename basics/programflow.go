package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 45
	for i < 50 {
		fmt.Println("Current Number: ", i)
		i++
	}

	for j := 0; j <= 3; j++ {
		fmt.Println("Current J Value: ", j)
	}

	for {
		fmt.Println("wef", i)
		i++
		if i == 53 {
			break
		}
	}

	for n := 0; n <= 9; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	day := time.Now().Weekday()
	fmt.Println(day)
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(day)
	whatAmI(1)
	whatAmI("hey")
}
