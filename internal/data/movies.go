package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"` // movie runtime in minutes
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // the version starts at 1 and will be incremented each time the movie is updated
}
