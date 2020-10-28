package main

import (
	"log"
	"strings"
	"time"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	"github.com/zufardhiyaulhaq/openstackweekly/handlers"
	"github.com/zufardhiyaulhaq/openstackweekly/models"
	"github.com/zufardhiyaulhaq/openstackweekly/pkg/scrappers"
	"gopkg.in/yaml.v2"
)

func main() {
	// start GitHub handler
	handler := handlers.Github{}
	handler.Start()

	// Init scrapper
	scrapper := scrappers.OpenstackWeekly{}

	// initialize content
	var content models.OpenStackContents

	// check current content file exist
	if !handler.FileExist("content.yaml") {
		newContent, err := yaml.Marshal(content)
		if err != nil {
			log.Fatal(err)
		}

		CreateFile(handler, "content.yaml", "init OpenStack Weekly content file", newContent)
	}

	// get current content
	contentTmpl := handler.GetFile("content.yaml")
	err := yaml.Unmarshal(contentTmpl, &content)
	if err != nil {
		log.Fatal(err)
	}
	currentContentLength := len(content.Content)

	// feed current content to scrapper
	openstackWeeklycontent := scrapper.GetOpenStackWeekly(content)
	openstackWeeklycontentByte, err := yaml.Marshal(openstackWeeklycontent)
	if err != nil {
		log.Fatal(err)
	}
	newContentLength := len(openstackWeeklycontent.Content)

	if newContentLength != currentContentLength {
		handler.UpdateFile("content.yaml", "update OpenStack Weekly content", openstackWeeklycontentByte)
	} else {
		log.Printf("[openstackweekly] no update about OpenStack")
		return
	}

	// communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	var weeklyName string
	var weeklyData []communityv1alpha1.ArticleSpec

	// populate weeklyData from openstackWeeklycontent
	for index, value := range openstackWeeklycontent.Content {
		if !value.IsDelivered {
			var data communityv1alpha1.ArticleSpec
			data.Title = value.Title
			data.Url = value.Url
			data.Type = value.Kind
			weeklyData = append(weeklyData, data)
			openstackWeeklycontent.Content[index].IsDelivered = true
		}
	}

	// push the updated content.yaml
	openstackWeeklycontentByte, err = yaml.Marshal(openstackWeeklycontent)
	if err != nil {
		log.Fatal(err)
	}
	handler.UpdateFile("content.yaml", "update OpenStack Weekly content", openstackWeeklycontentByte)

	//Init builder
	builder := Builder{}

	// Build
	location, _ := time.LoadLocation("Asia/Jakarta")
	time := time.Now().In(location).Format("02-01-2006")
	weeklyName = "OpenStack Weekly " + time

	builder.build(weeklyName, weeklyData)
	weeklyCRD, err := yaml.Marshal(builder)

	commitMessage := "Weekly: Add " + weeklyName
	CreateFile(handler, strings.ToLower(strings.ReplaceAll(weeklyName, " ", "-"))+".yaml", commitMessage, weeklyCRD)
}
