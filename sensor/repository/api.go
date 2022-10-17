package repository

import "github.com/SchoolGolang/multithreading-practice/sensor"

type Repository[T any] interface {
	AddSensor(sensor *sensor.Sensor[T])
	GetSensor(id string) sensor.Sensor[T]
	RemoveSensorByPlantID(plantID string)
	GetAll() []sensor.Sensor[T]
}
