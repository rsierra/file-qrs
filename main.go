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
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var port = flag.String("p", "8100", "port to serve on")
var root_folder = flag.String("d", ".", "the root folder of files to host")

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*root_folder))
	http.Handle("/_app/files/", http.StripPrefix("/_app/files/", fs))

	ss := http.FileServer(http.Dir("statics"))
	http.Handle("/_app/statics/", http.StripPrefix("/_app/statics/", ss))

	htppasswd := os.Getenv("HTPASSWD_FILE")
	if htppasswd == "" {
		http.HandleFunc("/", serveTemplate)
	} else {
		log.Printf("Using '%s' http basic auth file\n", htppasswd)
		h := auth.HtpasswdFileProvider(htppasswd)
		a := auth.NewBasicAuthenticator("File-qrs Realm", h)
		http.Handle("/", a.Wrap(serveAuthTemplate))
	}

	log.Printf("Serving %s on HTTP port: %s\n", *root_folder, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func serveAuthTemplate(w http.ResponseWriter, ar *auth.AuthenticatedRequest) {
	serveTemplate(w, &ar.Request)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	request_path := filepath.Clean(r.URL.Path)

	files, folders, err := ListFilesAndFolders(filepath.Clean(filepath.Join(*root_folder, request_path)))
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	data := struct {
		CurrentFolder string
		ParentFolder  string
		Files         []string
		Folders       []string
	}{FolderToWeb(request_path), FolderToWeb(filepath.Dir(request_path)), files, folders}

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", data)
}

func FolderToWeb(folder string) string {
	if folder == "/" {
		return folder
	} else {
		return folder + "/"
	}
}

func ListFilesAndFolders(root string) ([]string, []string, error) {
	var files []string
	var folders []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, folders, err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			folders = append(folders, file.Name())
		} else {
			files = append(files, file.Name())
		}
	}

	return files, folders, nil
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Resource not found")
	}
}
