package kml

import (
	"github.com/Ericwyn/GoTools/date"
	"github.com/Ericwyn/GoTools/file"
	"time"
)

func RanderLocationsToKml(locations []Location) {
	kml := Rander(locations, ColorRed)
	file.WriteAppend("track_"+date.Format(time.Now(), "yyyyMMdd_HHmmss")+".kmz", []string{kml})
}
