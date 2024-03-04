package models

import "time"

// ApiErr is the common structure for errors coming from MangaDex API.
type ApiErr struct {
	Id      string      `json:"id"`
	Status  int         `json:"status"`
	Title   string      `json:"title"`
	Detail  string      `json:"detail"`
	Context interface{} `json:"context"`
}

// ApiTag is the common structure for Tags used by MangaDex API.
type ApiTag struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name struct {
			En string `json:"en"`
		} `json:"name"`
		Description struct {
		} `json:"description"`
		Group   string `json:"group"`
		Version int    `json:"version"`
	} `json:"attributes"`
	Relationships []interface{} `json:"relationships"`
}

type MangaRequestParam struct {
	Order          string
	IncludedTags   string
	ExcludedTags   string
	Title          string
	Author         string
	AuthorOrArtist string
	Status         string
	Public         string
	ContentRating  string
	Limit          string
	Offset         string
}

type MangaRequest struct {
	OrderType      string
	OrderValue     string
	IncludedTags   []string
	ExcludedTags   []string
	Title          string
	Author         string
	AuthorOrArtist string
	Status         []string
	Public         []string
	Limit          int
	Offset         int
}

// ApiManga is the data structure for the Manga request to MangaDex API.
//
//	Request: GET https://api.mangadex.org/manga
//	Possible GET params:
//	- order[{title,year,createdAt,updatedAt,latestUploadedChapter,followedCount,relevance,rating}]={asc,desc}
//	- includedTags[]={tag-id}&={tag-id}&={tag-id}...
//	- excludedTags[]={tag-id}&={tag-id}&={tag-id}...
//	- contentRating[]=safe&=...
//	- title={title}
//	- author={author}
//	- authorOrArtist={authorOrArtist}
//	- status=ongoing&=completed&=hiatus&=cancelled
//	- publicationDemographic=shounen&=seinen&=shoujo&=josei&=none
//	- availableTranslatedLanguage[]=en&=...
//	- limit=10(default)
//	- offset={>=0}
type ApiManga struct {
	Result   string   `json:"result"`
	Errors   []ApiErr `json:"errors,omitempty"`
	Response string   `json:"response,omitempty"`
	Data     []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title struct {
				En string `json:"en"`
			} `json:"title"`
			AltTitles []struct {
				Zh   string `json:"zh,omitempty"`
				Ko   string `json:"ko,omitempty"`
				En   string `json:"en,omitempty"`
				Ja   string `json:"ja,omitempty"`
				JaRo string `json:"ja-ro,omitempty"`
				KoRo string `json:"ko-ro,omitempty"`
			} `json:"altTitles"`
			Description struct {
				En string `json:"en"`
			} `json:"description"`
			IsLocked bool `json:"isLocked"`
			Links    struct {
				Al    string `json:"al,omitempty"`
				Ap    string `json:"ap,omitempty"`
				Mu    string `json:"mu"`
				Raw   string `json:"raw,omitempty"`
				Bw    string `json:"bw,omitempty"`
				Kt    string `json:"kt,omitempty"`
				Amz   string `json:"amz,omitempty"`
				Ebj   string `json:"ebj,omitempty"`
				Mal   string `json:"mal,omitempty"`
				Nu    string `json:"nu,omitempty"`
				Engtl string `json:"engtl,omitempty"`
				Cdj   string `json:"cdj,omitempty"`
			} `json:"links"`
			OriginalLanguage               string    `json:"originalLanguage"`
			LastVolume                     string    `json:"lastVolume"`
			LastChapter                    string    `json:"lastChapter"`
			PublicationDemographic         *string   `json:"publicationDemographic"`
			Status                         string    `json:"status"`
			Year                           *int      `json:"year"`
			ContentRating                  string    `json:"contentRating"`
			Tags                           []ApiTag  `json:"tags"`
			State                          string    `json:"state"`
			ChapterNumbersResetOnNewVolume bool      `json:"chapterNumbersResetOnNewVolume"`
			CreatedAt                      time.Time `json:"createdAt"`
			UpdatedAt                      time.Time `json:"updatedAt"`
			Version                        int       `json:"version"`
			AvailableTranslatedLanguages   []string  `json:"availableTranslatedLanguages"`
			LatestUploadedChapter          string    `json:"latestUploadedChapter"`
		} `json:"attributes"`
		Relationships []struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			Related string `json:"related,omitempty"`
		} `json:"relationships"`
	} `json:"data,omitempty"`
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	Total  int `json:"total,omitempty"`
}

// ApiCover is the data structure for the Cover request to MangaDex API.
//
//	Request: GET https://api.mangadex.org/cover/{manga-id}
type ApiCover struct {
	Result   string   `json:"result"`
	Errors   []ApiErr `json:"errors,omitempty"`
	Response string   `json:"response"`
	Data     []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Volume      string `json:"volume"`
			FileName    string `json:"fileName"`
			Description string `json:"description"`
			Locale      string `json:"locale"`
			Version     int    `json:"version"`
			CreatedAt   string `json:"createdAt"`
			UpdatedAt   string `json:"updatedAt"`
		} `json:"attributes"`
		Relationships []struct {
			Id         string `json:"id"`
			Type       string `json:"type"`
			Related    string `json:"related"`
			Attributes struct {
			} `json:"attributes"`
		} `json:"relationships"`
	} `json:"data"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// ApiTags is the data structure for the Tag request to MangaDex API.
//
//	Request: GET https://api.mangadex.org/manga/tag
type ApiTags struct {
	Result   string   `json:"result"`
	Errors   []ApiErr `json:"errors,omitempty"`
	Response string   `json:"response"`
	Data     []ApiTag `json:"data"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
	Total    int      `json:"total"`
}

// ApiMangaFeed is the data structure for the Manga feed request to MangaDex API.
//
//	Request: GET https://api.mangadex.org/manga/{manga-id}/feed
type ApiMangaFeed struct {
	Result   string   `json:"result"`
	Errors   []ApiErr `json:"errors,omitempty"`
	Response string   `json:"response"`
	Data     []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title              string `json:"title"`
			Volume             string `json:"volume"`
			Chapter            string `json:"chapter"`
			Pages              int    `json:"pages"`
			TranslatedLanguage string `json:"translatedLanguage"`
			Uploader           string `json:"uploader"`
			ExternalUrl        string `json:"externalUrl"`
			Version            int    `json:"version"`
			CreatedAt          string `json:"createdAt"`
			UpdatedAt          string `json:"updatedAt"`
			PublishAt          string `json:"publishAt"`
			ReadableAt         string `json:"readableAt"`
		} `json:"attributes"`
		Relationships []struct {
			Id         string `json:"id"`
			Type       string `json:"type"`
			Related    string `json:"related"`
			Attributes struct {
			} `json:"attributes"`
		} `json:"relationships"`
	} `json:"data"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// ApiChapterScan is the data structure for the Chapter scan request to MangaDex API.
//
//	Request: GET https://api.mangadex.org/at-home/server/{chapter-id}
type ApiChapterScan struct {
	Result  string   `json:"result"`
	Errors  []ApiErr `json:"errors,omitempty"`
	BaseUrl string   `json:"baseUrl"`
	Chapter struct {
		Hash      string   `json:"hash"`
		Data      []string `json:"data"`
		DataSaver []string `json:"dataSaver"`
	} `json:"chapter"`
}

// ApiMangaStats is the data structure for the Chapter stats request to MangaDex API.
//
//	Request: GET https://api.mangadex.org/statistics/manga/{manga-id}
type ApiMangaStats struct {
	Result     string `json:"result"`
	Statistics struct {
		Id struct {
			Comments struct {
				ThreadId     int `json:"threadId"`
				RepliesCount int `json:"repliesCount"`
			} `json:"comments"`
			Rating struct {
				Average      float64 `json:"average"`
				Bayesian     float64 `json:"bayesian"`
				Distribution struct {
					Rating1  int `json:"1"`
					Rating2  int `json:"2"`
					Rating3  int `json:"3"`
					Rating4  int `json:"4"`
					Rating5  int `json:"5"`
					Rating6  int `json:"6"`
					Rating7  int `json:"7"`
					Rating8  int `json:"8"`
					Rating9  int `json:"9"`
					Rating10 int `json:"10"`
				} `json:"distribution"`
			} `json:"rating"`
			Follows int `json:"follows"`
		} `json:"-"`
	} `json:"statistics"`
}