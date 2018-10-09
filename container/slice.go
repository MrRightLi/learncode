package container

import "fmt"

func main() {
	// slice 是 array的 view
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[2:6]", arr[:6])
	fmt.Println("arr[2:6]", arr[2:])
	fmt.Println("arr[2:6]", arr[:])

	s := arr[2:6]
	s[0] = 10
	s[2] = 11
	fmt.Println(arr)

	// slice 的扩展
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 :=  arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	// 向slice添加元素
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// s4 and s5 no longer view arr.
	// 添加元素时如果超越cap，系统会重新分配更大的底层数组，原来的数组如果没有人用，原来的数组就会被回收
	// 由于值传递的关系，必须接受append的返回值
	fmt.Println("arr =", arr)

}
