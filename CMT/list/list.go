package list

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
)

type Decoder interface {
	sort.Interface
	Decode([][]string) error
}

type Encoder interface {
	Header() []string
	Encode() [][]string
}

type List interface {
	Encoder
	Decoder
}

// Sort sorts a decoded List.
func Sort(list Decoder) {
	sort.Sort(list)
}

// Read reads csv input and decodes the data into a given List reference.
func Read(reader io.Reader, list Decoder) error {
	data, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return err
	}
	if !(len(data) > 1) {
		return nil
	}

	var items [][]string
	if len(data) > 0 {
		for _, d := range data[1:] {
			var row []string
			for _, r := range d {
				row = append(row, r)
			}
			items = append(items, row)
		}
	}

	if err := list.Decode(items); err != nil {
		return err
	}

	Sort(list)

	return nil
}

// ReadFile reads csv input from a file and decodes elements into a given List reference.
func ReadFile(path string, list Decoder) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := Read(file, list); err != nil {
		return err
	}

	return nil
}

// Write encodes and writes a given List reference.
func Write(writer io.Writer, list Encoder) error {
	var rows [][]string

	res := list.Encode()
	if res == nil {
		return nil
	}

	w := csv.NewWriter(writer)
	rows = append(rows, list.Header())
	rows = append(rows, res...)
	w.WriteAll(rows)

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

// WriteFile encodes a given List reference into a csv file.
func WriteFile(path string, list Encoder) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := Write(file, list); err != nil {
		return err
	}

	return nil
}
