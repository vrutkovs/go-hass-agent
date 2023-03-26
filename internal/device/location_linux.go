package device

import (
	"github.com/joshuar/go-hass-agent/internal/logging"
	"github.com/maltegrosse/go-geoclue2"
)

type linuxLocation struct {
	latitude  float64
	longitude float64
	accuracy  float64
	speed     float64
	altitude  float64
}

func (l *linuxLocation) Gps() []float64 {
	return []float64{l.latitude, l.longitude}
}

func (l *linuxLocation) GpsAccuracy() int {
	return int(l.accuracy)
}

func (l *linuxLocation) Battery() int {
	return 0
}

func (l *linuxLocation) Speed() int {
	return int(l.speed)
}

func (l *linuxLocation) Altitude() int {
	return int(l.altitude)
}

func (l *linuxLocation) Course() int {
	return 0
}

func (l *linuxLocation) VerticalAccuracy() int {
	return 0
}

func LocationUpdater(appID string, locationInfoCh chan interface{}) {

	locationInfo := &linuxLocation{}

	gcm, err := geoclue2.NewGeoclueManager()
	logging.CheckError(err)

	client, err := gcm.GetClient()
	logging.CheckError(err)

	err = client.SetDesktopId(appID)
	logging.CheckError(err)

	err = client.Start()
	logging.CheckError(err)

	c := client.SubscribeLocationUpdated()
	for v := range c {
		_, location, err := client.ParseLocationUpdated(v)
		logging.CheckError(err)

		locationInfo.latitude, err = location.GetLatitude()
		logging.CheckError(err)

		locationInfo.longitude, err = location.GetLongitude()
		logging.CheckError(err)

		locationInfo.accuracy, err = location.GetAccuracy()
		logging.CheckError(err)

		locationInfo.speed, err = location.GetSpeed()
		logging.CheckError(err)

		locationInfo.altitude, err = location.GetAltitude()
		logging.CheckError(err)

		locationInfoCh <- locationInfo
	}

}