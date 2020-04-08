/*
 * @Author: your name
 * @Date: 2020-04-04 15:45:22
 * @LastEditTime: 2020-04-07 18:38:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go source\src\control\view.go
 */
package control

import (
	"io"
	"net/http"
	"os"
)

// IndexView
func IndexView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
	f.Close()
}

// list view
func ListView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/list.html")
	io.Copy(w, f)
	f.Close()
}

// edit view
func EditView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/edit.html")
	io.Copy(w, f)
	f.Close()
}

// detail view
func DetailView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/detail.html")
	io.Copy(w, f)
	f.Close()
}

// add a article
func ViewArticleAdd(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/add.html")
	io.Copy(w, f)
	f.Close()
}
