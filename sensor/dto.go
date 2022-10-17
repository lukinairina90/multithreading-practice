package sensor

type SensorData[T any] struct {
	PlantID string
	Data    T
}
