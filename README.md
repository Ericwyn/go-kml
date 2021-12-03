# go-kml
create a kml file(the file type which unzip from .kmz) with location message list

and then you can use GoogleEarth to visit all location track points

and if you want to parse the location message from nmea log, you can find the project [go-nema]()

## install

```shell
go get github.com/Ericwyn/go-kml
```

## demo code
```go
import kml "github.com/Ericwyn/go-kml"

func main() {
    locations := []kml.Location {
        {
            Longitude:115.0382390,
            Latitude:23.5106791,
            Altitude: 30.0,
            Time: "20190310 09:10:10",
        },
        {
            Longitude:115.0383390,
            Latitude:23.5107791,
            Altitude: 30.0,
            Time: "20190310 09:11:10",
        },
        {
            Longitude:115.0384390,
            Latitude:23.5108791,
            Altitude: 30.0,
            Time: "20190310 09:12:10",
        },
    }
    kmlString := kml.Render(locations, kml.ColorRed, "track-name")
}
```