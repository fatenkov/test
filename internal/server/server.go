package server

import "net/http"

type Server interface { // содержит методы, доступные в этом интерфейсе, сами методы описаны в httpserver.go
	Run()
	Sum(w http.ResponseWriter, r *http.Request)
	Sub(w http.ResponseWriter, r *http.Request)
	Div(w http.ResponseWriter, r *http.Request)
	Mult(w http.ResponseWriter, r *http.Request)
	Stop()
}
