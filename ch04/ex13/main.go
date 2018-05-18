// Open Movie DatabaseのJSONに基づくウェブサービスは、https://omdapi.com/から映画を名前で検索し、そのポスター画像をダウンロードさせてくれます。
// コマンドラインで指定された映画のポスター画像をダウンロードするツールposterを書きなさい

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"log"
)

const URL = "http://www.omdbapi.com/?t="

type Movie struct {
	Title    string
	Country  string
	Poster   string
	Response string
	Error    string // Request limit reached!
}

func searchMovie(keywords string) (Movie, error) {
	resp, err := http.Get(URL + keywords + "&apikey=" + os.Getenv("OMDB_API_KEY"))
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		log.Fatalln(err)
	}
	if movie.Response != "True" {
		log.Fatalln(movie.Error)
	}
	if movie.Poster == "" {
		log.Fatalln("ポスター画像がありません")
	}

	return movie, err
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("映画の名前を入力してください")
	}

	keywords := url.QueryEscape(strings.Join(os.Args[1:], " "))
	movie, err := searchMovie(keywords)

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := check.Get(movie.Poster)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}


	file, err := os.Create(movie.Title + ".jpg")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatalln(err)
	}
}

/*
{
	Title: "Blade",
	Year: "1998",
	Rated: "R",
	Released: "21 Aug 1998",
	Runtime: "120 min",
	Genre: "Action, Horror",
	Director: "Stephen Norrington",
	Writer: "David S. Goyer",
	Actors: "Wesley Snipes, Stephen Dorff, Kris Kristofferson, N'Bushe Wright",
	Plot: "A half-vampire, half-mortal man becomes a protector of the mortal race, while slaying evil vampires.",
	Language: "English, Russian",
	Country: "USA",
	Awards: "4 wins & 8 nominations.",
	Poster: "https://images-na.ssl-images-amazon.com/images/M/MV5BMTQ4MzkzNjcxNV5BMl5BanBnXkFtZTcwNzk4NTU0Mg@@._V1_SX300.jpg",
	Ratings: [],
	Metascore: "45",
	imdbRating: "7.1",
	imdbVotes: "212,014",
	imdbID: "tt0120611",
	Type: "movie",
	DVD: "22 Dec 1998",
	BoxOffice: "N/A",
	Production: "New Line Cinema",
	Website: "N/A",
	Response: "True"
}
*/