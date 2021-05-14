// func main() {

//     psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "password=%s dbname=%s sslmode=disable",
//     host, port, user, password, dbname)
//     db, err := sql.Open("postgres", psqlInfo)
//     if err != nil {
//     panic(err)
//   }
//   defer db.Close()
// sqlStatement := `

//     INSERT INTO users (fname, lname, age)
//     VALUES ('Raym', 'Mat', 25)
//     RETURNING id`
//     id := 0
//     err = db.QueryRow(sqlStatement).Scan(&id)
//     if err != nil {
//         panic(err)
//     }
// fmt.Println("New record ID is:", id)

// msqlStatement :=`
//     select fname, lname FROM users
//     WHERE id=1;
// `
//    // err = db.QueryRow(sqlStatement).Scan(&id)
//    //  if err != nil {
//    //      panic(err)
//    //  }

// }

package main

import (
	"fmt"
	router "go-postgres/routers/router"
	"log"
	"net/http"
	"os"
)

// const (
//   host     = "localhost"
//   port     = 5432
//   user     = "postgres"
//   password = "0663"
//   dbname   = "Demo"
// )

func main() {

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	var host string
	host = os.Getenv("DB_User")
	fmt.Printf("%s", host)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//https://levelup.gitconnected.com/building-a-simple-rest-api-in-go-ee29cf3e7334

//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//https://levelup.gitconnected.com/building-a-simple-rest-api-in-go-ee29cf3e7334





// in another folder view struct.go FileServerpackage utils

// type Response struct {
// 	Code int         `json:"code"`
// 	Body interface{} `json:"body"`
// }



