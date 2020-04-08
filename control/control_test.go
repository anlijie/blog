/*
 * @Author: your name
 * @Date: 2020-04-07 17:18:36
 * @LastEditTime: 2020-04-07 17:24:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go source\src\control\control_test.go
 */
package control

import (
	"log"
	"testing"
)

func TestRandStr(t *testing.T) {
	log.Println(RandStr(10))
}

func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}

func BenchmarkRandStrll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}
