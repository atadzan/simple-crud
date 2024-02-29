package controller

import (
	"crypto/sha512"
	"encoding/hex"
	"log"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/atadzan/simple-crud/pkg/repository/db"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
)

func generatePasswordHash(password *string, passwordSalt string) error {

	// convert password string to slice of bytes
	passwordBytes := []byte(*password)

	// creating sha-512 header
	sha512Header := sha512.New()

	// convert passwordSalt string to slice of bytes
	saltBytes := []byte(passwordSalt)

	// Append passwordSalt to password
	passwordBytes = append(passwordBytes, saltBytes...)

	// write password to sha-512 header
	if _, err := sha512Header.Write(passwordBytes); err != nil {
		return errors.New(err)
	}

	// get sha-512 hashed password
	hashedPassword := sha512Header.Sum(nil)

	// convert hashed password to HEX string
	*password = hex.EncodeToString(hashedPassword)

	return nil

}

func (ctl *Controller) register(c *fiber.Ctx) error {
	var input models.AuthParams

	// receiving input params from request
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusOK).JSON(newMessage(err.Error()))
	}

	// validate input params
	if err := input.Validate(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMessage(err.Error()))
	}

	// generate password hash for security
	if err := generatePasswordHash(&input.Password, ctl.authConfig.PasswordHashSalt); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}

	// pass input values to repository layer for further operations
	if err := ctl.repo.Register(c.Context(), input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}

	// if everything is OK, return "Success" message with 200 HTTP status
	return c.Status(fiber.StatusOK).JSON(newMessage(successMsg))
}

func (ctl *Controller) signIn(c *fiber.Ctx) error {
	var input models.AuthParams

	// receiving input params from request
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMessage(err.Error()))
	}

	// validate input params
	if err := input.Validate(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMessage(err.Error()))
	}

	// generate password hash for security
	if err := generatePasswordHash(&input.Password, ctl.authConfig.PasswordHashSalt); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
	}

	// pass input values to repository layer for further operations
	authorId, err := ctl.repo.GetAuthorId(c.Context(), input)
	if err != nil {
		log.Println(err)
		switch {
		case errors.Is(err, db.ErrNotFound):
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(newMessage(errInternalServerMsg))
		}
	}
	accessToken := generateAccessToken(authorId, ctl.authConfig.JWTSigningKey)
	// if everything is OK, return access token with 200 HTTP status
	return c.Status(fiber.StatusOK).JSON(newToken(accessToken))
}
