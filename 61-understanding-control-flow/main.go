package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//SEQUENCE
	fmt.Println("This is the first statment to run")
	fmt.Println("This is the srcond statment to run")
	x := 40
	y := 5
	fmt.Printf("x=%v \n y=%v \n", x, y)

	//CONDITIONAL
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to than the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || y < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 {
		fmt.Println("x is not 42")
	}

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater or equal x which is %v \n", z, x)
	} else {
		fmt.Printf("z is %v and that is less than x which is %v \n", z, x)
	}

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("This is default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("This is default case for x")
	}

	//no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is default case for x")
	}

	//no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of ALL OF THE fallthrough statments: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of ALL OF THE fallthrough statments: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of ALL OF THE fallthrough statments: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL OF THE fallthrough statments: this is default case for x")
	}

	// concurency
	//switch of channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("Value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("Value from channel 2", v2)
	}

	//LOOP / ITERATIVE
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first for loop \n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop \n", y)
		y++
	}

	// break

	for {
		fmt.Printf("y is %v \t\t third for loop \n", y)
		if y > 20 {
			break
		}
		y++
	}

	//continue

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

	// nested loop

	for i := 0; i < 5; i++ {
		fmt.Println("--")
		if i%2 != 0 {
			continue
		}
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v \n", i, j)
		}
	}

	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

}
