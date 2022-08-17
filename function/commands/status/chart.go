package status

import (
	"bytes"
	_ "embed"
	"function/config"
	"function/resources"
	"io"
	"time"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func generateChart(data []*nCovData) io.Reader {
	fontFile, _ := resources.Asset("PretendardVariable.ttf")
	font, _ := truetype.Parse(fontFile)
	graph := chart.Chart{
		Font: font,
		Series: []chart.Series{
			chart.TimeSeries{
				Name: "확진자",
				XValues: []time.Time{
					time.Date(2022, 0, 1, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 2, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 3, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 4, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 5, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 6, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 7, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 8, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 9, 9, 0, 0, 0, time.Local),
				},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0},
			},
			chart.TimeSeries{
				Name: "사망자",
				XValues: []time.Time{
					time.Date(2022, 0, 1, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 2, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 3, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 4, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 5, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 6, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 7, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 8, 9, 0, 0, 0, time.Local),
					time.Date(2022, 0, 9, 9, 0, 0, 0, time.Local),
				},
				YValues: []float64{3.0, 4.0, 8.0, 12.0, 16.0, 3.0, 24.0, 28.0, 32.0},
			},
		},
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			GridMajorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 80, G: 80, B: 80, A: 255}, StrokeWidth: 1.0},
			GridMinorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 160, G: 160, B: 160, A: 255}, StrokeWidth: 0.4},
		},
		YAxisSecondary: chart.YAxis{
			Style:          chart.StyleShow(),
			GridMajorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 80, G: 80, B: 80, A: 255}, StrokeWidth: 1.0},
			GridMinorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 160, G: 160, B: 160, A: 255}, StrokeWidth: 0.4},
		},
		Width:  600,
		Height: 400,
	}
	graph.Elements = append(
		graph.Elements,
		chart.Legend(&graph, chart.Style{Font: font}),
	)
	buffer := bytes.NewBuffer(nil)
	graph.Render(chart.PNG, buffer)
	return io.MultiReader(buffer)
}

func generateChartURL(data []*nCovData) string {
	client := webhook.New(config.LogWebhookID, config.LogWebhookToken)
	msg, err := client.ExecuteAndWait(webhook.ExecuteData{
		Files: []sendpart.File{{Name: data[0].date.Format("2006-01-02") + " chart.png", Reader: generateChart(data)}},
	})
	if err != nil {
		panic("Can't send chart image: " + err.Error())
	}
	return msg.Attachments[0].URL
}
