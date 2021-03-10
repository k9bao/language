package fileG

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Demo1() {
	if fileInfo, err := os.Stat("fileutil_test.go"); err == nil {
		log.Printf("ModTime:%v", fileInfo.ModTime())
		log.Printf("IsDir:%v", fileInfo.IsDir())
		log.Printf("Name:%v", fileInfo.Name())
		log.Printf("Size:%v", fileInfo.Size())

		fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		createTime := fileSys.CreationTime.Nanoseconds()   // 返回的是纳秒
		lastModiry := fileSys.LastWriteTime.Nanoseconds()  // 返回的是纳秒
		accessTime := fileSys.LastAccessTime.Nanoseconds() // 返回的是纳秒
		fileAttributes := fileSys.FileAttributes
		log.Printf("CreateTime:%v", createTime/1e9)
		log.Printf("ModifyTime:%v", lastModiry/1e9)
		log.Printf("AccessTime:%v", accessTime/1e9)
		log.Printf("FileAttr:%+v", fileAttributes)

		log.Printf("%+v", fileInfo)
	} else {
		log.Println("file/dir is not exist")
	}
}

func TestDemoCopy(t *testing.T) {
	written, err := CopyFile("C:\\work\\videos\\yanjing\\g-jd-out.h264", "C:\\work\\videos\\yanjing\\g-jd.h264", 0x0782A26E, -1)
	if err == nil {
		log.Println("copyFile Success.", written)
	} else {
		log.Println("copyFile fail.", err)
	}
}

func TestDelFileLine(t *testing.T) {
	//DelLines("C:\\Users\\Administrator\\Desktop\\core-err2.txt", "level=info msg=\"video")
	DelLines("C:\\Users\\Administrator\\Desktop\\core-err3.txt", "test2")
}

//go test -run ^(TestFileUtil)$
func TestFileUtil(t *testing.T) {
	DuplicateElement("D:\\work\\test\\ok.txt")
	DuplicateElement("D:\\work\\test\\err.txt")
	DelDumpText("D:\\work\\test\\pd1Videos.txt", "D:\\work\\test\\ok.txt")
	DelDumpText("D:\\work\\test\\pd1Videos.txt", "D:\\work\\test\\ignore.txt")
}

func DuplicateElement(file1 string) {
	fileList1 := GetFileList(file1)
	fileList1 = RmSameEle(fileList1)
	WriteFile(file1, fileList1)
}

func DelDumpText(file1, file2 string) {
	fileList1 := GetFileList(file1)
	fileList2 := GetFileList(file2)
	fileList1 = DeleteSlice2(fileList1, fileList2)
	WriteFile(file1, fileList1)
}

func TestReplace(t *testing.T) {
	fmt.Println(ReplaceSpecialExpr("abcdef", "", "a"))
	fmt.Println(ReplaceSpecialExpr("abcdef.ä¸‹è½½", "", ".ä¸‹è½½$"))
	var ymd = "^((([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29))\\s+([0-1]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"
	ymd = "\\d{4}-\\d{2}(-\\d{2})-"
	fmt.Println(ReplaceSpecialExpr("2018-01-25-test.file", "", ymd))

	fmt.Println(CheckSpecialExpr("2018-01-25-test.file", ymd))
}

func TestGetParentPath(t *testing.T) {
	file := "D:/work/knowledgebao/knowledgebao.github.io/_posts/2019-06-28-待整理列表.md"
	paths, fileName := filepath.Split(file)
	fileSuffix := path.Ext(fileName) //获取文件后缀
	fileName = strings.TrimSuffix(fileName, fileSuffix)
	parentPath := path.Base(paths)
	fmt.Println(parentPath, fileName)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
