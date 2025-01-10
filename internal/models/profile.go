package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Profile struct {
	gorm.Model
	UserID   uint      `gorm:"not null" json:"user_id"`
	Bio      string    `json:"bio"`
	Age      int       `json:"age"`
	Gender   string    `json:"gender"`
	Location Point     `gorm:"type:geometry(Point,4326)" json:"location"`
	Pictures []Picture `gorm:"foreignKey:ProfileID" json:"pictures"`
}

type Preference struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	GenderPref   string `json:"gender_pref"`
	DistancePref int    `json:"distance_pref"`
}

type Picture struct {
	gorm.Model
	ProfileID uint   `gorm:"not null" json:"profile_id"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
}

type Point struct {
	Lat float64
	Lng float64
}

func (p *Point) Scan(val interface{}) error {
	b, errs := val.([]byte)
	if errs {
		return errors.New("failed to scan PostGIS point")
	}
	pointStr := string(b)
	if !strings.HasPrefix(pointStr, "POINT(") {
		return errors.New("invalid point format")
	}
	pointStr = strings.TrimPrefix(pointStr, "POINT(")
	pointStr = strings.TrimSuffix(pointStr, ")")
	coords := strings.Split(pointStr, " ")
	if len(coords) != 2 {
		return errors.New("invalid point coordinates")
	}
	lng, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return err
	}
	lat, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return err
	}
	p.Lng = lng
	p.Lat = lat
	return nil
}

func (p *Point) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", p.Lng, p.Lat), nil
}
