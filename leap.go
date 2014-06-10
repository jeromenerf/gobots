package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-leapmotion"
)

func main() {
	leapAdaptor := new(gobotLeap.LeapAdaptor)
	leapAdaptor.Name = "leap"
	leapAdaptor.Port = "127.0.0.1:6437"

	leap := gobotLeap.NewLeap(leapAdaptor)
	leap.Name = "leap"

	work := func() {
		gobot.On(leap.Events["Message"], func(data interface{}) {
			printGestures(data.(gobotLeap.LeapFrame))
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{leapAdaptor},
		Devices:     []gobot.Device{leap},
		Work:        work,
	}

	robot.Start()
}

func printGestures(frame gobotLeap.LeapFrame) {
	for key, gesture := range frame.Gestures {
		fmt.Println("Gesture", key, gesture)
	}
}
