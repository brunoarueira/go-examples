package main

import (
        "fmt"
)

type Person struct {
        name string
        age int
        profession string
}

func (p *Person) greetings() string {
        return fmt.Sprintf("Hi my name is %v, I'm %v years old and I'm %v", p.name, p.age, p.profession)
}

func main() {
        john := Person{ name: "John", age: 28, profession: "Programmer" }

        fmt.Println(john.greetings())
}
