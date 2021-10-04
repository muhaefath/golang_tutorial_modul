package main

//go:generate ./modules/model/newList.sh Person

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
