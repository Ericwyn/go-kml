package go_kml

import (
	"github.com/Ericwyn/GoTools/date"
	"github.com/Ericwyn/GoTools/file"
	"github.com/Ericwyn/go-kml/kml"
	"testing"
	"time"
)

func TestRanderNmeaToKml(t *testing.T) {
	locations := []kml.Location{
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
	kmlString := kml.Rander(locations, kml.ColorMagenta, "test-tracking")
	file.WriteAppend("track_"+date.Format(time.Now(), "yyyyMMdd_HHmmss")+".kmz", []string{kmlString})
}
