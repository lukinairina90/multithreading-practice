package repository

import "github.com/SchoolGolang/multithreading-practice/sensor"

type SensorRepo[T any] struct {
	sensors map[string]sensor.Sensor[T]
}

func NewRepository[T any]() *SensorRepo[T] {
	return &SensorRepo[T]{
		sensors: make(map[string]sensor.Sensor[T]),
	}
}

func (r *SensorRepo[T]) AddSensor(sensor *sensor.Sensor[T]) {
	if _, ok := r.sensors[sensor.ID]; ok {
		return
	}
	r.sensors[sensor.ID] = *sensor
}

func (r *SensorRepo[T]) GetSensor(id string) sensor.Sensor[T] {
	return r.sensors[id]
}

func (r *SensorRepo[T]) GetSensorByPlantID(plantID string) *sensor.Sensor[T] {
	for _, v := range r.sensors {
		if v.PlantID == plantID {
			return &v
		}
	}
	return nil
}

func (r *SensorRepo[T]) RemoveSensorByPlantID(plantID string) {
	for key := range r.sensors {
		s := r.sensors[key]
		if s.PlantID == plantID {
			s.Disconnect()
			delete(r.sensors, key)
		}
	}
}

func (r *SensorRepo[T]) GetAll() []sensor.Sensor[T] {
	sensors := make([]sensor.Sensor[T], 0, len(r.sensors))

	for _, item := range r.sensors {
		sensors = append(sensors, item)
	}

	return sensors
}
