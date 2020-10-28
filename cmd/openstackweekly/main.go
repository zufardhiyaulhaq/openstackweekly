package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zufardhiyaulhaq/openstackweekly/handlers"
	"github.com/zufardhiyaulhaq/openstackweekly/models"
	"github.com/zufardhiyaulhaq/openstackweekly/scrappers"
	"gopkg.in/yaml.v2"
	// communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
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
		fmt.Println(err)
	}

	// feed current content to scrapper
	openstackWeekluData := scrapper.GetWeekly(content)

	// // get latest weekly from handlers
	// recentWeeklyNames := GetFiles(handler)

	// compare weekly logic
	// compare newest weekly from scrapper
	// with latest list of weekly from datastore
	// newestWeeklyName := scrapper.GetWeeklyName()
	// for _, v := range recentWeeklyNames {
	// 	if strings.ToLower(strings.ReplaceAll(newestWeeklyName, " ", "-"))+".yaml" == v {
	// 		log.Println("Weekly already in datastore")
	// 		return
	// 	}
	// }

	// Scapper logic
	// must return list fo ArticleSpec defined in community-operator
	// communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	// var weekly []communityv1alpha1.ArticleSpec
	// weekly = scrapper.GetWeekly()

	// Init builder
	// builder := Builder{}

	// Build
	location, _ := time.LoadLocation("Asia/Jakarta")
	time := time.Now().In(location).Format("02-01-2006")
	name := "OpenStack Weekly " + time
	fmt.Println(name)
	// builder.build(name, weekly)

	// Add to Github
	// crd, err := yaml.Marshal(builder)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// commitMessage := "Weekly: Add " + newestWeeklyName
	// CreateFile(handler, strings.ToLower(strings.ReplaceAll(newestWeeklyName, " ", "-"))+".yaml", commitMessage, crd)
}
