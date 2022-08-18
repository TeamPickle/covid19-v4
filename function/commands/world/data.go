package world

import (
	"context"
	"function/external/coronaboard"
	"math"
	"sort"
)

func getMainData(ctx context.Context) (data mainData, totalPages int) {
	coronaboardData, err := coronaboard.ParseBoard(ctx)
	if err != nil {
		panic(err)
	}
	data = mainData{
		active:         coronaboardData.ChartForGlobal.Global.Active[len(coronaboardData.ChartForGlobal.Global.Active)-1],
		confirmed:      coronaboardData.ChartForGlobal.Global.ConfirmedAcc[len(coronaboardData.ChartForGlobal.Global.ConfirmedAcc)-1],
		confirmedDelta: coronaboardData.ChartForGlobal.Global.Confirmed[len(coronaboardData.ChartForGlobal.Global.Confirmed)-1],
		released:       coronaboardData.ChartForGlobal.Global.ReleasedAcc[len(coronaboardData.ChartForGlobal.Global.ReleasedAcc)-1],
		releasedDelta:  coronaboardData.ChartForGlobal.Global.Released[len(coronaboardData.ChartForGlobal.Global.Released)-1],
		death:          coronaboardData.ChartForGlobal.Global.DeathAcc[len(coronaboardData.ChartForGlobal.Global.DeathAcc)-1],
		deathDelta:     coronaboardData.ChartForGlobal.Global.Death[len(coronaboardData.ChartForGlobal.Global.Death)-1],
		countries:      int64(len(coronaboardData.StatGlobalNow)),
	}
	totalPages = int(math.Ceil(float64(len(coronaboardData.StatGlobalNow))/10)) + 1
	return
}

func getListData(ctx context.Context, page int) (data []*detailData, totalPages int) {
	coronaboardData, err := coronaboard.ParseBoard(ctx)
	if err != nil {
		panic(err)
	}

	// sort statGlobalNow by confirmed
	sort.Slice(coronaboardData.StatGlobalNow, func(i, j int) bool {
		a, b := coronaboardData.StatGlobalNow[i], coronaboardData.StatGlobalNow[j]
		return a.Confirmed > b.Confirmed
	})

	for i := page * 10; i < page*10+10 && i < len(coronaboardData.StatGlobalNow); i++ {
		data = append(data, &detailData{
			flag:        coronaboardData.StatGlobalNow[i].Flag,
			countryName: coronaboardData.I18NAll.Ko[coronaboardData.StatGlobalNow[i].Cc],
			confirmed:   coronaboardData.StatGlobalNow[i].Confirmed,
			released:    coronaboardData.StatGlobalNow[i].Released,
			death:       coronaboardData.StatGlobalNow[i].Death,
		})
	}
	totalPages = int(math.Ceil(float64(len(coronaboardData.StatGlobalNow))/10)) + 1
	return
}
