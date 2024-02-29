package main

import (
	"net"

	"github.com/oschwald/maxminddb-golang"
)

type GeoIpDb struct {
  db *maxminddb.Reader
}

type GeoIpRecord struct {
  Location struct {
			TimeZone uintptr `maxminddb:"time_zone"`
		} `maxminddb:"location"`
}

func NewDb() (*GeoIpDb, error) {
	reader, err := maxminddb.Open("./geoip/GeoLite2-City.mmdb")
	if err != nil {
    return nil, err
	}

  db := GeoIpDb{}

  db.db = reader

  return &db, nil
}

func (db *GeoIpDb) Close() {
  db.db.Close()
}


func (db *GeoIpDb) TimezoneFromIp(ip net.IP) (string, error) {
	var record GeoIpRecord

  err := db.db.Lookup(ip, &record)
	if err != nil {
    return "", err
	}

  var timeZone string
  err = db.db.Decode(record.Location.TimeZone, &timeZone)

  if err != nil {
    return "", err
  }

  return timeZone, nil
}
