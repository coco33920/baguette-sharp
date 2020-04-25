package core

import (
	"../log"
	"../utils"
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
	AddGetter("FINANCIER", Fibonacci)
	AddGetter("PROFITEROLES", Sqrt)
	AddGetter("OPERA", Negative)
}

func Negative(params ...interface{}) string {
	a := ConvertNParams(params...)[0]
	return FormatNumber(float64(-a))
}

func Sqrt(params ...interface{}) string {
	a := ConvertNParams(params...)
	b := a[0]
	if b < 0 {
		return "Not real"
	}
	return FormatNumber(math.Sqrt(float64(b)+0.00))
}

func Add(params ...interface{}) string {
	a := ConvertNParams(params...)
	sum := 0
	for _, value := range a {
		sum += value
	}
	return strconv.Itoa(sum)
}

func Subtract(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a - b)
}

func Multiply(params ...interface{}) string {
	a := ConvertNParams(params...)
	total := 1
	for _, value := range a {
		total *= value
	}
	return strconv.Itoa(total)
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

func ConvertNParams(params ...interface{}) []int {
	var a []int
	i := 0
	for i < (len(params)) {
		intA, _ := strconv.Atoi(params[i].(string))
		a = append(a, intA)
		i++
	}
	return a
}

func FormatNumber(a float64) string {
	return fmt.Sprintf("%.0f", a)
}

func Loga(a int, b int) float64 {
	return math.Log(float64(a)) / math.Log(float64(b))
}

func Fibonacci(params ...interface{}) string {
	a := ConvertNParams(params...)[0]
	return strconv.Itoa(utils.Fibonacci(a))
}
