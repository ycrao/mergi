package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pngCharMap map[string]string
	pngCharMap = make(map[string]string)
	chars_1 := "abcdefghijklmnopqrstuvwxyz"
	chars_2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	chars_3 := "0123456789.+-%↑↓"
	for k, v := range chars_1 {
		ptY := k*28
		pngCharMap[string(v)] = "-i \"./chars_1.png\" -c \"" + strconv.Itoa(ptY) + " 0 28 64\" "
	}
	for k, v := range chars_2 {
		ptY := k*28
		pngCharMap[string(v)] = "-i \"./chars_2.png\" -c \"" + strconv.Itoa(ptY) + " 0 28 64\" "
	}
	for k, v := range chars_3 {
		ptY := k*28
		pngCharMap[string(v)] = "-i \"./chars_3.png\" -c \"" + strconv.Itoa(ptY) + " 0 28 64\" "
	}
	// fmt.Printf("%+v\n", pngCharMap)
	str := "349.56↑"
	bash := "./mergi -t \"TTTTTTT\" "
	for _, v := range str {
		bash = bash + pngCharMap[string(v)]
	}
	fmt.Println(bash)
}


func mapAlpha() {
	var pngCharMap map[string]string
	pngCharMap = make(map[string]string)
	chars_1 := "abcdefghijklmnopqrstuvwxyz"
	//var chars_2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//var chars_3 := "0123456789.+-%↑↓"
	for k, v := range chars_1 {
		fmt.Println(k)
		fmt.Println(string(v))
		ptY := (k - 1)*28
		pngCharMap[string(v)] = "chars_1.png 0 " + strconv.Itoa(ptY) + " 28 64"
	}
	fmt.Printf("%+v\n", pngCharMap)
}