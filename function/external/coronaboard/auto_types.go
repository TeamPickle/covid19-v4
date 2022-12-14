package coronaboard

type GlobalNowStatus struct {
	Cc              string `json:"cc"`
	Confirmed       int64  `json:"confirmed"`
	Death           int64  `json:"death"`
	Released        int64  `json:"released"`
	Tested          int64  `json:"tested"`
	Critical        int64  `json:"critical"`
	Active          int64  `json:"active"`
	ConfirmedPrev   int64  `json:"confirmed_prev"`
	DeathPrev       int64  `json:"death_prev"`
	ReleasedPrev    int64  `json:"released_prev"`
	ActivePrev      int64  `json:"active_prev"`
	TestedPrev      int64  `json:"tested_prev"`
	CriticalPrev    int64  `json:"critical_prev"`
	ConfirmedXest   int64  `json:"confirmed_xest"`
	DeathXest       int64  `json:"death_xest"`
	ReleasedXest    int64  `json:"released_xest"`
	ActiveXest      int64  `json:"active_xest"`
	TestedXest      int64  `json:"tested_xest"`
	CriticalXest    int64  `json:"critical_xest"`
	Population      int64  `json:"population"`
	Incidence       int64  `json:"incidence"`
	Flag            string `json:"flag"`
	EhpadConfirmed  int64  `json:"ehpadConfirmed,omitempty"`
	EhpadProbable   int64  `json:"ehpadProbable,omitempty"`
	EhpadDeath      int64  `json:"ehpadDeath,omitempty"`
	Testing         int64  `json:"testing,omitempty"`
	Negative        int64  `json:"negative,omitempty"`
	FromOversea     int64  `json:"fromOversea,omitempty"`
	TestingPrev     int64  `json:"testing_prev,omitempty"`
	NegativePrev    int64  `json:"negative_prev,omitempty"`
	TestingXest     int64  `json:"testing_xest,omitempty"`
	NegativeXest    int64  `json:"negative_xest,omitempty"`
	VaccinatedFully int64  `json:"vaccinatedFully,omitempty"`
}

type DomesticNowStatus struct {
	Region              string `json:"region"`
	Confirmed           int64  `json:"confirmed,omitempty"`
	Death               int64  `json:"death,omitempty"`
	Released            int64  `json:"released,omitempty"`
	VaccinatedOnce      int64  `json:"vaccinatedOnce,omitempty"`
	VaccinatedFully     int64  `json:"vaccinatedFully,omitempty"`
	Active              int64  `json:"active"`
	ConfirmedPrev       int64  `json:"confirmed_prev"`
	ReleasedPrev        int64  `json:"released_prev"`
	DeathPrev           int64  `json:"death_prev"`
	ActivePrev          int64  `json:"active_prev"`
	VaccinatedOncePrev  int64  `json:"vaccinatedOnce_prev"`
	VaccinatedFullyPrev int64  `json:"vaccinatedFully_prev"`
	VaccinatedBoostPrev int64  `json:"vaccinatedBoost_prev"`
	VaccinatedBoost     int64  `json:"vaccinatedBoost,omitempty"`
}

type CoronaBoardData struct {
	LastUpdated     int                 `json:"lastUpdated"`
	StatGlobalNow   []GlobalNowStatus   `json:"statGlobalNow"`
	StatDomesticNow []DomesticNowStatus `json:"statDomesticNow"`
	ChartForGlobal  struct {
		Kr struct {
			Date         []string      `json:"date"`
			Active       []int         `json:"active"`
			ConfirmedAcc []int         `json:"confirmed_acc"`
			DeathAcc     []int         `json:"death_acc"`
			ReleasedAcc  []int         `json:"released_acc"`
			CriticalAcc  []interface{} `json:"critical_acc"`
			Confirmed    []int         `json:"confirmed"`
			Death        []int         `json:"death"`
			Released     []int         `json:"released"`
			Critical     []interface{} `json:"critical"`
		} `json:"KR"`
		Global struct {
			Date         []string      `json:"date"`
			Active       []int64       `json:"active"`
			ConfirmedAcc []int64       `json:"confirmed_acc"`
			DeathAcc     []int64       `json:"death_acc"`
			ReleasedAcc  []int64       `json:"released_acc"`
			CriticalAcc  []interface{} `json:"critical_acc"`
			Confirmed    []int64       `json:"confirmed"`
			Death        []int64       `json:"death"`
			Released     []int64       `json:"released"`
			Critical     []interface{} `json:"critical"`
		} `json:"global"`
	} `json:"chartForGlobal"`
	ChartForDomestic struct {
		Date  []string `json:"date"`
		Seoul struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Busan struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Daegu struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Incheon struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Gwangju struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Daejeon struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Ulsan struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Sejong struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Gyeonggi struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Gangwon struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Chungbuk struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Chungnam struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Jeonbuk struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Jeonnam struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Gyeongbuk struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Gyeongnam struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Jeju struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
		Oversea struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"??????"`
	} `json:"chartForDomestic"`
	ChartForVaccineByRegion struct {
		Date []string `json:"date"`
		All  struct {
			Date               []string `json:"date"`
			VaccinatedOnceAcc  []int    `json:"vaccinatedOnce_acc"`
			VaccinatedOnce     []int    `json:"vaccinatedOnce"`
			VaccinatedFullyAcc []int    `json:"vaccinatedFully_acc"`
			VaccinatedFully    []int    `json:"vaccinatedFully"`
			VaccinatedBoostAcc []int    `json:"vaccinatedBoost_acc"`
			VaccinatedBoost    []int    `json:"vaccinatedBoost"`
		} `json:"??????"`
	} `json:"chartForVaccineByRegion"`
	ChartForVaccineByType struct {
		AstraZeneca struct {
			VaccinatedOnceAcc []int         `json:"vaccinatedOnce_acc"`
			MildAcc           []interface{} `json:"mild_acc"`
			AnaphylaxisAcc    []interface{} `json:"anaphylaxis_acc"`
			CriticalAcc       []interface{} `json:"critical_acc"`
			DeathAcc          []interface{} `json:"death_acc"`
			VaccinatedOnce    []int         `json:"vaccinatedOnce"`
			Mild              []interface{} `json:"mild"`
			Anaphylaxis       []interface{} `json:"anaphylaxis"`
			Critical          []interface{} `json:"critical"`
			Death             []interface{} `json:"death"`
		} `json:"?????????????????????"`
		Pfizer struct {
			VaccinatedOnceAcc []int         `json:"vaccinatedOnce_acc"`
			MildAcc           []interface{} `json:"mild_acc"`
			AnaphylaxisAcc    []interface{} `json:"anaphylaxis_acc"`
			CriticalAcc       []interface{} `json:"critical_acc"`
			DeathAcc          []interface{} `json:"death_acc"`
			VaccinatedOnce    []int         `json:"vaccinatedOnce"`
			Mild              []interface{} `json:"mild"`
			Anaphylaxis       []interface{} `json:"anaphylaxis"`
			Critical          []interface{} `json:"critical"`
			Death             []interface{} `json:"death"`
		} `json:"?????????"`
		Moderna struct {
			VaccinatedOnceAcc []int         `json:"vaccinatedOnce_acc"`
			MildAcc           []interface{} `json:"mild_acc"`
			AnaphylaxisAcc    []interface{} `json:"anaphylaxis_acc"`
			CriticalAcc       []interface{} `json:"critical_acc"`
			DeathAcc          []interface{} `json:"death_acc"`
			VaccinatedOnce    []int         `json:"vaccinatedOnce"`
			Mild              []interface{} `json:"mild"`
			Anaphylaxis       []interface{} `json:"anaphylaxis"`
			Critical          []interface{} `json:"critical"`
			Death             []interface{} `json:"death"`
		} `json:"?????????"`
		Janssen struct {
			VaccinatedOnceAcc []int         `json:"vaccinatedOnce_acc"`
			MildAcc           []interface{} `json:"mild_acc"`
			AnaphylaxisAcc    []interface{} `json:"anaphylaxis_acc"`
			CriticalAcc       []interface{} `json:"critical_acc"`
			DeathAcc          []interface{} `json:"death_acc"`
			VaccinatedOnce    []int         `json:"vaccinatedOnce"`
			Mild              []interface{} `json:"mild"`
			Anaphylaxis       []interface{} `json:"anaphylaxis"`
			Critical          []interface{} `json:"critical"`
			Death             []interface{} `json:"death"`
		} `json:"??????"`
		Novavax struct {
			VaccinatedOnceAcc []int         `json:"vaccinatedOnce_acc"`
			MildAcc           []interface{} `json:"mild_acc"`
			AnaphylaxisAcc    []interface{} `json:"anaphylaxis_acc"`
			CriticalAcc       []interface{} `json:"critical_acc"`
			DeathAcc          []interface{} `json:"death_acc"`
			VaccinatedOnce    []int         `json:"vaccinatedOnce"`
			Mild              []interface{} `json:"mild"`
			Anaphylaxis       []interface{} `json:"anaphylaxis"`
			Critical          []interface{} `json:"critical"`
			Death             []interface{} `json:"death"`
		} `json:"????????????"`
		Etc struct {
			VaccinatedOnceAcc []int         `json:"vaccinatedOnce_acc"`
			MildAcc           []interface{} `json:"mild_acc"`
			AnaphylaxisAcc    []interface{} `json:"anaphylaxis_acc"`
			CriticalAcc       []interface{} `json:"critical_acc"`
			DeathAcc          []interface{} `json:"death_acc"`
			VaccinatedOnce    []int         `json:"vaccinatedOnce"`
			Mild              []interface{} `json:"mild"`
			Anaphylaxis       []interface{} `json:"anaphylaxis"`
			Critical          []interface{} `json:"critical"`
			Death             []interface{} `json:"death"`
		} `json:"??????"`
		Date []string `json:"date"`
	} `json:"chartForVaccineByType"`
	VaccineDataForDashboard struct {
		AstraZeneca struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"?????????????????????"`
		Pfizer struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"?????????"`
		Moderna struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"?????????"`
		Janssen struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"??????"`
		Novavax struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"????????????"`
		Etc struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"??????"`
		All struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VaccinatedFully     int `json:"vaccinatedFully"`
			VaccinatedFullyPrev int `json:"vaccinatedFully_prev"`
			VaccinatedBoost     int `json:"vaccinatedBoost"`
			VaccinatedBoostPrev int `json:"vaccinatedBoost_prev"`
		} `json:"??????"`
	} `json:"vaccineDataForDashboard"`
	Kr struct {
		ChartTesting struct {
			Date           []string  `json:"date"`
			ConfirmRateAcc []float64 `json:"confirm_rate_acc"`
			ConfirmedAcc   []int     `json:"confirmed_acc"`
			NegativeAcc    []int     `json:"negative_acc"`
			Testing        []int     `json:"testing"`
			ConfirmRate    []float64 `json:"confirm_rate"`
			Confirmed      []int     `json:"confirmed"`
			Negative       []int     `json:"negative"`
		} `json:"chartTesting"`
		StatByKrLocation []struct {
			Region             string  `json:"region"`
			Confirmed          int     `json:"confirmed"`
			Death              int     `json:"death"`
			Released           int     `json:"released"`
			VaccinatedOnce     int     `json:"vaccinatedOnce,omitempty"`
			VaccinatedFully    int     `json:"vaccinatedFully,omitempty"`
			Active             int     `json:"active"`
			Population         int     `json:"population"`
			Incidence          float64 `json:"incidence"`
			ConfirmedPrev      int     `json:"confirmed_prev"`
			ReleasedPrev       int     `json:"released_prev"`
			DeathPrev          int     `json:"death_prev"`
			ActivePrev         int     `json:"active_prev"`
			ConfirmedSevenDays int     `json:"confirmed_sevenDays"`
			ReleasedSevenDays  int     `json:"released_sevenDays"`
			DeathSevenDays     int     `json:"death_sevenDays"`
			ConfirmedEightDays int     `json:"confirmed_eightDays"`
			ReleasedEightDays  int     `json:"released_eightDays"`
			DeathEightDays     int     `json:"death_eightDays"`
		} `json:"statByKrLocation"`
		BannedList struct {
			Banned struct {
				Count int `json:"count"`
				List  []struct {
					Continent string `json:"continent"`
					Countries string `json:"countries"`
				} `json:"list"`
			} `json:"banned"`
			BannedRemoved struct {
				Count int `json:"count"`
				List  []struct {
					Continent string `json:"continent"`
					Countries string `json:"countries"`
				} `json:"list"`
			} `json:"bannedRemoved"`
			Quarantined struct {
				Count int `json:"count"`
				List  []struct {
					Continent string `json:"continent"`
					Countries string `json:"countries"`
				} `json:"list"`
			} `json:"quarantined"`
			Vaccine struct {
				Count int `json:"count"`
				List  []struct {
					Continent string `json:"continent"`
					Countries string `json:"countries"`
				} `json:"list"`
			} `json:"vaccine"`
			Restricted struct {
				Count int `json:"count"`
				List  []struct {
					Continent string `json:"continent"`
					Countries string `json:"countries"`
				} `json:"list"`
			} `json:"restricted"`
			FlightBanned []interface{} `json:"flightBanned"`
			SourceTitle  string        `json:"sourceTitle"`
			SourceLink   string        `json:"sourceLink"`
			LastUpdated  string        `json:"lastUpdated"`
		} `json:"bannedList"`
		ChartBySource struct {
			All struct {
				Time string `json:"time"`
				Data struct {
					NAMING_FAILED struct {
						NAMING_FAILED   int `json:"????????????"`
						NAMING_FAILED0  int `json:"????????? ??????"`
						NAMING_FAILED1  int `json:"????????????????????????"`
						NAMING_FAILED2  int `json:"?????????"`
						NAMING_FAILED3  int `json:"?????? ?????? ??????"`
						NAMING_FAILED4  int `json:"???????????????????????????"`
						NAMING_FAILED5  int `json:"???????????????"`
						NAMING_FAILED6  int `json:"?????? "`
						NAMING_FAILED7  int `json:"???????????????????????????"`
						NAMING_FAILED8  int `json:"????????????"`
						NAMING_FAILED9  int `json:"???????????????"`
						NAMING_FAILED10 int `json:"??????????????????"`
						NAMING_FAILED11 int `json:"????????????????????????"`
						NAMING_FAILED12 int `json:"???????????????????????????"`
						NAMING_FAILED13 int `json:"??????????????? ????????????"`
						NAMING_FAILED14 int `json:"??????????????? ??????????????????"`
						NAMING_FAILED15 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED16 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED17 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED18 int `json:"?????? ????????? ????????? ??????"`
						NAMING_FAILED19 int `json:"?????? ????????? ??????????????????"`
						IM              int `json:"IM????????? ?????????????????????"`
						NAMING_FAILED20 int `json:"?????? ?????????/?????????/?????????/????????? ??????"`
						NAMING_FAILED21 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED22 int `json:"?????? ?????? ??????"`
						NAMING_FAILED23 int `json:"?????? ????????? ????????????"`
						BTJ             int `json:"?????? ?????? BTJ????????????"`
						Num2            int `json:"?????? ????????? ????????????2"`
						NAMING_FAILED24 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED25 int `json:"?????? ????????? ?????????"`
						NAMING_FAILED26 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED0 struct {
						NAMING_FAILED   int `json:"????????????"`
						NAMING_FAILED0  int `json:"????????? ??????"`
						NAMING_FAILED1  int `json:"????????????????????????"`
						NAMING_FAILED2  int `json:"?????????"`
						NAMING_FAILED3  int `json:"?????? ?????? ??????"`
						NAMING_FAILED4  int `json:"???????????????????????????"`
						NAMING_FAILED5  int `json:"???????????????"`
						NAMING_FAILED6  int `json:"?????? "`
						NAMING_FAILED7  int `json:"???????????????????????????"`
						NAMING_FAILED8  int `json:"????????????"`
						NAMING_FAILED9  int `json:"???????????????"`
						NAMING_FAILED10 int `json:"??????????????????"`
						NAMING_FAILED11 int `json:"???????????????????????????"`
						NAMING_FAILED12 int `json:"??????????????? ????????????"`
						NAMING_FAILED13 int `json:"??????????????? ??????????????????"`
						NAMING_FAILED14 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED15 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED16 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED17 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED1 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"???????????????"`
						NAMING_FAILED6 int `json:"?????? "`
						NAMING_FAILED7 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED8 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED2 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"???????????????"`
						NAMING_FAILED6 int `json:"?????? "`
						NAMING_FAILED7 int `json:"???????????????????????????"`
						NAMING_FAILED8 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED3 struct {
						NAMING_FAILED   int `json:"????????????"`
						NAMING_FAILED0  int `json:"????????? ??????"`
						NAMING_FAILED1  int `json:"????????????????????????"`
						NAMING_FAILED2  int `json:"?????????"`
						NAMING_FAILED3  int `json:"?????? ?????? ??????"`
						NAMING_FAILED4  int `json:"???????????????????????????"`
						NAMING_FAILED5  int `json:"???????????????"`
						NAMING_FAILED6  int `json:"?????? "`
						NAMING_FAILED7  int `json:"???????????????????????????"`
						NAMING_FAILED8  int `json:"????????????"`
						NAMING_FAILED9  int `json:"???????????????"`
						NAMING_FAILED10 int `json:"??????????????????"`
						NAMING_FAILED11 int `json:"???????????????????????????"`
						NAMING_FAILED12 int `json:"??????????????? ??????????????????"`
						NAMING_FAILED13 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED4 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????"`
						NAMING_FAILED5 int `json:"????????????????????????"`
						IM             int `json:"IM????????? ?????????????????????"`
						NAMING_FAILED6 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED7 int `json:"?????? ?????? ??????"`
						NAMING_FAILED8 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED5 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"???????????????"`
						IM             int `json:"IM????????? ?????????????????????"`
						NAMING_FAILED6 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED6 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????"`
						NAMING_FAILED5 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED7 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED8 struct {
						NAMING_FAILED   int `json:"????????????"`
						NAMING_FAILED0  int `json:"????????? ??????"`
						NAMING_FAILED1  int `json:"????????????????????????"`
						NAMING_FAILED2  int `json:"?????????"`
						NAMING_FAILED3  int `json:"?????? ?????? ??????"`
						NAMING_FAILED4  int `json:"???????????????????????????"`
						NAMING_FAILED5  int `json:"???????????????"`
						NAMING_FAILED6  int `json:"?????? "`
						NAMING_FAILED7  int `json:"???????????????????????????"`
						NAMING_FAILED8  int `json:"????????????"`
						NAMING_FAILED9  int `json:"???????????????"`
						NAMING_FAILED10 int `json:"??????????????????"`
						NAMING_FAILED11 int `json:"???????????????????????????"`
						NAMING_FAILED12 int `json:"??????????????? ??????????????????"`
						NAMING_FAILED13 int `json:"?????? ????????? ????????? ??????"`
						NAMING_FAILED14 int `json:"?????? ????????? ??????????????????"`
						IM              int `json:"IM????????? ?????????????????????"`
						NAMING_FAILED15 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED9 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"???????????????"`
						NAMING_FAILED6 int `json:"???????????????????????????"`
						NAMING_FAILED7 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED10 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"???????????????"`
						NAMING_FAILED6 int `json:"?????? "`
						NAMING_FAILED7 int `json:"?????? ?????????/?????????/?????????/????????? ??????"`
						NAMING_FAILED8 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED11 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????????????????????"`
						NAMING_FAILED1 int `json:"?????????"`
						NAMING_FAILED2 int `json:"?????? ?????? ??????"`
						NAMING_FAILED3 int `json:"???????????????????????????"`
						NAMING_FAILED4 int `json:"???????????????"`
						NAMING_FAILED5 int `json:"???????????????????????????"`
						NAMING_FAILED6 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED12 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"?????? "`
						NAMING_FAILED6 int `json:"???????????????????????????"`
						NAMING_FAILED7 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED8 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED13 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"????????????????????????"`
						NAMING_FAILED5 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED14 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????????????????"`
						NAMING_FAILED5 int `json:"???????????????"`
						NAMING_FAILED6 int `json:"???????????????????????????"`
						BTJ            int `json:"?????? ?????? BTJ????????????"`
						Num2           int `json:"?????? ????????? ????????????2"`
						NAMING_FAILED7 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED15 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????? ??????"`
						NAMING_FAILED1 int `json:"????????????????????????"`
						NAMING_FAILED2 int `json:"?????????"`
						NAMING_FAILED3 int `json:"?????? ?????? ??????"`
						NAMING_FAILED4 int `json:"???????????????"`
						NAMING_FAILED5 int `json:"?????? "`
						NAMING_FAILED6 int `json:"?????? ????????? ?????????"`
						NAMING_FAILED7 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED16 struct {
						NAMING_FAILED  int `json:"????????????"`
						NAMING_FAILED0 int `json:"????????????????????????"`
						NAMING_FAILED1 int `json:"?????????"`
						NAMING_FAILED2 int `json:"?????? ?????? ??????"`
						NAMING_FAILED3 int `json:"??????"`
					} `json:"??????"`
					NAMING_FAILED17 struct {
						NAMING_FAILED int `json:"????????????"`
					} `json:"??????"`
				} `json:"data"`
			} `json:"all"`
			Recent struct {
				Time string `json:"time"`
				Data struct {
					NAMING_FAILED struct {
						NAMING_FAILED   int `json:"????????? ?????????????????? "`
						NAMING_FAILED0  int `json:"????????? ????????? "`
						NAMING_FAILED1  int `json:"????????? ????????????"`
						NAMING_FAILED2  int `json:"????????? ?????????"`
						NAMING_FAILED3  int `json:"?????? ????????????"`
						NAMING_FAILED4  int `json:"?????? ????????? ?????? "`
						NAMING_FAILED5  int `json:"?????? ????????? ????????? ??????"`
						Num2            int `json:"?????? ????????? ?????????2 "`
						NAMING_FAILED6  int `json:"?????? ???????????? ????????????"`
						NAMING_FAILED7  int `json:"?????? ????????? ????????????"`
						NAMING_FAILED8  int `json:"?????? ????????? ??????"`
						IM              int `json:"IM ????????? ????????? ?????? ????????????"`
						NAMING_FAILED9  int `json:"?????? ?????? ??????"`
						NAMING_FAILED10 int `json:"?????? ?????? ?????????"`
						NAMING_FAILED11 int `json:"?????? ?????? ????????????"`
						NAMING_FAILED12 int `json:"?????? ?????? ????????????"`
						Num3            int `json:"?????? ???????????? ?????????3"`
					} `json:"??????"`
					NAMING_FAILED0 struct {
						NAMING_FAILED  int `json:"????????? ?????????????????? "`
						NAMING_FAILED0 int `json:"????????? ????????? "`
						NAMING_FAILED1 int `json:"????????? ????????????"`
						NAMING_FAILED2 int `json:"????????? ?????????"`
						NAMING_FAILED3 int `json:"?????? ????????????"`
					} `json:"??????"`
					NAMING_FAILED1 struct {
						NAMING_FAILED  int `json:"?????? ????????? ?????? "`
						NAMING_FAILED0 int `json:"?????? ????????? ????????? ??????"`
						Num2           int `json:"?????? ????????? ?????????2 "`
						NAMING_FAILED1 int `json:"?????? ???????????? ????????????"`
						NAMING_FAILED2 int `json:"?????? ????????? ????????????"`
						NAMING_FAILED3 int `json:"?????? ????????? ??????"`
						IM             int `json:"IM ????????? ????????? ?????? ????????????"`
						NAMING_FAILED4 int `json:"?????? ?????? ??????"`
						NAMING_FAILED5 int `json:"?????? ?????? ?????????"`
						NAMING_FAILED6 int `json:"?????? ?????? ????????????"`
						NAMING_FAILED7 int `json:"?????? ?????? ????????????"`
						Num3           int `json:"?????? ???????????? ?????????3"`
					} `json:"?????????"`
				} `json:"data"`
			} `json:"recent"`
		} `json:"chartBySource"`
		ChartByGender struct {
			Time string `json:"time"`
			Data struct {
				Confirmed struct {
					Male   int `json:"male"`
					Female int `json:"female"`
				} `json:"confirmed"`
				Death struct {
					Male   int `json:"male"`
					Female int `json:"female"`
				} `json:"death"`
				Released struct {
					Male   int `json:"male"`
					Female int `json:"female"`
				} `json:"released"`
			} `json:"data"`
		} `json:"chartByGender"`
		ChartByAge struct {
			Time string `json:"time"`
			Data struct {
				Confirmed struct {
					Zero9  int `json:"0-9???"`
					One0   int `json:"10???"`
					Two0   int `json:"20???"`
					Three0 int `json:"30???"`
					Four0  int `json:"40???"`
					Five0  int `json:"50???"`
					Six0   int `json:"60???"`
					Seven0 int `json:"70???"`
					Eight0 int `json:"80?????????"`
				} `json:"confirmed"`
				Death struct {
					Zero9  int `json:"0-9???"`
					One0   int `json:"10???"`
					Two0   int `json:"20???"`
					Three0 int `json:"30???"`
					Four0  int `json:"40???"`
					Five0  int `json:"50???"`
					Six0   int `json:"60???"`
					Seven0 int `json:"70???"`
					Eight0 int `json:"80?????????"`
				} `json:"death"`
				Released struct {
					Zero9  int `json:"0-9???"`
					One0   int `json:"10???"`
					Two0   int `json:"20???"`
					Three0 int `json:"30???"`
					Four0  int `json:"40???"`
					Five0  int `json:"50???"`
					Six0   int `json:"60???"`
					Seven0 int `json:"70???"`
					Eight0 int `json:"80?????????"`
				} `json:"released"`
			} `json:"data"`
		} `json:"chartByAge"`
		PatientLogUpdate []struct {
			Order      string `json:"order"`
			Open       string `json:"open"`
			Message    string `json:"message"`
			SourceText string `json:"source_text"`
			SourceLink string `json:"source_link"`
		} `json:"patientLogUpdate"`
		SdLevel struct {
			SdLevel15 string `json:"sdLevel15"`
			SdLevel2  string `json:"sdLevel2"`
			SdLevel25 string `json:"sdLevel25"`
			SdLevel3  string `json:"sdLevel3"`
			Stage1    string `json:"stage1"`
			Stage2    string `json:"stage2"`
			Stage3    string `json:"stage3"`
			Stage4    string `json:"stage4"`
		} `json:"sdLevel"`
		NoticeFeature struct {
			Notice []struct {
				Type string `json:"type"`
				HTML string `json:"html"`
			} `json:"notice"`
		} `json:"noticeFeature"`
		StatLastUpdated string `json:"statLastUpdated"`
	} `json:"KR"`
	Mers struct {
		Confirmed   int `json:"confirmed"`
		ConfirmedKR int `json:"confirmedKR"`
		Death       int `json:"death"`
		DeathKR     int `json:"deathKR"`
		Countries   int `json:"countries"`
	} `json:"mers"`
	Sars struct {
		Confirmed   int `json:"confirmed"`
		ConfirmedKR int `json:"confirmedKR"`
		Death       int `json:"death"`
		DeathKR     int `json:"deathKR"`
		Countries   int `json:"countries"`
	} `json:"sars"`
	I18NAll struct {
		Ko     map[string]string `json:"ko"`
		En     map[string]string `json:"en"`
		ZhHans map[string]string `json:"zh-Hans"`
		ZhHant map[string]string `json:"zh-Hant"`
		Ja     map[string]string `json:"ja"`
	} `json:"i18nAll"`
}
