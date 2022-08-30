/*
 * @Author: GG
 * @Date: 2022-08-30 15:45:29
 * @LastEditTime: 2022-08-30 16:12:52
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\httprouter\main.go
 *
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello,%s!\n", ps.ByName("name"))
}

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}
func main() {

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/user/:uid", getuser)
	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deleteuser)
	router.PUT("/moduser/:uid", modifyuser)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
