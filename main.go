package main

import (
	"fmt"
	"golang_timeformatter/kronos"
	"time"
)

func main() {

	t1, t2, err := kronos.Now().Turkish(time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println(t1, t2)

	e1, e2, err := kronos.Now().Deutsch(time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println(e1, e2)

	g1, err := kronos.Get().FindDay("Monday", "Turkish")
	if err != nil {
		panic(err)
	}
	fmt.Println(g1)

	g2, err := kronos.Get().FindMonth("October", "Deutsch")
	if err != nil {
		panic(err)
	}
	fmt.Println(g2)
}
