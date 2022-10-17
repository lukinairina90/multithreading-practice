package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

type PHProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[int]
	dronesRepo droneRepository.DroneRepo
}

func NewPHProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[int],
	dronesRepo droneRepository.DroneRepo,
) *PHProcessor {
	return &PHProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *PHProcessor) RunProcessor(ctx context.Context) {
	//TODO: implement process functionality
}
