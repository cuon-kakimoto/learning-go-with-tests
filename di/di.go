package main

import (
	"fmt"
	"io"
	"net/http"
)

// DI will motivate you to inject in a database dependency (via an interface)
func Greet(writer io.Writer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}

// HACK: w http.ResponseWriterがio.Writerのinterfaceをもっているのか。
func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
	Greet(w, "world")
}
func main(){
	http.ListenAndServe(":5050", http.HandlerFunc(MyGreeterHandler))
}