package th_xlsx

import (
	"testing"
)

func TestImport(t *testing.T) {
	Import("in_student.xlsx")
}

func TestExport(t *testing.T) {
	Export("out_student.xlsx")
}
