package convert

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/signintech/gopdf"
	"path/filepath"
	"net/http")

func PdfA4(c echo.Context) (err error) {
	path := c.QueryParam("path")
	extension := c.QueryParam("extension")
	err = convertImage2PDF(path, extension)
	if err != nil{
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"path": path, "extension": extension})
}

func convertImage2PDF(path string, extension string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89} }) //A4
	var pattern bytes.Buffer
	pattern.WriteString(path)
	pattern.WriteString("/*.")
	pattern.WriteString(extension)
	files, _ := filepath.Glob(pattern.String())
	for _, file := range files{
		pdf.AddPage()
		err := pdf.Image(file, 0, 0, gopdf.PageSizeA4)
		if err != nil {
			return err
		}
	}
	var resultName bytes.Buffer
	resultName.WriteString(path)
	resultName.WriteString("/result.pdf")
	_ = pdf.WritePdf(resultName.String())
	return nil
}
