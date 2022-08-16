package status

type CoronaboardI18n struct {
	LangCode                        string `json:"langCode"`
	YmdFormat                       string `json:"ymdFormat"`
	YmFormat                        string `json:"ymFormat"`
	MdFormat                        string `json:"mdFormat"`
	HeadTitle                       string `json:"headTitle"`
	HeadDescription                 string `json:"headDescription"`
	OgLocaleKR                      string `json:"ogLocaleKR"`
	CountryGlobal                   string `json:"countryGlobal"`
	CountryKR                       string `json:"countryKR"`
	CountryUS                       string `json:"countryUS"`
	CountryFR                       string `json:"countryFR"`
	CountryNL                       string `json:"countryNL"`
	LastUpdated                     string `json:"lastUpdated"`
	MainTitle                       string `json:"mainTitle"`
	Confirmed                       string `json:"confirmed"`
	Death                           string `json:"death"`
	Released                        string `json:"released"`
	Recovered                       string `json:"recovered"`
	Critical                        string `json:"critical"`
	FatalityRate                    string `json:"fatalityRate"`
	DeathRate                       string `json:"deathRate"`
	ReleasedRate                    string `json:"releasedRate"`
	RecoveryRate                    string `json:"recoveryRate"`
	AffectedCountry                 string `json:"affectedCountry"`
	TotalTested                     string `json:"totalTested"`
	Testing                         string `json:"testing"`
	Negative                        string `json:"negative"`
	Hospitalized                    string `json:"hospitalized"`
	IntensiveCare                   string `json:"intensiveCare"`
	Notice                          string `json:"notice"`
	NewFeature                      string `json:"newFeature"`
	MenuMobileWorld                 string `json:"menuMobileWorld"`
	MenuMobileKorea                 string `json:"menuMobileKorea"`
	MenuMobileFrance                string `json:"menuMobileFrance"`
	MenuMobileNetherlands           string `json:"menuMobileNetherlands"`
	MenuMobileUS                    string `json:"menuMobileUS"`
	MenuMobileChart                 string `json:"menuMobileChart"`
	MenuMobileChartWorld            string `json:"menuMobileChartWorld"`
	MenuMobileChartKorea            string `json:"menuMobileChartKorea"`
	MenuMobilePrevention            string `json:"menuMobilePrevention"`
	MenuMobilePatientLog            string `json:"menuMobilePatientLog"`
	MenuMobileNews                  string `json:"menuMobileNews"`
	MenuMobileCountrySelect         string `json:"menuMobileCountrySelect"`
	MenuPcWorld                     string `json:"menuPcWorld"`
	MenuPcKorea                     string `json:"menuPcKorea"`
	MenuPUS                         string `json:"menuPUS"`
	MenuPCFrance                    string `json:"menuPCFrance"`
	MenuPCNetherlands               string `json:"menuPCNetherlands"`
	MenuPcChart                     string `json:"menuPcChart"`
	MenuPcChartWorld                string `json:"menuPcChartWorld"`
	MenuPcChartKorea                string `json:"menuPcChartKorea"`
	MenuPcPrevention                string `json:"menuPcPrevention"`
	MenuPcPatientLog                string `json:"menuPcPatientLog"`
	MenuPcNews                      string `json:"menuPcNews"`
	MenuPcCountrySelect             string `json:"menuPcCountrySelect"`
	Rok                             string `json:"rok"`
	Usa                             string `json:"usa"`
	World                           string `json:"world"`
	Netherlands                     string `json:"netherlands"`
	SlideKorea                      string `json:"slideKorea"`
	SlideUS                         string `json:"slideUS"`
	SlideWorld                      string `json:"slideWorld"`
	SlideChart                      string `json:"slideChart"`
	SlideChartUS                    string `json:"slideChartUS"`
	SlideChartKr                    string `json:"slideChartKr"`
	SlideChartFr                    string `json:"slideChartFr"`
	SlideChartNl                    string `json:"slideChartNl"`
	SlidePrevention                 string `json:"slidePrevention"`
	SlideBriefing                   string `json:"slideBriefing"`
	SlidePatientLog                 string `json:"slidePatientLog"`
	SlideNews                       string `json:"slideNews"`
	SlideGlobalNews                 string `json:"slideGlobalNews"`
	SlideKoreaNews                  string `json:"slideKoreaNews"`
	SlideYoutube                    string `json:"slideYoutube"`
	SlideVirusComparision           string `json:"slideVirusComparision"`
	SlideReference                  string `json:"slideReference"`
	SlideCreator                    string `json:"slideCreator"`
	EmergencyContact                string `json:"emergencyContact"`
	SlideImmigrant                  string `json:"slideImmigrant"`
	DashboardLegend1                string `json:"dashboardLegend1"`
	DashboardLegend2                string `json:"dashboardLegend2"`
	Terms                           string `json:"terms"`
	TermGlobal                      string `json:"termGlobal"`
	TermKR                          string `json:"termKR"`
	WorldLegend2                    string `json:"worldLegend2"`
	WorldLegend3                    string `json:"worldLegend3"`
	WorldLegend4                    string `json:"worldLegend4"`
	WorldLegend5                    string `json:"worldLegend5"`
	WorldLegend6                    string `json:"worldLegend6"`
	KoreaLegend1                    string `json:"koreaLegend1"`
	KoreaLegend2                    string `json:"koreaLegend2"`
	KoreaLegend3                    string `json:"koreaLegend3"`
	KoreaLegend4                    string `json:"koreaLegend4"`
	KoreaTableLegend1               string `json:"koreaTableLegend1"`
	KrByRegion                      string `json:"krByRegion"`
	Daily                           string `json:"daily"`
	Monthly                         string `json:"monthly"`
	Accumulated                     string `json:"accumulated"`
	DailyAccumulatedAll             string `json:"dailyAccumulatedAll"`
	RecentSevenDays                 string `json:"recentSevenDays"`
	RecentDays                      string `json:"recentDays"`
	CovidScreeningCenter            string `json:"covidScreeningCenter"`
	PublicHealthCenter              string `json:"publicHealthCenter"`
	Call120                         string `json:"call120"`
	Call1339                        string `json:"call1339"`
	SelectEach                      string `json:"selectEach"`
	SelectSum                       string `json:"selectSum"`
	DeselectAll                     string `json:"deselectAll"`
	SelectAll                       string `json:"selectAll"`
	SelectCountries                 string `json:"selectCountries"`
	SelectRegions                   string `json:"selectRegions"`
	SelectRegion                    string `json:"selectRegion"`
	SelectColumn                    string `json:"selectColumn"`
	NAMING_FAILED                   string `json:"전국"`
	NAMING_FAILED0                  string `json:"서울"`
	NAMING_FAILED1                  string `json:"서울외"`
	NAMING_FAILED2                  string `json:"부산"`
	NAMING_FAILED3                  string `json:"대구"`
	NAMING_FAILED4                  string `json:"인천"`
	NAMING_FAILED5                  string `json:"광주"`
	NAMING_FAILED6                  string `json:"대전"`
	NAMING_FAILED7                  string `json:"울산"`
	NAMING_FAILED8                  string `json:"세종"`
	NAMING_FAILED9                  string `json:"경기"`
	NAMING_FAILED10                 string `json:"강원"`
	NAMING_FAILED11                 string `json:"충북"`
	NAMING_FAILED12                 string `json:"충남"`
	NAMING_FAILED13                 string `json:"전북"`
	NAMING_FAILED14                 string `json:"전남"`
	NAMING_FAILED15                 string `json:"경북"`
	NAMING_FAILED16                 string `json:"경남"`
	NAMING_FAILED17                 string `json:"제주"`
	NAMING_FAILED18                 string `json:"검역"`
	NAMING_FAILED19                 string `json:"기타"`
	ChartLegendKoreaMap             string `json:"chartLegendKoreaMap"`
	Region                          string `json:"region"`
	Incidence                       string `json:"incidence"`
	ActiveCase                      string `json:"activeCase"`
	AccConfirmed                    string `json:"accConfirmed"`
	AccReleased                     string `json:"accReleased"`
	AccRecovered                    string `json:"accRecovered"`
	AccDeath                        string `json:"accDeath"`
	AccNegative                     string `json:"accNegative"`
	ConfirmedShort                  string `json:"confirmedShort"`
	TotalTestedShort                string `json:"totalTestedShort"`
	DeathShort                      string `json:"deathShort"`
	HospitalizedShort               string `json:"hospitalizedShort"`
	IntensiveCareShort              string `json:"intensiveCareShort"`
	ChartTitleKoreaTimeseries       string `json:"chartTitleKoreaTimeseries"`
	ChartTitleKoreaTrendByRegion    string `json:"chartTitleKoreaTrendByRegion"`
	ConfirmedRate                   string `json:"confirmedRate"`
	ConfirmedRateAcc                string `json:"confirmedRateAcc"`
	ChartTitleKoreaTestStatus       string `json:"chartTitleKoreaTestStatus"`
	ChartLegendKoreaTestStatus      string `json:"chartLegendKoreaTestStatus"`
	ChartTitleGlobalTrend           string `json:"chartTitleGlobalTrend"`
	ChartLegendGlobalTrend          string `json:"chartLegendGlobalTrend"`
	Country                         string `json:"country"`
	DeathRateCol                    string `json:"deathRateCol"`
	RecoveryRateCol                 string `json:"recoveryRateCol"`
	ChartTitleGlobalDaily           string `json:"chartTitleGlobalDaily"`
	ChartTitleKoreaStatusByRegion   string `json:"chartTitleKoreaStatusByRegion"`
	ChartTitleKoreaBySource         string `json:"chartTitleKoreaBySource"`
	ChartLegendKoreaByRecentSource  string `json:"chartLegendKoreaByRecentSource"`
	ChartLegendKoreaBySource        string `json:"chartLegendKoreaBySource"`
	ChartTitleKoreaByGender         string `json:"chartTitleKoreaByGender"`
	ChartTitleKoreaByAge            string `json:"chartTitleKoreaByAge"`
	ChartLegendCDC                  string `json:"chartLegendCDC"`
	ChartLegendTotalItem            string `json:"chartLegendTotalItem"`
	ChartLegendRealtime             string `json:"chartLegendRealtime"`
	PeopleCountUnit                 string `json:"peopleCountUnit"`
	SourceChartFormatter            string `json:"sourceChartFormatter"`
	ChartLegendGenderReleased       string `json:"chartLegendGenderReleased"`
	ChartLegendAgeReleased          string `json:"chartLegendAgeReleased"`
	ChartLegendAgeDeathRate         string `json:"chartLegendAgeDeathRate"`
	ChartLegendAgeReleasedRate      string `json:"chartLegendAgeReleasedRate"`
	Male                            string `json:"male"`
	Female                          string `json:"female"`
	ViewCount                       string `json:"viewCount"`
	CountryOptionGlobal             string `json:"countryOptionGlobal"`
	France                          string `json:"france"`
	Covid19                         string `json:"covid19"`
	DiseaseName                     string `json:"diseaseName"`
	OccuredYear                     string `json:"occuredYear"`
	Population                      string `json:"population"`
	LogScale                        string `json:"logScale"`
	LinearScale                     string `json:"linearScale"`
	ChartTitleGlobalLog             string `json:"chartTitleGlobalLog"`
	GlobalLogChartXaxisDeath        string `json:"globalLogChartXaxisDeath"`
	GlobalLogChartXaxisConfirmed    string `json:"globalLogChartXaxisConfirmed"`
	SelectCountryTitle              string `json:"selectCountryTitle"`
	SelectCountrySub                string `json:"selectCountrySub"`
	Disclaimer                      string `json:"disclaimer"`
	Donation                        string `json:"donation"`
	CsvNotice                       string `json:"csvNotice"`
	ShareTitle                      string `json:"shareTitle"`
	SupportTitle                    string `json:"supportTitle"`
	ShowMore                        string `json:"showMore"`
	TooltipJpCruise                 string `json:"tooltipJpCruise"`
	Yesterday                       string `json:"yesterday"`
	Today                           string `json:"today"`
	ChartTitleGlobalPieChart        string `json:"chartTitleGlobalPieChart"`
	Etc                             string `json:"etc"`
	SocialDistancingRules           string `json:"socialDistancingRules"`
	Stage2                          string `json:"stage2"`
	Stage25                         string `json:"stage25"`
	Stage3                          string `json:"stage3"`
	SdLocalLevel                    string `json:"sdLocalLevel"`
	SdNationLevel                   string `json:"sdNationLevel"`
	Sd25Cafe                        string `json:"sd25-cafe"`
	Sd3Cafe                         string `json:"sd3-cafe"`
	Sd2Public                       string `json:"sd2-public"`
	Sd25Public                      string `json:"sd25-public"`
	Sd3BanDetail                    string `json:"sd3-ban_detail"`
	Sd2Enter                        string `json:"sd2-enter"`
	Sd25Enter                       string `json:"sd25-enter"`
	Sd3Ban                          string `json:"sd3-ban"`
	Sd2Event                        string `json:"sd2-event"`
	Sd25Event                       string `json:"sd25-event"`
	Sd3Event                        string `json:"sd3-event"`
	Sd2Academy                      string `json:"sd2-academy"`
	Sd25Academy                     string `json:"sd25-academy"`
	Sd2Waterpark                    string `json:"sd2-waterpark"`
	Sd25Waterpark                   string `json:"sd25-waterpark"`
	Sd2Mart                         string `json:"sd2-mart"`
	Sd25Mart                        string `json:"sd25-mart"`
	Sd3Mart                         string `json:"sd3-mart"`
	Sd2School                       string `json:"sd2-school"`
	Sd25School                      string `json:"sd25-school"`
	Sd3School                       string `json:"sd3-school"`
	Sd2Sports                       string `json:"sd2-sports"`
	Sd25Sports                      string `json:"sd25-sports"`
	Sd3Sports                       string `json:"sd3-sports"`
	Sd2Church                       string `json:"sd2-church"`
	Sd25Church                      string `json:"sd25-church"`
	Sd3Church                       string `json:"sd3-church"`
	SelectSeoul                     string `json:"selectSeoul"`
	SelectKorea                     string `json:"selectKorea"`
	SelectNone                      string `json:"selectNone"`
	Vaccinated                      string `json:"vaccinated"`
	VaccinatedNew                   string `json:"vaccinatedNew"`
	SlideChartVaccine               string `json:"slideChartVaccine"`
	MenuMobileVaccine               string `json:"menuMobileVaccine"`
	MenuPcVaccine                   string `json:"menuPcVaccine"`
	VaccineLegend1                  string `json:"vaccineLegend1"`
	VaccineLegend2                  string `json:"vaccineLegend2"`
	ChartTitleVaccineTrendByRegion  string `json:"chartTitleVaccineTrendByRegion"`
	ChartTitleVaccineTrend          string `json:"chartTitleVaccineTrend"`
	ChartLegendVaccineTrendByRegion string `json:"chartLegendVaccineTrendByRegion"`
	TableInfoVaccingTredByRegion    string `json:"tableInfoVaccingTredByRegion"`
	Astrazeneca                     string `json:"astrazeneca"`
	Pfizer                          string `json:"pfizer"`
	TableLegendVaccinated           string `json:"tableLegendVaccinated"`
	ChartInfoVaccineTrendByType     string `json:"chartInfoVaccineTrendByType"`
	VaccinatedTotal                 string `json:"vaccinatedTotal"`
	VaccinationStatus               string `json:"vaccinationStatus"`
	VaccinatedOnce                  string `json:"vaccinatedOnce"`
	VaccinatedFully                 string `json:"vaccinatedFully"`
	VaccinatedBoost                 string `json:"vaccinatedBoost"`
	Total                           string `json:"total"`
	VaccinatedOnceRatio             string `json:"vaccinatedOnceRatio"`
	VaccinatedFullyRatio            string `json:"vaccinatedFullyRatio"`
	VaccinatedBoostRatio            string `json:"vaccinatedBoostRatio"`
	WorldLegend8                    string `json:"worldLegend8"`
	Cn                              string `json:"CN"`
	Kr                              string `json:"KR"`
	Ir                              string `json:"IR"`
	It                              string `json:"IT"`
	JPCruise                        string `json:"JPCruise"`
	Jp                              string `json:"JP"`
	Fr                              string `json:"FR"`
	De                              string `json:"DE"`
	Es                              string `json:"ES"`
	Sg                              string `json:"SG"`
	Us                              string `json:"US"`
	Hk                              string `json:"HK"`
	Kw                              string `json:"KW"`
	Gb                              string `json:"GB"`
	Bh                              string `json:"BH"`
	Ch                              string `json:"CH"`
	Th                              string `json:"TH"`
	Tw                              string `json:"TW"`
	Au                              string `json:"AU"`
	My                              string `json:"MY"`
	No                              string `json:"NO"`
	Ca                              string `json:"CA"`
	Iq                              string `json:"IQ"`
	At                              string `json:"AT"`
	Nl                              string `json:"NL"`
	Se                              string `json:"SE"`
	Ae                              string `json:"AE"`
	Vn                              string `json:"VN"`
	Be                              string `json:"BE"`
	Lb                              string `json:"LB"`
	Il                              string `json:"IL"`
	Om                              string `json:"OM"`
	Is                              string `json:"IS"`
	Sm                              string `json:"SM"`
	Mo                              string `json:"MO"`
	Hr                              string `json:"HR"`
	Qa                              string `json:"QA"`
	Ec                              string `json:"EC"`
	Fi                              string `json:"FI"`
	Gr                              string `json:"GR"`
	Dk                              string `json:"DK"`
	In                              string `json:"IN"`
	Mx                              string `json:"MX"`
	Dz                              string `json:"DZ"`
	Cz                              string `json:"CZ"`
	Pk                              string `json:"PK"`
	Ro                              string `json:"RO"`
	Ph                              string `json:"PH"`
	Az                              string `json:"AZ"`
	Ge                              string `json:"GE"`
	Ru                              string `json:"RU"`
	Br                              string `json:"BR"`
	Eg                              string `json:"EG"`
	ID                              string `json:"ID"`
	Pt                              string `json:"PT"`
	Af                              string `json:"AF"`
	Ad                              string `json:"AD"`
	Am                              string `json:"AM"`
	By                              string `json:"BY"`
	Kh                              string `json:"KH"`
	Do                              string `json:"DO"`
	Ee                              string `json:"EE"`
	Ie                              string `json:"IE"`
	Jo                              string `json:"JO"`
	Lv                              string `json:"LV"`
	Lt                              string `json:"LT"`
	Lu                              string `json:"LU"`
	Mk                              string `json:"MK"`
	Mc                              string `json:"MC"`
	Ma                              string `json:"MA"`
	Np                              string `json:"NP"`
	Nz                              string `json:"NZ"`
	Ng                              string `json:"NG"`
	Sa                              string `json:"SA"`
	Sn                              string `json:"SN"`
	Lk                              string `json:"LK"`
	Tn                              string `json:"TN"`
	Ua                              string `json:"UA"`
	Ar                              string `json:"AR"`
	Li                              string `json:"LI"`
	Pl                              string `json:"PL"`
	Cl                              string `json:"CL"`
	Hu                              string `json:"HU"`
	Fo                              string `json:"FO"`
	Si                              string `json:"SI"`
	Ps                              string `json:"PS"`
	Ba                              string `json:"BA"`
	Za                              string `json:"ZA"`
	Cr                              string `json:"CR"`
	Gi                              string `json:"GI"`
	Bt                              string `json:"BT"`
	Rs                              string `json:"RS"`
	Sk                              string `json:"SK"`
	Cm                              string `json:"CM"`
	Va                              string `json:"VA"`
	Pe                              string `json:"PE"`
	Tg                              string `json:"TG"`
	Co                              string `json:"CO"`
	Mt                              string `json:"MT"`
	Mv                              string `json:"MV"`
	Md                              string `json:"MD"`
	Py                              string `json:"PY"`
	Bg                              string `json:"BG"`
	Gf                              string `json:"GF"`
	Bd                              string `json:"BD"`
	Al                              string `json:"AL"`
	Gy                              string `json:"GY"`
	Mq                              string `json:"MQ"`
	Mn                              string `json:"MN"`
	Cy                              string `json:"CY"`
	Mf                              string `json:"MF"`
	Bn                              string `json:"BN"`
	Bf                              string `json:"BF"`
	Notused                         string `json:"NOTUSED"`
	Pa                              string `json:"PA"`
	Bl                              string `json:"BL"`
	Sx                              string `json:"SX"`
	Bo                              string `json:"BO"`
	Cd                              string `json:"CD"`
	Jm                              string `json:"JM"`
	Tr                              string `json:"TR"`
	Hn                              string `json:"HN"`
	Cu                              string `json:"CU"`
	Re                              string `json:"RE"`
	Gh                              string `json:"GH"`
	Ci                              string `json:"CI"`
	Pf                              string `json:"PF"`
	Vc                              string `json:"VC"`
	Tt                              string `json:"TT"`
	Sr                              string `json:"SR"`
	Lc                              string `json:"LC"`
	Cw                              string `json:"CW"`
	Mr                              string `json:"MR"`
	Ke                              string `json:"KE"`
	Pr                              string `json:"PR"`
	Aw                              string `json:"AW"`
	Uy                              string `json:"UY"`
	Gt                              string `json:"GT"`
	Ve                              string `json:"VE"`
	Ag                              string `json:"AG"`
	Et                              string `json:"ET"`
	Sd                              string `json:"SD"`
	Gn                              string `json:"GN"`
	Ga                              string `json:"GA"`
	Gp                              string `json:"GP"`
	Ky                              string `json:"KY"`
	Kz                              string `json:"KZ"`
	Sz                              string `json:"SZ"`
	Na                              string `json:"NA"`
	Sc                              string `json:"SC"`
	Yt                              string `json:"YT"`
	Rw                              string `json:"RW"`
	Vi                              string `json:"VI"`
	Gq                              string `json:"GQ"`
	Xk                              string `json:"XK"`
	Uz                              string `json:"UZ"`
	Gu                              string `json:"GU"`
	Cf                              string `json:"CF"`
	Cg                              string `json:"CG"`
	Bs                              string `json:"BS"`
	Tz                              string `json:"TZ"`
	So                              string `json:"SO"`
	Gl                              string `json:"GL"`
	Lr                              string `json:"LR"`
	Bj                              string `json:"BJ"`
	Bb                              string `json:"BB"`
	Gm                              string `json:"GM"`
	Kg                              string `json:"KG"`
	Me                              string `json:"ME"`
	Ms                              string `json:"MS"`
	Zm                              string `json:"ZM"`
	Dj                              string `json:"DJ"`
	Nc                              string `json:"NC"`
	Mu                              string `json:"MU"`
	Bm                              string `json:"BM"`
	Sv                              string `json:"SV"`
	Fj                              string `json:"FJ"`
	Ni                              string `json:"NI"`
	Td                              string `json:"TD"`
	Im                              string `json:"IM"`
	Ne                              string `json:"NE"`
	Ht                              string `json:"HT"`
	Ao                              string `json:"AO"`
	Cv                              string `json:"CV"`
	Sj                              string `json:"SJ"`
	Pg                              string `json:"PG"`
	Mg                              string `json:"MG"`
	Zw                              string `json:"ZW"`
	Tl                              string `json:"TL"`
	Er                              string `json:"ER"`
	Ug                              string `json:"UG"`
	Mz                              string `json:"MZ"`
	Gd                              string `json:"GD"`
	Dm                              string `json:"DM"`
	Sy                              string `json:"SY"`
	Tc                              string `json:"TC"`
	Bz                              string `json:"BZ"`
	Mm                              string `json:"MM"`
	La                              string `json:"LA"`
	Ly                              string `json:"LY"`
	Gw                              string `json:"GW"`
	Ml                              string `json:"ML"`
	Vg                              string `json:"VG"`
	Xn                              string `json:"XN"`
	Ai                              string `json:"AI"`
	PACruise                        string `json:"PACruise"`
	Bw                              string `json:"BW"`
	Sl                              string `json:"SL"`
	Bi                              string `json:"BI"`
	NLCarib                         string `json:"NLCarib"`
	Mw                              string `json:"MW"`
	Fk                              string `json:"FK"`
	Eh                              string `json:"EH"`
	Ss                              string `json:"SS"`
	Pm                              string `json:"PM"`
	St                              string `json:"ST"`
	NorthAmerica                    string `json:"NorthAmerica"`
	Europe                          string `json:"Europe"`
	Asia                            string `json:"Asia"`
	SouthAmerica                    string `json:"SouthAmerica"`
	Oceania                         string `json:"Oceania"`
	Africa                          string `json:"Africa"`
	Ye                              string `json:"YE"`
	Tj                              string `json:"TJ"`
	Km                              string `json:"KM"`
	Ls                              string `json:"LS"`
	Sb                              string `json:"SB"`
	Wf                              string `json:"WF"`
	Mh                              string `json:"MH"`
	Vu                              string `json:"VU"`
	Ws                              string `json:"WS"`
	Fm                              string `json:"FM"`
	Sh                              string `json:"SH"`
	Pw                              string `json:"PW"`
	To                              string `json:"TO"`
	Ki                              string `json:"KI"`
	Ck                              string `json:"CK"`
	Nu                              string `json:"NU"`
	Nr                              string `json:"NR"`
	Kp                              string `json:"KP"`
	Tv                              string `json:"TV"`
	Two0220811T022806Z              string `json:"20220811T022806Z"`
	Two0220814T103255Z              string `json:"20220814T103255Z"`
	Two0220727T124850Z              string `json:"20220727T124850Z"`
}

type CoronaBoardData struct {
	LastUpdated   int `json:"lastUpdated"`
	StatGlobalNow []struct {
		Cc              string `json:"cc"`
		Confirmed       int    `json:"confirmed"`
		Death           int    `json:"death"`
		Released        int    `json:"released"`
		Tested          int    `json:"tested"`
		Critical        int    `json:"critical"`
		Active          int    `json:"active"`
		ConfirmedPrev   int    `json:"confirmed_prev"`
		DeathPrev       int    `json:"death_prev"`
		ReleasedPrev    int    `json:"released_prev"`
		ActivePrev      int    `json:"active_prev"`
		TestedPrev      int    `json:"tested_prev"`
		CriticalPrev    int    `json:"critical_prev"`
		ConfirmedXest   int    `json:"confirmed_xest"`
		DeathXest       int    `json:"death_xest"`
		ReleasedXest    int    `json:"released_xest"`
		ActiveXest      int    `json:"active_xest"`
		TestedXest      int    `json:"tested_xest"`
		CriticalXest    int    `json:"critical_xest"`
		Population      int    `json:"population"`
		Incidence       int    `json:"incidence"`
		Flag            string `json:"flag"`
		EhpadConfirmed  int    `json:"ehpadConfirmed,omitempty"`
		EhpadProbable   int    `json:"ehpadProbable,omitempty"`
		EhpadDeath      int    `json:"ehpadDeath,omitempty"`
		Testing         int    `json:"testing,omitempty"`
		Negative        int    `json:"negative,omitempty"`
		FromOversea     int    `json:"fromOversea,omitempty"`
		TestingPrev     int    `json:"testing_prev,omitempty"`
		NegativePrev    int    `json:"negative_prev,omitempty"`
		TestingXest     int    `json:"testing_xest,omitempty"`
		NegativeXest    int    `json:"negative_xest,omitempty"`
		VaccinatedFully int    `json:"vaccinatedFully,omitempty"`
	} `json:"statGlobalNow"`
	StatDomesticNow []struct {
		Region              string `json:"region"`
		Confirmed           int    `json:"confirmed,omitempty"`
		Death               int    `json:"death,omitempty"`
		Released            int    `json:"released,omitempty"`
		VaccinatedOnce      int    `json:"vaccinatedOnce,omitempty"`
		VaccinatedFully     int    `json:"vaccinatedFully,omitempty"`
		Active              int    `json:"active"`
		ConfirmedPrev       int    `json:"confirmed_prev"`
		ReleasedPrev        int    `json:"released_prev"`
		DeathPrev           int    `json:"death_prev"`
		ActivePrev          int    `json:"active_prev"`
		VaccinatedOncePrev  int    `json:"vaccinatedOnce_prev"`
		VaccinatedFullyPrev int    `json:"vaccinatedFully_prev"`
		VaccinatedBoostPrev int    `json:"vaccinatedBoost_prev"`
		VaccinatedBoost     int    `json:"vaccinatedBoost,omitempty"`
	} `json:"statDomesticNow"`
	ChartForGlobal struct {
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
			Active       []int         `json:"active"`
			ConfirmedAcc []int         `json:"confirmed_acc"`
			DeathAcc     []int         `json:"death_acc"`
			ReleasedAcc  []int         `json:"released_acc"`
			CriticalAcc  []interface{} `json:"critical_acc"`
			Confirmed    []int         `json:"confirmed"`
			Death        []int         `json:"death"`
			Released     []int         `json:"released"`
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
		Ko     CoronaboardI18n `json:"ko"`
		En     CoronaboardI18n `json:"en"`
		ZhHans CoronaboardI18n `json:"zh-Hans"`
		ZhHant CoronaboardI18n `json:"zh-Hant"`
		Ja     CoronaboardI18n `json:"ja"`
	} `json:"i18nAll"`
}
