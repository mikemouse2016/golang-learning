package main
import(
 "runtime/pprof"
 "os"
)
 
func main() {
        f,err:= os.Create("cpu.prof")
        if err != nil {
               // 可以输出一些出错信息
               return
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
        // other code …
}

