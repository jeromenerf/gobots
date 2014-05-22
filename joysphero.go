package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-joystick"
	"github.com/hybridgroup/gobot-sphero"
	"math"
)

// Angle
func Angle(x, y int) uint16 {
	angle := math.Atan2(float64(x), float64(y)) / (math.Pi * 2) * 360
	direction := math.Mod((360 + 180 + angle), 360)
	return uint16(direction)
}

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
	var (
		x, y      int
		direction uint16
		speed     uint8
	)
	x = 1
	y = 1
	direction = 0
	speed = 0

	work := func() {
		gobot.On(joystick.Events["right_x"], func(data interface{}) {
			x = int(data.(int16))
		})
		gobot.On(joystick.Events["right_y"], func(data interface{}) {
			y = int(data.(int16))
		})
		gobot.Every("0.01s", func() {
			direction = Angle(x, y)
			speed = uint8(math.Sqrt(float64(x*x+y*y)) / 128) //255
			fmt.Println(x, y, speed, direction)
			sphero.Roll(speed, direction)
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{joystickAdaptor, spheroAdaptor},
		Devices:     []gobot.Device{joystick, sphero},
		Work:        work,
	}

	robot.Start()
}
