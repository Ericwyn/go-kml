package go_kmz

import (
	"github.com/Ericwyn/go-kmz/kmz"
	"testing"
)

func TestRanderNmeaToKml(t *testing.T) {
	locations := []kmz.Location{
		{
			Longitude: 115.0382390,
			Latitude:  23.5106791,
			Altitude:  30.0,
			Time:      "20190310 09:10:10",
		},
		{
			Longitude: 115.0383390,
			Latitude:  23.5107791,
			Altitude:  30.0,
			Time:      "20190310 09:11:10",
		},
		{
			Longitude: 115.0384390,
			Latitude:  23.5108791,
			Altitude:  30.0,
			Time:      "20190310 09:12:10",
		},
	}
	kmz.Rander(locations, kmz.ColorRed)
}