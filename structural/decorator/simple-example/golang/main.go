package main

import (
	"log"
	"math/rand"
	"time"
)

type OperateFn func()

// Decorate the operation
func Decorate(opFn OperateFn) {
	defer func(s time.Time) {
		log.Printf("elapsed time %0.2d ms", time.Since(s).Nanoseconds()/1000000)
	}(time.Now())

	// real operation function
	opFn()
}

func main() {
	// decorate log a
	Decorate(OperateFn(DoActionA))
	// decorate log b
	Decorate(OperateFn(DoActionB))
}

func DoActionA() {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	log.Println("finish action a")
}

func DoActionB() {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	log.Println("finish action b")
}
