package cmt_testing

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/GeoNet/data/CMT/list"
)

const data = "."

func TestFiles_Contents(t *testing.T) {

	var tests = []struct {
		l list.List
		f string
	}{
		{&list.CMTs{}, list.CMTsFile},
	}

	for _, tt := range tests {

		t.Logf("decode list file: %s", tt.f)
		if err := list.ReadFile(filepath.Join(data, tt.f), tt.l); err != nil {
			t.Fatal(err)
		}

		t.Logf("compare list file: %s", tt.f)
		file, err := os.OpenFile(filepath.Join(data, tt.f), os.O_RDONLY, 0)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		raw, err := func() ([]byte, error) {
			r, e := ioutil.ReadAll(file)
			return r, e
		}()
		if err != nil {
			t.Fatal(err)
		}

		res := tt.l.Encode()
		if res == nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		w := csv.NewWriter(&buf)

		var rows [][]string
		rows = append(rows, tt.l.Header())
		rows = append(rows, res...)
		w.WriteAll(rows)
		if err := w.Error(); err != nil {
			t.Fatal(err)
		}

		if string(raw) != buf.String() {
			t.Errorf("List file mismatch: %s", tt.f)

			file, err := ioutil.TempFile(os.TempDir(), "tst")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(file.Name())
			file.Write(buf.Bytes())

			cmd := exec.Command("diff", "-c", filepath.Join(data, tt.f), file.Name())
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				t.Fatal(err)
			}
			err = cmd.Start()
			if err != nil {
				t.Fatal(err)
			}
			defer cmd.Wait()
			diff, err := ioutil.ReadAll(stdout)
			if err != nil {
				t.Fatal(err)
			}
			t.Error(string(diff))
		}
	}
}
