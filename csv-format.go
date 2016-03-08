package main

import (
        "encoding/csv"
        //"encoding/json"
        "fmt"
        "os"
        //"strconv"
        "log"
        "strings"
)

type Empresa struct {
        Rut             string          `json:"rut"`
        Razon           string          `json:"razon"`
        Resolucion      int             `json:"resolucion"`
        Fecha           string          `json:"fecha"`
        Mail            string          `json:"mail"`
        Url             string          `json:"url"`
}

func main() {
// read data from CSV file

        csvFile, err := os.Open("./empresas.csv")

        if err != nil {
                fmt.Println(err)
        }

        defer csvFile.Close()

        reader := csv.NewReader(csvFile)
        reader.Comma = 59
        reader.FieldsPerRecord = -1
        reader.LazyQuotes = true

        csvData, err := reader.ReadAll()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        file, err := os.Create("empresas-formatted.csv")

        if err != nil{
                log.Fatalln("error al intentar crear el archivo csv: ", err)
        }

        writer := csv.NewWriter(file)

        for _, record := range csvData {
                record[1] = strings.TrimSpace(record[1])
                record[4] = strings.TrimSpace(record[4])
                if err := writer.Write(record); err != nil {
                        log.Fatalln("error writing record to csv:", err)
                }
        }
}