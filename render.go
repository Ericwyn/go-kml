package go_kml

import (
	"fmt"
	"github.com/Ericwyn/GoTools/str"
	"strconv"
)

type Location struct {
	Latitude  float64 // 纬度
	Longitude float64 // 经度
	Altitude  float64 // 高度

	Course float64 // 方向
	Speed  float64 // 速度

	Time string // 时间

	Description string // 描述
}

func (obj Location) String() string {
	return fmt.Sprintln(
		"Longitude:", obj.Longitude, "\n",
		"Latitude:", obj.Latitude, "\n",
		"Altitude:", obj.Altitude, "\n",
		"time", obj.Time)
}

type Placemark struct {
	Name        string
	Snippet     string
	Description string

	Location Location
	Color    string
}

/*
往 model string 当中渲染 key 标记的位置， 使用 renderString 来替代
*/
func renderModel(model string, key string, renderString string) string {
	renderKey := key
	if !str.Continues(model, renderKey) {
		fmt.Println("can not render the key:", key, ", not find the render flag")
	} else {
		model = str.ReplaceAll(model, renderKey, renderString)
	}
	return model
}

func Render(locations []Location, color string, trackName string) string {
	// 先载入 model
	model := kmlModel
	// 渲染线条颜色
	model = renderModel(model, kmlModel_key_lines_color, color)

	// 渲染线
	lineXml := renderLine(locations)
	model = renderModel(model, kmlModel_key_lines, lineXml)

	// 渲染点
	pointXml := renderPlacemark(locations, color)
	model = renderModel(model, kmlModel_key_points, pointXml)

	// 渲染 track 名称
	model = renderModel(model, kmlModel_key_track_name, trackName)
	return model
}

func RenderLineOnly(locations []Location, color string, trackName string) string {
	// 先载入 model
	model := kmlModel
	// 渲染线条颜色
	model = renderModel(model, kmlModel_key_lines_color, color)

	// 渲染线
	lineXml := renderLine(locations)
	model = renderModel(model, kmlModel_key_lines, lineXml)

	// 渲染点
	pointXml := renderPlacemark([]Location{}, color)
	model = renderModel(model, kmlModel_key_points, pointXml)

	// 渲染 track 名称
	model = renderModel(model, kmlModel_key_track_name, trackName)
	return model
}

func RenderPointOnly(locations []Location, color string, trackName string) string {
	// 先载入 model
	model := kmlModel
	// 渲染线条颜色
	model = renderModel(model, kmlModel_key_lines_color, color)

	// 渲染线
	lineXml := renderLine([]Location{})
	model = renderModel(model, kmlModel_key_lines, lineXml)

	// 渲染点
	pointXml := renderPlacemark(locations, color)
	model = renderModel(model, kmlModel_key_points, pointXml)

	// 渲染 track 名称
	model = renderModel(model, kmlModel_key_track_name, trackName)
	return model
}

/**
拼接成线的显示
*/
func renderLine(locations []Location) string {
	model := placeLineModel
	renderString := ""
	for _, location := range locations {
		renderString += fmt.Sprintf("%f", location.Longitude) +
			"," + fmt.Sprintf("%f", location.Latitude)

		if location.Altitude != 0 {
			renderString += "," + fmt.Sprintf("%f", location.Altitude)
		}
		renderString += " "
	}
	model = renderModel(model, placeLineModel_key_locations, renderString)
	//model = renderModel(model, placeLineModel_key_locations, renderString)
	return model
}

/**
拼接成点的显示
*/
func renderPlacemark(locations []Location, color string) string {
	res := ""
	for i, location := range locations {
		tempModel := placeMarkModel
		// 颜色
		tempModel = renderModel(tempModel, placeMarkModel_key_color, color)
		// 经纬度
		locationMsg := fmt.Sprintf("%f", location.Longitude) +
			"," + fmt.Sprintf("%f", location.Latitude)
		if location.Altitude != 0 {
			locationMsg += "," + fmt.Sprintf("%f", location.Altitude)
		}
		tempModel = renderModel(tempModel, placeMarkModel_key_location, locationMsg)

		// 名字
		tempModel = renderModel(tempModel, placeMarkModel_key_name,
			"位置"+strconv.Itoa(i)+" "+location.Time)

		// 介绍？
		tempModel = renderModel(tempModel, placeMarkModel_key_description, location.Description)

		res += tempModel + "\n"
	}
	return res
}
