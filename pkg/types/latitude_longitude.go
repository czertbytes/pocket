package types

import "appengine"

type LatitudeLongitude struct {
	Latitude  float64            `json:"latitude,omitempty" datastore:"-"`
	Longitude float64            `json:"longitude,omitempty" datastore:"-"`
	GeoPoint  appengine.GeoPoint `json:"-" datastore:"geopoint"`
}

func (self *LatitudeLongitude) SetFormattedValues() {
	self.SetLatitudeLongitude()
}

func (self *LatitudeLongitude) SetLatitude() {
	self.Latitude = self.GeoPoint.Lat
}

func (self *LatitudeLongitude) SetLongitude() {
	self.Longitude = self.GeoPoint.Lng
}

func (self *LatitudeLongitude) SetLatitudeLongitude() {
	self.SetLatitude()
	self.SetLongitude()
}
