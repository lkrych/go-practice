package main

import "fmt"

//interfaces are a type in go
//they are a language construct that defines a functionality
//they do this by encapsulating a bunch of methods
type test interface {
	SayHello()
	Say(s string)
	Increment()
	GetValue()
}

type testInterfaceImpl struct {
	i int
}

//constructor
func newTestInterfaceImpl() *testInterfaceImpl {
	return new(testInterfaceImpl)
}

func newTestInterfaceWithValue(v int) *testInterfaceImpl {
	return &testInterfaceImpl{i: v}
}

func (test *testInterfaceImpl) SayHello() {
	fmt.Println("Hello!")
}

func (test *testInterfaceImpl) Say(s string) {
	fmt.Println(s)
}

func (test *testInterfaceImpl) Increment() {
	test.i++
}

func (test *testInterfaceImpl) GetValue() {
	fmt.Println(test.i)
}

//embedding is a way to encapsulate an object within another object
type testEmbedding struct { //we want this struct to have all features of testInterfaceImpl
	*testInterfaceImpl //embedding, I guess this is kind of like inheritance?
}

func testEmptyInterface(v interface{}) {
	// if i, ok := v.(int); ok { //assert type
	// 	fmt.Println(i)
	// } else {
	// 	fmt.Println("I am not an integer type!!")
	// }

	//an alternative is to use a type switch

	switch val := v.(type) {
	case int:
		fmt.Println("I am an int", val)
	case string:
		fmt.Println("I am a string", val)
	default:
		fmt.Println("I am neither an int or a string", val)
	}
}

func main() {

	//interface

	var tiface test
	tiface = newTestInterfaceWithValue(5) //&testInterfaceImpl{} // or you can use the new keyword new(testInterfaceImpl)
	tiface.SayHello()
	tiface.Say("blah blah blah")
	tiface.GetValue()
	tiface.Increment()
	tiface.GetValue()
	tiface.Increment()
	tiface.GetValue()

	// embedding

	te := testEmbedding{testInterfaceImpl: newTestInterfaceWithValue(50)}
	te.SayHello()
	te.GetValue()
	te.Increment()
	te.GetValue()
	te.Increment()
	te.GetValue()

	//empty interface

	testEmptyInterface(3) //an empty interface encapsulates all types in Go
	testEmptyInterface("Hammerhead")
	testEmptyInterface(tiface)
	testEmptyInterface(te)

}
