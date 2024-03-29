package service

import (
	"strconv"
)

// FizzBuzz service struct.
type FizzBuzz struct{}

// NewFizzBuzz returns an instance of FizzBuzz.
func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{}
}

// Process fizzbuzz.
func (f *FizzBuzz) Process(int1, int2, limit int, str1, str2 string) []string {
	var s string

	arr := make([]string, 0, limit)

	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			s = str1 + str2
		} else if i%int1 == 0 {
			s = str1
		} else if i%int2 == 0 {
			s = str2
		} else {
			s = strconv.Itoa(i)
		}

		arr = append(arr, s)
	}

	return arr
}
