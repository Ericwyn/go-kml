package kml

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
		tempModel = randerModel(tempModel, placeMarkModel_key_description, location.Description)

		res += tempModel + "\n"
	}
	return res
}
