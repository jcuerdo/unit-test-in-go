package calculator

import (
	"fmt"
)

type Calculator struct{

}

func (calculator Calculator) Division(x int64, y int64) (int64, error) {

	if y == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return x/y , nil
}