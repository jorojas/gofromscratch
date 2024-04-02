package webserver

import "net/http"

func MyWebServer() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":3255", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
