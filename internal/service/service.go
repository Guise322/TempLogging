package service

import (
	"temp-logging/internal/metrics"
	"temp-logging/internal/temp"
	"time"
)

type TempDataService struct {
	TempData chan float64
}

func NewTempDataService() TempDataService {
	return TempDataService{TempData: make(chan float64)}
}

func (s TempDataService) ServeTempData(
	dataGetter temp.DataGetter,
	dataWriter metrics.DataWriter,
	period time.Duration) error {
	var err error

	for err == nil {
		err = s.serve(dataGetter, dataWriter)

		time.Sleep(period)
	}

	return err
}

func (s TempDataService) serve(dataGetter temp.DataGetter, dataWriter metrics.DataWriter) error {
	temp, err := dataGetter.GetData()
	if err != nil {
		return err
	}

	dataWriter.WriteData(temp)

	return nil
}
