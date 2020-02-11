 package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Firstname string
	Lastname  string
	Table1    string
	Table2    string
	Table3    string
}

//var table declared globally because it is used in below func findTable()
var table = 1

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

	//Declaring variables that will be used throughout code.
	var assignment = ""
	//tablevar keeps track of how much each person changes tables between dinners
	var tablevar = 1
	//strings display table assignments at the end.
	var table1disp = ""
	var table2disp = ""
	var table3disp = ""

	//seeding so that it is random every time
	rand.Seed(time.Now().UnixNano())

	//tableFill keeps track of how full each table is
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//usedTables keeps track of tables that are full
	usedTables := []int{}
	//studentList := []Person{}

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
		//sets tablevar at the specific person's order of seating in the table. For example, fifth person at a table gets a tableVar value of 5.
		tablevar = tableFill[table] + 1
		//loopvar is set to manually run the for loop three times.
		var loopvar = 1

		for {
			if table < 32 {
				//only runs on the first loop.
				if loopvar == 1 {
					//continues to fill table until there are 8 people.
					if tableFill[table] < 8 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//once 8 people are at the table, adds one more and adds the table to usedTables to ignore
						usedTables = append(usedTables, table)
					}
				}
				//this converts the int table to a string for assignment
				assignment = strconv.Itoa(table)
			} else if table == 32 {
				//again only runs on the first loop
				if loopvar == 1 {
					//table 32 is kitchen crew, so same deal as above but with 6 people instead of 8.
					if tableFill[table] < 6 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//again, once hits 6 then add one more and close it off.
						usedTables = append(usedTables, table)
					}
				}
				//set assignment for Kitchen Crew
				assignment = "Kitchen Crew"
			} else if table == 33 {
				//only runs in first loop again.
				if loopvar == 1 {
					//table 33 will be assigned as waiters.
					if tableFill[table] < 30 {
						tableFill[table]++
					} else {
						tableFill[table]++
						//same deal, go until full.
						usedTables = append(usedTables, table)
					}
				}
				assignment = "Waiter"
			}
			//Each run-through of the loop sets the assignment for one night.
			if loopvar == 1 {
				table1disp = assignment
			} else if loopvar == 2 {
				table2disp = assignment
			} else if loopvar == 3 {
				table3disp = assignment
			}

			//move to the next table assignment and repeat loop.
			table = table + tablevar
			//make sure to avoid index out of range error.
			for {
				if table > 33 {
					table = table - 33
				} else {
					break
				}
			}

			//manually continue the loop
			if loopvar < 3 {
				loopvar++
			} else if loopvar == 3 {
				break
			}

		}
		//Once all three nights are assigned, add specific person to struct Person
		people := Person{
			Firstname: line[1],
			Lastname:  line[0],
			Table1:    table1disp,
			Table2:    table2disp,
			Table3:    table3disp,
		}
		fmt.Println(people)

	}
}
