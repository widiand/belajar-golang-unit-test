package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("before")

	m.Run()

	fmt.Println("after")
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("Khannedy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Khannedy")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	tests := []string {
		"Eko",
		"Khannedy",
		"Budi",
	}

	for _, test := range tests {
		b.Run(test, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test)
			}
		})
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Eko)",
			request: "Eko",
			expected: "Hello Eko",
		},
		{
			name: "HelloWorld(Khannedy)",
			request: "Khannedy",
			expected: "Hello Khannedy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "harusnya Hello Eko")
	})
	t.Run("Khannedy", func(t *testing.T) {
		result := HelloWorld("Khannedy")
		require.Equal(t, "Hello Khannedy", result, "harusnya Hello Khannedy")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")//wil skip the test below
	}

	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Ekoo", result, "harusnya Hello Eko")//Fail if fail (still continnue)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "harusnya Hello Eko")//FailNow if fail (not continue)
}

func TestHelloWOrldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "harusnya Hello Eko")//Fail if fail (still continnue)
}

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