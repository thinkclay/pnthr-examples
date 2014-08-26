package main

import (
  "github.com/pnthr/pnthr"
	"fmt"
  "log"
  "os"
  "net/http"
)

func main() {
	/**
	 * POST /
	 *
	 * All requests will come through the root, via post
	 * Each request should have an app id and payload that has been encrypted with the app secret
	 * We want to take this payload, decrypt it with the app secret
	 * Once we have the raw payload, we encrypt first with the app password
	 * Secondly with the encrypt with the app secret, for transport back to the requestor
	 */
	http.HandleFunc("/", root)

	log.Println("Listening for connections...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
  /**
   * In the open command window set the following for Heroku:
   * heroku config:set MONGOHQ_URL=mongodb://user:pass@somesite.com:10027
   */
  var URI string = os.Getenv("DB_HOST")
  var DBName string = os.Getenv("DB_NAME")

  if URI == "" {
    fmt.Println("no connection string provided, using localhost")
    URI = "localhost"
  }

  if DBName == "" {
    fmt.Println("no database name provided, bailing!")
    os.Exit(1)
  }
  
  conn := pnthr.MongoConnection{DB_HOST: os.Getenv("DB_HOST"), DB_NAME: os.Getenv("DB_NAME")}

  pnthr.Server(w, r, conn)
}