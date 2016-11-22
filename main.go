package main

import "fmt"
import "database/sql"
import "net/http"
import _ "github.com/go-sql-driver/mysql"
import "html/template"

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "root:@/praweb")
	if err != nil {
		fmt.Println("koneksi tidak terjangkau")

	}
	return db
}

func tampil(res http.ResponseWriter, req *http.Request) {
	db := connect()
	defer db.Close()

	t, _ := template.New("t1").Parse(indeks)

	var id, nama string

	row, _ := db.Query("select * from mhs")
	for row.Next() {
		row.Scan(&id, &nama)
		data := map[string]string{
			"id":   id,
			"nama": nama,
		}
		t.Execute(res, data)
	}
}

const indeks = `<table width='100%' border='1'>
				<tr><td width='40%'>{{.id}}</td><td width='60%'>{{.nama}}</td></tr>
				</table>`

func main() {

	http.HandleFunc("/", tampil)

	fmt.Println("localhost:8080 running now....")
	http.ListenAndServe(":8080", nil)

}
