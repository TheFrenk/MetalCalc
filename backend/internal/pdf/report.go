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

// drawSectionHeader малює заголовок секції з фоновою смугою
func drawSectionHeader(pdf *fpdf.Fpdf, text string) {
	pdf.SetFillColor(30, 80, 160)   // темно-синій фон
	pdf.SetTextColor(255, 255, 255) // білий текст
	pdf.SetFont("DejaVu", "B", 11)
	pdf.CellFormat(0, 9, text, "", 1, "L", true, 0, "")
	pdf.SetTextColor(30, 30, 30)
	pdf.Ln(2)
}

// drawRow малює рядок таблиці з чергуванням кольорів
func drawRow(pdf *fpdf.Fpdf, label, value string, even bool) {
	if even {
		pdf.SetFillColor(240, 245, 255) // світло-блакитний
	} else {
		pdf.SetFillColor(255, 255, 255) // білий
	}

	pdf.SetFont("DejaVu", "B", 10)
	pdf.CellFormat(85, 7, label+":", "LB", 0, "L", true, 0, "")

	pdf.SetFont("DejaVu", "", 10)
	pdf.CellFormat(0, 7, value, "RB", 1, "L", true, 0, "")
}

func GenerateForgingReport(data ReportData, fontDir string) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")

	// Підключення TTF-шрифту з підтримкою UTF-8 (кирилиця)
	// Розмістіть DejaVuSans.ttf та DejaVuSans-Bold.ttf у папці fonts/
	pdf.AddUTF8Font("DejaVu", "", fontDir+"/DejaVuSans.ttf")
	pdf.AddUTF8Font("DejaVu", "B", fontDir+"/DejaVuSans-Bold.ttf")

	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 15)
	pdf.SetMargins(15, 15, 15)

	// ── Шапка ──────────────────────────────────────────────
	pdf.SetFillColor(15, 55, 130)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("DejaVu", "B", 18)
	pdf.CellFormat(0, 16, data.Title, "", 1, "C", true, 0, "")
	pdf.Ln(4)

	// Дата генерації
	pdf.SetFont("DejaVu", "", 9)
	pdf.SetTextColor(100, 100, 100)
	pdf.CellFormat(0, 6,
		fmt.Sprintf("Дата формування: %s", time.Now().Format("02.01.2006 15:04")),
		"", 1, "R", false, 0, "")
	pdf.Ln(6)

	// ── Секція: Вхідні параметри ───────────────────────────
	pdf.SetTextColor(30, 30, 30)
	drawSectionHeader(pdf, "  Вхідні параметри")

	for i, p := range data.Params {
		val := p.Value
		if p.Unit != "" {
			val += "  " + p.Unit
		}
		drawRow(pdf, p.Label, val, i%2 == 0)
	}
	pdf.Ln(8)

	// ── Секція: Результати розрахунків ─────────────────────
	drawSectionHeader(pdf, "  Результати розрахунків")

	for i, r := range data.Results {
		val := r.Value
		if r.Unit != "" {
			val += "  " + r.Unit
		}
		drawRow(pdf, r.Label, val, i%2 == 0)
	}
	pdf.Ln(8)

	// ── Підвал ─────────────────────────────────────────────
	pdf.SetFont("DejaVu", "", 8)
	pdf.SetTextColor(160, 160, 160)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(15, pdf.GetY(), 195, pdf.GetY())
	pdf.Ln(2)
	pdf.CellFormat(0, 5, "Звіт згенеровано автоматично — Система розрахунку кування", "", 0, "C", false, 0, "")

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
