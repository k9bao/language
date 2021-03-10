package fileG

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//github.io发布的时候需要特定的格式，这里实现对指定文件
//夹的文件进行处理，使其符合发布条件

//AddTagData combine func
func AddTagData(absFile string) {
	if path.Ext(absFile) == ".md" {
		CheckHaveData(absFile)
		AddTag(absFile)
		AddDataTime(absFile)
	}
}

//AddDataTime 给文件名添加时间,比如：源文件名叫oldFile.txt,则处理后格式为2019-08-08-oldFile.txt
func AddDataTime(absFile string) {
	// fmt.Println(filepath.Base(absFile)) //获取路径中的文件名test.txt
	paths, fileName := filepath.Split(absFile)
	if CheckSpecialExpr(fileName, "\\d{4}-\\d{2}(-\\d{2})-") {
		return
	}
	newFilename := GetFileModTime(absFile).Format("2006-01-02-") + fileName
	newAbsFile := paths + newFilename
	os.Rename(absFile, newAbsFile)
}

//GetValue return value from "key:value"
func GetValue(key string, text string) string {
	a := strings.Split(text, ":")
	if len(a) >= 2 && strings.TrimSpace(a[0]) == key {
		return strings.TrimSpace(a[1])
	}
	return ""
}

//AddTag 判断文件是否有标题，如果没有添加标题，如果有跳过，标题格式如下：
//矫正filename和folderName，因为文件名和父目录可能会发生变化。
/*
---
layout: post
title: filename
date: 2016-01-09 11:15:06
description: filename
tag: folderName

---
*/
func AddTag(absFile string) {
	f, err := os.OpenFile(absFile, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()

	//GetInfo
	fi, err := f.Stat()
	if err != nil {
		fmt.Println("stat fileinfo error")
		return
	}
	br := bufio.NewReader(f)
	buffer := make([]byte, fi.Size())

	//get file info
	//date := GetCreateTime(fi).Format("2006-01-02 15:04:05")
	date := fi.ModTime().Format("2006-01-02 15:04:05")
	paths, title := filepath.Split(absFile)
	title = strings.TrimSuffix(title, path.Ext(title))
	title = ReplaceSpecialExpr(title, "", "\\d{4}-\\d{2}(-\\d{2})-")
	tag := path.Base(paths)

	state := 0 //0:noTag,1:haveTag but need rewrite,2:haveTag and not need rewrite
	//Check is OK? if not ok,leave file text only.
	a, _, c := br.ReadLine()
	if c == io.EOF {
		//
	} else if (string(a)) == "---" {
		br.ReadLine()           // layout
		a, _, _ = br.ReadLine() //title
		title2 := GetValue("title", string(a))
		a, _, _ = br.ReadLine()              //date
		date2 := GetValue("date", string(a)) //get y-m-d
		br.ReadLine()                        //description
		a, _, _ = br.ReadLine()              //tag
		tag2 := GetValue("tag", string(a))
		if (title2 != title) || !strings.Contains(date, date2) || (tag2 != tag) {
			state = 1
		} else {
			state = 2
		}
	}

	if state == 2 {
		return
	}
	//Read file text
	f.Seek(0, 0)
	if _, err = f.Read(buffer); err != nil {
		log.Fatal("read file fail")
		return
	}
	if state == 1 {
		for i := 5; i < len(buffer); i++ {
			if buffer[i-2] == '-' && buffer[i-1] == '-' && buffer[i] == '-' {
				buffer = buffer[i+1:]
				break
			}
		}
	}

	//set file size is zero
	f.Truncate(0)

	//write info
	f.Seek(0, 0)
	bw := bufio.NewWriter(f)
	bw.WriteString("---\nlayout: post")
	bw.WriteString("\ntitle: ")
	bw.WriteString(title)
	bw.WriteString("\ndate: ")
	bw.WriteString(date)
	bw.WriteString("\ndescription: ")
	bw.WriteString(title)
	bw.WriteString("\ntag: ")
	bw.WriteString(tag)
	bw.WriteString("\n\n---\n")

	//write file text
	bw.Write(buffer)
	bw.Flush()
}

var needWrite = 0

//CheckHaveData 检测是否有数据,如果文件小于10行,就认为文件是空的
func CheckHaveData(absFile string) {
	reqNums := 15
	count, _ := LineCounter(absFile)
	if count < reqNums {
		needWrite++
		fmt.Println(absFile, " LineCounter = ", count, " < ", reqNums, "sum=", needWrite)
	}
}
