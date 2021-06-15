package imdb

import (
	"strings"
	"time"

	"github.com/grumpypixel/gofu/stringslice"
)

type IMDB struct {
	AllPosterResolutions bool
	WaitBetweenRequests  time.Duration
	Verbose              bool
}

func NewIMDB(allPosterResolutions bool) *IMDB {
	return &IMDB{
		AllPosterResolutions: allPosterResolutions,
	}
}

func (db *IMDB) FetchPoster(movie string) []string {
	posters, _ := db.fetchPoster(movie)
	var urls []string
	for _, poster := range posters {
		urls = append(urls, poster.ImageURL)
	}
	return urls
}

// fetch the movie title
func (db *IMDB) FetchTitle(movie string) (string, string) {
	url, ok := db.validateMovieSource(movie)
	if !ok {
		return "", ""
	}
	titleID, _ := db.TitleIDFromURL(url)
	title, _ := db.findMovieTitle(url)
	return titleID, title
}

// get the title ID from an IMDB URL
func (db *IMDB) TitleIDFromURL(imdbURL string) (string, bool) {
	if !strings.Contains(imdbURL, "title/tt") {
		return "", false
	}

	s := strings.Split(imdbURL, "/")
	s = stringslice.TrimElements(s)
	s = stringslice.RemoveEmptyElements(s)

	indexTitle := stringslice.IndexOfElement("title", s)
	if indexTitle == -1 || indexTitle == len(s)-1 {
		return "", false
	}
	title := s[indexTitle+1]
	return title, true
}

func (db *IMDB) URLFromSource(source string) (string, bool) {
	return db.validateMovieSource(source)
}
