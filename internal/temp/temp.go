package temp

import (
	"os/exec"
	"regexp"
	"strconv"
)

type DataGetter interface {
	GetData() (data float64, err error)
}

type TemperatureDataGetter struct{}

func (TemperatureDataGetter) GetData() (data float64, err error) {
	app := "vcgencmd"
	arg := "measure_temp"

	cmd := exec.Command(app, arg)
	stdout, err := cmd.Output()

	if err != nil {
		return 0.0, err
	}

	rawTemp := extractTempData(string(stdout))
	temp, err := strconv.ParseFloat(rawTemp, 32)

	return temp, err
}

func extractTempData(rawTemp string) string {
	re := regexp.MustCompile("[0-9][0-9].[0-9]")

	return re.FindString(rawTemp)
}
