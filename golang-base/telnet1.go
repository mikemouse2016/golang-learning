package main

import (
    "net"
    "os"
    "fmt"
    "time"
    "strings"
)

func main() {
    //  println(os.Args[0])
    //  println("Hello world!")

    tcpAddr , err := net.ResolveTCPAddr("tcp4", ":8080")
    checkError(err, 1)

    listener , err := net.ListenTCP("tcp", tcpAddr)
    checkError(err, 2)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }

    os.Exit(0)
}

func handleClient(conn net.Conn) {
    conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // 设置两分钟超时。
    //conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // 设置两分钟超时。
    user_cmd := make([]byte, 128) //设置用户输入的命令
    defer conn.Close()

    for {
        read_len, err := conn.Read(user_cmd)

        if err != nil {
            fmt.Println(err)
            break
        }
        if read_len == 0 {
            break // connection already closed by client
        }

        //fmt.Println(string(user_cmd))
        //fmt.Println(len(string(user_cmd)))//长度是128，而不是time字符串的长度。
        //####################需要特殊处理字符串，找到每一个字符，然后累加。####################
        cmd_str := ""
        for i := 0; i < len(user_cmd); i ++ {
            //println(cmd_str[i])
            if user_cmd[i] == 0 {
                break
            }else {
                cmd_str += string(user_cmd[i])
            }
        }

        //替换字符
        cmd_str = strings.Replace(cmd_str, "\r\n", "", -1)

        if cmd_str == "time" {
            time_now := time.Now().String() + "\n"
            conn.Write([]byte(time_now))
        }else if cmd_str == "exit" {//退出命令。
            conn.Close()
        }else{
            if cmd_str == ""{
                fmt.Println("aaa\n")
            }
            help(conn)
        }

        user_cmd = make([]byte, 128) // clear last read content
    }
}

func help(conn net.Conn){
    conn.Write([]byte("help"))
}

func checkError(err error, num int) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%d) Fail Error：%s", num, err.Error())
        os.Exit(1)
    }
}
