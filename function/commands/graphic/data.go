package graphic

import "github.com/fogleman/gg"

type locationInfo struct {
	x     int
	y     int
	textX int
	textY int
	align gg.Align
}

var (
	locationInfos = map[string]locationInfo{
		"서울": {
			x:     345,
			y:     342,
			textX: 270,
			textY: 23,
			align: gg.AlignRight,
		},
		"부산": {
			x:     633,
			y:     755,
			textX: 1400,
			textY: 775,
			align: gg.AlignLeft,
		},
		"대구": {
			x:     570,
			y:     648,
			textX: 1400,
			textY: 513,
			align: gg.AlignLeft,
		},
		"인천": {
			x:     276,
			y:     353,
			textX: 270,
			textY: 200,
			align: gg.AlignRight,
		},
		"광주": {
			x:     331,
			y:     776,
			textX: 270,
			textY: 915,
			align: gg.AlignRight,
		},
		"대전": {
			x:     413,
			y:     562,
			textX: 270,
			textY: 675,
			align: gg.AlignRight,
		},
		"울산": {
			x:     662,
			y:     700,
			textX: 1400,
			textY: 685,
			align: gg.AlignLeft,
		},
		"세종": {
			x:     396,
			y:     519,
			textX: 270,
			textY: 520,
			align: gg.AlignRight,
		},
		"경기": {
			x:     244,
			y:     234,
			textX: 270,
			textY: 300,
			align: gg.AlignRight,
		},
		"강원": {
			x:     389,
			y:     172,
			textX: 1400,
			textY: 167,
			align: gg.AlignLeft,
		},
		"충북": {
			x:     413,
			y:     426,
			textX: 1400,
			textY: 253,
			align: gg.AlignLeft,
		},
		"충남": {
			x:     248,
			y:     454,
			textX: 270,
			textY: 410,
			align: gg.AlignRight,
		},
		"전북": {
			x:     293,
			y:     618,
			textX: 270,
			textY: 775,
			align: gg.AlignRight,
		},
		"전남": {
			x:     143,
			y:     739,
			textX: 270,
			textY: 1000,
			align: gg.AlignRight,
		},
		"경북": {
			x:     489,
			y:     366,
			textX: 1400,
			textY: 423,
			align: gg.AlignLeft,
		},
		"경남": {
			x:     463,
			y:     664,
			textX: 1400,
			textY: 933,
			align: gg.AlignLeft,
		},
		"제주": {
			x:     610,
			y:     898,
			textX: 1400,
			textY: 1023,
			align: gg.AlignLeft,
		},
	}
)
