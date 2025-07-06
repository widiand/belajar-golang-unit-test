package helper

import "testing"

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		//unit test failed
		panic("Result is not Hello Eko")
	}
}

func TestHelloWorldKhanndey(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		//unit test failed
		panic("Result is not Hello Khannedy")
	}
}