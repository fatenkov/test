package httpserver

import (
	"net/http"
	"strconv"
	"test/internal/service"
	"test/internal/service/exponentiation"
	//"test/internal/service/calculator"
)

type Server struct {
	server         *http.Server
	service        service.Service
	exponentiation service.Exponentiation
}

func New(port int, service service.Service) *Server {
	// Имплементировать конструктор
	return &Server{
		server:         &http.Server{},
		service:        service,
		exponentiation: &exponentiation.Exponentiation{},
		//service: &calculator.Calculator{}, // как будто так тоже правильно? или таки велосипед?
	}

}

func (s *Server) Run() {
	// Имплементировать метод
}

func (s *Server) Sum(w http.ResponseWriter, r *http.Request) {
	// Имплементировать метод
	// приняли запрос, вытащили аргументы, провалидировали
	// Получаем значения x и y из URL и отрабатываем ошибку если если x или y не являются числами
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))

	if errX != nil || errY != nil {
		http.Error(w, "Некорректные параметры x или y", http.StatusBadRequest)
		return
	}

	// отдали в соответствующий метод s.service, получили результат выполнения
	z := s.service.Sum(x, y)

	//возвращаем обратно пользователю
	w.Write([]byte(strconv.Itoa(z)))
}

func (s *Server) Sub(w http.ResponseWriter, r *http.Request) {
	// Имплементировать метод
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))

	if errX != nil || errY != nil {
		http.Error(w, "Некорректные параметры x или y", http.StatusBadRequest)
		return
	}

	z := s.service.Sub(x, y)

	w.Write([]byte(strconv.Itoa(z)))
}

func (s *Server) Div(w http.ResponseWriter, r *http.Request) {
	// Имплементировать метод
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))

	if errX != nil || errY != nil {
		http.Error(w, "Некорректные параметры x или y", http.StatusBadRequest)
		return
	}

	z := s.service.Div(x, y)

	w.Write([]byte(strconv.Itoa(z)))
}

func (s *Server) Mult(w http.ResponseWriter, r *http.Request) {
	// Имплементировать метод
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))

	if errX != nil || errY != nil {
		http.Error(w, "Некорректные параметры x или y", http.StatusBadRequest)
		return
	}

	z := s.service.Mult(x, y)

	w.Write([]byte(strconv.Itoa(z)))
}

func (s *Server) Exp(w http.ResponseWriter, r *http.Request) {
	// Имплементировать метод
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))

	if errX != nil || errY != nil {
		http.Error(w, "Некорректные параметры x или y", http.StatusBadRequest)
		return
	}

	z := s.exponentiation.Exp(x, y)

	w.Write([]byte(strconv.Itoa(z)))
}

func (s *Server) Stop() {
	// Имплементировать метод
}
