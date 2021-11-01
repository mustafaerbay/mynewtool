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
	title  string
	author string
}

// rssParserCmd represents the parser command
var rssParserCmd = &cobra.Command{
	Use:   "rssparser",
	Short: "list unresolved issues",
	Long:  `this rss parser for get some issues from gitlab from rss resources`,
	Run: func(cmd *cobra.Command, args []string) {
		fp := gofeed.NewParser()
		gitlabURL := "https://rnd-gitlab-eu.huawei.com/htrdc-isd/ebg/octopus/-/issues.atom?feed_token=crmDdCRB8H_HwasY6WRS&not%5Blabel_name%5D%5B%5D=Fixed+-+Done&scope=all&state=opened&utf8=%E2%9C%93"
		feed, err := fp.ParseURL(gitlabURL)
		checkError("gitlabURL not reachable", err)
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
		title2err := f.SetCellValue("Sheet1", "B1", title)
		if title2err != nil {
			fmt.Printf("error setting title")
		}

		author2err := f.SetCellValue("Sheet1", "C1", author)
		if author2err != nil {
			fmt.Printf("error setting author")
		}

		for i := 0; i < len(feed.Items); i++ {
			// fmt.Printf("%s|%s", feed.Items[i].Title, feed.Items[i].Author.Name)
			// 	fmt.Println("")
			newissue := issue{
				title:  feed.Items[i].Title,
				author: feed.Items[i].Author.Name,
			}
			// val1=s+strconv.Itoa(i+1)
			val2string := strconv.Itoa(i+1) + "|" + newissue.title + "|" + newissue.author
			err := f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), strconv.Itoa(i+1))
			if err != nil {
				fmt.Printf("error setting number of raw:%s", strconv.Itoa(i+1))
			}
			titleErr := f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), newissue.title)
			if titleErr != nil {
				fmt.Printf("error setting issue title:%s", newissue.title)
			}
			authorerr := f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), newissue.author)
			if authorerr != nil {
				fmt.Printf("error setting issue author:%s", newissue.author)
			}
			f.SetActiveSheet(index)
			if err := f.SaveAs("Book1.xlsx"); err != nil {
				fmt.Println(err)
			}
			fmt.Println(val2string)
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
