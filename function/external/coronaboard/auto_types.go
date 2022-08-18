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
		} `json:"서울"`
		Busan struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"부산"`
		Daegu struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"대구"`
		Incheon struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"인천"`
		Gwangju struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"광주"`
		Daejeon struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"대전"`
		Ulsan struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"울산"`
		Sejong struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"세종"`
		Gyeonggi struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"경기"`
		Gangwon struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"강원"`
		Chungbuk struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"충북"`
		Chungnam struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"충남"`
		Jeonbuk struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"전북"`
		Jeonnam struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"전남"`
		Gyeongbuk struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"경북"`
		Gyeongnam struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"경남"`
		Jeju struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"제주"`
		Oversea struct {
			Date         []string `json:"date"`
			Active       []int    `json:"active"`
			ConfirmedAcc []int    `json:"confirmed_acc"`
			DeathAcc     []int    `json:"death_acc"`
			ReleasedAcc  []int    `json:"released_acc"`
			Confirmed    []int    `json:"confirmed"`
			Death        []int    `json:"death"`
			Released     []int    `json:"released"`
		} `json:"검역"`
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
		} `json:"합계"`
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
		} `json:"아스트라제네카"`
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
		} `json:"화이자"`
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
		} `json:"모더나"`
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
		} `json:"얀센"`
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
		} `json:"노바백스"`
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
		} `json:"기타"`
		Date []string `json:"date"`
	} `json:"chartForVaccineByType"`
	VaccineDataForDashboard struct {
		AstraZeneca struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"아스트라제네카"`
		Pfizer struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"화이자"`
		Moderna struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"모더나"`
		Janssen struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"얀센"`
		Novavax struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"노바백스"`
		Etc struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VacciatendFully     int `json:"vacciatendFully"`
			VacciantedFullyPrev int `json:"vacciantedFully_prev"`
		} `json:"기타"`
		All struct {
			VaccinatedOnce      int `json:"vaccinatedOnce"`
			VaccinatedOncePrev  int `json:"vaccinatedOnce_prev"`
			VaccinatedFully     int `json:"vaccinatedFully"`
			VaccinatedFullyPrev int `json:"vaccinatedFully_prev"`
			VaccinatedBoost     int `json:"vaccinatedBoost"`
			VaccinatedBoostPrev int `json:"vaccinatedBoost_prev"`
		} `json:"합계"`
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
						NAMING_FAILED   int `json:"해외유입"`
						NAMING_FAILED0  int `json:"신천지 관련"`
						NAMING_FAILED1  int `json:"기존해외유입관련"`
						NAMING_FAILED2  int `json:"조사중"`
						NAMING_FAILED3  int `json:"그외 집단 발병"`
						NAMING_FAILED4  int `json:"성북구사랑제일교회"`
						NAMING_FAILED5  int `json:"광화문집회"`
						NAMING_FAILED6  int `json:"클럽 "`
						NAMING_FAILED7  int `json:"용인시우리제일교회"`
						NAMING_FAILED8  int `json:"리치웨이"`
						NAMING_FAILED9  int `json:"구로콜센터"`
						NAMING_FAILED10 int `json:"쿠팡물류센터"`
						NAMING_FAILED11 int `json:"광주방문판매모임"`
						NAMING_FAILED12 int `json:"수도권개척교회모임"`
						NAMING_FAILED13 int `json:"서울강서구 댄스교습"`
						NAMING_FAILED14 int `json:"서울노량진 임용단기학원"`
						NAMING_FAILED15 int `json:"서울 용산구 건설현장"`
						NAMING_FAILED16 int `json:"서울 강서구 종교시설"`
						NAMING_FAILED17 int `json:"서울 송파구 교정시설"`
						NAMING_FAILED18 int `json:"경기 용인시 수지구 교회"`
						NAMING_FAILED19 int `json:"경기 양주시 육류가공업체"`
						IM              int `json:"IM선교회 미인가교육시설"`
						NAMING_FAILED20 int `json:"충북 괴산군/음성군/진천군/안성시 병원"`
						NAMING_FAILED21 int `json:"광주 광산구 요양병원"`
						NAMING_FAILED22 int `json:"광주 서구 교회"`
						NAMING_FAILED23 int `json:"전북 순창군 요양병원"`
						BTJ             int `json:"경북 상주 BTJ열방센터"`
						Num2            int `json:"경북 구미시 종교시설2"`
						NAMING_FAILED24 int `json:"부산 금정구 요양병원"`
						NAMING_FAILED25 int `json:"경남 진주시 기도원"`
						NAMING_FAILED26 int `json:"기타"`
					} `json:"전국"`
					NAMING_FAILED0 struct {
						NAMING_FAILED   int `json:"해외유입"`
						NAMING_FAILED0  int `json:"신천지 관련"`
						NAMING_FAILED1  int `json:"기존해외유입관련"`
						NAMING_FAILED2  int `json:"조사중"`
						NAMING_FAILED3  int `json:"그외 집단 발병"`
						NAMING_FAILED4  int `json:"성북구사랑제일교회"`
						NAMING_FAILED5  int `json:"광화문집회"`
						NAMING_FAILED6  int `json:"클럽 "`
						NAMING_FAILED7  int `json:"용인시우리제일교회"`
						NAMING_FAILED8  int `json:"리치웨이"`
						NAMING_FAILED9  int `json:"구로콜센터"`
						NAMING_FAILED10 int `json:"쿠팡물류센터"`
						NAMING_FAILED11 int `json:"수도권개척교회모임"`
						NAMING_FAILED12 int `json:"서울강서구 댄스교습"`
						NAMING_FAILED13 int `json:"서울노량진 임용단기학원"`
						NAMING_FAILED14 int `json:"서울 용산구 건설현장"`
						NAMING_FAILED15 int `json:"서울 강서구 종교시설"`
						NAMING_FAILED16 int `json:"서울 송파구 교정시설"`
						NAMING_FAILED17 int `json:"기타"`
					} `json:"서울"`
					NAMING_FAILED1 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"광화문집회"`
						NAMING_FAILED6 int `json:"클럽 "`
						NAMING_FAILED7 int `json:"부산 금정구 요양병원"`
						NAMING_FAILED8 int `json:"기타"`
					} `json:"부산"`
					NAMING_FAILED2 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"광화문집회"`
						NAMING_FAILED6 int `json:"클럽 "`
						NAMING_FAILED7 int `json:"용인시우리제일교회"`
						NAMING_FAILED8 int `json:"기타"`
					} `json:"대구"`
					NAMING_FAILED3 struct {
						NAMING_FAILED   int `json:"해외유입"`
						NAMING_FAILED0  int `json:"신천지 관련"`
						NAMING_FAILED1  int `json:"기존해외유입관련"`
						NAMING_FAILED2  int `json:"조사중"`
						NAMING_FAILED3  int `json:"그외 집단 발병"`
						NAMING_FAILED4  int `json:"성북구사랑제일교회"`
						NAMING_FAILED5  int `json:"광화문집회"`
						NAMING_FAILED6  int `json:"클럽 "`
						NAMING_FAILED7  int `json:"용인시우리제일교회"`
						NAMING_FAILED8  int `json:"리치웨이"`
						NAMING_FAILED9  int `json:"구로콜센터"`
						NAMING_FAILED10 int `json:"쿠팡물류센터"`
						NAMING_FAILED11 int `json:"수도권개척교회모임"`
						NAMING_FAILED12 int `json:"서울노량진 임용단기학원"`
						NAMING_FAILED13 int `json:"기타"`
					} `json:"인천"`
					NAMING_FAILED4 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"광화문집회"`
						NAMING_FAILED5 int `json:"광주방문판매모임"`
						IM             int `json:"IM선교회 미인가교육시설"`
						NAMING_FAILED6 int `json:"광주 광산구 요양병원"`
						NAMING_FAILED7 int `json:"광주 서구 교회"`
						NAMING_FAILED8 int `json:"기타"`
					} `json:"광주"`
					NAMING_FAILED5 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"광화문집회"`
						IM             int `json:"IM선교회 미인가교육시설"`
						NAMING_FAILED6 int `json:"기타"`
					} `json:"대전"`
					NAMING_FAILED6 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"광화문집회"`
						NAMING_FAILED5 int `json:"기타"`
					} `json:"울산"`
					NAMING_FAILED7 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"기타"`
					} `json:"세종"`
					NAMING_FAILED8 struct {
						NAMING_FAILED   int `json:"해외유입"`
						NAMING_FAILED0  int `json:"신천지 관련"`
						NAMING_FAILED1  int `json:"기존해외유입관련"`
						NAMING_FAILED2  int `json:"조사중"`
						NAMING_FAILED3  int `json:"그외 집단 발병"`
						NAMING_FAILED4  int `json:"성북구사랑제일교회"`
						NAMING_FAILED5  int `json:"광화문집회"`
						NAMING_FAILED6  int `json:"클럽 "`
						NAMING_FAILED7  int `json:"용인시우리제일교회"`
						NAMING_FAILED8  int `json:"리치웨이"`
						NAMING_FAILED9  int `json:"구로콜센터"`
						NAMING_FAILED10 int `json:"쿠팡물류센터"`
						NAMING_FAILED11 int `json:"수도권개척교회모임"`
						NAMING_FAILED12 int `json:"서울노량진 임용단기학원"`
						NAMING_FAILED13 int `json:"경기 용인시 수지구 교회"`
						NAMING_FAILED14 int `json:"경기 양주시 육류가공업체"`
						IM              int `json:"IM선교회 미인가교육시설"`
						NAMING_FAILED15 int `json:"기타"`
					} `json:"경기"`
					NAMING_FAILED9 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"광화문집회"`
						NAMING_FAILED6 int `json:"용인시우리제일교회"`
						NAMING_FAILED7 int `json:"기타"`
					} `json:"강원"`
					NAMING_FAILED10 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"광화문집회"`
						NAMING_FAILED6 int `json:"클럽 "`
						NAMING_FAILED7 int `json:"충북 괴산군/음성군/진천군/안성시 병원"`
						NAMING_FAILED8 int `json:"기타"`
					} `json:"충북"`
					NAMING_FAILED11 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"기존해외유입관련"`
						NAMING_FAILED1 int `json:"조사중"`
						NAMING_FAILED2 int `json:"그외 집단 발병"`
						NAMING_FAILED3 int `json:"성북구사랑제일교회"`
						NAMING_FAILED4 int `json:"광화문집회"`
						NAMING_FAILED5 int `json:"용인시우리제일교회"`
						NAMING_FAILED6 int `json:"기타"`
					} `json:"충남"`
					NAMING_FAILED12 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"클럽 "`
						NAMING_FAILED6 int `json:"용인시우리제일교회"`
						NAMING_FAILED7 int `json:"전북 순창군 요양병원"`
						NAMING_FAILED8 int `json:"기타"`
					} `json:"전북"`
					NAMING_FAILED13 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"광주방문판매모임"`
						NAMING_FAILED5 int `json:"기타"`
					} `json:"전남"`
					NAMING_FAILED14 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"성북구사랑제일교회"`
						NAMING_FAILED5 int `json:"광화문집회"`
						NAMING_FAILED6 int `json:"용인시우리제일교회"`
						BTJ            int `json:"경북 상주 BTJ열방센터"`
						Num2           int `json:"경북 구미시 종교시설2"`
						NAMING_FAILED7 int `json:"기타"`
					} `json:"경북"`
					NAMING_FAILED15 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"신천지 관련"`
						NAMING_FAILED1 int `json:"기존해외유입관련"`
						NAMING_FAILED2 int `json:"조사중"`
						NAMING_FAILED3 int `json:"그외 집단 발병"`
						NAMING_FAILED4 int `json:"광화문집회"`
						NAMING_FAILED5 int `json:"클럽 "`
						NAMING_FAILED6 int `json:"경남 진주시 기도원"`
						NAMING_FAILED7 int `json:"기타"`
					} `json:"경남"`
					NAMING_FAILED16 struct {
						NAMING_FAILED  int `json:"해외유입"`
						NAMING_FAILED0 int `json:"기존해외유입관련"`
						NAMING_FAILED1 int `json:"조사중"`
						NAMING_FAILED2 int `json:"그외 집단 발병"`
						NAMING_FAILED3 int `json:"기타"`
					} `json:"제주"`
					NAMING_FAILED17 struct {
						NAMING_FAILED int `json:"해외유입"`
					} `json:"검역"`
				} `json:"data"`
			} `json:"all"`
			Recent struct {
				Time string `json:"time"`
				Data struct {
					NAMING_FAILED struct {
						NAMING_FAILED   int `json:"중랑구 아동관련시설 "`
						NAMING_FAILED0  int `json:"강북구 사우나 "`
						NAMING_FAILED1  int `json:"성동구 대학병원"`
						NAMING_FAILED2  int `json:"광진구 음식점"`
						NAMING_FAILED3  int `json:"중구 복지시설"`
						NAMING_FAILED4  int `json:"경기 군포시 교회 "`
						NAMING_FAILED5  int `json:"경기 수원시 권선구 교회"`
						Num2            int `json:"경기 수원시 요양원2 "`
						NAMING_FAILED6  int `json:"경기 남양주시 보육시설"`
						NAMING_FAILED7  int `json:"충남 당진시 유통업체"`
						NAMING_FAILED8  int `json:"충남 서천군 교회"`
						IM              int `json:"IM 선교회 미인가 대안 교육시설"`
						NAMING_FAILED9  int `json:"광주 서구 교회"`
						NAMING_FAILED10 int `json:"대구 북구 사무실"`
						NAMING_FAILED11 int `json:"부산 중구 재활병원"`
						NAMING_FAILED12 int `json:"부산 서구 항운노조"`
						Num3            int `json:"부산 해운대구 일가족3"`
					} `json:"전국"`
					NAMING_FAILED0 struct {
						NAMING_FAILED  int `json:"중랑구 아동관련시설 "`
						NAMING_FAILED0 int `json:"강북구 사우나 "`
						NAMING_FAILED1 int `json:"성동구 대학병원"`
						NAMING_FAILED2 int `json:"광진구 음식점"`
						NAMING_FAILED3 int `json:"중구 복지시설"`
					} `json:"서울"`
					NAMING_FAILED1 struct {
						NAMING_FAILED  int `json:"경기 군포시 교회 "`
						NAMING_FAILED0 int `json:"경기 수원시 권선구 교회"`
						Num2           int `json:"경기 수원시 요양원2 "`
						NAMING_FAILED1 int `json:"경기 남양주시 보육시설"`
						NAMING_FAILED2 int `json:"충남 당진시 유통업체"`
						NAMING_FAILED3 int `json:"충남 서천군 교회"`
						IM             int `json:"IM 선교회 미인가 대안 교육시설"`
						NAMING_FAILED4 int `json:"광주 서구 교회"`
						NAMING_FAILED5 int `json:"대구 북구 사무실"`
						NAMING_FAILED6 int `json:"부산 중구 재활병원"`
						NAMING_FAILED7 int `json:"부산 서구 항운노조"`
						Num3           int `json:"부산 해운대구 일가족3"`
					} `json:"서울외"`
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
					Zero9  int `json:"0-9세"`
					One0   int `json:"10대"`
					Two0   int `json:"20대"`
					Three0 int `json:"30대"`
					Four0  int `json:"40대"`
					Five0  int `json:"50대"`
					Six0   int `json:"60대"`
					Seven0 int `json:"70대"`
					Eight0 int `json:"80대이상"`
				} `json:"confirmed"`
				Death struct {
					Zero9  int `json:"0-9세"`
					One0   int `json:"10대"`
					Two0   int `json:"20대"`
					Three0 int `json:"30대"`
					Four0  int `json:"40대"`
					Five0  int `json:"50대"`
					Six0   int `json:"60대"`
					Seven0 int `json:"70대"`
					Eight0 int `json:"80대이상"`
				} `json:"death"`
				Released struct {
					Zero9  int `json:"0-9세"`
					One0   int `json:"10대"`
					Two0   int `json:"20대"`
					Three0 int `json:"30대"`
					Four0  int `json:"40대"`
					Five0  int `json:"50대"`
					Six0   int `json:"60대"`
					Seven0 int `json:"70대"`
					Eight0 int `json:"80대이상"`
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
