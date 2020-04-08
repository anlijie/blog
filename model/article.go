/*
 * @Author: your name
 * @Date: 2020-04-03 16:16:51
 * @LastEditTime: 2020-04-07 20:09:03
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go source\src\model\article.go
 */
package model

import "time"

type Article struct {
	Id      int64     `json:"id"` // tag
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Hits    int       `json:"hits"`
	Utime   time.Time `json:"utime"`
}

// query one data
func ArticleGet(id int64) (Article, error) {
	mod := Article{}
	err := Db.Unsafe().Get(&mod, "select * from Article where id = ? limit 1", id)
	return mod, err
}

// query a list data
func ArticleList() ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := Db.Unsafe().Select(&mods, "select * from Article order by id desc limit 10")
	return mods, err
}

// delete data By id
func ArticleDel(id int64) bool {
	res, _ := Db.Exec("delete from Article where id  = ?", id)
	if res == nil {
		return false
	}
	rows, _ := res.RowsAffected()
	if rows >= 1 {
		return true
	}
	return false
}

// add a data
func ArticleAdd(mod *Article) error {
	_, err := Db.Exec("insert into Article (title, author, content, hits, utime) value(?,?,?,?,?)", mod.Title, mod.Author, mod.Content, mod.Hits, mod.Utime)
	return err
}

// add a data
func ArticleEdit(mod *Article) error {
	_, err := Db.Exec("update article set title = ?, author = ?, content = ?, hits = ? where id = ?", mod.Title, mod.Author, mod.Content, mod.Hits, mod.Id)
	return err
}

// article page
func ArticlePage(pi int, ps int) ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := Db.Unsafe().Select(&mods, "select * from Article order by id desc limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

// article page sum
func ArticlePageCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from article")
	return count
}
