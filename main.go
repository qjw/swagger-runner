package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const SWAGGER_UI_PATH = "/ui"
const SWAGGER_SPEC_PATH = "/spec"

//go:generate go-bindata-assetfs swaggerui/...

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, fmt.Sprintf(
		"%s/?url=%s",
		SWAGGER_UI_PATH,
		SWAGGER_SPEC_PATH,
	), http.StatusFound)
}


func openFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func spec(content string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, content)
	}
}

func main() {
	specPath := flag.String("spec", "spec.json", "spec file path.")
	port := flag.Uint("port", 12345, "bind port.")
	flag.Parse()

	fmt.Printf("start server (:%d), using spec file [%s]\n", *port, *specPath)

	s := openFile(*specPath)

	var router = mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc(SWAGGER_SPEC_PATH, spec(s))
	router.PathPrefix(SWAGGER_UI_PATH).Handler(http.StripPrefix("/ui", http.FileServer(assetFS())))

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), router)
}
