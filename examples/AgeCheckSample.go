package examples

type MyPoGo struct {
}

func (p *MyPoGo) GetStringLength(sarg string) int {
	return len(sarg)
}

type User struct {
	Name string
	Age  int
	Male bool
}
