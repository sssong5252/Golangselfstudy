package main
import "fmt"

func main() {
	var s sampleinterface //interface 변수에 값이 설정되어 있지 않아서 런타임 에러(nil 포인터 에러)
	fmt.Println(s)
	s.samplemethod()
}

type sampleinterface interface {
	samplemethod()
}

