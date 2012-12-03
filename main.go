package main

import (
	"fmt"
	"proudlygeek/goscii/encoder"
	"proudlygeek/mongodb/manager"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	asciiEncoder    = &encoder.Encoder{}
	mongoArtManager = &manager.MongoArtManager{Encoder: asciiEncoder}
)

func logException(res http.ResponseWriter, req *http.Request) {
	str := recover()
	fmt.Println("[ERROR]:", str)
	http.Redirect(res, req, "/", 302)
}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("views/layout.html", "views/index.html")
	t.Execute(res, nil)
}

func upload(res http.ResponseWriter, req *http.Request) {
	defer logException(res, req)

	if req.Method == "POST" {
		file, _, err := req.FormFile("pic")

		if err != nil {
			log.Panic(err)
		}
		defer file.Close()

		wr := &manager.MongoWriter{}
		image, err := asciiEncoder.DecodeImage(file)
		if err != nil {
			log.Panic(err)
		}

		err = asciiEncoder.Asciify(image, wr)
		if err != nil {
			log.Panic(err)
		}

		uri := mongoArtManager.Save(wr)
		http.Redirect(res, req, "/art/"+uri, 302)
	}
}

func show(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("views/layout.html", "views/upload.html")
	t.Execute(res, nil)
	fmt.Println("Loading Art", req.URL.Path[5:])
	art := mongoArtManager.Load(req.URL.Path[5:])
	fmt.Fprintf(res, "%s", art)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/art/", show)
	http.HandleFunc("/upload", upload)
	fmt.Println("Running on 127.0.0.1:", os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}
