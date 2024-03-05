package main

type Pokemon struct {
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Height                 int     `json:"height"`
	ID                     int     `json:"id"`
	IsDefault              bool    `json:"is_default"`
	LocationAreaEncounters string  `json:"location_area_encounters"`
	Name                   string  `json:"name"`
	Order                  int     `json:"order"`
	Weight                 int     `json:"weight"`
	Stats                  []Stats `json:"stats"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
