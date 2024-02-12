package main

import "fmt"

func main() {
	var valid = make([]string, 1)
	valid[0] = "4533"
	valid[1] = "1234"
	fmt.Println(checkPin(valid, "4#5534"))
}

func checkPin(validPins []string, rawPin string) bool {
	var shortPin = shortenPin(rawPin)

	for i := 0; i < len(validPins); i++ {
		if validPins[i] == shortPin {
			return true
		}
	}

	return false
}

func shortenPin(pin string) string {

	var returnVal = make([]byte, len(pin))
	for i := 0; i < len(pin); i++ {
		if pin[i] == '#' {
			returnVal = returnVal[:len(returnVal)-2]
		} else {
			returnVal = append(returnVal, pin[i])
		}
	}

	return string(returnVal)
}
