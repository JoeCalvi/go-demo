package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Ninja struct {
		Name string `json:"full_name"`
		Weapon Weapon `json:"weapon"`
		Level int `json:"level"`
	}

type Weapon struct {
	Name string `json:"weapon_name"`
	Level int `json:"weapon_level"`
}

func main()  {
	fmt.Println("Hello world!")

	// var name type = expression
	var integer int = 1
	fmt.Println(integer)
	fmt.Println()

	// SECTION declaring variables
	// either the type or the expression can be ommitted
	// multiple variables can be declared in a single line
	var integer1, integer2 = 1, 2
	fmt.Println(integer1, integer2)
	fmt.Println()

	// or you can use "short declaration"
	// name := expression
	boolean := false
	fmt.Println(boolean)
	fmt.Println()

	// SECTION pointers
	x := 1
	p := &x

	// points to the ADDRESS of variable x
	fmt.Println(p)
	fmt.Println()
	// points to the actual VALUE of variable x
	fmt.Println(*p)
	fmt.Println()

	// SECTION types
	// fahrenheit and celsius

	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	// converts fahrenheit to celsius
	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)
	fmt.Println()

	// SECTION floating point numbers
	var pi float32 = 3.141
	fmt.Println(pi)
	fmt.Println()
	// prints type of var pi
	fmt.Printf("%T", pi)
	fmt.Println()

	// SECTION math library options
	fmt.Println(math.Ceil(3.6)) // 4
	fmt.Println(math.Floor(3.003)) // 3
	fmt.Println(math.Min(1, 0)) // 0
	fmt.Println(math.Max(1, 0)) // 1
	fmt.Println(math.Abs(-1)) // 1 (absolute value)
	fmt.Println(math.Sqrt(100)) // 10 (square root)
	fmt.Println(math.Pow(2, 3)) // 8 (2 to the 3rd power)

	// SECTION strings
	s := "hello world"
	fmt.Println(s)
	fmt.Println(len(s)) // 11
	fmt.Println(s[0]) // value of the character "h"
	fmt.Printf("%c", s[0]) // the character "h"
	fmt.Println(s[0:6]) // "hello"

	fmt.Println(strings.ToUpper("Hello world")) // HELLO WORLD
	fmt.Println()
	fmt.Println(strings.ToLower("HELLO WORLD")) // hello world
	fmt.Println()

	var sb strings.Builder
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println()
	sb.WriteString("Hello")
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println()

	z := 123
	y := strconv.Itoa(z) //converts to string
	fmt.Println(y)
	fmt.Println()

	// SECTION constants
	const five int = 5
	fmt.Println(five)
	fmt.Println()

	const (
		a = 1
		b = 2
		d = 4
		e = 5
	)
	fmt.Println(a, b, d, e)
	fmt.Println()

	// SECTION for loops
	
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println()
	}

	// how to write a for each
	hw := "Hello World"
	for i, c := range hw {
		fmt.Printf("%d %c", i, c)
		fmt.Println()
	}

	// SECTION arrays
	var array [3]int
	fmt.Println(array)
	fmt.Println()

	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)
	fmt.Println()

	array2 := [4]int{1, 2, 3, 4}
	fmt.Println(array2)
	fmt.Println()

	// SECTION structs 
	type ninja struct {
		name string
		weapons []string
		level int
	}

	wallace := ninja{name: "Wallace"}
	wallace = ninja{
		name: "Wallace",
		weapons: []string{"Ninja Star"},
		level: 1,
	}
	fmt.Println(wallace)
	fmt.Println()
	fmt.Println(wallace.name)
	fmt.Println()
	fmt.Println(wallace.weapons[0])
	fmt.Println()
	fmt.Println(wallace.level)
	fmt.Println()
	wallace.weapons = append(wallace.weapons, "Kitana")
	fmt.Println(wallace.weapons)
	fmt.Println()

	type dojo struct {
		name string
		ninja ninja
	}

	golangDojo := dojo{
		name: "Golang Dojo",
		ninja: wallace,
	}
	fmt.Println(golangDojo)
	fmt.Println()
	fmt.Println(golangDojo.ninja.name)
	fmt.Println()

	type addressedDojo struct {
		name string
		ninja *ninja 
	}
	addressed := addressedDojo{"Addressed Golang Dojo", &wallace}
	// prints the ADDRESS of wallace
	fmt.Println(addressed)
	fmt.Println()
	// prints wallace's info
	fmt.Println(*addressed.ninja)
	fmt.Println()

	// SECTION maps
	var m map[string]string
	fmt.Println(m == nil)
	fmt.Println()
	m = make(map[string]string, 5)
	m = map[string]string{"Wallace": "Programmer"}
	fmt.Println(m)
	fmt.Println()
	m["Johnny"] = "Not a programmer"
	fmt.Println(m)
	fmt.Println()

	// delete(m, "Johnny")
	// fmt.Println(m)
	// fmt.Println()

	for name, title := range m {
		fmt.Println(name, title)
		fmt.Println()
	}

	// SECTION handling JSONs

	sFrom := `{"full_name": "Wallace", "weapon": {"weapon_name": "Ninja Star", "weapon_level": 1}, "level": 1}`
	fmt.Println(sFrom)
	fmt.Println()

	var wallaceJSON Ninja
	err := json.Unmarshal([]byte(sFrom), &wallaceJSON)
	if err != nil {
		fmt.Println("Sad boy.")
		fmt.Println(err)
		fmt.Println()
	}
	fmt.Println(wallaceJSON)
	fmt.Println()


	sTo, err := json.Marshal(wallaceJSON)
	if err != nil {
		fmt.Println("Sad boy.")
		fmt.Println(err)
		fmt.Println()

	}
	fmt.Printf("%s\n", sTo)
	fmt.Println()

}	
