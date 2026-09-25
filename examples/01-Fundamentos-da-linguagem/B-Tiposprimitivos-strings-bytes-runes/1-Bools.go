// booleano
package main

import "fmt"

func boolLearning() {
	const alwaysTrue = true
	var maybeFalse bool = true
	maybeTrue := true

	if maybeTrue {
		fmt.Println("it's false bro")
	} else if maybeFalse == true {
		fmt.Println("it's true bro")
	}
}
