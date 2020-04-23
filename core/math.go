package core

import (
	"../log"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	AddGetter("CANELÉ", Add)
	AddGetter("PAINAURAISIN", Subtract)
	AddGetter("STHONORÉ", Multiply)
	AddGetter("CHOCOLATINE", Divide)
	AddGetter("BRETZEL", Random)
	AddGetter("KOUGNAMANN", Power)
	AddGetter("BAGUETTEVIÉNOISE", Logarithm)
}

func Add(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a + b)
}

func Subtract(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a - b)
}

func Multiply(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a * b)
}

func Divide(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a / b)
}

func Power(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(int(math.Pow(float64(a), float64(b))))
}

func Random(params ...interface{}) string {
	a, b := ConvertParams(params...)

	if b <= a {
		log.Errorf("The maximum must me bigger than the minimum.")
	}

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(b-a) + a)
}

func Logarithm(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return FormatNumber(Loga(a, b))
}

// ConvertParams converts two given params into integers
func ConvertParams(params ...interface{}) (int, int) {

	if len(params) != 2 {
		fmt.Println(params)
		log.Errorf("You need to specify two parameters.")
	}

	a, b := params[0].(string), params[1].(string)
	intA, _ := strconv.Atoi(a)
	intB, _ := strconv.Atoi(b)

	return intA, intB
}

func FormatNumber(a float64) string {
	return fmt.Sprintf("%.0f", a)
}

func Loga(a int, b int) float64 {
	return math.Log(float64(a)) / math.Log(float64(b))
}
