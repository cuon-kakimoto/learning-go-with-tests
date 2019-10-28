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

func Countdown(out io.Writer){
	for i := countdownStart; i >0; i--{
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main(){
	Countdown(os.Stdout)
}