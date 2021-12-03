package go_kml

import (
	"github.com/Ericwyn/GoTools/file"
	"testing"
)

func TestKmlColor(t *testing.T) {
	locations := []Location{
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
	testColorMap := map[string]string {
		"白色 ":"#FFFFFFFF",
		"红色":"#FFFF0000",
		"绿色":"#FF00FF00",
		"蓝色":"#FF0000FF",
		"牡丹红":"#FFFF00FF",
		"青色":"#FF00FFFF",
		"黄色":"#FFFFFF00",
		"黑色":"#FF000000",
		"海蓝":"#FF70DB93",
		"巧克力色":"#FF5C3317",
		"蓝紫色":"#FF9F5F9F",
		"黄铜色":"#FFB5A642",
		"亮金色":"#FFD9D919",
		"棕色":"#FFA67D3D",
		"青铜色":"#FF8C7853",
		"2号青铜色":"#FFA67D3D",
		"士官服蓝色":"#FF5F9F9F",
		"冷铜色":"#FFD98719",
		"铜色":"#FFB87333",
		"珊瑚红":"#FFFF7F00",
		"紫蓝色":"#FF42426F",
		"深棕":"#FF5C4033",
		"深绿":"#FF2F4F2F",
		"深铜绿色":"#FF4A766E",
		"深橄榄绿":"#FF4F4F2F",
		"深兰花色":"#FF9932CD",
		"深紫色":"#FF871F78",
		"深石板蓝":"#FF6B238E",
		"深铅灰色":"#FF2F4F4F",
		"深棕褐色":"#FF97694F",
		"深绿松石色":"#FF7093DB",
		"暗木色":"#FF855E42",
		"淡灰色":"#FF545454",
		"土灰玫瑰红色":"#FF856363",
		"长石色":"#FFD19275",
		"火砖色":"#FF8E2323",
		"森林绿":"#FF238E23",
		"金色":"#FFCD7F32",
		"鲜黄色":"#FFDBDB70",
		"灰色":"#FFC0C0C0",
		"铜绿色":"#FF527F76",
		"青黄色":"#FF93DB70",
		"猎人绿":"#FF215E21",
		"印度红":"#FF4E2F2F",
		"土黄色":"#FF9F9F5F",
		"浅蓝色":"#FFC0D9D9",
		"浅灰色":"#FFA8A8A8",
		"浅钢蓝色":"#FF8F8FBD",
		"浅木色":"#FFE9C2A6",
		"石灰绿色":"#FF32CD32",
		"桔黄色":"#FFE47833",
		"褐红色":"#FF8E236B",
		"中海蓝色":"#FF32CD99",
		"中蓝色":"#FF3232CD",
		"中森林绿":"#FF6B8E23",
		"中鲜黄色":"#FFEAEAAE",
		"中兰花色":"#FF9370DB",
		"中海绿色":"#FF426F42",
		"中石板蓝色":"#FF7F00FF",
	}
	for key := range testColorMap {
		se := key
		kmlString:= Render(locations, testColorMap[se], se)
		file.WriteAppend("track_" + se + ".kml", []string{kmlString})
	}
}
