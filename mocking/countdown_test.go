package main

import (
	"bytes"
	"testing"
	"reflect"
	
)


func TestCountdown(t *testing.T) {

	// buffer := &bytes.Buffer{}
	// spySleeper := &SpySleeper{}

	// Countdown(buffer, spySleeper)

	// got := buffer.String()
	// want := "3\n2\n1\nGo!"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	// if spySleeper.Calls != 3 {

	// 	t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	// }

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func (t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)


		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
