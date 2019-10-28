package main

import (
	"net/http"
	"time"
	"fmt"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string)(winner string, error error){
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error){
	select {
	// What select lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	// HACK: 最初に帰ってきたpingを勝者として扱う
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)

	// if aDuration < bDuration{
	// 	return a
	// }
	// return b
}

func ping(url string) chan struct{} {
	// Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool
	// メモリ消費量が一番小さいらしい
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}