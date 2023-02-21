package route

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github/malekradhouane/test-cdi/errs"
	"github/malekradhouane/test-cdi/service"
	. "github/malekradhouane/test-cdi/store"
	"net/http"
)

//AstronautActions represents astronauts controller actions
type AstronautActions struct {
	astronautService *service.AstronautService
}

//NewAstronautActions constructor
func NewAstronautActions(us *service.AstronautService) *AstronautActions {
	return &AstronautActions{
		astronautService: us,
	}
}

//Create _
func (ua AstronautActions) Create(c *gin.Context) {
	req := new(Astronaut)
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = ua.astronautService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "astronaut added"})

}

//Get _
func (ua AstronautActions) Get(c *gin.Context) {
	id := c.Param("id")
	astronaut, err := ua.astronautService.Get(c.Request.Context(), id)
	if err != nil {
		if errs.IsNoSuchEntityError(err) {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	c.JSON(http.StatusOK, astronaut)
}

//List lists all astronaut
func (ua AstronautActions) List(c *gin.Context) {
	astronauts, err := ua.astronautService.List(c.Request.Context())
	if err != nil {
		if errs.IsNoSuchEntityError(err) {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	c.JSON(http.StatusOK, astronauts)
}

//Update _
func (ua AstronautActions) Update(c *gin.Context) {
	id := c.Param("id")
	req := new(Astronaut)
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = ua.astronautService.Update(c.Request.Context(), req, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "astronaut updated"})

}

//Delete _
func (ua AstronautActions) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ua.astronautService.Delete(c.Request.Context(), id)
	if err != nil {
		if err == errs.ErrNoSuchEntity {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("no entity with id : %s", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "astronaut deleted with success"})
}
