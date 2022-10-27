package processor

import (
	"context"
	"fmt"
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
	for {
		select {
		case hydrationMsg := <-p.input:
			if hydrationMsg.PlantID == "" {
				continue
			}

			p.plantsRepo.SetHydration(hydrationMsg.PlantID, hydrationMsg.Data)

			ng := p.plantsRepo.GetNormalHydration(hydrationMsg.PlantID)

			plant := p.plantsRepo.GetPlant(hydrationMsg.PlantID)
			switch {
			case hydrationMsg.Data <= ng:
				h := calculateNeededHydration(hydrationMsg.Data)
				p.dronesRepo.Hydrate(hydrationMsg.PlantID, h)

				fmt.Printf("Plant %s with ID %s | Hydration fell below the permissible norm %f and the plant was flooded to %f\n", plant.Name, plant.ID, ng, h)
			default:
				fmt.Printf("Plant %s with ID %s | Hydration is OK\n", plant.Name, plant.ID)
			}

		case <-ctx.Done():
			return
		}
	}
}

func calculateNeededHydration(fact float64) float64 {
	return ((1.0 - fact) * 0.8) + fact
}
