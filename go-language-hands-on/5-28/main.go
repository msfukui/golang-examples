package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// open SQL and return DB
	setDB := func() *sql.DB {
		con, er := sql.Open("sqlite3", "data.sqlite3")
		if er != nil {
			log.Fatal(er)
		}
		return con
	}

	// get Web data function.
	wf := func() {
		fstr := "https://www.biglobe.ne.jp" // crowering sample text
		if !strings.HasPrefix(fstr, "http") {
			fstr = "http://" + fstr
		}
		dc, er := goquery.NewDocument(fstr)
		if er != nil {
			return
		}
		ttl := dc.Find("title")
		html, er := dc.Html()
		if er != nil {
			return
		}
		cvtr := md.NewConverter("", true, nil)
		mkdn, er := cvtr.ConvertString(html)
		if er != nil {
			return
		}
	}
	ff := func() {
		var qry string = "select * from md_data where title like ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs, er := con.Query(qry, "%"+"https://www.biglobe.ne.jp"+"%")
		if er != nil {
			return
		}
		res := ""
		for rs.Next() {
			var ID int
			var TT string
			var UR string
			var MR string
			er := rs.Scan(&ID, &TT, &UR, &MR)
			if er != nil {
				return
			}
			res += strconv.Itoa(ID) + ":" + TT + "\n"
		}
	}

	// find by id function.
	idf := func(id int) {
		var qry string = "select * from md_data where id = ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs := con.QueryRow(qry, id)
		var ID int
		var TT string
		var UR string
		var MR string
		rs.Scan(&ID, &TT, &UR, &MR)
	}

	// save function.
	sf := func() {
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		qry := "insert into md_data (title, url, markdown) values (?, ?, ?)"
		_, er := con.Exec(qry, "BIGLOBE", "https://www.biglobe.ne.jp", "sample")
	}

	// Export data function.
	xf := func() {
		fn := "BIGLOBE" + ".md"
		ctt := "# " + "BIGLOBE" + "\n\n"
		ctt += "## " + "BIGLOBE" + "\n\n"
		ctt += "sample"
		er := ioutil.WriteFile(fn, []byte(ctt), os.ModePerm)
		if er != nil {
			return
		}
	}
}
