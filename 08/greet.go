package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet greetings a person
// io.Writer is an interface implemented by bytes.Buffer
// func Greet(writer *bytes.Buffer, name string)
// now the function is more generic
func Greet(writer io.Writer, name string) {

	// Fprintf takes a Writer to send the string to
	// Printf defaults it to stdout
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler handle a http request
// http.ResponseWriter is an interface that implements io.Writer
// (that's why it doesn't need the `*`)
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	PORT := ":5000"
	fmt.Printf("Server running at: http://localhost%s\n", PORT)
	http.ListenAndServe(PORT, http.HandlerFunc(MyGreeterHandler))
}
