package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-joystick"
)

func main() {
	joystickAdaptor := new(gobotJoystick.JoystickAdaptor)
	joystickAdaptor.Name = "x52"
	joystickAdaptor.Params = map[string]interface{}{
		"config": "./saitek-x52-reverse.json",
	}

	joystick := gobotJoystick.NewJoystick(joystickAdaptor)
	joystick.Name = "x52"

	work := func() {
		/*
			gobot.On(joystick.Events["left_throttle"], func(data interface{}) {
				fmt.Println("left_throttle", data)
			})
			gobot.On(joystick.Events["left_wheel"], func(data interface{}) {
				fmt.Println("left_wheel", data)
			})
			gobot.On(joystick.Events["right_x"], func(data interface{}) {
				fmt.Println("right_x", data)
			})
			gobot.On(joystick.Events["right_y"], func(data interface{}) {
				fmt.Println("right_y", data)
			})
		*/
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{joystickAdaptor},
		Devices:     []gobot.Device{joystick},
		Work:        work,
	}

	robot.Start()
}
