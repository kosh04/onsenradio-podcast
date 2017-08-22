package main

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/eduncan911/podcast"
)

const baseURL = "http://www.onsen.ag"

var now = time.Now()

func podcastWriter(w io.Writer) {
	doc, err := goquery.NewDocument("http://www.onsen.ag/easy/easy.php")
	if err != nil {
		log.Fatal(err)
	}
	p := podcast.New(
		"インターネットラジオステーション＜音泉＞",
		baseURL,
		"Podcast onsen.ag internet radio (unofficial)",
		&now, &now,
	)
	p.Language = "ja"

	doc.Find(".programConts").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".programTitle").Text()
		pageLink, _ := s.Find(".programPageLink > a").Attr("href")
		logo, _ := s.Find(".programLogo > img").Attr("src")
		mail, found := s.Find(".programMail > a[href^='mailto:']").Attr("href")
		if found {
			mail = strings.TrimPrefix(mail, "mailto:")
		}
		date, _ := parseUpdate(s.Find(".update").Text())
		item := podcast.Item{
			Title:       fmt.Sprintf("%s %s", s.Find(".update").Text(), title),
			Link:        baseURL + pageLink,
			Description: "TODO",
			Author: &podcast.Author{
				Name:  s.Find(".programPersonality").Text(),
				Email: mail,
			},
			PubDate: &date,
		}
		item.AddImage(baseURL + logo)

		audioLink, found := s.Find("form[action$='.mp3']").Attr("action")
		if found {
			// TODO: unknown audio size
			item.AddEnclosure(audioLink, podcast.MP3, 0)
		}

		if _, err := p.AddItem(item); err != nil {
			log.Printf("podcast.AddItem(%q) error: %v\n", item.Title, err)
			return
		}
	})

	p.Encode(w)
}

// parseUpdate parse time from html ".update" class field text
// e.g. s: "8/21 UP" => "2017/08/21" => "2017-08-21 00:00:00 +0900 JST"
func parseUpdate(s string) (time.Time, error) {
	matched, err := regexp.MatchString(`\d{1,2}/\d{1,2} UP`, s)
	if !matched {
		return time.Time{}, err
	}
	ss := fmt.Sprintf("%d/%s", now.Year(), strings.TrimSuffix(s, " UP")) // "&nbsp;UP"
	return time.ParseInLocation("2006/1/2", ss, time.Local)
}

func description(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}
	return doc.Find("#introductionWrap").Text(), nil
}
