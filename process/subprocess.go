package process


import (
	"fmt"
)

type SubProcessInterface interface{
	RunSubProcess(id int)(Result, error)
}

type SubProcess struct{

}

func (subProcess *SubProcess) RunSubProcess(id int) (Result, error){
	fmt.Print("Hi, im a subprocess ¡¡¡¡")

	return Result{}, nil
}
