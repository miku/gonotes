package main

type Greeter interface {
	Greet() string
}

type English struct{}

func (e *English) Greet() string {
	return "Hello"
}

type German struct{}

func (g German) Greet() string {
	return "Guten Tag"
}

func main() {
	// var e Greeter = English{}
	// var ep Greeter = &English{}
	// var g Greeter = German{}
	// var gp Greeter = &German{}

	// log.Println(e, ep, g, gp)
}
