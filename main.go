package main

import (
	"fmt"
	"golang_timeformatter/kronos"
	"time"
)

func main() {

	t1, t2 := kronos.Now().Turkish(time.Now())

	fmt.Println(t1, t2)

	e1, e2 := kronos.Now().English(time.Now())

	fmt.Println(e1, e2)

}
