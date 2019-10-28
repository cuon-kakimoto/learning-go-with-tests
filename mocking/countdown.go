package main

// TODO: 自動でimportが切り替わらないのはなぜだ
import (
	"io"
	"os"
	"fmt"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}
// type DefaultSleeper struct{}

type ConfigurableSleeper struct{
	duration time.Duration
  // HACK: structの中に関数も設定できるのね。
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}

// Since we are using the ConfigurableSleeper, it is now safe to delete the DefaultSleeper implementation. Wrapping up our program and having a more generic Sleeper with arbitrary long countdowns.
// func (d *DefaultSleeper) Sleep(){
// 	time.Sleep(1 * time.Second)
// }

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i--{
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main(){
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}