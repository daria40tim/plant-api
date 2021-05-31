package handler

import (
	"net/http"

	plantapi "github.com/daria40tim/plant-api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCard(c *gin.Context) {
	var input plantapi.Card
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Card.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"c_id": id,
	})

}

type getAllCardsResponse struct {
	Data []plantapi.CardAll `json:"data"`
}

func (h *Handler) getAllCards(c *gin.Context) {
	cards, err := h.services.Card.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCardsResponse{
		Data: cards,
	})

}

func (h *Handler) getCardById(c *gin.Context) {

}

func (h *Handler) updateCardById(c *gin.Context) {

}
