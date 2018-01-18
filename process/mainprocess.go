package process


import (
	"fmt"
)

type Process struct{
	SubProcess SubProcessInterface
}

func (process *Process) RunProcess(times int) (Result, error){
	fmt.Print("Hi, im a process ¡¡¡¡")
	for i := 0 ; i < times; i++ {
		_, error := process.SubProcess.RunSubProcess(i)
		if error != nil{
			return Result{}, fmt.Errorf("We have a problem")
		}
	}

	return Result{"All went ok"}, nil
}
