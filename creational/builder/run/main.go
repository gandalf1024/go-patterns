package main

import "go-patterns/creational/builder/car"

//链式编程 -- 调用函数 返回对象本身
func main() {
	assembly := car.NewCarBuilder().Paint(car.RedColor)

	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Drive()
}
