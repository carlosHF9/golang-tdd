package  main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return 
}

type Sleeper interface {
	Sleep()
}



type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {}


func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}