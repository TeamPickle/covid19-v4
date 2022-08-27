package status

import "github.com/wcharczuk/go-chart"

type chartTextSeries struct {
	chart.AnnotationSeries
}

func (ts chartTextSeries) chartStyleDefaults(defaults chart.Style) chart.Style {
	return chart.Style{
		FontColor:   chart.DefaultTextColor,
		Font:        defaults.Font,
		FillColor:   chart.DefaultAnnotationFillColor,
		FontSize:    chart.DefaultAnnotationFontSize,
		StrokeColor: defaults.StrokeColor,
		StrokeWidth: defaults.StrokeWidth,
		Padding:     chart.DefaultAnnotationPadding,
	}
}

func (ts chartTextSeries) Render(r chart.Renderer, canvasBox chart.Box, xrange, yrange chart.Range, defaults chart.Style) {
	if ts.Style.IsZero() || ts.Style.Show {
		seriesStyle := ts.Style.InheritFrom(ts.chartStyleDefaults(defaults))
		for _, a := range ts.Annotations {
			style := a.Style.InheritFrom(seriesStyle)
			lx := canvasBox.Left + xrange.Translate(a.XValue)
			ly := canvasBox.Bottom - yrange.Translate(a.YValue)
			box := chart.Draw.MeasureText(r, a.Label, style)
			width := box.Width()
			chart.Draw.Text(r, a.Label, lx-width/2, ly-10, style)
		}
	}
}
