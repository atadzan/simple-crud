package controller

import (
	"log"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/atadzan/simple-crud/pkg/repository/db"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
)

func (ctl *Controller) getBookGenres(c *fiber.Ctx) error {
	genres, err := ctl.repo.GetGenres(c.Context())
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}
	return c.Status(fiber.StatusOK).JSON(genres)
}

func (ctl *Controller) createBook(c *fiber.Ctx) error {
	authorId := getAuthorIDFromCtx(c)
	var input models.CreateBookParams
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("invalid input params"))
	}

	input.AuthorId = authorId
	if err := ctl.repo.Create(c.Context(), input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))

	}

	return c.Status(fiber.StatusOK).JSON(newMessage(successMsg))
}

func (ctl *Controller) getBooks(c *fiber.Ctx) error {
	paginationParams := getPaginationParams(c)
	filter := getFilterParams(c)

	books, err := ctl.repo.GetAll(c.Context(), models.BooksParams{
		Filter:        filter,
		Pagination:    paginationParams,
		CreatedAtSort: c.Query("sort", "ASC"),
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMessage(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (ctl *Controller) getBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("invalid id"))
	}
	book, err := ctl.repo.GetById(c.Context(), id)
	if err != nil {
		log.Println(err)
		switch {
		case errors.Is(err, db.ErrNotFound):
			return c.Status(fiber.StatusBadRequest).JSON(newMessage(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
		}
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func (ctl *Controller) searchBook(c *fiber.Ctx) error {
	searchWord := c.Query("searchWord")
	if len(searchWord) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("empty search word"))
	}
	books, err := ctl.repo.Search(c.Context(), models.SearchParams{
		SearchWord:    searchWord,
		CreatedAtSort: c.Query("sort", "ASC"),
		Pagination:    getPaginationParams(c),
		Filter:        getFilterParams(c),
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusOK).JSON(newMessage(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (ctl *Controller) updateBook(c *fiber.Ctx) error {
	var input models.UpdateBookParams
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("invalid input params"))
	}
	if err := ctl.repo.Update(c.Context(), input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(newMessage(successMsg))
}

func (ctl *Controller) deleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id == 0 {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("invalid input params"))
	}
	authorId := getAuthorIDFromCtx(c)
	if err = ctl.repo.Delete(c.Context(), uint32(id), authorId); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}
	return c.Status(fiber.StatusOK).JSON(newMessage(successMsg))
}

func (ctl *Controller) downloadBookIMG(c *fiber.Ctx) error {
	path := c.Path("path")
	if len(path) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("empty path"))
	}
	file, err := ctl.repo.GetFile(c.Context(), path)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).SendStream(file.Reader, int(file.Size))
}
