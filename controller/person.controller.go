package controller

import (
	"eCommerce/model/dto"
	"eCommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	service service.PersonService
}

func NewPersonController(service service.PersonService) PersonController {
	return PersonController{
		service: service,
	}
}

// GetPersons             godoc
// @Summary      Get persons array
// @Description  Responds with the list of all books as JSON.
// @Tags         Persons
// @Produce      json
// @Success      200  {array}  dto.PersonGetDto
// @Router       /persons [get]
func (ctr *PersonController) GetAll(ctx *gin.Context) {
	response, err := ctr.service.GetAll()
	if err != nil {
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetPersonById             godoc
// @Summary      Get single person by id
// @Description  Returns the person whose ID value matches the id.
// @Tags         Persons
// @Produce      json
// @Param        id  path      string  true  "search person by id"
// @Success      200  {object}  dto.PersonGetDto
// @Failure      404  {object}  errs.ErrResponse
// @Router       /persons/{id} [get]
func (ctr *PersonController) GetPersonById(ctx *gin.Context) {
	id := ctx.Param("id")
	person, err := ctr.service.Get(&id)
	if err != nil {
		ctx.Abort()
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, person)
}

// CreatePerson             godoc
// @Summary      Create a new person
// @Description  Takes a person JSON and store in DB. Return saved ID.
// @Tags         Persons
// @Produce      json
// @Param        person  body      dto.PersonCreateDto  true  "PersonCreateDto JSON"
// @Success      201   {object}  string
// @Failure      404  {object}  errs.ErrResponse
// @Failure      400  {object}  errs.ErrResponse
// @Router       /persons [post]
func (ctr *PersonController) CreatePerson(ctx *gin.Context) {
	var dto dto.PersonCreateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.Abort()
		ctx.Error(err)
		return
	}

	id, err := ctr.service.Create(&dto)
	if err != nil {
		ctx.Abort()
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": id})
}

// UpdatePerson             godoc
// @Summary      Update an existed person
// @Description  Takes a person JSON and store in DB. Return No Content.
// @Tags         Persons
// @Produce      json
// @Param        person  body      dto.PersonUpdateDto  true  "PersonUpdateDto JSON"
// @Success      204
// @Failure      404  {object}  errs.ErrResponse
// @Failure      400  {object}  errs.ErrResponse
// @Router       /persons [put]
func (ctr *PersonController) UpdatePerson(ctx *gin.Context) {
	var id = ctx.Param("id")
	var dto dto.PersonUpdateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.Abort()
		ctx.Error(err)
		return
	}

	if err := ctr.service.Update(&id, &dto); err != nil {
		ctx.Abort()
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// DeletePerson             godoc
// @Summary      Delete an existed person
// @Description  Takes a person ID and removes it from the database. Return No Content.
// @Tags         Persons
// @Produce      json
// @Success      204
// @Failure      404  {object}  errs.ErrResponse
// @Failure      400  {object}  errs.ErrResponse
// @Router       /persons [delete]
func (ctr *PersonController) DeletePerson(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ctr.service.Delete(&id)
	if err != nil {
		ctx.Abort()
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (ctr *PersonController) RegisterRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/persons")
	route.GET("/:id", ctr.GetPersonById)
	route.GET("/", ctr.GetAll)
	route.POST("/", ctr.CreatePerson)
	route.PUT("/:id", ctr.UpdatePerson)
	route.DELETE("/:id", ctr.DeletePerson)
}
