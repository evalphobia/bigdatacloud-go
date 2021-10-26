package bigdatacloud

type ErrData struct {
	StatusCode  int    `json:"status_code"`
	Status      int64  `json:"status"`
	Description string `json:"description"`
}

func (r *ErrData) SetStatusCode(code int) {
	r.StatusCode = code
}

func (r ErrData) HasError() bool {
	return r.StatusCode >= 400
}

// IPResponse is a response data of IP Geolocation APIs.
// ref: https://www.bigdatacloud.com/ip-geolocation-apis
type IPResponse struct {
	ErrData

	IP                        string       `json:"ip"`
	LocalityLanguageRequested string       `json:"localityLanguageRequested"`
	IsReachableGlobally       bool         `json:"isReachableGlobally"`
	Country                   Country      `json:"country"`
	Location                  Location     `json:"location"`
	LastUpdated               string       `json:"lastUpdated"`
	Network                   Network      `json:"network"`
	Confidence                string       `json:"confidence"`
	ConfidenceArea            []Area       `json:"confidenceArea"`
	SecurityThreat            string       `json:"securityThreat"`
	HazardReport              HazardReport `json:"hazardReport"`
}

type Country struct {
	ISOAlpha2         string             `json:"isoAlpha2"`
	ISOAlpha3         string             `json:"isoAlpha3"`
	M49Code           int64              `json:"m49Code"`
	Name              string             `json:"name"`
	ISOName           string             `json:"isoName"`
	ISONameFull       string             `json:"isoNameFull"`
	ISOAdminLanguages []ISOAdminLanguage `json:"isoAdminLanguages"`
	UNRegion          string             `json:"unRegion"`
	Currency          Currency           `json:"currency"`
	WBRegion          WBRegion           `json:"wbRegion"`
	WBIncomeLevel     WBIncomeLevel      `json:"wbIncomeLevel"`
	CallingCode       string             `json:"callingCode"`
	CountryFlagEmoji  string             `json:"countryFlagEmoji"`
	WikidataID        string             `json:"wikidataId"`
	GeonameID         int64              `json:"geonameId"`
	Continents        []Continent        `json:"continents"`
}

type ISOAdminLanguage struct {
	ISOAlpha3  string `json:"isoAlpha3"`
	ISOAlpha2  string `json:"isoAlpha2"`
	ISOName    string `json:"isoName"`
	NativeName string `json:"nativeName"`
}

type Currency struct {
	NumericCode int64  `json:"numericCode"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	MinorUnits  int64  `json:"minorUnits"`
}

type WBRegion struct {
	ID       string `json:"id"`
	ISO2code string `json:"iso2code"`
	Value    string `json:"value"`
}

type WBIncomeLevel struct {
	ID       string `json:"id"`
	ISO2code string `json:"iso2code"`
	Value    string `json:"value"`
}

type Continent struct {
	Continent     string `json:"continent"`
	ContinentCode string `json:"continentCode"`
}

type Location struct {
	Continent                   string       `json:"continent"`
	ContinentCode               string       `json:"continentCode"`
	ISOPrincipalSubdivision     string       `json:"isoPrincipalSubdivision"`
	ISOPrincipalSubdivisionCode string       `json:"isoPrincipalSubdivisionCode"`
	City                        string       `json:"city"`
	LocalityName                string       `json:"localityName"`
	PostCode                    string       `json:"postcode"`
	Latitude                    float64      `json:"latitude"`
	Longitude                   float64      `json:"longitude"`
	PlusCode                    string       `json:"plusCode"`
	TimeZone                    TimeZone     `json:"timeZone"`
	LocalityInfo                LocalityInfo `json:"localityInfo"`
}

type TimeZone struct {
	IANATimeID             string `json:"ianaTimeId"`
	DisplayName            string `json:"displayName"`
	EffectiveTimeZoneFull  string `json:"effectiveTimeZoneFull"`
	EffectiveTimeZoneShort string `json:"effectiveTimeZoneShort"`
	UTCOffsetSeconds       int64  `json:"utcOffsetSeconds"`
	UTCOffset              string `json:"utcOffset"`
	IsDaylightSavingTime   bool   `json:"isDaylightSavingTime"`
	LocalTime              string `json:"localTime"`
}

type LocalityInfo struct {
	Administrative []LocalityInfoData `json:"administrative"`
	Informative    []LocalityInfoData `json:"informative"`
}

type LocalityInfoData struct {
	Order       int64  `json:"order"`
	AdminLevel  int64  `json:"adminLevel"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ISOName     string `json:"isoName"`
	ISOCode     string `json:"isoCode"`
	WikidataID  string `json:"wikidataId"`
	GeonameID   int64  `json:"geonameId"`
}

type Network struct {
	Registry                string    `json:"registry"`
	RegistryStatus          string    `json:"registryStatus"`
	RegisteredCountry       string    `json:"registeredCountry"`
	RegisteredCountryName   string    `json:"registeredCountryName"`
	Organisation            string    `json:"organisation"`
	IsReachableGlobally     bool      `json:"isReachableGlobally"`
	IsBogon                 bool      `json:"isBogon"`
	BGPPrefix               string    `json:"bgpPrefix"`
	BGPPrefixNetworkAddress string    `json:"bgpPrefixNetworkAddress"`
	BGPPrefixLastAddress    string    `json:"bgpPrefixLastAddress"`
	TotalAddresses          int64     `json:"totalAddresses"`
	Carriers                []Carrier `json:"carriers"`
	ViaCarriers             []Carrier `json:"viaCarriers"`
}

type Carrier struct {
	ASN                    string `json:"asn"`
	ASNNumeric             int64  `json:"asnNumeric"`
	Organisation           string `json:"organisation"`
	Name                   string `json:"name"`
	TotalIpv4Addresses     int64  `json:"totalIpv4Addresses"`
	TotalIpv4Prefixes      int64  `json:"totalIpv4Prefixes"`
	TotalIpv4BogonPrefixes int64  `json:"totalIpv4BogonPrefixes"`
	Rank                   int64  `json:"rank"`
	RankText               string `json:"rankText"`
}

type Area struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type HazardReport struct {
	IsKnownAsTorServer       bool  `json:"isKnownAsTorServer"`
	IsKnownAsVPN             bool  `json:"isKnownAsVpn"`
	IsKnownAsProxy           bool  `json:"isKnownAsProxy"`
	IsSpamhausDrop           bool  `json:"isSpamhausDrop"`
	IsSpamhausEdrop          bool  `json:"isSpamhausEdrop"`
	IsSpamhausAsnDrop        bool  `json:"isSpamhausAsnDrop"`
	IsBlacklistedUceprotect  bool  `json:"isBlacklistedUceprotect"`
	IsBlacklistedBlocklistDe bool  `json:"isBlacklistedBlocklistDe"`
	IsKnownAsMailServer      bool  `json:"isKnownAsMailServer"`
	IsKnownAsPublicRouter    bool  `json:"isKnownAsPublicRouter"`
	IsBogon                  bool  `json:"isBogon"`
	IsUnreachable            bool  `json:"isUnreachable"`
	HostingLikelihood        int64 `json:"hostingLikelihood"` // [0 - 10]
	IsHostingASN             bool  `json:"isHostingAsn"`
	IsCellular               bool  `json:"isCellular"`
}
