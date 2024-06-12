package Artist_struct

type Artist struct {
	Id           int         `json:"id"`
	Image        interface{} `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate int         `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
	Relations    interface{} `json:"relations"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationsIndex struct {
	Index []Relation `json:"index"`
}
