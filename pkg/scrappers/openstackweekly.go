package scrappers

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/zufardhiyaulhaq/openstackweekly/models"
)

type OpenstackWeekly struct{}

func (s *OpenstackWeekly) GetOpenStackWeekly(currentContent models.OpenStackContents) models.OpenStackContents {
	content := getOpenStackNews(currentContent)
	content = getOpenStackSuperUserNews(content)

	return content
}

func getOpenStackNews(currentContent models.OpenStackContents) models.OpenStackContents {
	res, err := http.Get("https://www.openstack.org/news/")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".recentBox").Each(func(i int, data *goquery.Selection) {
		urlQuery, _ := data.Find("a").Attr("href")
		url := "https://www.openstack.org/" + urlQuery

		title := data.Find("a").Text()
		doAppend := true

		for _, v := range currentContent.Content {
			if v.Url == url {
				doAppend = false
			}
		}

		if doAppend {
			content := models.Content{
				Title: title, Url: url, Kind: "OpenStack News", IsDelivered: false,
			}
			currentContent.Content = append(currentContent.Content, content)
		}
	})

	return currentContent
}

func getOpenStackSuperUserNews(currentContent models.OpenStackContents) models.OpenStackContents {
	res, err := http.Get("https://superuser.openstack.org/")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".article-card").Each(func(i int, data *goquery.Selection) {
		url, _ := data.Find("a").Attr("href")
		title := data.Find("h3").Text()
		doAppend := true

		for _, v := range currentContent.Content {
			if v.Url == url {
				doAppend = false
			}
		}

		if doAppend {
			content := models.Content{
				Title: title, Url: url, Kind: "OpenStack Super User News", IsDelivered: false,
			}
			currentContent.Content = append(currentContent.Content, content)
		}
	})

	return currentContent
}
