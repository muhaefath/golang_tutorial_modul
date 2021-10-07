package main

func SayHello(name string) string {

	return "hello world" + name
}

func SayHelloName(name string) string {

	return "hello world" + name
}

func SayHelloFunc(name ClientFunction) string {

	return "hello world"
}

func NewOrder(webhook WebhookHandler) {

	webhook.NewOrder(nil, nil)
}

func SayHelloFuncName(name ClientFunction) string {

	return "hello world"
}
