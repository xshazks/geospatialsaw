package geospatialsaw

type GeoBorder struct {
	Type        string        `bson:"type"`
	Coordinates [][][]float64 `bson:"coordinates"`
}

type LocationData struct {
	ID          string    `bson:"_id"`
	Province    string    `bson:"province"`
	District    string    `bson:"district"`
	SubDistrict string    `bson:"sub_district"`
	Village     string    `bson:"village"`
	Border      GeoBorder `bson:"border"`
}
