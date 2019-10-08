package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	maxRetries := 2
	retries := 0
	var err error
	for retries <= maxRetries {
		sleepTime := calcSleepTime(retries) * time.Second
		fmt.Println("Let's sleep for", sleepTime)
		time.Sleep(sleepTime)
		result := try()
		if result {
			fmt.Println("Try success!")
			err = nil
			break
		} else {
			fmt.Println("Try failure!")
			err = errors.New("Tried but failed")
			retries++
		}
	}
	if err != nil {
		panic("Procedure FAILED ^_^;")
	}
	fmt.Println("Procedure SUCCEEDED ^O^")
}

func calcSleepTime(retries int) time.Duration {
	if retries == 0 {
		return time.Duration(0)
	}
	retries--
	return time.Duration(math.Pow(2, float64(retries)))
}

func try() bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(4) < 3 {
		return false
	}
	return true
}
