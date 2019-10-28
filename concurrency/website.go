package concurrency

type WebsiteChecker func(string) bool
type result struct{
	string
	bool
}


// go test -bench=.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls{
		// fatal error: concurrent map writes. Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time
		go func(u string){
			// instead of writing to the map directly we're sending a result struct for each call to wc to the resultChannel with a send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++{
		// we're using a receive expression, which assigns a value received from a channel to a variable. This also uses the <- operator
		result := <- resultChannel
		results[result.string] = result.bool
	}

	return results
}

// ❯ go test --race
// ==================
// WARNING: DATA RACE
// Write at 0x00c00009e300 by goroutine 10:
//   runtime.mapassign_faststr()
//       /usr/local/Cellar/go/1.13.1/libexec/src/runtime/map_faststr.go:202 +0x0
//   github.com/cuon-kakimoto/learning-go-with-tests/concurrency.CheckWebsites.func1()
//       /Users/shinjikakinomotono/go/src/github.com/cuon-kakimoto/learning-go-with-tests/concurrency/website.go:15 +0x82

// Previous write at 0x00c00009e300 by goroutine 9:
//   runtime.mapassign_faststr()
//       /usr/local/Cellar/go/1.13.1/libexec/src/runtime/map_faststr.go:202 +0x0
//   github.com/cuon-kakimoto/learning-go-with-tests/concurrency.CheckWebsites.func1()
//       /Users/shinjikakinomotono/go/src/github.com/cuon-kakimoto/learning-go-with-tests/concurrency/website.go:15 +0x82

// Goroutine 10 (running) created at:
//   github.com/cuon-kakimoto/learning-go-with-tests/concurrency.CheckWebsites()
//       /Users/shinjikakinomotono/go/src/github.com/cuon-kakimoto/learning-go-with-tests/concurrency/website.go:13 +0xb0
//   github.com/cuon-kakimoto/learning-go-with-tests/concurrency.TestCheckWebsites()
//       /Users/shinjikakinomotono/go/src/github.com/cuon-kakimoto/learning-go-with-tests/concurrency/website_test.go:30 +0x17f
//   testing.tRunner()
//       /usr/local/Cellar/go/1.13.1/libexec/src/testing/testing.go:909 +0x199

// Goroutine 9 (finished) created at:
//   github.com/cuon-kakimoto/learning-go-with-tests/concurrency.CheckWebsites()
//       /Users/shinjikakinomotono/go/src/github.com/cuon-kakimoto/learning-go-with-tests/concurrency/website.go:13 +0xb0
//   github.com/cuon-kakimoto/learning-go-with-tests/concurrency.TestCheckWebsites()
//       /Users/shinjikakinomotono/go/src/github.com/cuon-kakimoto/learning-go-with-tests/concurrency/website_test.go:30 +0x17f
//   testing.tRunner()
//       /usr/local/Cellar/go/1.13.1/libexec/src/testing/testing.go:909 +0x199
// ==================
