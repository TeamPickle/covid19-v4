package status

import (
	"bytes"
	_ "embed"
	"fmt"
	"function/config"
	"function/resources"
	"io"
	"time"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func generateChart(data []*nCovData) io.Reader {
	caseSeries := chart.TimeSeries{Name: "신규확진"}
	deathSeries := chart.TimeSeries{Name: "사망"}

	for _, v := range data {
		date := v.date.Truncate(time.Hour * 24)
		caseSeries.XValues = append(caseSeries.XValues, date)
		deathSeries.XValues = append(deathSeries.XValues, date)
		caseSeries.YValues = append(caseSeries.YValues, float64(v.confirmedDelta))
		deathSeries.YValues = append(deathSeries.YValues, float64(v.deathDelta))
	}

	graph := chart.Chart{
		Width:  600,
		Height: 400,
		Font:   resources.Pretendard,
		Series: []chart.Series{caseSeries, deathSeries},
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			GridMajorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 80, G: 80, B: 80, A: 255}, StrokeWidth: 1.0},
			GridMinorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 160, G: 160, B: 160, A: 255}, StrokeWidth: 0.4},
		},
		YAxis: chart.YAxis{
			Style:          chart.StyleShow(),
			GridMajorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 80, G: 80, B: 80, A: 255}, StrokeWidth: 1.0},
			GridMinorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 160, G: 160, B: 160, A: 255}, StrokeWidth: 0.4},
			AxisType:       chart.YAxisSecondary,
			ValueFormatter: func(v interface{}) string {
				return fmt.Sprintf("%.f", v.(float64))
			},
		},
	}
	graph.Elements = append(
		graph.Elements,
		chart.Legend(&graph, chart.Style{Font: resources.Pretendard}),
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
