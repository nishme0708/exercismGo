package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
import "fmt"
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}
type Italian struct{}

func (italian Italian) LanguageName() string {
	return "Italian"
}
func (italian Italian) Greet(name string) string {
	return "Ciao " + name
}

type Portuguese struct{}

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}
func (portuguese Portuguese) Greet(name string) string {
	return "Olá " + name
}



func SayHello(visitor string, hindiGreeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", hindiGreeter.LanguageName(), hindiGreeter.Greet(visitor))
}
