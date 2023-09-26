package config

import "time"

type Entrepreneurship struct {
	ID                 string      `gorm:"type:uuid;primaryKey;column:id_entrepreneurship" json:"idEntrepreneurship"`
	Name               string      `gorm:"type:varchar(255);not null;column:name" json:"name"`
	EntityOrgCommunity string      `gorm:"type:varchar(255);column:entity_org_community" json:"entityOrgCommunity"`
	Address            string      `gorm:"type:varchar(255);column:address" json:"address"`
	Phone              string      `gorm:"type:varchar(255);column:phone" json:"phone"`
	Description        string      `gorm:"type:varchar(255);column:description" json:"description"`
	Sector             string      `gorm:"type:varchar(255);column:sector" json:"sector"`
	StageCie           string      `gorm:"type:varchar(255);column:stage_cie" json:"stageCie"`
	State              bool        `gorm:"type:boolean;column:state" json:"state"`
	ImgURL             string      `gorm:"type:varchar(255);column:img_url" json:"imgUrl"`
	CreatedAt          time.Time   `gorm:"autoCreateTime"`
	UpdatedAt          time.Time   `gorm:"autoUpdateTime"`
	IDEntrepreneur     string      `gorm:"column:id_entrepreneur" json:"idEntrepreneur"`
	Resources          []Resources `gorm:"foreignKey:IDEntrepreneurship;references:ID"`
}

type Entrepreneur struct {
	ID                string             `gorm:"type:uuid;primaryKey" json:"idEntrepreneur"`
	Entrepreneurships []Entrepreneurship `gorm:"foreignKey:IDEntrepreneur;references:ID"`
}

type Resources struct {
	ID                 string    `gorm:"type:uuid;primaryKey;column:id_resource" json:"idResource"`
	Name               string    `gorm:"type:varchar(255);not null;column:name" json:"name"`
	TotalCost          float64   `gorm:"type:numeric(10,2);column:total_cost" json:"totalCost"`
	Description        string    `gorm:"type:varchar(255);column:description" json:"description"`
	Category           string    `gorm:"type:varchar(255);column:category" json:"category"`
	Availability       bool      `gorm:"type:boolean;column:availability" json:"availability"`
	Quantity           int       `gorm:"type:int;column:quantity" json:"quantity"`
	State              bool      `gorm:"type:boolean;column:state" json:"state"`
	ImgURL             string    `gorm:"type:varchar(255);column:img_url" json:"imgUrl"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
	IDEntrepreneurship string    `gorm:"column:id_entrepreneurship" json:"idEntrepreneurship"`
	//Tags               []Tags    `gorm:"many2many:resources_tags;foreignKey:ID;references:ID;joinForeignKey:IDResource;joinReferences:IDTag;"`
	Costs         []Costs         `gorm:"foreignKey:IDResource;references:ID"`
	ResourcesTags []ResourcesTags `gorm:"foreignKey:IDResource"`
}

type Tags struct {
	ID            string          `gorm:"type:uuid;primaryKey;column:id_tag" json:"idTag"`
	Name          string          `gorm:"type:varchar(255);not null;column:name" json:"name"`
	ResourcesTags []ResourcesTags `gorm:"foreignKey:IDTag"`
}

//tabla intermedia
type ResourcesTags struct {
	IDResource string `gorm:"type:uuid;primaryKey;column:id_resource" json:"idResource"`
	IDTag      string `gorm:"type:uuid;primaryKey;column:id_tag" json:"idTag"`
}

type Costs struct {
	ID              string           `gorm:"type:uuid;primaryKey;column:id_costs" json:"idCost"`
	Name            string           `gorm:"type:varchar(255);not null;column:name" json:"name"`
	Quantity        int              `gorm:"type:int;column:quantity" json:"quantity"`
	UnitMeasure     string           `gorm:"type:varchar(255);column:unit_measure" json:"unitMeasure"`
	UnitCost        float64          `gorm:"type:numeric(10,2);column:unit_cost" json:"unitCost"`
	TotalCost       float64          `gorm:"type:numeric(10,2);column:total_cost" json:"totalCost"`
	TypeCost        string           `gorm:"type:varchar(255);column:type_cost" json:"typeCost"`
	Detail          string           `gorm:"type:varchar(255);column:detail" json:"detail"`
	Provider        string           `gorm:"type:varchar(255);column:provider" json:"provider"`
	State           bool             `gorm:"type:boolean;column:state" json:"state"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"`
	UpdatedAt       time.Time        `gorm:"autoUpdateTime"`
	IDResource      string           `gorm:"column:id_resource" json:"idResource"`
	HistoricalCosts []HistoricalCost `gorm:"foreignKey:IDCost;references:ID"`
}

type HistoricalCost struct {
	ID        string    `gorm:"type:uuid;primaryKey;column:id_historical_cost" json:"idHistoricalCost"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Amount    float64   `gorm:"type:numeric(10,2);not null;column:amount" json:"amount"`
	Detail    string    `gorm:"type:varchar(255);not null;column:detail" json:"detail"`
	IDCost    string    `gorm:"column:id_cost" json:"idCost"`
}
