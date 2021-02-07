nogo
====

Nogo helps to compile non-go files into go binaries. It might be helpful if you want to get a kickstart while writing your own, small library. If you are looking for a well-written, feature-complete solution, please have a look at: https://github.com/markbates/pkger.

The nogo-method will only work if you are using go modules.

# 1) generate nogo.go
```bash
# install nogogen
go get github.com/rbicker/nogo/cmd/nogogen

# run nogogen to generate a a nogo file within your golang project
nogogen

# by default, nogogen will include a folder called "assets" and all of it's subfolders and -files
# if you want to include other (maybe multiple) directories, use the NOGO_DIRS env variable
NOGO_DIRS="/templates /public" nogogen
# please make sure to use absolute paths, using your project directory as root
# the command generates a file called "nogo.go" under "internal/nogo"
```

# 2) use nogo

```golang
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/rbicker/nogo-playground/internal/nogo"
)

func main() {
	// serve template
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
		}
		f, err := nogo.Get("/assets/templates/test.html")
		if err != nil {
			log.Printf("error while opening test html file: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		buf := new(strings.Builder)
		_, err = io.Copy(buf, f)
		if err != nil {
			log.Printf("error while reading test html file: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t, err := template.New("").Parse(buf.String())
		t.Execute(w, struct {
			Foo string
		}{
			Foo: "Bar",
		})
	})

	// serve static files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(nogo.Dir("/assets/public"))))
	http.ListenAndServe(":3000", nil)
}

```