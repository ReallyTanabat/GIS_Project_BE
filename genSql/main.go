package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	_, err2 := f.WriteString("INSERT INTO AirPollutionPM25(Country,City,Year,Pm25,Latitude,Longtitude,Population,Wbinc16_text,Region,Conc_pm25,Color_pm25,geom) VALUES\n")

	if err2 != nil {
		log.Fatal(err2)
	}
	records := readCsvFile("./data.csv")
	isTrue := false
	count := 0
	for col, row := range records {
		if isTrue {
			if count == 999 {
				_, err2 = f.WriteString(";\n")

				if err2 != nil {
					log.Fatal(err2)
				}
				_, err2 := f.WriteString("INSERT INTO AirPollutionPM25(Country,City,Year,Pm25,Latitude,Longtitude,Population,Wbinc16_text,Region,Conc_pm25,Color_pm25,geom) VALUES\n")

				if err2 != nil {
					log.Fatal(err2)
				}

				count = 0
			}
			var city strings.Builder

			if strings.Contains(row[1], "'") {
				for _, data := range row[1] {
					city.WriteRune(data)
					if data == 39 {
						city.WriteRune(data)
					}
				}

			} else {
				city.WriteString(row[1])
			}
			_, err2 := f.WriteString("('" + row[0] + "','" + city.String() + "'," + row[2] + "," + row[3] + "," + row[4] + "," + row[5] + "," + row[6] + ",'" + row[7] + "','" + row[8] + "','" + row[9] + "','" + row[10] + "',geometry::Parse('POINT(" + row[4] + " " + row[5] + ")'))")
			if err2 != nil {
				log.Fatal(err2)
			}
			if len(records) == col+1 {
				count = 998
			}
			if count != 998 {
				_, err2 := f.WriteString(",\n")
				if err2 != nil {
					log.Fatal(err2)
				}
			}
			// if len(records) != col+1 {
			// 	_, err2 := f.WriteString(",\n")
			// 	if err2 != nil {
			// 		log.Fatal(err2)
			// 	}
			// }
			count++
		}
		isTrue = true
	}
	_, err2 = f.WriteString(";\n")

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("done")
	// fmt.Println(records)
}
