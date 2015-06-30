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

func Fi(path string) os.FileInfo {
	f, err := os.Open(path)
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

func Fsz(path string) int64 {
	fi, err := os.Open(path)
	NoError(err)
	return Fszf(fi)
}

func Fremove(path string) (err error) {
	err = os.Remove(path)
	return err
}

func Fnew(path string) error {
	if Exists(path) {
		return errors.New("file does not exist" + " " + path)
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 通过文件结尾读取类型
func FileType(path string) string {
	fileName := strings.Split(path, ".")
	return fileName[len(fileName)-1]
}

//删除一个文件，或者文件夹，如果该路径不存在，返回 false 如果是文件夹，递归删除
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

//确保某个路径的父目录存在,不存在创建
func FcheckParents(path string) {
	pth = path.Dir(path)
	err := os.MkdirAll(pth, os.ModeDir|0x755)
	if nil != err {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello World!")
}
