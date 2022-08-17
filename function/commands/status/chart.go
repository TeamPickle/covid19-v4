package status

import (
	"bytes"
	"function/config"
	"io"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
	"github.com/wcharczuk/go-chart"
)

func generateChart(data []*nCovData) io.Reader {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
		XAxis: chart.XAxis{
			TickStyle:      chart.Style{Show: true},
			GridMajorStyle: chart.Style{Show: true},
		},
		YAxis: chart.YAxis{
			TickStyle:      chart.Style{Show: true},
			GridMajorStyle: chart.Style{Show: true},
		},
		Background: chart.Style{},
	}
	graph.Elements = append(graph.Elements, chart.Legend(&graph))
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
