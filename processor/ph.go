package processor

import (
	"context"
	"fmt"
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
	for {
		select {
		case phMsg := <-p.input:
			if phMsg.PlantID == "" {
				continue
			}

			p.plantsRepo.SetPh(phMsg.PlantID, phMsg.Data)
			nLph, nUph := p.plantsRepo.GetNormalPh(phMsg.PlantID)

			plant := p.plantsRepo.GetPlant(phMsg.PlantID)

			switch {
			case phMsg.Data < nLph:
				s := (nLph + nUph) / 2
				p.dronesRepo.AdjustSoils(phMsg.PlantID, s)

				fmt.Printf("Plant %s with ID %s | The PH level fell below the acceptable minimum norm %d to %d. PH level restored to %d\n", plant.Name, plant.ID, nLph, phMsg.Data, s)
			case phMsg.Data > nUph:
				s := (nLph + nUph) / 2
				p.dronesRepo.AdjustSoils(phMsg.PlantID, s)

				fmt.Printf("Plant %s with ID %s | The PH level has exceeded the acceptable maximum rate %d to %d. PH level restored to %d\n", plant.Name, plant.ID, nUph, phMsg.Data, s)
			default:
				fmt.Printf("Plant %s with ID %s | PH is OK\n", plant.Name, plant.ID)
			}
		case <-ctx.Done():
			return
		}
	}
}
