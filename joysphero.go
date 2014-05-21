package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-joystick"
	"github.com/hybridgroup/gobot-sphero"
	"math"
)

func main() {
	joystickAdaptor := new(gobotJoystick.JoystickAdaptor)
	joystickAdaptor.Name = "x52"
	joystickAdaptor.Params = map[string]interface{}{
		"config": "./saitek-x52.json",
	}

	joystick := gobotJoystick.NewJoystick(joystickAdaptor)
	joystick.Name = "x52"

	spheroAdaptor := new(gobotSphero.SpheroAdaptor)
	spheroAdaptor.Name = "Sphero"
	spheroAdaptor.Port = "/dev/tty.Sphero-PWG-RN-SPP"

	sphero := gobotSphero.NewSphero(spheroAdaptor)
	sphero.Name = "Sphero"

	speed := 0.0
	direction := 0.0

	work := func() {
		gobot.On(joystick.Events["right_x"], func(data interface{}) {
			direction = math.Abs(float64(data.(int16) / 128)) //255
		})
		gobot.On(joystick.Events["right_y"], func(data interface{}) {
			speed = math.Abs(float64(data.(int16) / 91)) // 360
		})
		gobot.Every("0.01s", func() {
			fmt.Println(speed, direction)
			sphero.Roll(uint8(speed), uint16(direction))
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{joystickAdaptor, spheroAdaptor},
		Devices:     []gobot.Device{joystick, sphero},
		Work:        work,
	}

	robot.Start()
}
