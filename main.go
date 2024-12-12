package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"stefgodin/queudo/web"
	"strconv"
)

func main() {
	web.DefaultRouter.RegisterController(&web.ChannelCtrl{})

	log.Fatal(http.ListenAndServe(":42069", nil))
}

type Peepee struct {
	Title string
	Req   *http.Request
}

func Raw(s string) template.HTML {
	return template.HTML(s)
}

func handleYourMom(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/bigdeedeen.html")
	if err != nil {
		fmt.Fprint(w, "Don work")
		return
	}

	p := &Peepee{
		Title: "Fukumuku",
		Req:   r,
	}

	tpl.Execute(w, p)
}

func handleYourDad(w http.ResponseWriter, r *http.Request) {
	isPost := r.Method == "POST"

	q := r.URL.Query()
	if !q.Has("id") {
		w.Header().Add("Location", "/")
		w.WriteHeader(302)
		return
	}

	id, err := strconv.Atoi(q.Get("id"))
	if err != nil {
		w.Header().Add("Location", "/")
		w.WriteHeader(302)
		return
	}

	if isPost {
		fmt.Print("THIS IS A FOOKIN POST")
	}

	temp := template.New("templates/bigdaddy.html")
	// TODO: DIS DON WORK - PLZ FIX CUZ U SUK
	temp.Funcs(template.FuncMap{
		"raw": Raw,
	})
	temp, err2 := temp.ParseFiles("templates/bigdaddy.html")
	if err2 != nil {
		fmt.Fprint(w, "Don work")
		return
	}

	p := &Peepee{
		Title: "<h2>Fukumuku " + strconv.Itoa(id) + "</h2>\"'",
		Req:   r,
	}

	temp.Execute(w, p)
}
