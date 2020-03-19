package kmz

import (
	"fmt"
	"github.com/Ericwyn/GoTools/str"
	"strconv"
)

type Location struct {
	Latitude  float64 // 纬度
	Longitude float64 // 经度
	Altitude  float64 // 高度

	Time string // 时间
}

type Placemark struct {
	Name        string
	Snippet     string
	Description string

	Location Location
	Color    string
}

const kmlModel_key_lines = `${lines}`
const kmlModel_key_points = `${points}`
const kmlModel_key_lines_color = `${color}`
const kmlModel = `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<kml xmlns="http://www.opengis.net/kml/2.2">
    <Document>
        <name><![CDATA[track1]]></name>
        <visibility>1</visibility>
        <open>1</open>
        <Snippet></Snippet>
        <Style id="gv_waypoint_normal">
            <IconStyle>
                <color>ffffffff</color>
                <scale>1</scale>
                <Icon>
                    <href>https://maps.google.com/mapfiles/kml/pal4/icon56.png</href>
                </Icon>
                <hotSpot x="0.5" xunits="fraction" y="0.5" yunits="fraction" />
            </IconStyle>
            <LabelStyle>
                <color>ffffffff</color>
                <scale>1</scale>
            </LabelStyle>
            <BalloonStyle>
                <text><![CDATA[<div style="font-family:Arial,sans-serif; min-width:200px;"><h3>$[name]</h3> <div style="margin-top:8px;">$[description]</div></div>]]></text>
            </BalloonStyle>
        </Style>
        <Style id="gv_waypoint_highlight">
            <IconStyle>
                <color>ffffffff</color>
                <scale>1.2</scale>
                <Icon>
                    <href>https://maps.google.com/mapfiles/kml/pal4/icon56.png</href>
                </Icon>
                <hotSpot x="0.5" xunits="fraction" y="0.5" yunits="fraction" />
            </IconStyle>
            <LabelStyle>
                <color>ffffffff</color>
                <scale>1</scale>
            </LabelStyle>
            <BalloonStyle>
                <text><![CDATA[<div style="font-family:Arial,sans-serif; min-width:200px;"><h3>$[name]</h3> <div style="margin-top:8px;">$[description]</div></div>]]></text>
            </BalloonStyle>
        </Style>
        <Style id="gv_trackpoint_normal">
            <IconStyle>
                <scale>0.3</scale>
                <Icon>
                    <href>http://maps.google.com/mapfiles/kml/pal2/icon26.png</href>
                </Icon>
            </IconStyle>
            <LabelStyle>
                <scale>0</scale>
            </LabelStyle>
            <BalloonStyle>
                <text><![CDATA[<div style="font-family:Arial,sans-serif; min-width:200px;"><h3>$[name]</h3> <div style="margin-top:8px;">$[description]</div></div>]]></text>
            </BalloonStyle>
        </Style>
        <Style id="gv_trackpoint_highlight">
            <IconStyle>
                <scale>0.4</scale>
                <Icon>
                    <href>http://maps.google.com/mapfiles/kml/pal2/icon26.png</href>
                </Icon>
            </IconStyle>
            <LabelStyle>
                <scale>1.2</scale>
            </LabelStyle>
            <BalloonStyle>
                <text><![CDATA[<div style="font-family:Arial,sans-serif; min-width:200px;"><h3>$[name]</h3> <div style="margin-top:8px;">$[description]</div></div>]]></text>
            </BalloonStyle>
        </Style>
        <StyleMap id="gv_waypoint">
            <Pair>
                <key>normal</key>
                <styleUrl>#gv_waypoint_normal</styleUrl>
            </Pair>
            <Pair>
                <key>highlight</key>
                <styleUrl>#gv_waypoint_highlight</styleUrl>
            </Pair>
        </StyleMap>
        <StyleMap id="gv_trackpoint">
            <Pair>
                <key>normal</key>
                <styleUrl>#gv_trackpoint_normal</styleUrl>
            </Pair>
            <Pair>
                <key>highlight</key>
                <styleUrl>#gv_trackpoint_highlight</styleUrl>
            </Pair>
        </StyleMap>
        <Folder id="Tracks">
            <name>Tracks</name>
            <visibility>1</visibility>
            <open>0</open>
            <Folder id="track 1">
                <name><![CDATA[track1]]></name>
                <Snippet></Snippet>
                <description><![CDATA[&nbsp;]]></description>
                <Placemark>
                    <name><![CDATA[track1]]></name>
                    <Snippet></Snippet>
                    <description><![CDATA[&nbsp;]]></description>
                    <Style>
                        <LineStyle>
                            <color>` + kmlModel_key_lines_color + `</color>
                            <width>4</width>
                        </LineStyle>
                    </Style>
                    <MultiGeometry>
` + kmlModel_key_lines + `
                    </MultiGeometry>
                </Placemark>
                <Folder id="track 1 points">
                    <name>Points</name>
` + kmlModel_key_points + `
                </Folder>
            </Folder>
        </Folder>
    </Document>
</kml>
`

const placeLineModel_key_locations = "${locations}"
const placeLineModel = `
                        <LineString>
                            <tessellate>1</tessellate>
                            <altitudeMode>clampToGround</altitudeMode>
                            <coordinates>` + placeLineModel_key_locations + `</coordinates>
                        </LineString>
`

const placeMarkModel_key_name = "${name}"
const placeMarkModel_key_color = "${color}"
const placeMarkModel_key_location = `${location}`
const placeMarkModel_key_description = `${description}`
const placeMarkModel = `
					<Placemark>
                        <name>` + placeMarkModel_key_name + `</name>
                        <Snippet></Snippet>
                        <description>` + placeMarkModel_key_description + `</description>
                        <styleUrl>#gv_trackpoint</styleUrl>
                        <Style>
                            <IconStyle>
                                <color>` + placeMarkModel_key_color + `</color>
                            </IconStyle>
                            <LabelStyle>
                                <color>` + placeMarkModel_key_color + `</color>
                            </LabelStyle>
                        </Style>
                        <Point>
                            <altitudeMode>clampToGround</altitudeMode>
                            <coordinates>` + placeMarkModel_key_location + `</coordinates>
                        </Point>
                    </Placemark>`

/*
往 model string 当中渲染 key 标记的位置， 使用 randerString 来替代
*/
func randerModel(model string, key string, randerString string) string {
	randerKey := key
	if !str.Continues(model, randerKey) {
		fmt.Println("can not rander the key:", key, ", not find the rander flag")
	} else {
		model = str.ReplaceAll(model, randerKey, randerString)
	}
	return model
}

func Rander(locations []Location, color string) string {
	// 先载入 model
	model := kmlModel
	// 渲染线条颜色
	model = randerModel(model, kmlModel_key_lines_color, color)

	// 渲染线
	lineXml := randerLine(locations)
	model = randerModel(model, kmlModel_key_lines, lineXml)

	// 渲染点
	pointXml := randerPlacemark(locations, color)
	model = randerModel(model, kmlModel_key_points, pointXml)

	return model
}

/**
拼接成线的显示
*/
func randerLine(locations []Location) string {
	model := placeLineModel
	randerString := ""
	for _, location := range locations {
		randerString += fmt.Sprintf("%f", location.Longitude) +
			"," + fmt.Sprintf("%f", location.Latitude)

		if location.Altitude != 0 {
			randerString += "," + fmt.Sprintf("%f", location.Altitude)
		}
		randerString += " "
	}
	model = randerModel(model, placeLineModel_key_locations, randerString)
	//model = randerModel(model, placeLineModel_key_locations, randerString)
	return model
}

/**
拼接成点的显示
*/
func randerPlacemark(locations []Location, color string) string {
	res := ""
	for i, location := range locations {
		tempModel := placeMarkModel
		// 颜色
		tempModel = randerModel(tempModel, placeMarkModel_key_color, color)
		// 经纬度
		locationMsg := fmt.Sprintf("%f", location.Longitude) +
			"," + fmt.Sprintf("%f", location.Latitude)
		if location.Altitude != 0 {
			locationMsg += "," + fmt.Sprintf("%f", location.Altitude)
		}
		tempModel = randerModel(tempModel, placeMarkModel_key_location, locationMsg)

		// 名字
		tempModel = randerModel(tempModel, placeMarkModel_key_name,
			"位置"+strconv.Itoa(i)+" "+location.Time)

		// 介绍？
		tempModel = randerModel(tempModel, placeMarkModel_key_description, "")

		res += tempModel + "\n"
	}
	return res
}
