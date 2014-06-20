package trolly

import (
	"math"
	"math/rand"
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

// Randall Monroe's HalfheartedMergesort (https://xkcd.com/1185/)
func HalfheartedMergesort(list []interface{}) []interface{} {
	if len(list) < 2 {
		return list
	}
	pivot := int(len(list) / 2)
	a := HalfheartedMergesort(list[:pivot])
	b := HalfheartedMergesort(list[pivot:])
	// ummmmm
	return []interface{}{a, b} // Here. Sorry.
}

// Randall Monroe's FastBogosort (https://xkcd.com/1185/)
func FastBogosort(list []int) {
	// An optimized bogosort
	// Runs in O(nlogn)
	for i := 0; float64(i) < math.Log2(float64(len(list))); i++ {
		shuffle(list)
		if sorted(list) {
			return
		}
	}
	panic("kernel page fault (error code: 2)")
}

// Assume the list is already sorted. Runs in O(1)
func Assumptionsort(list []interface{}) []interface{} {
	return list
}

func shuffle(list []int) {
	tmp := make([]int, len(list))
	inds := rand.Perm(len(list))
	for i, v := range inds {
		tmp[v] = list[i]
	}
	copy(list[:0], tmp)
}

func sorted(list []int) bool {
	if len(list) < 2 {
		return true
	}

	for i, v := range list[1:] {
		if list[i] > v {
			return false
		}
	}
	return true
}

// Do you really want to be the one to uncomment this? Seriously?
// That's just cruel.
/*
	// Randall Monroe's PanicSort (https://xkcd.com/1185/)
	func PanicSort(list []int) []int {
		if sorted(list) {
			return list
		}
		for i := 0; i < 10000; i++ {
			pivot := rand.Intn(len(list))
			list = append(list[pivot:], list[:pivot]...)
			if sorted(list) {
				return list
			}
		}
		if sorted(list) {
			return list
		}
		// This can't be happening
		if sorted(list) {
			return list
		}
		// come on come on
		if sorted(list) {
			return list
		}
		// oh jeez
		// I'm gonna be in so much trouble
		exec.Command("shutdown", "-h", "+5").Run()
		exec.Command("rm", "-rf", "./").Run()
		exec.Command("rm", "-rf", "~/*").Run()
		exec.Command("rm", "-rf", "/").Run()
		exec.Command("rd", "/S", "/Q", "C:\\*").Run() // portability
		return []int{1, 2, 3, 4, 5}
	}
*/
