	package main

	import (
	"net/http"
)

	func main()  {

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHomeHandler)

	//Start the server
	http.ListenAndServe("127.0.0.1:3000", nil)

}

	func  homeHandler(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello World"))
	}

	func contactHomeHandler(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Contact Page"))
	}

