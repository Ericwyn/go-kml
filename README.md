# go-kmz
export nmea log to KMZ file, which you can open with googleMap

## 使用示例
将一个 Location，输出成 kml 文件

```go
func main() {
    locations := []kmz.Location {
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
    kmz.Rander(locations, kmz.ColorRed)
}
```