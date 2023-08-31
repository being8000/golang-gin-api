package model

type RegionCountry struct {
	Name      string           `gorm:"type:varchar(50)"`
	Initials  string           `gorm:"type:char(3)"`
	Code      string           `gorm:"uniqueIndex;type:char(10);not null;"`
	Provinces []RegionProvince `gorm:"foreignKey:CountryCode;references:Code"`
	Model
}

type RegionProvince struct {
	CountryCode string        `gorm:"type:varchar(10)"`
	Country     RegionCountry `gorm:"foreignKey:CountryCode;references:Code"`
	Name        string        `gorm:"type:varchar(75)"`
	Initials    string        `gorm:"type:char(3)"`
	Code        string        `gorm:"uniqueIndex;type:varchar(10);not null;"`
	Cities      []RegionCity  `gorm:"foreignKey:ProvinceCode;references:Code"`
	Model
}
type RegionCity struct {
	ProvinceCode string           `gorm:"type:varchar(10)"`
	Province     RegionProvince   `gorm:"foreignKey:ProvinceCode;references:Code"`
	Name         string           `gorm:"type:varchar(75)"`
	Initials     string           `gorm:"type:char(3)"`
	Code         string           `gorm:"uniqueIndex;type:varchar(10);not null;"`
	Districts    []RegionDistrict `gorm:"foreignKey:CityCode;references:Code"`
	Model
}

type RegionDistrict struct {
	CityCode string         `gorm:"type:varchar(10)"`
	City     RegionCity     `gorm:"foreignKey:CityCode;references:Code"`
	Name     string         `gorm:"type:varchar(75)"`
	Initials string         `gorm:"type:char(3)"`
	Code     string         `gorm:"uniqueIndex;type:varchar(10);not null;"`
	Streets  []RegionStreet `gorm:"foreignKey:DistrictCode;references:Code"`
	Model
}

type RegionStreet struct {
	DistrictCode string         `gorm:"type:varchar(10)"`
	District     RegionDistrict `gorm:"foreignKey:DistrictCode;references:Code"`
	Name         string         `gorm:"type:varchar(75)"`
	ZipCode      string         `gorm:"type:varchar(75)"`
	Type         string         `gorm:"type:varchar(20)"`
	Code         string         `gorm:"uniqueIndex;type:varchar(10);not null;"`
	Model
}
