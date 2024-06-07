package main

import (
	"fmt"
	"net/http"
	"strconv"
	"test/internal/calc"
	"text/template" // пакет необходим для парсинга файла index.html
)

func main() {

	// Создаём файловый сервер
	fs := http.FileServer(http.Dir("web/calc"))
	// Отдаём index.html по запросу на http://localhost:8080
	http.Handle("/", fs)
	// Регистрируем обработчики запросов пользователя
	// специально для Гита
	http.HandleFunc("/doCulc", doCulc)
	http.HandleFunc("/sum", calc.Sum)
	http.HandleFunc("/sub", calc.Sub)
	http.HandleFunc("/mult", calc.Mult)
	http.HandleFunc("/div", calc.Div)
	// Запускаем локальный сервер на порту :8080
	//http.ListenAndServe(":8080", nil) - то что было
	err := http.ListenAndServe(":8080", nil) // запустим с отработкой возможной ошибки
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

type dataForCulc struct { // создадим структуру
	Answer        int    // Ответ нашей операции
	Error         string // сообщения об ошибках
	IsAnswerExist bool   // если ответ есть = true
	Num1          int    // 1-е число
	Num2          int    // 2-е число
	Operation     string // тип операции
}

func doCulc(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("web/calc/index.html") // Parse template file
	//t.Execute(w, dataForCulc{Num1: 6, Num2: 10})

	r.ParseForm() // парсим нашу форму

	var err error

	data := dataForCulc{}

	if len(r.Form["number1"]) > 0 { //Проверяем, что массив не пустой
		data.Num1, err = strconv.Atoi(r.Form["number1"][0]) //Конвертируем string --> INT + присваиваем значение полю структуры
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
	}
	if len(r.Form["number2"]) > 0 { // все тоже самое со 2-м числом
		data.Num2, err = strconv.Atoi(r.Form["number2"][0])
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
	}
	if len(r.Form["operation"]) > 0 { //Проверяем, что массив не пустой
		data.Operation = r.Form["operation"][0]
	}
	// Выведем ту же операцию в терминал
	fmt.Printf("num1 %v num2 %v %v\n", data.Num1, data.Num2, data.Operation)

	switch data.Operation { // определяем какая операция прошла при нажатии кнопки
	case "sum":
		data.Answer = data.Num1 + data.Num2
		data.IsAnswerExist = true
	case "sub":
		data.Answer = data.Num1 - data.Num2
		data.IsAnswerExist = true
	case "mult":
		data.Answer = data.Num1 * data.Num2
		data.IsAnswerExist = true
	case "div":
		if data.Num2 != 0 {
			data.Answer = data.Num1 / data.Num2
			data.IsAnswerExist = true
		} else {
			data.Error = " Делить на ноль нельзя!"
		}
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

}
