package service

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Service) CreateNote(c echo.Context) error {
	note := new(Note)
	err := c.Bind(&note)
	fmt.Println(c.Request().Header)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	resp, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		s.logger.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Error("Неудачный запрос, статус:", resp.StatusCode)
	}

	var jsonResponse QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		s.logger.Error(err)
	}

	quoteBody := jsonResponse.Quote.Body

	repo := s.notesRepo
	body := fmt.Sprintf("%s %s", note.Body, quoteBody)
	err = repo.CreateNewNote(note.Title, body)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}
