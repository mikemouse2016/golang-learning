package main

import (
	"log"
	"strconv"
)

//捕获因未知输入导致的程序异常
func catch(nums ...int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	return nums[1] * nums[2] * nums[3] //index out of range
}

//主动抛出 panic，不推荐使用，可能会导致性能问题
func toFloat64(num string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[W]", r)
		}
	}()
	if num == "" {
		panic("param is null") //主动抛出 panic
		//log.Panicf("param is null")
	}
	return strconv.ParseFloat(num, 10)
}
func main() {
	catch(2, 8)
	toFloat64("")
}

/*
Panic
是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数F
调用panic，函数F 的执行被中断，并且F 中的延迟函数会正常执行，然后F 返回到调
用它的地方。在调用的地方，F 的行为就像调用了panic。这一过程继续向上，直到程
序崩溃时的所有goroutine 返回。
恐慌可以直接调用panic 产生。也可以由运行时错误产生，例如访问越界的数组。
Recover
是一个内建的函数，可以让进入令人恐慌的流程中的goroutine 恢复过来。Recover
仅在延迟函数中有效。
在正常的执行过程中，调用recover 会返回nil 并且没有其他任何效果。如果当前
的goroutine 陷入恐慌，调用recover 可以捕获到panic 的输入值，并且恢复正常的执
行。
*/
