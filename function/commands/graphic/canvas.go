package graphic

import (
	"bytes"
	"fmt"
	"function/resources"
	"image"
	"image/color"
	_ "image/png"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func makeNewCanvas() *gg.Context {
	return gg.NewContext(1665, 1125)
}

func fillBackground(dc *gg.Context) *gg.Context {
	dc.Push()
	defer dc.Pop()
	dc.SetFillStyle(gg.NewSolidPattern(color.RGBA{30, 30, 30, 255}))
	dc.DrawRectangle(0, 0, 1665, 1125)
	dc.Fill()
	return dc
}

func drawBackgroundImage(dc *gg.Context, fileName string, offsetX, offsetY int) *gg.Context {
	dc.Push()
	defer dc.Pop()
	backgroundImage, err := resources.GraphicImages.ReadFile("graphic/" + fileName)
	if err != nil {
		panic(err)
	}
	backgroundReader := bytes.NewReader(backgroundImage)
	image, _, err := image.Decode(backgroundReader)
	if err != nil {
		panic(err)
	}
	dc.DrawImage(image, offsetX, offsetY)
	return dc
}

func getPrefix(per float64) string {
	if per > 20 {
		return ""
	}
	if per > 1.5 {
		return "h"
	}
	if per > 0.5 {
		return "hh"
	}
	return "hhh"
}

/*
  locations.slice(1).forEach((v) => {
    context.font = 'bold 48px NotoSans';
    const measure = context.measureText(`${data[v].confirmed}`);
    context.textAlign = graphicData[v].align;
    context.fillText(
      `${data[v].confirmed}`,
      graphicData[v].textX,
      graphicData[v].textY,
    );
    context.font = '24px NotoSans';
    context.fillText(
      `(총 ${data[v].confirmedAcc})`,
      graphicData[v].align === 'left'
        ? graphicData[v].textX + measure.width + 10
        : graphicData[v].textX - measure.width - 10,
      graphicData[v].textY,
    );
  });
*/

func drawRegion(dc *gg.Context, per float64, regionName string) *gg.Context {
	dc.Push()
	defer dc.Pop()

	imageBytes, err := resources.GraphicImages.Open("graphic/" + getPrefix(per) + regionName + ".png")
	if err != nil {
		panic(err)
	}
	image, _, err := image.Decode(imageBytes)
	if err != nil {
		panic(err)
	}

	locationInfo := locationInfos[regionName]

	dc.DrawImage(image, locationInfo.x+358, locationInfo.y+44)
	return dc
}

func drawRegionText(dc *gg.Context, confirmed, confirmedAcc int64, regionName string) *gg.Context {
	dc.Push()
	defer dc.Pop()

	locationInfo := locationInfos[regionName]

	font := truetype.NewFace(resources.Pretendard, &truetype.Options{Size: 48})
	dc.SetFontFace(font)
	dc.SetRGB255(0, 0, 0)
	ax := 0.0
	w, _ := dc.MeasureString(fmt.Sprintf("%d", confirmed))
	if locationInfo.align == gg.AlignRight {
		ax = 1
		w = -w
	}
	dc.DrawStringWrapped(
		fmt.Sprintf("%d", confirmed),
		float64(locationInfo.textX), float64(locationInfo.textY),
		ax, 0, 500, 16,
		locationInfo.align,
	)

	font = truetype.NewFace(resources.Pretendard, &truetype.Options{Size: 24})
	dc.SetFontFace(font)

	dc.DrawStringWrapped(
		fmt.Sprintf("(총 %d)", confirmedAcc),
		float64(locationInfo.textX)+w, float64(locationInfo.textY),
		ax, 0, 500, 16,
		locationInfo.align,
	)

	return dc
}

func makeImage(dc *gg.Context) *bytes.Reader {
	dc.Push()
	defer dc.Pop()

	var buf bytes.Buffer
	dc.EncodePNG(&buf)

	return bytes.NewReader(buf.Bytes())
}
