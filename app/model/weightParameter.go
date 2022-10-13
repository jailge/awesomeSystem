package model

type WeightParameter struct {
	Craft          []string `bson:"craft"`
	Texture        []string `bson:"texture"`
	Process        []string `bson:"process"`
	PurchaseStatus []string `bson:"purchase_status"`
}

type Craft struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type Texture struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type Process struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type PurchaseStatus struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type BelongTo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type WeightStatus struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type Material struct {
	Id            string `bson:"_id"`
	InnerCode     string `bson:"inner_code"`
	Name          string `bson:"name"`
	Code          string `bson:"code"`
	Specification string `bson:"specification"`
	Status        bool   `bson:"status"`
}

type Supplier struct {
	Id        string `bson:"_id"`
	InnerCode string `bson:"inner_code"`
	Name      string `bson:"name"`
	Code      string `bson:"code"`
	TypeCode  string `bson:"type_code"`
	Status    bool   `bson:"status"`
}
