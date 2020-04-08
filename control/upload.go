/*
 * @Author: your name
 * @Date: 2020-04-07 15:18:23
 * @LastEditTime: 2020-04-07 17:29:07
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go source\src\control\upload.go
 */
package control

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
	"unsafe"
)

func ApiUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20 * 20) // 20MB
	f, h, err := r.FormFile("upfile")
	if err != nil {
		Fail(w, "上传失败", err.Error())
	}
	log.Println(f, h, err)
	os.MkdirAll("static/upload", 0666)
	ext := path.Ext(h.Filename)
	name := "static/upload" + RandStr(10) + ext
	dst, _ := os.Create(name)
	io.Copy(dst, f)
	f.Close()
	dst.Close()
	// Succ(w, "上传成功", "xxx")
	w.Header().Set("Content-Type", "application/json")
	// original: "mobileBg.jpg", state: "SUCCESS", title: "mobileBg.jpg",…
	// "{\"original\":\"mobileBg.jpg\",\"state\":\"SUCCESS\",\"title\":\"mobileBg.jpg\",\"url\":\"" + ("/" + name) + "\"}"
	//w.Write([]byte("{\"original\": \"mobileBg.jpg\", \"state\": \"SUCCESS\", \"title\": \"mobileBg.jpg\",\"url\":\"" + ("/" + name) + "\"}"))
	mod := editorReply{
		Original: h.Filename,
		State:    "SUCCESS",
		Title:    h.Filename,
		Url:      "/" + name,
	}
	w.Write(mod.Json())
}

var all = "abcdefghijklmnopqrstuvwxyz1234567890"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStr(ln int) string {
	res := make([]byte, ln, ln)
	for i := 0; i < ln; i++ {
		res[i] = all[rand.Intn(36)]
	}
	return *(*string)(unsafe.Pointer(&res))
}

type editorReply struct {
	Original string `json:"original"`
	State    string `json:"state"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}

// method
func (er *editorReply) Json() []byte {
	buf, _ := json.Marshal(er)
	return buf
}
