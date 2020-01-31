package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var table = 1
var assignment = ""
var waiterTable = 1

func findTable() {
	//find a new table assignment number
	table = rand.Intn(33)
	//gets rid of table 0
	table++

}

func contains(s []int, e int) bool {
	//searches for an integer through a slice. Used to check if a table is already filled before assigning.
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	//seeding so that it is random every time
	rand.Seed(time.Now().UnixNano())

	//tableFill keeps track of how full each table is
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//usedTables keeps track of tables that are full
	usedTables := []int{}

	csvFile, _ := os.Open("Dinner.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//picks a random number to assign the current person to a table
		findTable()

		for {
			//continues to find a new number until it finds a table or assignment that isn't filled up yet.
			if contains(usedTables, table) {
				findTable()
			} else {
				break
			}
		}

		if table < 32 {
			//continues to fill table until there are 8 people.
			if tableFill[table] < 8 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//once 8 people are at the table, adds one more and adds the table to usedTables to ignore.
				usedTables = append(usedTables, table)
			}

			fmt.Println(line[0], line[1], table)

		} else if table == 32 {
			//table 32 is kitchen crew, so same deal as above but with 6 people instead of 8.
			if tableFill[table] < 6 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//again, once hits 6 then add one more and close it off.
				usedTables = append(usedTables, table)
			}
			assignment = "Kitchen Crew"
			fmt.Println(line[0], line[1], assignment)
		} else if table == 33 {
			//table 33 will be assigned as waiters.
			if tableFill[table] < 30 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//same deal, go until full.
				usedTables = append(usedTables, table)
			}
			assignment = "Waiter"
			waiterTable++
			fmt.Println(line[0], line[1], assignment, waiterTable)
		}
	}
}
