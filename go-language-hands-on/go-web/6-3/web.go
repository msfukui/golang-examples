package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// Temps is template structure.
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	return tmps
}

// setup template function.
func setupTemp() *Temps {
	temps := new(Temps)
	temps.notemp = notemp()
	// set index template.
	indx, er := template.ParseFiles("templates/index.html")
	if er != nil {
		indx = temps.notemp
	}
	temps.indx = indx
	//set hello template.
	helo, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo
	return temps
}

// index handler.
func index(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "index",
		Title:    "Index",
		Message:  "This is Top page.",
	}
	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler.
func hello(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	id1 := rq.FormValue("id")
	nm1 := rq.FormValue("name")
	msg1 := "id: " + id1 + ", Name: " + nm1

	msg2 := "type name and password:"
	if rq.Method == "POST" {
		nm2 := rq.PostFormValue("name")
		pw2 := rq.PostFormValue("pass")
		msg2 = "name: " + nm2 + ", password: " + pw2
	}

	msg3 := "login name & password:"

	ses, _ := cs.Get(rq, "hello-session")

	if rq.Method == "POST" {
		ses.Values["login"] = nil
		ses.Values["name"] = nil
		nm3 := rq.PostFormValue("name")
		pw3 := rq.PostFormValue("pass")
		if nm3 == pw3 {
			ses.Values["login"] = true
			ses.Values["name"] = nm3
		}
		ses.Save(rq, w)
	}

	flg, _ := ses.Values["login"].(bool)
	lname, _ := ses.Values["name"].(string)
	if flg {
		msg3 = "logined: " + lname
	}

	item := struct {
		Title    string
		Message1 string
		Message2 string
		Message3 string
	}{
		Title:    "Send values",
		Message1: msg1,
		Message2: msg2,
		Message3: msg3,
	}
	er := tmp.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func hello2(w http.ResponseWriter, rq *http.Request) {
	data := []string{
		"One", "Two", "Three",
	}
	item := struct {
		Title string
		Data  []string
	}{
		Title: "Hello",
		Data:  data,
	}
	er := page("hello2").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	temps := setupTemp()
	// index handling.
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq, temps.indx)
	})
	// hello handling.
	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq, temps.helo)
	})
	// hello2 handling.
	http.HandleFunc("/hello2", func(w http.ResponseWriter, rq *http.Request) {
		hello2(w, rq)
	})
	if err := http.ListenAndServe("", nil); err != nil {
		log.Fatal(err)
	}
}
