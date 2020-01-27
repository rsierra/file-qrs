/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var port = flag.String("p", "8100", "port to serve on")
var directory = flag.String("d", ".", "the directory of files to host")

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*directory))
	http.Handle("/files/", http.StripPrefix("/files/", fs))

	ss := http.FileServer(http.Dir("statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", ss))

	http.HandleFunc("/", serveTemplate)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	var template_name string
	if template_name = filepath.Clean(r.URL.Path); template_name == "/" {
		template_name = "/index.html"
	}

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", template_name)

	files, err := IOReadDir(*directory)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Template %s\n", template_name)
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", files)
}

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
		}
	return files, nil
}
