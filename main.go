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
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	auth "github.com/abbot/go-http-auth"
)

var port = flag.String("p", "8100", "port to serve on")
var root_folder = flag.String("d", ".", "the root folder of files to host")

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*root_folder))
	http.Handle("/_/files/", http.StripPrefix("/_/files/", fs))

	ss := http.FileServer(http.Dir("web/assets"))
	http.Handle("/_/web/assets/", http.StripPrefix("/_/web/assets/", ss))

	htppasswd := os.Getenv("HTPASSWD_FILE")
	if htppasswd == "" {
		http.HandleFunc("/", serveTemplate)
	} else {
		log.Printf("Using '%s' http basic auth file\n", htppasswd)
		h := auth.HtpasswdFileProvider(htppasswd)
		a := auth.NewBasicAuthenticator("File-qrs Realm", h)
		http.Handle("/", a.Wrap(serveAuthTemplate))
	}

	log.Printf("Serving %s on HTTP port %s\n", *root_folder, *port)
	log.Fatal(http.ListenAndServe(":"+*port, logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
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

	lp := filepath.Join("web", "templates", "layout.html")
	fp := filepath.Join("web", "templates", "index.html")
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
		log.Printf("Resource not found\n")
		fmt.Fprint(w, "Resource not found")
	}
}
