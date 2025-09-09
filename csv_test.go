package filange_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vieolo/filange"
)

func TestReadCSV(t *testing.T) {

	restult, err := filange.ReadCsvFile("./files/weekday.csv")

	assert.Nil(t, err)

	expected := [][]string{
		{"Name", "Abbreviation", "Numeric", "Numeric-2"},
		{"Monday", "Mon.", "1", "01"},
		{"Tuesday", "Tue.", "2", "02"},
		{"Wednesday", "Wed.", "3", "03"},
		{"Thursday", "Thu.", "4", "04"},
		{"Friday", "Fri.", "5", "05"},
		{"Saturday", "Sat.", "6", "06"},
		{"Sunday", "Sun.", "7", "07"},
	}

	assert.Equal(t, expected, restult)

}
