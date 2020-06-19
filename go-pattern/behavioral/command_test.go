package behavioral

import "testing"

func TestCommand(t *testing.T) {
	//init receiver
	air := NewAirCond()
	light := NewLight()
	//new cmd
	lightOn := NewLightOnCMD(light)
	lightOff := NewLightOffCMD(light)
	airOn := NewAirCondOnCMD(air)
	airOff := NewAirCondOffCMD(air)
	//add cmd
	invoker := NewInvoker()
	invoker.addOnCMD(airOn, lightOn)
	invoker.addOffCMD(airOff, lightOff)
	invoker.execOn()
	invoker.execOff()
}
