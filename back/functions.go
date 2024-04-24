package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Index struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Loc struct {
	Ind []Index `json:"index"`
}

type Relation struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}

type Everything struct {
	Everyone []Artist
	Location Loc
}

type ArtistInfo struct {
	Artist
	Relation
}

func GetAllArtists() ([]Artist, error) {
	var artists []Artist
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return artists, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return artists, err
	}
	err = json.Unmarshal(bytes, &artists)
	if err != nil {
		return artists, err
	}
	defer response.Body.Close()
	return artists, nil
}

func GetAllLocations() (Loc, error) {
	var location Loc
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return location, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return location, err
	}
	defer response.Body.Close()
	return location, nil
}

func OneArtist(id int) (Artist, error) {
	var artist Artist
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id))
	if err != nil {
		return artist, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return artist, err
	}

	err = json.Unmarshal(bytes, &artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

func GetLocation(id int) (Loc, error) {
	var location Loc // see the structure
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id))
	if err != nil {
		return location, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return location, err
	}
	return location, nil
}

func Relations(id int) (Relation, error) {
	var rel Relation
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id))
	if err != nil {
		return rel, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return rel, err
	}
	err = json.Unmarshal(bytes, &rel)
	if err != nil {
		return rel, err
	}
	return rel, nil
}
