package list

import (
	"fmt"
	"strconv"
	"time"
)

const (
	cmtEVENT_ID int = iota
	cmtDate
	cmtLatitude
	cmtLongitude
	cmtStr1
	cmtDp1
	cmtRake1
	cmtStr2
	cmtDp2
	cmtRake2
	cmtML
	cmtMw
	cmtMo
	cmtCD
	cmtNS
	cmtDC
	cmtMxx
	cmtMxy
	cmtMxz
	cmtMyy
	cmtMyz
	cmtMzz
	cmtVR
	cmtTva
	cmtTpl
	cmtTaz
	cmtNva
	cmtNpl
	cmtNaz
	cmtPva
	cmtPpl
	cmtPaz
	cmtLast
)

// CMT represents a single CMT solution.
type CMT struct {
	EVENT_ID  string
	Date      time.Time
	Latitude  float32
	Longitude float32
	Str1      int16
	Dp1       int16
	Rake1     int16
	Str2      int16
	Dp2       int16
	Rake2     int16
	ML        float32
	Mw        float32
	Mo        float64
	CD        int16
	NS        int16
	DC        int16
	Mxx       float32
	Mxy       float32
	Mxz       float32
	Myy       float32
	Myz       float32
	Mzz       float32
	VR        int16
	Tva       float32
	Tpl       int16
	Taz       int16
	Nva       float32
	Npl       int16
	Naz       int16
	Pva       float32
	Ppl       int16
	Paz       int16
}

// Date: yyyymmddhhmm00
const DateFormat string = "20060102150400"

func (c *CMT) decode(row []string) error {
	if len(row) != cmtLast {
		return fmt.Errorf("incorrect item length")
	}
	t, err := time.Parse(DateFormat, parseRaw(row[cmtDate]))
	if err != nil {
		return err
	}

	tLatitude, err := parseFloat32(row[cmtLatitude])
	if err != nil {
		return err
	}
	tLongitude, err := parseFloat32(row[cmtLongitude])
	if err != nil {
		return err
	}
	tStr1, err := parseInt16(row[cmtStr1])
	if err != nil {
		return err
	}
	tDp1, err := parseInt16(row[cmtDp1])
	if err != nil {
		return err
	}
	tRake1, err := parseInt16(row[cmtRake1])
	if err != nil {
		return err
	}
	tStr2, err := parseInt16(row[cmtStr2])
	if err != nil {
		return err
	}
	tDp2, err := parseInt16(row[cmtDp2])
	if err != nil {
		return err
	}
	tRake2, err := parseInt16(row[cmtRake2])
	if err != nil {
		return err
	}
	tML, err := parseFloat32(row[cmtML])
	if err != nil {
		return err
	}
	tMw, err := parseFloat32(row[cmtMw])
	if err != nil {
		return err
	}
	tMo, err := parseFloat64(row[cmtMo])
	if err != nil {
		return err
	}
	tCD, err := parseInt16(row[cmtCD])
	if err != nil {
		return err
	}
	tNS, err := parseInt16(row[cmtNS])
	if err != nil {
		return err
	}
	tDC, err := parseInt16(row[cmtDC])
	if err != nil {
		return err
	}
	tMxx, err := parseFloat32(row[cmtMxx])
	if err != nil {
		return err
	}
	tMxy, err := parseFloat32(row[cmtMxy])
	if err != nil {
		return err
	}
	tMxz, err := parseFloat32(row[cmtMxz])
	if err != nil {
		return err
	}
	tMyy, err := parseFloat32(row[cmtMyy])
	if err != nil {
		return err
	}
	tMyz, err := parseFloat32(row[cmtMyz])
	if err != nil {
		return err
	}
	tMzz, err := parseFloat32(row[cmtMzz])
	if err != nil {
		return err
	}
	tVR, err := parseInt16(row[cmtVR])
	if err != nil {
		return err
	}
	tTva, err := parseFloat32(row[cmtTva])
	if err != nil {
		return err
	}
	tTpl, err := parseInt16(row[cmtTpl])
	if err != nil {
		return err
	}
	tTaz, err := parseInt16(row[cmtTaz])
	if err != nil {
		return err
	}
	tNva, err := parseFloat32(row[cmtNva])
	if err != nil {
		return err
	}
	tNpl, err := parseInt16(row[cmtNpl])
	if err != nil {
		return err
	}
	tNaz, err := parseInt16(row[cmtNaz])
	if err != nil {
		return err
	}
	tPva, err := parseFloat32(row[cmtPva])
	if err != nil {
		return err
	}
	tPpl, err := parseInt16(row[cmtPpl])
	if err != nil {
		return err
	}
	tPaz, err := parseInt16(row[cmtPaz])
	if err != nil {
		return err
	}

	*c = CMT{
		EVENT_ID:  row[cmtEVENT_ID],
		Date:      t,
		Latitude:  *tLatitude,
		Longitude: *tLongitude,
		Str1:      *tStr1,
		Dp1:       *tDp1,
		Rake1:     *tRake1,
		Str2:      *tStr2,
		Dp2:       *tDp2,
		Rake2:     *tRake2,
		ML:        *tML,
		Mw:        *tMw,
		Mo:        *tMo,
		CD:        *tCD,
		NS:        *tNS,
		DC:        *tDC,
		Mxx:       *tMxx,
		Mxy:       *tMxy,
		Mxz:       *tMxz,
		Myy:       *tMyy,
		Myz:       *tMyz,
		Mzz:       *tMzz,
		VR:        *tVR,
		Tva:       *tTva,
		Tpl:       *tTpl,
		Taz:       *tTaz,
		Nva:       *tNva,
		Npl:       *tNpl,
		Naz:       *tNaz,
		Pva:       *tPva,
		Ppl:       *tPpl,
		Paz:       *tPaz,
	}

	return nil
}

func (c CMT) encode() []string {
	var row []string

	row = append(row, c.EVENT_ID)
	row = append(row, c.Date.Format(DateFormat))
	row = append(row, strconv.FormatFloat(float64(c.Latitude), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Longitude), 'f', -1, 32))
	row = append(row, strconv.FormatInt(int64(c.Str1), 10))
	row = append(row, strconv.FormatInt(int64(c.Dp1), 10))
	row = append(row, strconv.FormatInt(int64(c.Rake1), 10))
	row = append(row, strconv.FormatInt(int64(c.Str2), 10))
	row = append(row, strconv.FormatInt(int64(c.Dp2), 10))
	row = append(row, strconv.FormatInt(int64(c.Rake2), 10))
	row = append(row, strconv.FormatFloat(float64(c.ML), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Mw), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(c.Mo, 'E', -1, 32))
	row = append(row, strconv.FormatInt(int64(c.CD), 10))
	row = append(row, strconv.FormatInt(int64(c.NS), 10))
	row = append(row, strconv.FormatInt(int64(c.DC), 10))
	row = append(row, strconv.FormatFloat(float64(c.Mxx), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Mxy), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Mxz), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Myy), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Myz), 'f', -1, 32))
	row = append(row, strconv.FormatFloat(float64(c.Mzz), 'f', -1, 32))
	row = append(row, strconv.FormatInt(int64(c.VR), 10))
	row = append(row, strconv.FormatFloat(float64(c.Tva), 'f', -1, 32))
	row = append(row, strconv.FormatInt(int64(c.Tpl), 10))
	row = append(row, strconv.FormatInt(int64(c.Taz), 10))
	row = append(row, strconv.FormatFloat(float64(c.Nva), 'f', -1, 32))
	row = append(row, strconv.FormatInt(int64(c.Npl), 10))
	row = append(row, strconv.FormatInt(int64(c.Naz), 10))
	row = append(row, strconv.FormatFloat(float64(c.Pva), 'f', -1, 32))
	row = append(row, strconv.FormatInt(int64(c.Ppl), 10))
	row = append(row, strconv.FormatInt(int64(c.Paz), 10))

	return row
}

type CMTs []CMT

func (c CMTs) Len() int      { return len(c) }
func (c CMTs) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CMTs) Less(i, j int) bool {
	switch {
	case c[i].Date.Before(c[j].Date):
		return true
	case c[i].Date.After(c[j].Date):
		return false
	default:
		return c[i].EVENT_ID < c[j].EVENT_ID
	}
}

func (c CMTs) Header() []string {
	return []string{"EVENT_ID", "Date", "Latitude", "Longitude", "Str1", "Dp1", "Rake1", "Str2", "Dp2", "Rake2", "ML", "Mw", "Mo", "CD", "NS", "DC", "Mxx", "Mxy", "Mxz", "Myy", "Myz", "Mzz", "VR", "Tva", "Tpl", "Taz", "Nva", "Npl", "Naz", "Pva", "Ppl", "Paz"}
}
func (c CMTs) Encode() [][]string {
	var items [][]string

	for _, addr := range c {
		items = append(items, addr.encode())
	}

	return items
}
func (c *CMTs) Decode(data [][]string) error {
	for _, v := range data {
		var cmt CMT
		if err := cmt.decode(v); err != nil {
			return err
		}
		*c = append(*c, cmt)
	}
	return nil
}
