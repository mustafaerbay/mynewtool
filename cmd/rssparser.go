/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"encoding/csv"
	"os"
	"strconv"

	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

type issue struct {
	title string
	author string
	assignee string
}

// rssParserCmd represents the parser command
var rssParserCmd = &cobra.Command{
	Use:   "rssparser",
	Short: "list unresolved issues",
	Long: `this rss parser for get some issues from gitlab from rss resources`,
	Run: func(cmd *cobra.Command, args []string) {
		fp := gofeed.NewParser()
		GITLAB_URL := "https://rnd-gitlab-eu.huawei.com/htrdc-isd/ebg/octopus/-/issues.atom?feed_token=crmDdCRB8H_HwasY6WRS&not%5Blabel_name%5D%5B%5D=Fixed+-+Done&scope=all&state=opened&utf8=%E2%9C%93"
		feed, err := fp.ParseURL(GITLAB_URL)
		checkError("Gitlab_url not reachable",err)
		f := excelize.NewFile()
		index := f.NewSheet("Sheet1")
		file, err := os.Create("result.csv")
		checkError("Cannot create file", err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		// a := make([]string, 5)
		// a = feed.Items()
		title := "Issues Title"
		author := "Issue Author"
		// fmt.Println("")
		// fmt.Printf("%s|%s",title,author)
		// fmt.Printf()
		f.SetCellValue("Sheet1", "B1", title)
		f.SetCellValue("Sheet1", "C1", author)
		for i := 0; i < len(feed.Items); i++ {
		 	// fmt.Printf("%s|%s", feed.Items[i].Title, feed.Items[i].Author.Name)
		// 	fmt.Println("")
			newissue := issue{
				title: feed.Items[i].Title,
				author: feed.Items[i].Author.Name,
			}
			// val1=s+strconv.Itoa(i+1)
			val2_string := strconv.Itoa(i+1)+"|"+newissue.title+"|"+newissue.author
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), strconv.Itoa(i+1))
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), newissue.title)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), newissue.author)
			f.SetActiveSheet(index)
			if err := f.SaveAs("Book1.xlsx"); err != nil {
				fmt.Println(err)
			}
			fmt.Println(val2_string)
		}
	},
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatalln(message, err)
    }
}
func init() {
	rootCmd.AddCommand(rssParserCmd)
}
