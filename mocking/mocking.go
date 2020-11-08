package main
import (
	"io"
	"fmt"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}


var countdownStart = 3

type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(time.Duration)
}

func (c * ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0 ; i -- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, "Go!")
}


func main() {
    sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}


