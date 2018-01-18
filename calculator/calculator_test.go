package calculator_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jcuerdo/unit-test-in-go/calculator"
)


func TestDivisionByZero(t *testing.T) {

	calculator := calculator.Calculator{}
	_, err := calculator.Division(5,0)

	assert.Error(t,err,"")
}


func TestDivision(t *testing.T) {

	var expected int64 = 2
	calculator := calculator.Calculator{}
	result, _ := calculator.Division(5,2)

	assert.Equal(t,expected,result)
}


func TestBigDivision(t *testing.T) {

	var expected int64 = 2
	calculator := calculator.Calculator{}
	result, _ := calculator.Division(50000000000,20000000000)

	assert.Equal(t,expected,result)
}