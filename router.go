package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"petwork.core/util"
	"github.com/gorilla/context"
)

func Router() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexHandler)

	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return context.ClearHandler(Interceptor(CorsSupport(router)))
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	util.Response{
		Message: "Welcome to Vault-Tec" }.Writer(w).Status(108, http.StatusOK).Json()
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(
		w,
		util.Response{ Code: http.StatusNotFound, Message: "Not Found" }.ToJSON(),
		http.StatusNotFound)
}

func CorsSupport(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			} else {
				h.ServeHTTP(w, r)
			}
		})
}

func Interceptor(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start :=  time.Now()
			h.ServeHTTP(w, r)
			elapsed := time.Since(start)
			util.Info.Println("URL:", r.URL, "-", "Elapsed Time:", elapsed.String())
		})
}
