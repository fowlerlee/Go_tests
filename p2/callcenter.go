package main

import (
	"fmt"
)

type Call struct {
	callerNumber int64
	duration     int64
}

type CallCenter struct {
	callsInQueue map[int64]string
	operator     FirstLineSupport
}

func NewCallCenter() *CallCenter {
	return &CallCenter{
		callsInQueue: make(map[int64]string),
		operator:     "lee",
	}
}

func (c *CallCenter) HandleCall(operator FirstLineSupport, call Call) {
	if operator.Away() {
		c.AddToQueue(call)
		fmt.Printf("call: %d is in the waiting line", call.callerNumber)
	} else {
		fmt.Printf("call: %d has been handled by: %s\n", call.callerNumber, c.operator)
	}
}

func (callCenter *CallCenter) AddToQueue(call Call) {
	keys := make([]int64, 0, len(callCenter.callsInQueue))
	for k := range callCenter.callsInQueue {
		keys = append(keys, k)
	}

}

type FirstLineSupport string

func (operator FirstLineSupport) OperatorStatusChanged(status bool) bool {
	if status != operator.Away() {
		return operator.OnLine()
	} else {
		return operator.Away()
	}
}

func (operator FirstLineSupport) OnLine() bool {
	return true
}

func (operator FirstLineSupport) Away() bool {
	return false
}

func main() {
	var lee FirstLineSupport
	lee = "lee"
	lee.OperatorStatusChanged(lee.OnLine())
	call1 := &Call{1210021, 1}
	lee.OperatorStatusChanged(lee.Away())
	call2 := &Call{1321001, 1}
	lee.OperatorStatusChanged(lee.OnLine())
	call3 := &Call{1431331, 1}
	callCenter := *NewCallCenter()
	fmt.Printf("the call1: %+v to %+v at call-center: %#v\n", call1.callerNumber, lee, callCenter.callsInQueue)

	callCenter.HandleCall(lee, *call1)
	callCenter.HandleCall(lee, *call2)
	callCenter.HandleCall(lee, *call3)
	fmt.Printf("the calls handled: %d\n", callCenter.callsInQueue)
}
