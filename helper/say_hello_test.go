package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before..")

	m.Run()

	fmt.Println("After...")

}
func TestSayHello(t *testing.T) {
	result := SayHelloTo("Fazrul")
	assert.Equal(t, "Hello Fazrul", result, "Result must be `Hello Fazrul`")
	fmt.Println("Execute")

}

func TestSayHelloFail(t *testing.T) {
	result := SayHelloTo("Fazrul")
	require.Equal(t, "Hello Fazrul", result, "Result Must be `Hello Fazrul`")
	fmt.Println("Execute 2")
}

func TestSubSayHello(t *testing.T) {
	myName := InsertName("Fazrul")
	t.Run("Sub Test Success", func(t *testing.T) {
		result := SayHelloTo(myName)
		require.Equal(t, "Hello Fazrul", result, "Result Must be `Hello Fazrul`")
	})

	t.Run("Sub Test Fail", func(t *testing.T) {
		result := SayHelloTo(myName)
		require.NotEqual(t, "Hello Ajul", result, "Result Must not be `Hello Fazrul`")
	})
}

func TestTableTestSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "HelloWorld(Fazrul)", request: "Fazrul", expected: "Hello Fazrul"},
		{name: "HelloWorld(John)", request: "John", expected: "Hello John"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			myName := InsertName(test.request)
			result := SayHelloTo(myName)
			require.Equal(t, test.expected, result, "Result Must not be "+test.expected)

		})

	}
}
