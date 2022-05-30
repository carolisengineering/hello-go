package main

import "reflect"
import "testing"


func TestSum(t *testing.T) {

	checkSum := func(t testing.TB, collection []int, want int) {
		got := Sum(collection)
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, collection)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		checkSum(t, numbers, 15)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		checkSum(t, numbers, 6)
	})
	
}

func TestSumAllTails(t *testing.T) {

	t.Run("calculate the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0,9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely calculate the sum of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}