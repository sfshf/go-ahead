package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(`素数筛选法（Eratosthenes Seive）：
例：使用eratosthens筛选法列出0~1000范围内的素数。`)
	var limit int64 = 1000
	primes, _ := sieve(limit)
	fmt.Println(primes)

}

func sieve(limit int64) (primes []int64, limitError error) {

	if limit <= 1 {
		return nil, errors.New("main: invalid upper limit.")
	}

	var bs [1000000]bool
	primes = make([]int64, 0, limit)
	bs[0] = true
	bs[1] = true
	for i := int64(0); i <= limit; i++ {
		if !bs[i] {
			for j := i * i; j <= limit; j += i {
				bs[j] = true
			}
			primes = append(primes, i)
		}
	}
	return primes, nil

}
