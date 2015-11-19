package main

import (
	"os"
	"flag"
	"strconv"
	"petwork.core/util"
	"net/http"
)


var (
	port string = ""
)

func main() {


	// Initialize Logger
	util.InitLogger()

	FetchPort(&port)

	util.Info.Println("App Listening to: ", port)

	util.Error.Fatal(http.ListenAndServe(":" + port, Router()))


}


func FetchPort(port *string) {

	if *port != "" {
		return
	}

	*port = os.Getenv("PORT")

	if *port == "" {
		flagPort := flag.Int("PORT", 3000, "Port Number to serve the app")
		flag.Parse()
		*port = strconv.Itoa(*flagPort)
	}
}