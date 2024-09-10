package main

type Albumb struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albumbs = []Albumb{
	{1, "The Dark Side of the Moon", "Pink Floyd", 1.99},
	{2, "Thriller", "Michael Jackson", 1.99},
	{3, "Back in Black", "AC/DC", 1.99},
}
