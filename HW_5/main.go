package main

import (
	"fmt"
	"math/rand"
)

type Generator interface {
	Generate() interface{}
}

type GeneratorSlice interface {
	GenerateSlice() interface{}
}

type GenerorWithParam interface {
	GenerateWithParam(param int) interface{}
}

//integer

type intGenerator struct {
	val int
}

func (ig intGenerator) Generate() interface{} {
	return 0
}

func (ig intGenerator) GenerateSlice() interface{} {
	return []int{0}
}

func (ig intGenerator) GenerateWithParam(param int) interface{} {
	return rand.Intn(param + 1)
}

//string

type strGenerator struct {
	val string
}

func (sg strGenerator) Generate() interface{} {
	return " "
}

func (sg strGenerator) GenerateSlice() interface{} {
	return []string{" "}
}

func (sg strGenerator) GenerateWithParam(param int) interface{} {
	var str string = ""
	for i := 0; i < param; i++ {
		str += " "
	}
	return str
}

//bool

type boolGenerator struct {
	val bool
}

func (bg boolGenerator) Generate() interface{} {
	return false
}

func (bg boolGenerator) GenerateSlice() interface{} {
	return []bool{false}
}

func (bg boolGenerator) GenerateWithParam(param int) interface{} {
	if param == 0 {
		return false
	}
	if param == 1 {
		return true
	}
	return true
}

//float

type floatGenerator struct {
	val float32
}

func (fg floatGenerator) Generate() interface{} {
	return 0.01
}

func (fg floatGenerator) GenerateSlice() interface{} {
	return []float32{0.1}
}

func (fg floatGenerator) GenerateWithParam(param int) interface{} {
	random := float32(rand.Intn(param))
	random = random + 0.1
	return random
}

func GenerateSome(generator Generator) interface{} {
	return generator.Generate()
}

func main() {
	one := floatGenerator{}
	a := one.GenerateWithParam(150)
	fmt.Printf("%dasd\n", a)

}
