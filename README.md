# GoldenDaemon
Run golang app as background program.

# Install
> go get github.com/Baozisoftware/GoldenDaemon

# Example
```go
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
```
**Precautions**
1. Avoid using loop output or get input code, otherwise it will make CPU occupancy high;
2. For triggering parameters, do not need to add "-"

# Screenshots
> example.exe

![normal](https://github.com/Baozisoftware/GoldenDaemon/blob/master/example/screenshots/normal.png)
> example.exe -d

![restarted](https://github.com/Baozisoftware/GoldenDaemon/blob/master/example/screenshots/restarted.png)