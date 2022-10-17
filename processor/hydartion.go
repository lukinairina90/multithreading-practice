package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

type HydrationProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[float64]
	dronesRepo droneRepository.DroneRepo
}

func NewHydrationProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[float64],
	dronesRepo droneRepository.DroneRepo,
) *HydrationProcessor {
	return &HydrationProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *HydrationProcessor) RunProcessor(ctx context.Context) {
	//TODO: implement process functionality
}
