package main

import (
	"reflect"
	"testing"
	
)

func TestSum(t *testing.T) {

	checkSums := func(t testing.TB, got, want int) {


		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	checkSliceSums := func(t testing.TB, got, want []int) {

		t.Helper()

		if !reflect.DeepEqual(got, want) {

			t.Errorf("Got %v Want %v", got, want )
		}
	}

	

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		checkSums(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		checkSums(t, got, want)
	})


	t.Run("Sum multiple collections", func(t *testing.T) {

		array_1 := []int{1, 2, 3}
		array_2 := []int{9, 2}

		got := SumAll(array_1, array_2)
		want := []int{6, 11}

		
		
		checkSliceSums(t, got, want)
	})

	t.Run("Sum multiple collections, but only their tail", func(t *testing.T) {
		array_1 := []int{1, 2}
		array_2 := []int{9, 2}

		
		got := SumAllTails(array_1, array_2)
		want := []int{2, 2}

		
		checkSliceSums(t, got, want)
	})

}
