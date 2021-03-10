package fileG

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"syscall"
	"time"
)

//IsDir 判断文件是否是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//DelLines 检测absFile中的每一行内容,使用正则表达式expr进行匹配,把匹配行保存到absFile_back文件中.
func DelLines(absFile, expr string) error {
	f, err := os.OpenFile(absFile, os.O_RDWR, 0)
	if err != nil {
		log.Println("open file error")
		return errors.New("open file fail")
	}
	defer f.Close()
	fw, err := os.Create(fmt.Sprintf("%v_back", absFile))
	if err != nil {
		log.Println("open wfile error")
		return errors.New("open wfile fail")
	}
	defer fw.Close()
	//Read and write
	br := bufio.NewReader(f)
	wr := bufio.NewWriter(fw)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			log.Println("file is over")
			break
		}
		if CheckSpecialExpr(string(line), expr) {
			continue
		} else {
			wr.Write(line)
			wr.WriteString("\n")
			wr.Flush()
		}
	}
	return nil
}

//GetFileList return TextList,used ReadLine
func GetFileList(path string) []string {
	fileList := make([]string, 0, 100)
	f, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		log.Println("open file error")
		return fileList
	}
	defer f.Close()
	//Read
	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		fileList = append(fileList, string(line))
	}
	return fileList
}

func LineCounter(absFile string) (int, error) {
	var readSize int
	var err error
	var count int

	f, err := os.OpenFile(absFile, os.O_RDONLY, 0)
	if err != nil {
		return 0, errors.New("open file error")
	}
	defer f.Close()
	r := bufio.NewReader(f)

	buf := make([]byte, 1024)
	for {
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}
		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || readSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
	}
	if readSize > 0 && count == 0 || count > 0 {
		count++
	}
	if err == io.EOF {
		return count, nil
	}
	return count, err
}

//RmSameEle remove same string
func RmSameEle(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//DeleteSlice2 del same text form src
func DeleteSlice2(src []string, del []string) []string {
	j := 0
	for _, val := range src {
		find := false
		for _, val2 := range del {
			if val == val2 {
				find = true
				break
			}
		}
		if find == false {
			src[j] = val
			j++
		}
	}
	return src[:j]
}

//WriteFile write fileList to file,if file exist,delete
func WriteFile(file string, fileList []string) {
	f, _ := os.Create(file)
	defer f.Close()
	bw := bufio.NewWriter(f)
	for _, s := range fileList {
		bw.WriteString(s)
		bw.WriteString("\n")
	}
	bw.Flush()
}

//CopyFile copy from srcName to dstName,from beg pos,pos must >=0, if count <=0, copy to src end
func CopyFile(dstName, srcName string, beg, count int64) (written int64, err error) {
	if beg < 0 {
		beg = 0
	}
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	src.Seek(beg, 0)
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	if count > 0 {
		return io.CopyN(dst, src, count)
	}
	return io.Copy(dst, src)
}

//ReplaceSpecialExpr 用repl替换满足正则表达式的内容
func ReplaceSpecialExpr(src, repl, expr string) string {
	reg := regexp.MustCompile(expr)
	return reg.ReplaceAllString(src, repl)
}

//CheckSpecialExpr 查找filename是否有复合正则要求的内容
func CheckSpecialExpr(src, expr string) bool {
	reg := regexp.MustCompile(expr)
	return reg.MatchString(src)
}

//TrimSpecialSuffix 文件删除指定后缀
func TrimSpecialSuffix(absFile string, suffix string) {
	// fmt.Println(filepath.Base(absFile)) //获取路径中的文件名test.txt
	if path.Ext(absFile) == suffix {
		newFileName := strings.TrimSuffix(absFile, suffix)
		os.Rename(absFile, newFileName)
	}
}

//GetFileModTime 获取文件修改时间 返回time.Time对象
func GetFileModTime(path string) time.Time {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error")
		return time.Now()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("stat fileinfo error")
		return time.Now()
	}

	return fi.ModTime()
}

//DEPTH 遍历的最深层次
var DEPTH = 10

//WalkDir 遍历目录dirpath,如果发现文件，则使用f进行处理
func WalkDir(dirpath string, depth int, f func(string)) {
	if depth > DEPTH { //大于设定的深度
		return
	}
	files, err := ioutil.ReadDir(dirpath) //读取目录下文件
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			WalkDir(dirpath+"/"+file.Name(), depth+1, f)
			continue
		} else {
			f(dirpath + "/" + file.Name())
		}
	}
}

//GetCreateTime 获取文件创建时间
func GetCreateTime(info os.FileInfo) time.Time {
	sysinfo := info.Sys()
	if stat, ok := sysinfo.(*syscall.Win32FileAttributeData); ok {
		return time.Unix(0, stat.CreationTime.Nanoseconds())
	}
	// if stat, ok := sysinfo.(*syscall.Stat_t); ok {
	// 	t := time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
	// 	fmt.Println(t)
	// }
	return time.Time{}
}
