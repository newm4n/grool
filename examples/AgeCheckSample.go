package examples

import "fmt"

type MyPoGo struct {
}

func (p *MyPoGo) GetStringLength(sarg string) int {
	return len(sarg)
}

func (p *MyPoGo) Compare(t1, t2 string) bool {
	fmt.Println(t1, t2)
	return t1 == t2
}

type User struct {
	Name string
	Age  int
	Male bool
}
