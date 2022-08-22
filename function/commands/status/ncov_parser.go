package status

import (
	"bytes"
	"context"
	"function/config"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type nCovItemData struct {
	activeCount    int64
	confirmedCount int64
	date           time.Time
	deathCount     int64
	releasedCount  int64
	testingCount   int64
}

type nCovData struct {
	activeAccumulated    int64
	activeDelta          int64
	confirmedAccumulated int64
	confirmedDelta       int64
	deathAccumulated     int64
	deathDelta           int64
	releasedAccumulated  int64
	releasedDelta        int64
	testingCount         int64
	date                 time.Time
}

func parseNCov(ctx context.Context) ([]*nCovData, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("ServiceKey", config.OpenAPIkey).
		SetQueryParam("pageNo", "1").
		SetQueryParam("numOfRows", "8").
		SetQueryParam("startCreateDt", time.Now().Add(-8*24*time.Hour).Format("20060102")).
		SetQueryParam("endCreateDt", time.Now().Format("20060102")).
		SetContext(ctx).
		Get("http://openapi.data.go.kr/openapi/service/rest/Covid19/getCovid19InfStateJson")
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, err
	}
	list := []nCovItemData{}
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		data := nCovItemData{}
		data.activeCount, _ = strconv.ParseInt(s.Find("activeCnt").Text(), 10, 64)
		data.confirmedCount, _ = strconv.ParseInt(s.Find("decideCnt").Text(), 10, 64)
		data.date, _ = time.Parse("2006-01-02 15:04:05.000", s.Find("createDt").Text())
		data.deathCount, _ = strconv.ParseInt(s.Find("deathCnt").Text(), 10, 64)
		data.releasedCount, _ = strconv.ParseInt(s.Find("clearCnt").Text(), 10, 64)
		data.testingCount, _ = strconv.ParseInt(s.Find("examCnt").Text(), 10, 64)
		list = append(list, data)
	})

	nCovDataList := []*nCovData{}

	for i := 0; i < len(list)-1; i++ {
		data := nCovData{}
		data.activeAccumulated = list[i].activeCount
		data.activeDelta = list[i].activeCount - list[i+1].activeCount
		data.confirmedAccumulated = list[i].confirmedCount
		data.confirmedDelta = list[i].confirmedCount - list[i+1].confirmedCount
		data.deathAccumulated = list[i].deathCount
		data.deathDelta = list[i].deathCount - list[i+1].deathCount
		data.releasedAccumulated = list[i].releasedCount
		data.releasedDelta = list[i].releasedCount - list[i+1].releasedCount
		data.testingCount = list[i].testingCount
		data.date = list[i].date
		nCovDataList = append(nCovDataList, &data)
	}

	return nCovDataList, nil
}
