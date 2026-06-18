package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main1() {

	router := mux.NewRouter()

	router.HandleFunc("/", basefunc)
	router.HandleFunc("/test", basefunc1)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal("Err")
	}

}

func basefunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Accept-Type", "html/text")
	w.Write([]byte(`
	<!DOCTYPE html>
<html>
<body>

<h2>HTML Forms</h2>

<form action="/test">
  <label for="fname">First name:</label><br>
  <input type="text" id="fname" name="fname" value="John"><br>
  <label for="lname">Last name:</label><br>
  <input type="text" id="lname" name="lname" value="Doe"><br><br>
  <input type="submit" value="Submit">
</form> 

<p>If you click the "Submit" button, the form-data will be sent to a page called "/action_page.php".</p>

</body>
</html>


	`))
}

func basefunc1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`Bss kr bhai`))
}
