package main

import "fmt"

// Note this isn't concurrent safe!
type Singleton struct {
	// ... other fields here
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance != nil {
		return instance
	}

	instance = new(Singleton)
	return instance
}

// GO111MODULE="off" go run singleton.go
func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("Singleton works, both variables contain the same instance.")
	} else {
		panic("Singleton instances should be the same")
	}
}
