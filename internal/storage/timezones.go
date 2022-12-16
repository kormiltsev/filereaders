package storage

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type TZ struct {
	FileAddress string  `json:"file_address"`
	CurrentTZ   string  `json:"current_timezone"`
	HistoryTZ   History `json:"history"`
	ListOfTZ    map[string]string
}
type History []struct {
	StartTime int64 //seconds (10 chars)
	Timezone  string
}

var Parsedzones = TZ{
	FileAddress: "data/timezone.json",
	CurrentTZ:   "Europe/Berlin",
	HistoryTZ:   nil,
	ListOfTZ:    a, //make(map[string]string, 0),
}

func ParseJsonTimeZones() (*TZ, error) {
	file := "data/timezone.json"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		log.Println("not open file ", file, err)
		return CreateNewFile()
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&Parsedzones)
	if err == io.EOF {
		log.Println("New .json created")
		return CreateNewFile()
	}
	if err != nil {
		log.Println("cand decode from json", err)
	}
	return &Parsedzones, nil
}

// for empty
var Hist = History{
	{1662239320, "Europe/Berlin"},
	{1000000000, "Europe/Moscow"},
}
var Exportzones = TZ{
	FileAddress: "data/timezone.json",
	CurrentTZ:   "Europe/Berlin",
	HistoryTZ:   Hist,
	ListOfTZ:    a, //make(map[string]string, 0),
}

func CreateNewFile() (*TZ, error) {
	f, err := os.OpenFile(Exportzones.FileAddress, os.O_CREATE|os.O_WRONLY, 0777) //os.O_APPEND|
	if err != nil {
		log.Println("not directory ", Exportzones.FileAddress, err)
		return nil, err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	enc.Encode(&Exportzones)
	if err != nil {
		log.Println("cand encode to json", err)
		return nil, err
	}
	return &Exportzones, nil
}

var a = map[string]string{
	"Africa/Cairo":           "Egypt Standard Time",
	"Africa/Casablanca":      "Morocco Standard Time",
	"Africa/Johannesburg":    "South Africa Standard Time",
	"Africa/Khartoum":        "Sudan Standard Time",
	"Africa/Lagos":           "W. Central Africa Standard Time",
	"Africa/Nairobi":         "E. Africa Standard Time",
	"Africa/Sao_Tome":        "Sao Tome Standard Time",
	"Africa/Tripoli":         "Libya Standard Time",
	"Africa/Windhoek":        "Namibia Standard Time",
	"America/Adak":           "Aleutian Standard Time",
	"America/Anchorage":      "Alaskan Standard Time",
	"America/Araguaina":      "Tocantins Standard Time",
	"America/Asuncion":       "Paraguay Standard Time",
	"America/Bahia":          "Bahia Standard Time",
	"America/Bogota":         "SA Pacific Standard Time",
	"America/Buenos_Aires":   "Argentina Standard Time",
	"America/Cancun":         "Eastern Standard Time (Mexico)",
	"America/Caracas":        "Venezuela Standard Time",
	"America/Cayenne":        "SA Eastern Standard Time",
	"America/Chicago":        "Central Standard Time",
	"America/Chihuahua":      "Mountain Standard Time (Mexico)",
	"America/Cuiaba":         "Central Brazilian Standard Time",
	"America/Denver":         "Mountain Standard Time",
	"America/Godthab":        "Greenland Standard Time",
	"America/Grand_Turk":     "Turks And Caicos Standard Time",
	"America/Guatemala":      "Central America Standard Time",
	"America/Halifax":        "Atlantic Standard Time",
	"America/Havana":         "Cuba Standard Time",
	"America/Indianapolis":   "US Eastern Standard Time",
	"America/La_Paz":         "SA Western Standard Time",
	"America/Los_Angeles":    "Pacific Standard Time",
	"America/Mexico_City":    "Central Standard Time (Mexico)",
	"America/Miquelon":       "Saint Pierre Standard Time",
	"America/Montevideo":     "Montevideo Standard Time",
	"America/New_York":       "Eastern Standard Time",
	"America/Phoenix":        "US Mountain Standard Time",
	"America/Port-au-P:ince": "Haiti Standard Time",
	"America/Punta_Arenas":   "Magallanes Standard Time",
	"America/Regina":         "Canada Central Standard Time",
	"America/Santiago":       "Pacific SA Standard Time",
	"America/Sao_Paulo":      "E. South America Standard Time",
	"America/St_Johns":       "Newfoundland Standard Time",
	"America/Tijuana":        "Pacific Standard Time (Mexico)",
	"Asia/Almaty":            "Central Asia Standard Time",
	"Asia/Amman":             "Jordan Standard Time",
	"Asia/Baghdad":           "Arabic Standard Time",
	"Asia/Baku":              "Azerbaijan Standard Time",
	"Asia/Bangkok":           "SE Asia Standard Time",
	"Asia/Barnaul":           "Altai Standard Time",
	"Asia/Beirut":            "Middle East Standard Time",
	"Asia/Calcutta":          "India Standard Time",
	"Asia/Chita":             "Transbaikal Standard Time",
	"Asia/Colombo":           "Sri Lanka Standard Time",
	"Asia/Damascus":          "Syria Standard Time",
	"Asia/Dhaka":             "Bangladesh Standard Time",
	"Asia/Dubai":             "Arabian Standard Time",
	"Asia/Hebron":            "West Bank Standard Time",
	"Asia/Hovd":              "W. Mongolia Standard Time",
	"Asia/Irkutsk":           "North Asia East Standard Time",
	"Asia/Jerusalem":         "Israel Standard Time",
	"Asia/Kabul":             "Afghanistan Standard Time",
	"Asia/Kamchatka":         "Russia Time Zone 11",
	"Asia/Karachi":           "Pakistan Standard Time",
	"Asia/Katmandu":          "Nepal Standard Time",
	"Asia/Krasnoyarsk":       "North Asia Standard Time",
	"Asia/Magadan":           "Magadan Standard Time",
	"Asia/Novosibirsk":       "N. Central Asia Standard Time",
	"Asia/Omsk":              "Omsk Standard Time",
	"Asia/Pyongyang":         "North Korea Standard Time",
	"Asia/Qyzylorda":         "Qyzylorda Standard Time",
	"Asia/Rangoon":           "Myanmar Standard Time",
	"Asia/Riyadh":            "Arab Standard Time",
	"Asia/Sakhalin":          "Sakhalin Standard Time",
	"Asia/Seoul":             "Korea Standard Time",
	"Asia/Shanghai":          "China Standard Time",
	"Asia/Singapore":         "Singapore Standard Time",
	"Asia/Srednekolymsk":     "Russia Time Zone 10",
	"Asia/Taipei":            "Taipei Standard Time",
	"Asia/Tashkent":          "West Asia Standard Time",
	"Asia/Tbilisi":           "Georgian Standard Time",
	"Asia/Tehran":            "Iran Standard Time",
	"Asia/Tokyo":             "Tokyo Standard Time",
	"Asia/Tomsk":             "Tomsk Standard Time",
	"Asia/Ulaanbaatar":       "Ulaanbaatar Standard Time",
	"Asia/Vladivostok":       "Vladivostok Standard Time",
	"Asia/Yakutsk":           "Yakutsk Standard Time",
	"Asia/Yekaterinburg":     "Ekaterinburg Standard Time",
	"Asia/Yerevan":           "Caucasus Standard Time",
	"Atlantic/Azores":        "Azores Standard Time",
	"Atlantic/Cape_Verde":    "Cape Verde Standard Time",
	"Atlantic/Reykjavik":     "Greenwich Standard Time",
	"Australia/Adelaide":     "Cen. Australia Standard Time",
	"Australia/Brisbane":     "E. Australia Standard Time",
	"Australia/Darwin":       "AUS Central Standard Time",
	"Australia/Eucla":        "Aus Central W. Standard Time",
	"Australia/Hobart":       "Tasmania Standard Time",
	"Australia/Lord_Howe":    "Lord Howe Standard Time",
	"Australia/Perth":        "W. Australia Standard Time",
	"Australia/Sydney":       "AUS Eastern Standard Time",
	"Etc/GMT":                "UTC",
	"Etc/GMT+11":             "UTC-11",
	"Etc/GMT+12":             "Dateline Standard Time",
	"Etc/GMT+2":              "UTC-02",
	"Etc/GMT+8":              "UTC-08",
	"Etc/GMT+9":              "UTC-09",
	"Etc/GMT-12":             "UTC+12",
	"Etc/GMT-13":             "UTC+13",
	"Europe/Astrakhan":       "Astrakhan Standard Time",
	"Europe/Berlin":          "W. Europe Standard Time",
	"Europe/Bucharest":       "GTB Standard Time",
	"Europe/Budapest":        "Central Europe Standard Time",
	"Europe/Chisinau":        "E. Europe Standard Time",
	"Europe/Istanbul":        "Turkey Standard Time",
	"Europe/Kaliningrad":     "Kaliningrad Standard Time",
	"Europe/Kiev":            "FLE Standard Time",
	"Europe/London":          "GMT Standard Time",
	"Europe/Minsk":           "Belarus Standard Time",
	"Europe/Moscow":          "Russian Standard Time",
	"Europe/Paris":           "Romance Standard Time",
	"Europe/Samara":          "Russia Time Zone 3",
	"Europe/Saratov":         "Saratov Standard Time",
	"Europe/Volgograd":       "Volgograd Standard Time",
	"Europe/Warsaw":          "Central European Standard Time",
	"Indian/Mauritius":       "Mauritius Standard Time",
	"Pacific/Apia":           "Samoa Standard Time",
	"Pacific/Auckland":       "New Zealand Standard Time",
	"Pacific/Bougainville":   "Bougainville Standard Time",
	"Pacific/Chatham":        "Chatham Islands Standard Time",
	"Pacific/Easter":         "Easter Island Standard Time",
	"Pacific/Fiji":           "Fiji Standard Time",
	"Pacific/Guadalcanal":    "Central Pacific Standard Time",
	"Pacific/Honolulu":       "Hawaiian Standard Time",
	"Pacific/Kiritimati":     "Line Islands Standard Time",
	"Pacific/Marquesas":      "Marquesas Standard Time",
	"Pacific/Norfolk":        "Norfolk Standard Time",
	"Pacific/Port_Moresby":   "West Pacific Standard Time",
	"Pacific/Tongatapu":      "Tonga Standard Time",
}
