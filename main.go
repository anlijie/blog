/*
 * @Author: your name
 * @Date: 2020-04-03 16:32:49
 * @LastEditTime: 2020-04-07 20:29:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go source\main.go
 */
package main

import (
	"control"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", control.IndexView)                      // view
	http.HandleFunc("/add", control.ViewArticleAdd)              // view
	http.HandleFunc("/list", control.ListView)                   // view
	http.HandleFunc("/edit", control.EditView)                   // edit view
	http.HandleFunc("/detail", control.DetailView)               // detail view
	http.HandleFunc("/api/list/data", control.ListData)          // data
	http.HandleFunc("/api/index/data", control.IndexData)        // data
	http.HandleFunc("/api/article/page", control.ApiArticlePage) // page data
	http.HandleFunc("/api/list/del", control.ListDel)            // do
	http.HandleFunc("/api/article/add", control.ApiArticleAdd)   // add
	http.HandleFunc("/api/article/edit", control.ApiArticleEdit) // edit
	http.HandleFunc("/api/upload", control.ApiUpload)            // upload image
	http.ListenAndServe(":8080", nil)
}
