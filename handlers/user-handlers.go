package handlers

import "net/http"

func CreateUser() func(w http.ResponseWriter, req * http.Request) { return func(w http.ResponseWriter, req * http.Request){}}

func UpdateUser() func(w http.ResponseWriter, req * http.Request) {return func(w http.ResponseWriter, req * http.Request){}}
func DeleteUser() func(w http.ResponseWriter, req * http.Request) {return func(w http.ResponseWriter, req * http.Request){}}
func SoftDeleteUser() func(w http.ResponseWriter, req * http.Request) {return func(w http.ResponseWriter, req * http.Request){}}
func GetUser() func(w http.ResponseWriter, req * http.Request) {return func(w http.ResponseWriter, req * http.Request){}}
func ListUser() func(w http.ResponseWriter, req * http.Request) {return func(w http.ResponseWriter, req * http.Request){}}