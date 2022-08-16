package status

import (
	"bytes"
	"function/config"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type nCovItemData struct {
	activeCount    int
	confirmedCount int
	date           time.Time
	deathCount     int
	releasedCount  int
	testingCount   int
}

type nCovData struct {
	activeAccumulated    int
	activeDelta          int
	confirmedAccumulated int
	confirmedDelta       int
	deathAccumulated     int
	deathDelta           int
	releasedAccumulated  int
	releasedDelta        int
	testingCount         int
	date                 time.Time
}

func parseNCov() (*[]*nCovData, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("ServiceKey", config.OpenAPIkey).
		SetQueryParam("pageNo", "1").
		SetQueryParam("numOfRows", "7").
		SetQueryParam("startCreateDt", time.Now().Add(-8*24*time.Hour).Format("20060102")).
		SetQueryParam("endCreateDt", time.Now().Format("20060102")).
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
		data.activeCount, _ = strconv.Atoi(s.Find("activeCnt").Text())
		data.confirmedCount, _ = strconv.Atoi(s.Find("decideCnt").Text())
		data.date, _ = time.Parse("2006-01-02 15:04:05.000", s.Find("createDt").Text())
		data.deathCount, _ = strconv.Atoi(s.Find("deathCnt").Text())
		data.releasedCount, _ = strconv.Atoi(s.Find("clearCnt").Text())
		data.testingCount, _ = strconv.Atoi(s.Find("examCnt").Text())
		list = append(list, data)
	})

	nCovDataList := make([]*nCovData, 0)

	for i := 1; i < len(list); i++ {
		data := nCovData{}
		data.activeAccumulated = list[i].activeCount - list[i-1].activeCount
		data.activeDelta = list[i].activeCount - list[i-1].activeCount
		data.confirmedAccumulated = list[i].confirmedCount - list[i-1].confirmedCount
		data.confirmedDelta = list[i].confirmedCount - list[i-1].confirmedCount
		data.deathAccumulated = list[i].deathCount - list[i-1].deathCount
		data.deathDelta = list[i].deathCount - list[i-1].deathCount
		data.releasedAccumulated = list[i].releasedCount - list[i-1].releasedCount
		data.releasedDelta = list[i].releasedCount - list[i-1].releasedCount
		data.testingCount = list[i].testingCount
		data.date = list[i].date
		nCovDataList = append(nCovDataList, &data)
	}

	return &nCovDataList, nil
}
