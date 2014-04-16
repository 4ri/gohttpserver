// resource: http://golang.org/doc/articles/wiki/#tmp_3
// Remember to reset, not just 'Refresh'
// Problem: browser was displaying html, got header
// and added a Content-Type to give the browser a heads up.

package main
import (
    "fmt"
	"os"
	"io"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    header := w.Header()
	header.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	file, err := os.Open("input.html")
	if err != nil {
		fmt.Fprintf(w, "Cannot find input file, %s", err)
		return
	}
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Fprintf(w, "File cannot be read, %s", err)
		return
	}
	
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}