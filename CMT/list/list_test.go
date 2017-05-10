package list

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"testing"
)

func TestFiles_Contents(t *testing.T) {

	var tests = []struct {
		l List
		f string
		d string
	}{
		{
			&Aliases{},
			"testdata/aliases.csv",
			"-Header,Address\n+Device,Address",
		},
		{
			&Models{},
			"testdata/models.csv",
			"-Header,Make,Type,Metrics,States,Tags\n+Model,Make,Type,Metrics,States,Tags",
		},
		{
			&Localities{},
			"testdata/localities.csv",
			"-Header,Name,Code,Latitude,Longitude,Height\n+Locality,Name,Code,Latitude,Longitude,Height",
		},
		{
			&Devices{},
			"testdata/devices.csv",
			"-Header,Model,Address,Code,Tags,Notes\n+Hostname,Model,Address,Code,Tags,Notes",
		},
		{
			&Linknets{},
			"testdata/linknets.csv",
			"-Header,Local,Remote,Name,Tags,Cost\n+Linknet,Local,Remote,Name,Tags,Cost",
		},
		{
			&Runnets{},
			"testdata/runnets.csv",
			"-Header,Locality,Name,Tags\n+Runnet,Locality,Name,Tags",
		},
		{
			&Providers{},
			"testdata/providers.csv",
			"-Header,Name,Notes\n+Provider,Name,Notes",
		},
		{
			&Subnets{},
			"testdata/subnets.csv",
			"-Header,Provider,Name\n+Subnet,Provider,Name",
		},
		{
			&Tunnels{},
			"testdata/tunnels.csv",
			"-Header,Local,Remote,Name,Tags,Cost\n+Tunnel,Local,Remote,Name,Tags,Cost",
		},
		{
			&Supernets{},
			"testdata/supernets.csv",
			"-Header,Name,Provider,Sparse\n+Supernet,Name,Provider,Sparse",
		},
		{
			&Circuits{},
			"testdata/circuits.csv",
			"-Header,Provider,Circuit,Service,Number,Notes\n+Hostname,Provider,Circuit,Service,Number,Notes",
		},
		{
			&Tags{},
			"testdata/tags.csv",
			"-Header,Service,Metrics,Public\n+Tag,Service,Metrics,Public",
		},
		{
			&Thresholds{},
			"testdata/thresholds.csv",
			"-Header,Tags,Lower,Upper\n+Metric,Tags,Lower,Upper",
		},
		{
			&Mappings{},
			"testdata/mappings.csv",
			"-Header,Aliases\n+CName,Aliases",
		},
		{
			&Cellulars{},
			"testdata/cellulars.csv",
			"-Header,Provider,Username,Password,SIMCode,SIMNumber,Tags\n+Address,Provider,Username,Password,SIMCode,SIMNumber,Tags",
		},
	}

	for _, tt := range tests {

		t.Logf("decode list file: %s", tt.f)
		if err := ReadFile(tt.f, tt.l); err != nil {
			t.Fatal(err)
		}

		file, err := os.OpenFile(tt.f, os.O_RDONLY, 0)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		raw, err := func() (string, error) {
			r, e := ioutil.ReadAll(file)
			return string(r), e
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

		//Diff is a simple textual difference displayer.
		if diff := func(a, b string) string {
			var diff []string

			if a != b {
				seq1, seq2 := strings.Split(a, "\n"), strings.Split(b, "\n")

				var start, end int
				for start < len(seq1) && start < len(seq2) && seq1[start] == seq2[start] {
					start++
				}
				i, j := len(seq1)-1, len(seq2)-1
				for i > start && j > start && seq1[i] == seq2[j] {
					i--
					j--
					end++
				}

				subseq1, subseq2 := seq1[start:len(seq1)-end], seq2[start:len(seq2)-end]

				matrix := make([][]int, len(subseq1)+1)
				for i := 0; i < len(subseq1)+1; i++ {
					matrix[i] = make([]int, len(subseq2)+1)
				}
				for i := 1; i < len(matrix); i++ {
					for j := 1; j < len(matrix[i]); j++ {
						if subseq1[len(subseq1)-i] == subseq2[len(subseq2)-j] {
							matrix[i][j] = matrix[i-1][j-1] + 1
						} else {
							matrix[i][j] = int(math.Max(float64(matrix[i-1][j]),
								float64(matrix[i][j-1])))
						}
					}
				}

				i, j = len(subseq1), len(subseq2)
				for i > 0 || j > 0 {
					if i > 0 && matrix[i][j] == matrix[i-1][j] {
						diff = append(diff, "-"+subseq1[len(subseq1)-i])
						i--
					} else if j > 0 && matrix[i][j] == matrix[i][j-1] {
						diff = append(diff, "+"+subseq2[len(subseq2)-j])
						j--
					} else if i > 0 && j > 0 {
						i--
						j--
					}
				}
			}

			return strings.Join(diff, "\n")
		}(raw, buf.String()); diff != tt.d {
			if diff == "" {
				t.Errorf("empty list file diff %s:\n===\n%s\n===", tt.f, tt.d)
			} else if string(diff) != tt.d {
				t.Errorf("invalid list file diff %s:\n===\n%s\n===\n%s\n===", tt.f, string(diff), tt.d)
			}
		}
	}
}
