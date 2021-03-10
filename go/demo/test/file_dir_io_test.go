package test

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"../fileG"
)

var DEPTH = 10

func walkDir(dirpath string, depth int, f func(string)) {
	if depth > DEPTH { //大于设定的深度
		return
	}
	files, err := ioutil.ReadDir(dirpath) //读取目录下文件
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			walkDir(dirpath+"/"+file.Name(), depth+1, f)
			continue
		} else {
			f(dirpath + "/" + file.Name())
		}
	}
}

//Delete Ext .下载
func ProcessFun(absFile string) {
	// fmt.Println(filepath.Base(absFile)) //获取路径中的文件名test.txt
	if path.Ext(absFile) == ".下载" {
		paths, fileName := filepath.Split(absFile)
		filenameOnly := strings.TrimSuffix(fileName, ".下载")
		newFileName := paths + filenameOnly
		os.Rename(absFile, newFileName)
	}
}

//
func DelBeginString(absFile string) {
	paths, fileName := filepath.Split(absFile)
	newFileName := fileG.ReplaceSpecialExpr(fileName, "OpenCV3.1.0", "^(OpenCV 3.1.0).*?")
	os.Rename(absFile, paths+newFileName)
}

func ReadFileALL(filePath string) ([]byte, error) {
	if f, err := os.Open(filePath); err == nil {
		bs, err := ioutil.ReadAll(f)
		if err == nil {
			return nil, err
		}
		f.Close()
		return bs, err
	}
	return nil, errors.New("fail")
}

func TestRename(t *testing.T) {
	DEPTH = 1
	walkDir("D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\opencv\\图像处理教程配套PPT", 0, DelBeginString)
}
func TestReadDir1(t *testing.T) {
	//walkDir("C:/Users/Administrator/Desktop/websocket/websocket.org Echo Test - Powered by Kaazing_files", 0, ProcessFun)
	ReadFileALL("info.json")
}

func TestWriteAll(t *testing.T) {
	//walkDir("C:/Users/Administrator/Desktop/websocket/websocket.org Echo Test - Powered by Kaazing_files", 0, ProcessFun)
	b := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ioutil.WriteFile("test.txt", b, os.ModePerm)
}
