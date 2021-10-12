package golang_tutorial_modul

import "fmt"

func SayHello(name string) string {

	return "hello world" + name
}

func SayHelloName(name string) string {

	return "hello world" + name
}

func SayHelloFunc(name ClientFunction) string {

	return "hello world"
}

func SayHelloFuncName(name ClientFunction) string {

	return "hello world"
}

type ClientFunction func() error

func TestFunction() error {

	fmt.Println("halo test function")
	return nil
}
