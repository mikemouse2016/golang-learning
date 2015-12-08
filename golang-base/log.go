package main
import "log"
/*
Logging库一般提供不同的log等级。与这些logging库不同，Go中log包在你调用它的Fatal*()和Panic*()函数时，可以做的不仅仅是log。当你的应用调用这些函数时，Go也将会终止应用
*/
func main() {  
    log.Fatalln("Fatal Level: log entry") //app exits here
    log.Println("Normal Level: log entry")
}
