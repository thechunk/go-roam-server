package database

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type Restaurant struct {
	Name         string `redis:"rName"`
	SetId        string `redis:"rSetId"`
	Address1     string `redis:"rAddress1"`
	PriceRangeId string `redis:"rPriceRangeId"`
	Price        string `redis:"rPrice"`
	OpenRiceUrl  string `redis:"rOpenRiceUrl"`
}

type Coords struct {
	Lat float64
	Lng float64
}

func parsePositions(results []interface{}) (*map[string]Coords, error) {
	m := make(map[string]Coords)

	for i := 0; i < len(results); i++ {
		// restaurantId := []string{}
		result := results[i].([]interface{})
		key := string(result[0].([]uint8))
		coordPair := result[1].([]interface{})

		var err error
		c := Coords{}
		c.Lng, err = strconv.ParseFloat(string(coordPair[0].([]uint8)), 64)
		c.Lat, err = strconv.ParseFloat(string(coordPair[1].([]uint8)), 64)

		if err != nil {
			return nil, err
		}

		m[key] = c
	}

	return &m, nil
}

func RestaurantsNearby(lat float64, lng float64, radius float64) (*[]Restaurant, error) {
	conn := Conn()
	defer conn.Close()

	results, err := redis.Values(conn.Do("GEORADIUS", "restaurantLocations",
		lng,
		lat,
		radius,
		"km", "WITHCOORD"))
	if err != nil {
		return nil, err
	}

	positions, err := parsePositions(results)
	conn.Send("MULTI")
	for restaurantId, _ := range *positions {
		conn.Send("HGETALL", restaurantId)
	}

	values, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return nil, err
	}

	restaurants := []Restaurant{}
	for i := 0; i < len(values); i++ {
		restaurant := Restaurant{}
		raw := values[i].([]interface{})
		if err := redis.ScanStruct(raw, &restaurant); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}

	return &restaurants, nil
}

func RestaurantById(id int64) (*Restaurant, error) {
	conn := Conn()
	defer conn.Close()

	results, err := redis.Values(conn.Do("HGETALL",
		"restaurant:"+strconv.FormatInt(id, 10)))
	if err != nil {
		return nil, err
	}

	restaurant := Restaurant{}
	if err := redis.ScanStruct(results, &restaurant); err != nil {
		return nil, err
	}

	return &restaurant, nil
}
