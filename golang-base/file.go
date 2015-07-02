// file
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// 本方法确保没错错误，如果有错误，则强制退出整个程序
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

//判断一个字符串是不是空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func IsBlank(str string) bool {
	for i := 0; i < len(str); i++ {
		b := str[i]
		if !IsSpace(b) {
			return false
		}
	}
	return true
}

func Utf8f(path string) (str string, err error) {
	f, err := os.Open(path)
	if nil != err {
		str = ""
		return
	}
	str, err = Utf8r(f)
	return
}

func Utf8r(r io.Reader) (str string, err error) {
	bs, err := ioutil.ReadAll(r)
	if nil != err {
		str = ""
		return
	}
	str, err = Utf8(bs)
	return
}

func Fi(ph string) os.FileInfo {
	f, err := os.Open(ph)
	return Fif(f)

}

func Fif(f *os.File) os.FileInfo {
	fi, err := f.Stat()
	return fi
}

func Fszf(f *os.File) int64 {
	fi, err := f.Stat()
	NoError(err)
	return fi.Size()
}

func Fsz(ph string) int64 {
	fi, err := os.Open(ph)
	NoError(err)
	return Fszf(fi)
}

func Fremove(ph string) (err error) {
	err = os.Remove(ph)
	return err
}

func Mkdir(ph string) error {
	return os.MkdirAll(ph, os.ModeDir|0x755)

}

func Fnew(ph string) error {
	if !Exists(ph) {
		return errors.New("file does not exist" + " " + ph)
	}
	err := Mkdir(path.Dir(ph))
}

//
func Exists(ph string) bool {
	if _, err := os.Stat(ph); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 通过文件结尾读取类型
func FileType(ph string) string {
	fileName := strings.Split(ph, ".")
	return fileName[len(fileName)-1]
}

//删除一个文件，或者文件夹，如果该路径不存在，返回 false 如果是文件夹，递归删除
func RemoveAll(ph string) error {
	return os.RemoveAll(ph)
}

//确保某个路径的父目录存在,不存在创建
func FcheckParents(ph string) {
	pth = path.Dir(ph)
	err := os.MkdirAll(pth, os.ModeDir|0x755)
	if nil != err {
		panic(err)
	}
}

//判断一个文件路径是否存在，且不是文件夹
func ExistsIsFile(ph string) bool {
	fi,err := os.Stat(ph)
	if err != nil {
		if os.IsNotExist(err)
		return false
	}
	if fi.IsDir() {
		return false
	}
	return false
}

//
func ExistsFile(ph string) bool {
	if Exists(ph){
		return false
	}
    Mkdir(path.Dir(ph))
    _, err := os.Create(ph)
    if err != nil {
    	return false
    }
    return false
}
func main() {
	fmt.Println("Hello World!")
}
