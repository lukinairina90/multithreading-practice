package processor

import (
	"context"
	"fmt"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

const (
	leavesHealthThreshold = 50
	rootHealthThreshold   = 50
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
	for {
		select {
		case healthMsg := <-p.input:
			if healthMsg.PlantID == "" {
				continue
			}

			p.plantsRepo.SetHealth(healthMsg.PlantID, plant.HealthData{
				LeavesState: healthMsg.Data.LeavesState,
				RootsState:  healthMsg.Data.RootsState,
			})

			hd := p.plantsRepo.GetHealth(healthMsg.PlantID)

			pl := p.plantsRepo.GetPlant(healthMsg.PlantID)
			switch {
			case hd.LeavesState < leavesHealthThreshold:
				p.dronesRepo.ReplacePlant(healthMsg.PlantID)

				fmt.Printf("Plant %s with ID %s | The leaf quality of the plant has fallen below 50 percents. The plant has been replaced. \n", pl.Name, pl.ID)
			case hd.RootsState < rootHealthThreshold:
				p.dronesRepo.ReplacePlant(healthMsg.PlantID)

				fmt.Printf("Plant %s with ID %s | The root quality of the plant has fallen below 50 percents. The plant has been replaced. \n", pl.Name, pl.ID)
			default:
				fmt.Printf("Plant %s with ID %s | Health is OK\n", pl.Name, pl.ID)
			}

		case <-ctx.Done():
			return

		}
	}

}
