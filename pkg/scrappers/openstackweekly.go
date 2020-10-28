package scrappers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/zufardhiyaulhaq/openstackweekly/models"
)

type OpenstackWeekly struct{}

func (s *OpenstackWeekly) GetOpenStackWeekly(currentContent models.OpenStackContents) models.OpenStackContents {
	content := getOpenStackNews(currentContent)

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

	fmt.Println(doc)

	return currentContent

	// doc.Find(".upcoming-webinars").Each(func(i int, s0 *goquery.Selection) {
	// 	s0.Find("article").Not("a.button-like").Each(func(i int, s1 *goquery.Selection) {
	// 		url, _ := s1.Find("a").Attr("href")
	// 		title := strings.Replace(s1.Find("a").Text(),"Find Out More","",-1)
	// 		fmt.Println(title)
	// 		doAppend := true

	// 		for _, v := range currentContent.Content {
	// 			if (v.Url == url){
	// 				doAppend = false
	// 			}
	// 		}

	// 		if doAppend {
	// 			month := s1.Find(".upcoming-date.upcoming-date-mobile").ChildrenFiltered(".month").Text()
	// 			day := s1.Find(".upcoming-date.upcoming-date-mobile").ChildrenFiltered(".day").Text()
	// 			year := s1.Find(".upcoming-date.upcoming-date-mobile").ChildrenFiltered(".year").Text()

	// 			date := month+" "+day+" "+year
	// 			time := s1.Find(".details").ChildrenFiltered(".time").Text()

	// 			singleContent := models.WebinarCNCFContent{Title: title, Url: url, Date: date, Time: time, IsDelivered: false}
	// 			newContent.Content = append(newContent.Content,singleContent)
	// 		}
	// 	})
	// })
}

// func (s *OpenstackWeekly) GetWeekly() []communityv1alpha1.ArticleSpec {
// 	var articles []communityv1alpha1.ArticleSpec

// 	// scrapper logic here
// 	// your scrapper must populate articles

// 	response, err := http.Get("https://OpenstackWeekly.com/issues/latest?layout=bare")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer response.Body.Close()

// 	document, err := goquery.NewDocumentFromReader(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	document.Find("div#content").Each(func(i int, div *goquery.Selection) {
// 		div.Find("span.mainlink").Each(func(i int, span *goquery.Selection) {
// 			var article communityv1alpha1.ArticleSpec
// 			article.Title = span.Find("a").Text()
// 			article.Type = "News"
// 			article.Url, _ = span.Find("a").Attr("href")

// 			articles = append(articles, article)
// 		})
// 	})

// 	return articles
// }

// func (s *OpenstackWeekly) GetWeeklyName() string {
// 	var weeklyName string

// 	// scrapper logic here
// 	// your scrapper must populate weeklyName

// 	response, err := http.Get("https://OpenstackWeekly.com/issues/latest?layout=bare")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer response.Body.Close()

// 	document, err := goquery.NewDocumentFromReader(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	table := document.Find("div#content").ChildrenFiltered("table:first-of-type").First()
// 	td := table.Find("td:first-of-type").First()
// 	text := td.Find("p").Text()
// 	regex, _ := regexp.Compile(`\d+`)

// 	weeklyName = "OpenStack Weekly " + regex.FindString(text)
// 	return weeklyName
// }
