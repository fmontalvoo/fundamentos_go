package main

import "fmt"

// Break y Continue
func main() {

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}

		if i >= 8 {
			break
		}
		
		fmt.Println(i)
	}

}