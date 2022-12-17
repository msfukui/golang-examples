package main

import (
	"database/sql"
	"fmt"
	"golang-example/go-language-hands-on/src/hello"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Mydata is json structure.
type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

// Str get string value.
func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail +
		", " + strconv.Itoa(m.Age) + ">"
}

func main() {
	con, er := sql.Open("sqlite3", "../data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	ids := hello.Input("update ID")
	id, _ := strconv.Atoi(ids)
	qry := "select * from mydata where id = ?"
	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)
	ae := strconv.Itoa(tgt.Age)
	nm := hello.Input("name(" + tgt.Name + ")")
	ml := hello.Input("mail(" + tgt.Mail + ")")
	ge := hello.Input("age(" + ae + ")")
	ag, _ := strconv.Atoi(ge)

	if nm == "" {
		nm = tgt.Name
	}
	if ml == "" {
		ml = tgt.Mail
	}
	if ge == "" {
		ag = tgt.Age
	}

	qry = "update mydata set name = ?, mail = ?, age = ? where id = ?"
	con.Exec(qry, nm, ml, ag, id)
	showRecord(con)
}

// print all record.
func showRecord(con *sql.DB) {
	qry := "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

// get Mydata from Rows.
func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

// get Mydata from Row.
func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
