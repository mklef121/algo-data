package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var CSVFILE string = "/Users/m.nwabueze/Documents/software-projects/person/algo-data/golang/learn/projects/phone-csv.csv"

type PhoneBook []Entry

var data = PhoneBook{}

func (a PhoneBook) Len() int {
	return len(a)
}

func (a PhoneBook) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}

	return a[i].Surname < a[j].Surname
}

func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//An index used to map phone numbers to the index of the phonebook slice
var index map[string]int

var MIN = 97
var MAX = 122
var CSV_DELIMETER = '\t'

func main() {
	arguments := os.Args

	// fmt.Println(data, cap(data), len(data), "hia")
	if len(arguments) == 1 {
		// path.Base
		// The argument is like /var/folders/vj/xsf0z1495_915__5r226_phh0000gp/T/go-build2798432677/b001/exe/phonebook
		// So Base funtion gives us the last member which is phonebook
		exe := path.Base(arguments[0])
		fmt.Println("Usage: insert|delete|search|list <arguments>", exe)
		return
	}

	if inArray(arguments[1], []string{"search", "list", "delete", "insert"}) == -1 {
		fmt.Println("Passed argument muse be `search`, `list`, `delete`,`insert`")
		return
	}

	if setCSVFILE() != nil {
		return
	}

	// The contents of CSVFILE are kept in the data global variable that is defined as []Entry{},
	if err := readCSVFile(CSVFILE); err != nil {
		fmt.Println(err)
		return
	}
	if err := createIndex(); err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	// fmt.Println(data, "hia")
	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := initS(arguments[2], arguments[3], t)

		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		err := deleteEntry(t)

		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}

//run export PHONEBOOK="/tmp/csv.file"
func setCSVFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		CSVFILE = filepath
	}

	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating ", CSVFILE, "...")
		f, err := os.Create(CSVFILE)
		if err != nil {
			return err
		}
		defer f.Close()

		return nil
	}

	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", CSVFILE)
	}

	return nil
}

func initS(name, surname, tel string) *Entry {
	if name == "" || surname == "" {
		return nil
	}

	//convert time to string
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: name, Surname: surname, Tel: tel, LastAccess: LastAccess}
}

func insert(phone *Entry) error {
	_, ok := index[(*phone).Tel]
	if ok {
		return fmt.Errorf("%s already exists", phone.Tel)
	}

	data = append(data, *phone)

	// Update the index
	createIndex()
	err := saveCSVFile(CSVFILE)

	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(tel string) error {
	val, ok := index[tel]

	if !ok {
		return fmt.Errorf("%s cannot be found!", tel)
	}

	//if element is last in the queue
	if val == (len(data) - 1) {
		data = data[:val]
	} else {
		data = append(data[:val], data[(val+1):]...)

	}

	delete(index, tel)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

//The version I was using before without environmental variable
func createCSV() error {
	// If the CSVFILE does not exist, create an empty one
	fileInfo, err := os.Stat(CSVFILE)

	if err != nil {
		fmt.Println("Creating file ", CSVFILE)

		f, err := os.Create(CSVFILE)
		defer f.Close()

		if err != nil {
			fmt.Println("error creating file", err)

			return err
		}

		return nil

	} else if mode := fileInfo.Mode(); !mode.IsRegular() {
		fmt.Println(CSVFILE, " not a regular file!")
		return errors.New("Input file is not a regular file")
	}

	return nil
}

func search(key string) *Entry {
	val, ok := index[key]

	if !ok {
		return nil
	}

	data[val].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)

	return &data[val]
}

func list() {
	sort.Sort(PhoneBook(data))
	for _, v := range data {
		fmt.Println(v)
	}
}
func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func inArray(val interface{}, array interface{}) (index int) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	// startChar := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func populate(count int, data []Entry) []Entry {
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	for i := 0; i < count; i++ {
		name := getString(4)
		surname := getString(5)

		tel := strconv.Itoa(random(100, 199))

		data = append(data, Entry{Name: name, Surname: surname, Tel: tel})
	}
	// fmt.Println(data, "fool")
	return data
}

//CH3

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {

		fmt.Println("File does not exist")
		return err
	}

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Cannot read file. Reason: ", err)
		return err
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println("Error reading CSV file with CSV reader. Reason: ", err)
		return nil
	}

	for _, line := range records {
		items := strings.Split(line[0], string(CSV_DELIMETER))
		// fmt.Println(line, len(line), cap(line), line[0], items)
		temp := Entry{
			Name:       items[0],
			Surname:    items[1],
			Tel:        items[2],
			LastAccess: items[3],
		}
		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filepath string) error {

	csvfile, err := os.Create(filepath)

	if err != nil {
		fmt.Println("Could not create file")

		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	csvwriter.Comma = CSV_DELIMETER
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil

}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil
}
