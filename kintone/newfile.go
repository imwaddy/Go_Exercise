package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	kintone "github.com/kintone-labs/go-kintone"
)

func main() {

	appIDs := []uint64{}

	for _, appID := range appIDs {

		// Set URL, username and password
		app := &kintone.App{
			Domain:   "ke.com",
			User:     "",
			Password: "",
			AppId:    appID,
		}

		if app == nil {
			fmt.Printf("App ID %v is null", appID)
			fmt.Println("------------------------------------------------------------------------------")
			continue
		}

		// Get all records from app
		records, err := app.GetAllRecords(nil)
		if err != nil {
			fmt.Printf("Error while get reocrds %+v", err)
			fmt.Println("------------------------------------------------------------------------------")
			continue
		}

		fmt.Println("Records in app ID ", appID, " is ", len(records))

		// get current working directory
		path, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error while get working dir %+v", err)
			fmt.Println("------------------------------------------------------------------------------")
			continue
		}
		fmt.Println("Current Path : ", path)

		// create folder with appID
		folder := path + string(os.PathSeparator) + strconv.Itoa(int(appID)) + string(os.PathSeparator)
		err = os.Mkdir(folder, 0777)
		if err != nil && !strings.Contains(err.Error(), "file exists") {
			fmt.Printf("failed creating dir: %s", err.Error())
			fmt.Println("------------------------------------------------------------------------------")
			continue
		}

		// Create file in app folder
		fileName := strconv.Itoa(int(appID)) + ".csv"
		csvFile, err := os.Create(folder + fileName)
		if err != nil {
			fmt.Printf("failed creating file: %s", err)
			fmt.Println("------------------------------------------------------------------------------")
			continue
		}
		defer csvFile.Close()

		csvwriter := csv.NewWriter(csvFile)

		headers := []string{"Record ID", "ID", "Date", "text", "Creator Code", "Creator Name"}

		var ss [][]string
		var mentionsLength int
		totalComments := 0
		// Loop through records
		for _, record := range records {

			var comments []kintone.Comment

			limit := uint64(10)
			page := uint64(0)
			for {

				skip := page * limit
				// Get record comments till error occurs
				commentsArr, err := app.GetRecordComments(record.Id(), "asc", skip, limit)
				if err != nil {
					fmt.Println("Error get comment record ID: ", record.Id(), " Error %+v", err)
					break
				}
				if len(commentsArr) > 0 {
					comments = append(comments, commentsArr...)
				} else {
					break
				}
				page++

			}

			fmt.Println("Comment Length ", len(comments), " for Record ID ", record.Id())
			totalComments += len(comments)

			for _, e := range comments {
				sss := []string{}
				sss = append(sss, strconv.Itoa(int(record.Id())))
				sss = append(sss, e.Id)
				sss = append(sss, e.CreatedAt)
				sss = append(sss, e.Text)
				if e.Creator == nil {
					sss = append(sss, "", "")
				} else {
					sss = append(sss, e.Creator.Code)
					sss = append(sss, e.Creator.Name)
				}

				if e.Mentions != nil {

					if mentionsLength < len(e.Mentions) {
						mentionsLength = len(e.Mentions)
					}

					for _, mention := range e.Mentions {
						sss = append(sss, mention.Code, mention.Type)
					}

				}

				ss = append(ss, sss)

			}

		}

		for i := 1; i <= mentionsLength; i++ {
			headers = append(headers, "mentions "+strconv.Itoa(i)+" code")
			headers = append(headers, "mentions "+strconv.Itoa(i)+" type")
		}

		_ = csvwriter.Write(headers)
		for _, r := range ss {
			recLen := len(r)
			if recLen < (mentionsLength + len(headers)) {
				loop := recLen - (mentionsLength + len(headers))
				for i := 0; i < loop; i++ {
					r = append(r, "", "")
				}
			}
			_ = csvwriter.Write(r)
		}

		csvwriter.Flush()

		fmt.Println("Total comments for appID is ", totalComments)

		fmt.Println("------------------------------------------------------------------------------")

	}

	fmt.Println("-----------------------------------FINISHED-------------------------------------------")

}
