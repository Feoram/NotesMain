package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Service) CreateNote(c echo.Context) error {
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

	note := new(Note)
	err = c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.notesRepo
	body := quoteBody + " " + note.Body
	err = repo.CreateNewNote(note.Title, body)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}
