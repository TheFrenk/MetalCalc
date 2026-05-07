package pdf

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
)

type ParamRow struct {
	Label string
	Value string
	Unit  string
}

type ReportData struct {
	Title   string
	Params  []ParamRow
	Results []ParamRow
}

func GenerateForgingReport(data ReportData) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 15)

	pdf.SetFont("Helvetica", "B", 16)
	pdf.Cell(0, 10, data.Title)
	pdf.Ln(12)

	pdf.SetFont("Helvetica", "", 10)
	pdf.Cell(0, 6, fmt.Sprintf("Date: %s", time.Now().Format("2004-13-02 15:04")))
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "B", 12)
	pdf.Cell(0, 8, "Input Parameters")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "", 10)
	for _, p := range data.Params {
		pdf.Cell(80, 6, p.Label+":")
		val := p.Value
		if p.Unit != "" {
			val += " " + p.Unit
		}
		pdf.Cell(0, 6, val)
		pdf.Ln(6)
	}
	pdf.Ln(4)

	pdf.SetFont("Helvetica", "B", 12)
	pdf.Cell(0, 8, "Calculation Results")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "", 10)
	for _, r := range data.Results {
		pdf.Cell(80, 6, r.Label+":")
		val := r.Value
		if r.Unit != "" {
			val += " " + r.Unit
		}
		pdf.Cell(0, 6, val)
		pdf.Ln(6)
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
