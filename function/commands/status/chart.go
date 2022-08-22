package status

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"function/config"
	"function/resources"
	"io"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func generateChart(data []*nCovData) io.Reader {
	caseSeries := chart.TimeSeries{Name: "신규확진"}
	deathSeries := chart.TimeSeries{Name: "사망"}

	for _, v := range data {
		caseSeries.XValues = append(caseSeries.XValues, v.date)
		deathSeries.XValues = append(deathSeries.XValues, v.date)
		caseSeries.YValues = append(caseSeries.YValues, float64(v.confirmedDelta))
		deathSeries.YValues = append(deathSeries.YValues, float64(v.deathDelta))
	}
	a, _ := json.Marshal(data)
	fmt.Println(string(a))

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
		YAxisSecondary: chart.YAxis{
			Style:          chart.StyleShow(),
			GridMajorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 80, G: 80, B: 80, A: 255}, StrokeWidth: 1.0},
			GridMinorStyle: chart.Style{Show: true, StrokeColor: drawing.Color{R: 160, G: 160, B: 160, A: 255}, StrokeWidth: 0.4},
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
