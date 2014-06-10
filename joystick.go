package main

import (
	"fmt"
	"math"

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
	offset := 32767.0

	work := func() {
		gobot.On(joystick.Events["left_throttle"], func(data interface{}) {
			val := float64(data.(int16))
			fmt.Println("left_throttle", validatePitch(val, offset), val)
		})
		gobot.On(joystick.Events["right_rotate"], func(data interface{}) {
			val := float64(data.(int16))
			fmt.Println("right_rotate", validatePitch(val, offset), val)
		})
		gobot.On(joystick.Events["right_x"], func(data interface{}) {
			val := float64(data.(int16))
			fmt.Println("right_x", validatePitch(val, offset), val)
		})
		gobot.On(joystick.Events["right_y"], func(data interface{}) {
			val := float64(data.(int16))
			fmt.Println("right_y", validatePitch(val, offset), val)
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{joystickAdaptor},
		Devices:     []gobot.Device{joystick},
		Work:        work,
	}

	robot.Start()
}

func validatePitch(data float64, offset float64) float64 {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return float64(int(value*100)) / 100
		} else {
			return 1.0
		}
	} else {
		return 0.0
	}
}
