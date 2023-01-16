package lib

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

// These two come from a different package
type One struct{}
type Two struct{}
type Three struct{}

// With this method does not run
func (o *One) Method1(t *Two) {}

// With this method runs, but I can modify the firm of the method because is in another package
// that my code has a dependency
//func (o *One) Method1(t Two) {}

// For testing purposes just comment out the first Method1 and use the second one
// and you can see that the code runs without panic

// In my project I do this
// types.go
type typeRegister map[string]reflect.Type

func (t typeRegister) Set(i interface{}) {
	typ := reflect.TypeOf(i).Elem()
	t[typ.Name()] = typ
}

func (t typeRegister) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		return reflect.New(typ).Elem().Interface(), nil
	}

	return nil, errors.New("not valid type registered: " + name)
}

var typeRegistry = make(typeRegister)

func init() {
	// These two because they come from another package
	// Something like:
	// typeRegistry.Set(new(another_package.One))
	// typeRegistry.Set(new(another_package.Two))

	typeRegistry.Set(new(Two))
	typeRegistry.Set(new(Three))
	runtime.GC()
}

// main.go
func ReflectOutput() {
	// Dynamically retrieve the string "Two"
	// that is the name of the struct to initialize
	// in this example a string declarated here
	structNameThatIsParameter := "Two"

	// Dynamically retrieve the string "Method1"
	// that is the name of the method to call on the
	// structName value retrieved dynamically
	// here, hard coded
	methodOfStructInstantiated := "Method1"

	i, err := typeRegistry.Get(structNameThatIsParameter)
	if err != nil {
		// Handle the error, but here just panic
		panic(err)
	}

	// I actually call this because I know what struct I want to call
	// In the actual code this is something like:
	// instance := new(other_package.One)
	instance := new(One)

	method := reflect.ValueOf(instance).MethodByName(methodOfStructInstantiated)

	// I can't call like this because the firm of the Method1 that is run against
	// a pointer to the struct Two
	// func (o *One) Method1(t *Two) {} in the package that I'm using
	// if the firm of that method is: func (o *One) Method1(t Two) {}
	// without a pointer but a struct the method beneath runs
	called := method.Call([]reflect.Value{reflect.ValueOf(i)})

	fmt.Printf("%+v\n", called)
}
