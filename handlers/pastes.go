package handlers

import (
	"www-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
	"www-th3-z-xyz/storage"
)

func Pastes(c echo.Context) error {
	data := struct {
		Page models.Page
		Pastes []models.Paste
	} {
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Pastes",
			Id:          "pastes",
		},
		Pastes: models.GetPastes(),
	}

	return c.Render(http.StatusOK, "pastes/index", data)
}

func NewPaste(c echo.Context) error {
	content := []byte(c.FormValue("content"))
	uploaderId := "test"

	paste, err := models.NewPaste(storage.Db, content, uploaderId)
	if err != nil {
		return err
	}

	return c.Redirect(302, "files/"+paste.Filename)
}