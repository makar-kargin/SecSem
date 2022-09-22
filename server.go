package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "testDB.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	status_map := map[int]string{0: "AFK", 1: "Active"}

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	}

	usersHandler := func(w http.ResponseWriter, r *http.Request) {
		var (
			id          int
			login       string
			moneyAmount int
			cardNumber  string
			status      int
		)

		rows, err := db.Query("select * from users where status = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		isAny := false
		for rows.Next() {
			isAny = true
			err := rows.Scan(&id, &login, &moneyAmount, &cardNumber, &status)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(w,
				"Active user with login "+login+", "+
					"id: "+strconv.Itoa(id)+", "+
					"money amount: "+strconv.Itoa(moneyAmount)+", "+
					"cardNumber: "+cardNumber+", "+
					"\n")
		}
		if !isAny {
			http.ServeFile(w, r, "user_not_found.html")
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		var (
			id          int
			login       string
			moneyAmount int
			cardNumber  string
			status      int
		)

		query := r.URL.Query()
		value, isIn := query["login"]
		if !isIn {
			http.ServeFile(w, r, "home.html")
		}

		rows, err := db.Query("select * from users where login = ?", value[0])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		isAny := false
		for rows.Next() {
			isAny = true
			err := rows.Scan(&id, &login, &moneyAmount, &cardNumber, &status)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(w,
				"User with login "+login+", "+
					"id: "+strconv.Itoa(id)+", "+
					"money amount: "+strconv.Itoa(moneyAmount)+", "+
					"cardNumber: "+cardNumber+", "+
					"status: "+status_map[status]+
					"\n")
		}
		if !isAny {
			http.ServeFile(w, r, "user_not_found.html")
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	idHandler := func(w http.ResponseWriter, r *http.Request) {
		var (
			id          int
			login       string
			moneyAmount int
			cardNumber  string
			status      int
		)

		query := r.URL.Query()
		value, isIn := query["id"]
		if !isIn {
			http.ServeFile(w, r, "home.html")
		}

		rows, err := db.Query("select * from users where id = ?", value[0])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		isAny := false
		for rows.Next() {
			isAny = true
			err := rows.Scan(&id, &login, &moneyAmount, &cardNumber, &status)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(w,
				"User with id "+strconv.Itoa(id)+", "+
					"login: "+login+", "+
					"money amount: "+strconv.Itoa(moneyAmount)+", "+
					"cardNumber: "+cardNumber+", "+
					"status: "+status_map[status]+
					"\n")
		}
		if !isAny {
			http.ServeFile(w, r, "user_not_found.html")
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/by-login", loginHandler)
	http.HandleFunc("/by-id", idHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
