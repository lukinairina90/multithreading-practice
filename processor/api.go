package processor

import "github.com/SchoolGolang/multithreading-practice/sensor"

type Processor[U, T sensor.SensorData[U]] interface {
	ProcessMessage(T)
}
