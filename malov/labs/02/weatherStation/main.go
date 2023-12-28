package weatherStation

func Main() {
	wd := NewWeatherData()

	display := NewDisplay()
	wd.RegisterObserver(display)

	statsDisplay := NewStatsDisplay()
	wd.RegisterObserver(statsDisplay)

	wd.SetMeasurements(3, 0.7, 760)
	wd.SetMeasurements(4, 0.8, 761)

	wd.RemoveObserver(statsDisplay)

	wd.SetMeasurements(10, 0.8, 761)
	wd.SetMeasurements(-10, 0.8, 761)
}
