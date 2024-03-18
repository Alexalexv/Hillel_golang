package main

import (
	"errors"
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

func (ig intGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 {
		return nil, errors.New("param can`t be less than zero")
	}
	return rand.Intn(param + 1), nil
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

func (sg strGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 {
		return nil, errors.New("param can`t be less than zero")
	}
	var str string = ""
	for i := 0; i < param; i++ {
		str += " "
	}
	return str, nil
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

func (bg boolGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param == 0 {
		return false, nil
	}
	if param == 1 {
		return true, nil
	}
	return nil, errors.New("in bool generator param 1 or 2")
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

func (fg floatGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 {
		return nil, errors.New("param can`t be less than zero")
	}
	random := float32(rand.Intn(param))
	random = random + 0.1
	return random, nil
}

func GenerateSome(generator Generator) interface{} {
	return generator.Generate()
}

func main() {

}
