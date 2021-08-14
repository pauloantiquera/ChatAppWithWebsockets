package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Serving at :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's chat!
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAnServe:", err)
	}
}
