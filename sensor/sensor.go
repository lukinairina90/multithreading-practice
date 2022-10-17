package sensor

type Sensor[T any] struct {
	ID      string
	PlantID string
	ch      chan SensorData[T]
}

func (s *Sensor[T]) Connect() chan SensorData[T] {
	//Some connection work...
	return s.ch
}

func (s *Sensor[T]) Disconnect() {
	close(s.ch)
}

func NewSensor[T any](ID, plantID string) *Sensor[T] {
	return &Sensor[T]{
		ID:      ID,
		PlantID: plantID,
		ch:      make(chan SensorData[T], 2),
	}
}
