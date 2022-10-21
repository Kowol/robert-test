package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show Time!", name)
}

func (a *App) GetChart() []vehicleLog {
	//engineSpeedChart, _, _ := loadCSV()

	f, err := os.Open("./log.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	vLog := createVehicleLog(data)

	return vLog
}

type vehicleLog struct {
	Time           float64
	EngineSpeed    int
	PressureActual float64
	PressureWanted float64
}

type intLog struct {
	Time  float64
	Value int
}

type floatLog struct {
	Time  float64
	Value float64
}

func createVehicleLog(data [][]string) []vehicleLog {
	var logs []vehicleLog
	for i, line := range data {
		if i < 4 {
			continue
		}

		var singleLog vehicleLog

		for j, field := range line {

			switch j {
			case 1:
				v, err := strconv.ParseFloat(field, 32)
				if err != nil {
					log.Fatal(err)
				}
				singleLog.Time = v
				break

			case 2:
				v, err := strconv.Atoi(field)
				if err != nil {
					log.Fatal(err)
				}
				singleLog.EngineSpeed = v
				break

			case 4:
				v, err := strconv.ParseFloat(field, 32)
				if err != nil {
					log.Fatal(err)
				}
				singleLog.PressureActual = v
				break

			case 6:
				v, err := strconv.ParseFloat(field, 32)
				if err != nil {
					log.Fatal(err)
				}
				singleLog.PressureWanted = v
				break
			}
		}

		logs = append(logs, singleLog)
	}

	return logs
}

func loadCSV() ([]intLog, []floatLog, []floatLog) {
	f, err := os.Open("./log.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	vLog := createVehicleLog(data)

	var engineSpeedLog []intLog
	var pressureActualLog []floatLog
	var pressureWantedLog []floatLog

	for _, singleLog := range vLog {
		engineSpeedLog = append(engineSpeedLog, intLog{
			Time:  singleLog.Time,
			Value: singleLog.EngineSpeed,
		})

		pressureActualLog = append(pressureActualLog, floatLog{
			Time:  singleLog.Time,
			Value: singleLog.PressureActual,
		})

		pressureActualLog = append(pressureWantedLog, floatLog{
			Time:  singleLog.Time,
			Value: singleLog.PressureWanted,
		})
	}

	return engineSpeedLog, pressureWantedLog, pressureWantedLog
}
