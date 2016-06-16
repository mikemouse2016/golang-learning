package main

import (
    "os"
    "os/exec"
    "path/filepath"
    "fmt"
)

/*获取当前文件执行的路径*/
func main() {
    file, _ := exec.LookPath(os.Args[0])
    fmt.Println(file)
    path, _ := filepath.Abs(file)
    fmt.Println(path)
}

