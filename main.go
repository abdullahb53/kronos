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

}
