package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

type HealthProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[plant.HealthData]
	dronesRepo droneRepository.DroneRepo
}

func NewHealthProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[plant.HealthData],
	dronesRepo droneRepository.DroneRepo,
) *HealthProcessor {
	return &HealthProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *HealthProcessor) RunProcessor(ctx context.Context) {
	//TODO: implement process functionality
}
