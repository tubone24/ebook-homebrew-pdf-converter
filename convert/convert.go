package convert

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
)

func PdfA4(c echo.Context) (err error) {
	path := c.QueryParam("path")
	extension := c.QueryParam("extension")
	err = convertImage2PDF(path, extension)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"path": path, "extension": extension})
}

func convertImage2PDF(path string, extension string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	var pattern bytes.Buffer
	pattern.WriteString(path)
	pattern.WriteString("/*.")
	pattern.WriteString(extension)
	files, _ := filepath.Glob(pattern.String())
	for _, file := range files {
		pdf.AddPage()
		pdf.ImageOptions(file, 0, 0, 210, 297, false, gofpdf.ImageOptions{ImageType: extension, ReadDpi: true}, 0, "")
	}
	var resultName bytes.Buffer
	resultName.WriteString(path)
	resultName.WriteString("/result.pdf")
	err := pdf.OutputFileAndClose(resultName.String())
	if err != nil {
		return err
	}
	return nil
}
