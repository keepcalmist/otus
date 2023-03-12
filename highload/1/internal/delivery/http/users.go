package http

import (
	"github.com/gin-gonic/gin"
	"github.com/keepcalmist/otus/highload/1/internal/manager/users"
	"log"
	"net/http"
)

type usersRouter struct {
	manager users.UserService
}

func NewUsersHandler(usersManager users.UserService) *usersRouter {
	return &usersRouter{manager: usersManager}
}

func (h *usersRouter) Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := context.GetRawData()
		if err != nil {
			context.JSON(http.StatusBadRequest, Error{Error: err.Error()})
			return
		}
		req, err := getBody[UserRegisterReq](data)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusBadRequest, "Невалидные данные")
			return
		}

		resp, err := h.manager.Register(context, &users.UserRegisterReq{
			FirstName:  req.FirstName,
			SecondName: req.SecondName,
			Age:        req.Age,
			Biography:  req.Biography,
			City:       req.City,
			Password:   []byte(req.Password),
		})
		if err != nil {
			context.JSON(http.StatusInternalServerError, Error{Error: err.Error()})
			return
		}
		context.JSON(http.StatusOK, &UserRegisterResp{UserID: resp.UserID})
		return
	}
}

func (h *usersRouter) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r GetByIDReq
		if err := c.BindUri(&r); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		if err := r.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := h.manager.GetByID(c, r.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
		return
	}
}

func (h *usersRouter) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{Error: err.Error()})
			return
		}
		req, err := getBody[Login](data)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, "Невалидные данные")
			return
		}

		token, err := h.manager.Login(c, &users.Login{
			ID:       req.ID,
			Password: []byte(req.Password),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, LoginResp{Token: token})
	}
}
