package CSVReadWriter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	HomeAddress string `json:"homeaddress"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalcode"`
}

func convertJSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var ranking []Person
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"Firstname", "Lastname", "Address"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.Firstname, r.Lastname, fmt.Sprint(r.Address))
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func CSVReadWriter() {
	csvFileName := "addresses.csv"
	csvFileRead, err := os.Open(csvFileName)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(csvFileRead)

	var people []Person

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				HomeAddress: line[2],
				City:        line[3],
				State:       line[4],
				PostalCode:  line[5],
			},
		})
	}

	peopleJson, _ := json.MarshalIndent(people, "", " ")

	_ = ioutil.WriteFile("test.json", peopleJson, 0644)

	convertJSONToCSV("test.json", "data.csv")
	fmt.Println(string(peopleJson))
}

// things to do:
//1. change name of package and function of the package
//2. parameterize name of the input and output file
