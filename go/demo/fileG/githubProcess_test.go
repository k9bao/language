package fileG

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestProcess(t *testing.T) {
	WalkDir("C:/gwork/knowledgebao/knowledgebao.github.io/_posts", 0, AddTagData)
}

func TestWrite(t *testing.T) {
	// absFile, _ := GbkToUtf8(StringUtil.String2bytes("D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\2019-06-28-待整理列表.md"))
	// f, err := os.OpenFile(StringUtil.Bytes2String(absFile), os.O_RDWR, 0)
	// absFile := "D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\test.md"
	absFile := "D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\2019-06-28-待整理列表.md"
	f, err := os.OpenFile(absFile, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err.Error())
		return
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	bw.WriteString("---\nlayout: post")
	bw.WriteString("\ntitle: ")
	bw.WriteString("\n---\n")
	bw.Flush()
}
