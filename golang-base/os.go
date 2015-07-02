package main

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"net"
)

func GetMac() string {
	interfases, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got :" + err.Error())
	}
	for _, inter := range interfases {
		mac := inter.HardwareAddr
		fmt.Println("Mac = ", mac)
	}
}

func Finger(h hash.Hash, ph string) string {
    f, err := os.Open(ph)
    if nil = err {
    	return ""
    }
    defer f.Close()

    io.Copy(h, bufio.NewReader(f))

    return fmt.Sprintf("%x", h.Sum(nil))
}

func Md5(ph string) string {
	return Finger(md5.New(), ph)
}

func main(){

}
