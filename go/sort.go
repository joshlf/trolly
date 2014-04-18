package trolly

import (
	"math"
	"time"
)

func SleepSort(ints []int, accuracy time.Duration) {
	sleepSort(ints, accuracy, func(i int) int { return i })
}

func SleepSortAccurate(ints []int, accuracy time.Duration) {
	sleepSort(ints, accuracy, func(i int) int { return i * i })
}

func SleepSortVeryAccurate(ints []int, accuracy time.Duration) {
	sleepSort(ints, accuracy, func(i int) int { return int(math.Pow(2, float64(i))) })
}

func sleepSort(ints []int, accuracy time.Duration, scheme func(int) int) {
	c := make(chan int)

	for _, v := range ints {
		go func(i int) {
			time.Sleep(time.Duration(i) * accuracy)
			c <- i
		}(scheme(v))
	}

	for i := range ints {
		ints[i] = <-c
	}
}
