package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func Solution(A []int) int {
	// write your code in Go 1.4
	return 1
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	// var ctx context.Context
	// context.TODO()
	// context.WithCancel()
	// mapKeep:= make(map[int]bool)

	// me, her := mapKeep[5]

	if true {

	} else if false {

	}
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"

	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Using port number: ", PORT)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
