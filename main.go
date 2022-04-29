// Package website Основной пакет для работы сайта.
package website

import (
	"html/template"
	"log"
	"net/http"
)

// Address хранит адрес веб сайта по которому можно обратиться.
var Address = "localhost:8080"

// Hello функция для вывода страницы Приветствия.
// В самой функции понадобиться путь до html файла с самой страничкой.
func Hello(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("/web/home.html")
	if err != nil {
		log.Panic(err)
	}
	file.Execute(w, nil)
}

// Login функция для вывода страницы логина, где пользователь должен авторизоваться.
// В самой функции понадобиться путь до html файла с самой страничкой.
func Login(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("/web/login.html")
	if err != nil {
		log.Panic(err)
	}
	file.Execute(w, nil)
}

// Registration функция для вывода страницы регистрации, где пользователь должен зарегистрироваться.
// В самой функции понадобиться путь до html файла с самой страничкой.
func Registration(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("/web/registration.html")
	if err != nil {
		log.Panic(err)
	}
	file.Execute(w, nil)
}

// StartUp функция для вывода страницы стартапов, где пользователи могут познакомиться со всеми возможными стартапами.
// В самой функции понадобиться путь до html файла с самой страничкой.
func StartUp(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("/web/startUp.html")
	if err != nil {
		log.Panic(err)
	}
	file.Execute(w, nil)
}

// Account функция для вывода страницы личного кабинета, где пользователь может увидеть свой профиль.
// В самой функции понадобиться путь до html файла с самой страничкой.
func Account(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("/web/account.html")
	if err != nil {
		log.Panic(err)
	}
	file.Execute(w, nil)
}

// Main основная функция, которая отлавливает запросы пользователей.
func Main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/startUp", StartUp)
	http.HandleFunc("/account", Account)
	log.Panic(http.ListenAndServe(Address, nil))
}
