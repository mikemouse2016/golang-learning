package main

import (
	"fmt"
	"math"
)

const (
	ca = iota
	cb = 'A'
	cc = iota
)

func main() {
	var (
		a int
		b bool
		c string
		d []int
		e [2]int8
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(ca)
	fmt.Println(cb)
	fmt.Println(cc)

	if ia := 1; ia > 2 {
		fmt.Println(ia)
	}
	//fmt.Println(ia)

	var sa = 3
	switch {
	case a < 2:
		fmt.Println("a<2")
		fallthrough
	case sa > 2:
		fmt.Println("sa>2")
		fallthrough
	default:
		fmt.Println("default")
	}

	ia := []int{10: 2, 3, 4, 5}
	fmt.Println(ia)

	var ib [10]int
	ib[1] = 10
	fmt.Println(ib)

	ic := new([10]int)
	ic[1] = 10
	fmt.Println(ic)
	fmt.Println(*ic)

	ma := make([]int, 3, 10)
	fmt.Println(len(ma), cap(ma))
	fmt.Printf("%v, %p\n", ma, ma)

	mb := append(ma, 1, 2, 3, 4, 5, 6)
	fmt.Printf("%v, %p\n", mb, mb)

	mc := append(mb, 1, 2, 3, 4, 5, 6)
	fmt.Printf("%v, %p\n", mc, mc)

	ii := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ii)
	fmt.Println(ii[3:5])
	fmt.Println(ii[3:])
	fmt.Println(ii[:4])

	ca := []int{1, 2, 3, 4, 5, 6, 7, 8}
	cb := []int{11, 12, 13, 14, 15, 16}
	//copy(ca[1:5], cb[1:4])
	copy(ca, cb)
	fmt.Println(ca)

	m := make(map[string]int) //使用make创建一个空的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)      //输出 map[three:3 two:2 one:1] (顺序在运行时可能不一样)
	fmt.Println(len(m)) //输出 3

	v := m["two"]  //从map里取值
	fmt.Println(v) // 输出 2

	delete(m, "two")
	fmt.Println(m) //输出 map[three:3 one:1]

	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1) //输出 map[two:2 three:3 one:1] (顺序在运行时可能不一样)

	for key, val := range m1 {
		fmt.Printf("%s => %d \n", key, val)
		/*输出：(顺序在运行时可能不一样)
		  three => 3
		  one => 1
		  two => 2*/
	}

	var i int = 1
	var pInt *int = &i
	//输出：i=1     pInt=0xf8400371b0       *pInt=1
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	*pInt = 2
	//输出：i=2     pInt=0xf8400371b0       *pInt=2
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	i = 3
	//输出：i=3     pInt=0xf8400371b0       *pInt=3
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	//下面的代码分配了一个整型数组，长度为10，容量为100，并返回前10个数组的切片
	//aa := make([]int, 10, 100)

	//var p *[]int = new([]int)     // 为切片结构分配内存；*p == nil；很少使用
	//var v []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用

	// 不必要地使问题复杂化：
	var p *[]int = new([]int)
	fmt.Println(p) //输出：&[]
	*p = make([]int, 10, 10)
	fmt.Println(p)       //输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) //输出： 0

	sum(1, 2)
	sum(1, 2, 3)
	//传数组
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//初始化
	person := Person{"Tom", 30, "tom@gmail.com"}
	person = Person{name: "Tom", age: 30, email: "tom@gmail.com"}
	fmt.Println(person) //输出 {Tom 30 tom@gmail.com}
	pPerson := &person
	fmt.Println(pPerson) //输出 &{Tom 30 tom@gmail.com}
	pPerson.age = 40
	person.name = "Jerry"
	fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}

	r := rect{width: 10, height: 15}
	fmt.Println("面积: ", r.area())
	fmt.Println("周长: ", r.perimeter())
	rp := &r
	fmt.Println("面积: ", rp.area())
	fmt.Println("周长: ", rp.perimeter())
}

func sum(nums ...int) {
	fmt.Print(nums, " ") //输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)

	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}

	//延期函数
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}

}

func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

type Person struct {
	name  string
	age   int
	email string
}

type rect struct {
	width, height int
}

func (r rect) area() int { //求面积
	return r.width * r.height
}

func (r *rect) perimeter() int { //求周长
	return 2 * (r.width + r.height)
}

//---------- 接 口 --------//
type shape interface {
	area() float64      //计算面积
	perimeter() float64 //计算周长
}

//--------- 长方形 ----------//
type rectf struct {
	width, height float64
}

func (r *rectf) area() float64 { //面积
	return r.width * r.height
}

func (r *rectf) perimeter() float64 { //周长
	return 2 * (r.width + r.height)
}

//----------- 圆  形 ----------//
type circle struct {
	radius float64
}

func (c *circle) area() float64 { //面积
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 { //周长
	return 2 * math.Pi * c.radius
}

// ----------- 接口的使用 -----------//
func interface_test() {
	r := rectf{width: 2.9, height: 4.8}
	c := circle{radius: 4.3}

	s := []shape{&r, &c} //通过指针实现

	for _, sh := range s {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
	}
}
