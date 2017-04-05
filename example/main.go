package main

import (
	"github.com/Baozisoftware/GoldenDaemon"
	"flag"
	"net/http"
)

func main() {
	GoldenDaemon.RegisterTrigger("d", "-restarted")
	re := flag.Bool("restarted", false, "test")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		data := []byte("Hello,世界！")
		if *re {
			data = append(data, []byte("\n(Restarted)")...)
		}
		rw.Write(data)
	})
	http.ListenAndServe(":8080", mux)
}
