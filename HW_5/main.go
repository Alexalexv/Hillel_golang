package main

import (
	"errors"
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type Generator interface {
	Generate() interface{}
	GenerateSlice() interface{}
	GenerateWithParam(param int) (interface{}, error)
}

//integer

type intGenerator struct{}

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
	integer := rand.Intn(param + 1)
	log.Infof("new int has been created: %d", integer)
	return integer, nil
}

//string

type strGenerator struct{}

func (sg strGenerator) Generate() interface{} {
	return "_"
}

func (sg strGenerator) GenerateSlice() interface{} {
	return []string{"_"}
}

func (sg strGenerator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 {
		return nil, errors.New("param can`t be less than zero")
	}
	var str string = ""
	for i := 0; i < param; i++ {
		str += "_"
	}
	return str, nil
}

//bool

type boolGenerator struct{}

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

type floatGenerator struct{}

func (fg floatGenerator) Generate() interface{} {
	return 0.1
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

func SomeGenerator(generator Generator, param int) {
	fmt.Println("Your default value is:", generator.Generate())
	fmt.Println("Your slice value is:", generator.GenerateSlice())
	result, err := generator.GenerateWithParam(param)
	logger, _ := zap.NewProduction()
	if err != nil {
		logger.Sugar().Warnf("something went wrong - %s", err)
		return
	}
	resultInt, ok := result.(int)
	if ok {
		fmt.Printf("It`s definitely int, you generated value is %d\n", resultInt)
	}
	resultStr, ok := result.(string)
	if ok {
		fmt.Printf("It`s definitely string, you generated value is %s\n", resultStr)
	}
	resultBool, ok := result.(bool)
	if ok {
		fmt.Printf("It`s definitely string, you generated value is %t\n", resultBool)
	}
	resultFloat, ok := result.(float32)
	if ok {
		fmt.Printf("It`s definitely float32, you generated value is %.1f\n", resultFloat)
	}

}

func main() {

	integer := intGenerator{}
	SomeGenerator(integer, 123)

	str := strGenerator{}
	SomeGenerator(str, 25)

	boolVal := boolGenerator{}
	SomeGenerator(boolVal, 1)

	float := floatGenerator{}
	SomeGenerator(float, 125)

}
