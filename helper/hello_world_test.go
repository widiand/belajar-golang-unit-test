package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		//unit test failed
		t.Fail() //masih lanjut ke print
	}

	fmt.Println("TestHelloWorldEko done")
}

func TestHelloWorldKhanndey(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		//unit test failed
		t.FailNow() //langsung berhenti
	}

	fmt.Println("TestHelloWorldKhanndey done")
}

func TestHelloWorldEko1(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		//unit test failed
		t.Error("harusnya Hello Eko") //after logging the error will call Fail()
	}

	fmt.Println("dieksekusi")
}

func TestHelloWorldEko2(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		//unit test failed
		t.Fatal("harusnya Hello Eko") //after logging the error will call FailNow()
	}

	fmt.Println("tidak dieksekusi")
}