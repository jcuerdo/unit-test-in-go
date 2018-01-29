package process_test


import (
	"github.com/stretchr/testify/mock"
	"github.com/jcuerdo/unit-test-in-go/process"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type SubProcessMock struct {
	mock.Mock
}

func (subProcess *SubProcessMock) RunSubProcess(id int) (process.Result, error){
	args := subProcess.Called(id)
	return args.Get(0).(process.Result), args.Error(1)
}


func TestProcessOneTime(t *testing.T){
	subProcessMock := &SubProcessMock{}

	pr :=  &process.Process{SubProcess:subProcessMock}

	//PHP
	// $this->expects($this->once())->method("RunSubProcess")->with(0)->willReturn(new Result())

	subProcessMock.On("RunSubProcess",0).Return(process.Result{}, nil)

	result , error := pr.RunProcess(1)

	subProcessMock.AssertCalled(t,"RunSubProcess", 0)
	subProcessMock.AssertNumberOfCalls(t,"RunSubProcess", 1)

	assert.Equal(t,result,process.Result{"All went ok"})
	assert.Nil(t,error)

}

func TestProcessTwoTimesWithError(t *testing.T){
	subProcessMock := &SubProcessMock{}

	pr :=  &process.Process{SubProcess:subProcessMock}

	subProcessMock.On("RunSubProcess",0).Return(process.Result{}, nil).Times(1)
	subProcessMock.On("RunSubProcess",1).Return(process.Result{}, fmt.Errorf("Something went wrong")).Times(1)

	result , error := pr.RunProcess(3)

	subProcessMock.AssertCalled(t,"RunSubProcess",0)
	subProcessMock.AssertCalled(t,"RunSubProcess",1)
	subProcessMock.AssertNotCalled(t,"RunSubProcess",2)
	subProcessMock.AssertNumberOfCalls(t,"RunSubProcess", 2)

	assert.Equal(t,result,process.Result{})
	assert.EqualError(t,error,"We have a problem")
}