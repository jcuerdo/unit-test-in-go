package main

import "github.com/jcuerdo/unit-test-in-go/process"

func main(){
	subProcess := &process.SubProcess{}
	pr := process.Process{SubProcess:subProcess}

	pr.RunProcess(5)
}
