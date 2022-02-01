package controller

import (
	"time"

	"github.com/DasoTD/fiber-jwt/data"
	"github.com/DasoTD/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWTToken(user data.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
func SignUp(c *fiber.Ctx) error {
	post := new(models.SignupRequest)
	if err := c.BodyParser(post); err != nil {
		return err
	}
	if post.Email == "" || post.Name == "" || post.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "all fields are required")
	}
	//save the info in the database
	hash, err := bcrypt.GenerateFromPassword([]byte(post.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &data.User{
		Name:     post.Name,
		Email:    post.Email,
		Password: string(hash),
	}
	engine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	_, err = engine.Insert(user)
	if err != nil {
		return err
	}
	token, exp, err := CreateJWTToken(*user)
	if err != nil {
		return err
	}
	//create a jwt token
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func Login(c *fiber.Ctx) error {
	var post models.LoginRequest
	if err := c.BodyParser(post); err != nil {
		return err
	}
	if post.Email == "" || post.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "all fields are required")
	}
	engine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}
	user := new(data.User)
	has, err := engine.Where("email = ?", post.Email).Desc("id").Get(&user)
	if err != nil {
		return err
	}
	if !has {
		return fiber.NewError(fiber.StatusNotFound, "email not found")
	}

	token, exp, err := CreateJWTToken(*user)
	if err != nil {
		return err
	}
	//create a jwt token
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})

}
