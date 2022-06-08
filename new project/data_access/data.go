package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/go-sql-driver/mysql"
	"github.com/toqueteos/webbrowser"
	_ "github.com/toqueteos/webbrowser"
)

type student struct {
	Id   int
	Name string
}

type teacher struct {
	Id   int
	Name string
}

type course struct {
	Id         int
	Id_student int
	Id_teacher int
	Name       string
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "admin",
		Passwd:               "N9IaqrrA",
		Net:                  "tcp",
		Addr:                 "181.215.242.80:12973",
		DBName:               "School",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("connected!")

	webbrowser.Open("http://localhost:8080/")

	http.HandleFunc("/", startfunc)
	http.HandleFunc("/select", selectfunc)
	http.HandleFunc("/insert", insertfunc)
	http.HandleFunc("/command", commandfunc)
	http.HandleFunc("/commandpage", anycommandfunc)
	http.HandleFunc("/student", studentfunc)
	http.HandleFunc("/teacher", teacherfunc)
	http.HandleFunc("/course", coursefunc)
	http.HandleFunc("/istudent", insertstudentfunc)
	http.HandleFunc("/insertstudent", serveinsertstudentfunc)
	http.HandleFunc("/iteacher", insertteacherfunc)
	http.HandleFunc("/insertteacher", serveinsertteacherfunc)
	http.HandleFunc("/icourse", insertcoursefunc)
	http.HandleFunc("/insertcourse", serveinsertcoursefunc)
	http.HandleFunc("/link", linkselectfunc)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Select(commandString string) *sql.Rows {

	studentsData, err := db.Query(commandString)
	if err != nil {
		panic(err.Error())
	}
	return studentsData

}

func startfunc(w http.ResponseWriter, r *http.Request) {
	test := "\n\t"
	f, _ := os.Open("startpage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)

}

func selectfunc(w http.ResponseWriter, r *http.Request) {

	test := "\n\t"
	f, _ := os.Open("selectpage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func insertfunc(w http.ResponseWriter, r *http.Request) {
	test := "\n\t"
	f, _ := os.Open("insertpage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func studentfunc(w http.ResponseWriter, r *http.Request) {
	var cm = "select * from Student"
	var studentData student
	studentsData := Select(cm)
	var studentsDtaList []student

	for studentsData.Next() {

		if err := studentsData.Scan(&studentData.Id, &studentData.Name); err != nil {

			panic(err.Error())

		}
		studentsDtaList = append(studentsDtaList, studentData)

	}
	var t1 = template.Must(template.ParseFiles("studenttemplate.html"))
	t1.Execute(w, studentsDtaList)

}

func teacherfunc(w http.ResponseWriter, r *http.Request) {
	var cm = "select * from Teacher"
	var teacherData teacher
	teachersData := Select(cm)
	var teachersDtaList []teacher
	for teachersData.Next() {

		if err := teachersData.Scan(&teacherData.Id, &teacherData.Name); err != nil {

			panic(err.Error())

		}
		teachersDtaList = append(teachersDtaList, teacherData)

	}
	var t1 = template.Must(template.ParseFiles("teachertemplate.html"))
	t1.Execute(w, teachersDtaList)

}

func coursefunc(w http.ResponseWriter, r *http.Request) {
	var cm = "select * from Course"
	var courseData course
	coursesData := Select(cm)
	var coursesDtaList []course

	for coursesData.Next() {

		if err := coursesData.Scan(&courseData.Id, &courseData.Name, &courseData.Id_teacher, &courseData.Id_student); err != nil {

			panic(err.Error())

		}
		coursesDtaList = append(coursesDtaList, courseData)
	}
	var t1 = template.Must(template.ParseFiles("coursetemplate.html"))
	t1.Execute(w, coursesDtaList)
}

func insertstudentfunc(w http.ResponseWriter, r *http.Request) {
	test := "\n\t"
	f, _ := os.Open("insertstudentpage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func serveinsertstudentfunc(w http.ResponseWriter, r *http.Request) {
	db.Exec("insert into Student values(" + r.URL.Query()["id"][0] + "," + "'" + r.URL.Query()["name"][0] + "')")
	test := "\n\t"
	f, _ := os.Open("success.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func insertteacherfunc(w http.ResponseWriter, r *http.Request) {
	test := "\n\t"
	f, _ := os.Open("insertteacherpage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func serveinsertteacherfunc(w http.ResponseWriter, r *http.Request) {
	db.Exec("insert into Teacher values(" + r.URL.Query()["id"][0] + "," + "'" + r.URL.Query()["name"][0] + "')")
	test := "\n\t"
	f, _ := os.Open("success.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)

}

func insertcoursefunc(w http.ResponseWriter, r *http.Request) {
	test := "\n\t"
	f, _ := os.Open("insertcoursepage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func serveinsertcoursefunc(w http.ResponseWriter, r *http.Request) {
	db.Exec("insert into Course values(" + r.URL.Query()["id"][0] + "," + "'" + r.URL.Query()["name"][0] + "'" + "," + r.URL.Query()["id_teacher"][0] + "," + r.URL.Query()["id_student"][0] + ")")
	test := "\n\t"
	f, _ := os.Open("success.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

//////////////////////////////////////////////////

func commandfunc(w http.ResponseWriter, r *http.Request) {
	test := "\n\t"
	f, _ := os.Open("commandpage.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func anycommandfunc(w http.ResponseWriter, r *http.Request) {
	db.Exec(r.URL.Query()["command"][0])
	test := "\n\t"
	f, _ := os.Open("success.html")
	input := bufio.NewScanner(f)
	for input.Scan() {
		test += input.Text()
		test += "\n\t"
	}
	fmt.Fprintln(w, test)
}

func linkselectfunc(w http.ResponseWriter, r *http.Request) {
	var cm string
	if r.URL.Query()["table"][0] == "Student" {
		cm = "select * from Course where id_student = " + r.URL.Query()["id"][0]

	}
	if r.URL.Query()["table"][0] == "Teacher" {
		cm = "select * from Course where id_teacher = " + r.URL.Query()["id"][0]

	}
	var courseData course
	coursesData := Select(cm)
	var coursesDtaList []course

	for coursesData.Next() {

		if err := coursesData.Scan(&courseData.Id, &courseData.Name, &courseData.Id_teacher, &courseData.Id_student); err != nil {

			panic(err.Error())

		}
		coursesDtaList = append(coursesDtaList, courseData)
	}
	var t1 = template.Must(template.ParseFiles("coursetemplate.html"))
	t1.Execute(w, coursesDtaList)
}
