package handler

import (
	"net/http"
	"strconv"

	plantapi "github.com/daria40tim/plant-api"
	"github.com/gin-gonic/gin"
)

type getAllDirsResponse struct {
	Data []plantapi.Dir `json:"data"`
}

func (h *Handler) getDirs(c *gin.Context) {

	dirs, err := h.services.OCR.GetDirs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllDirsResponse{
		Data: dirs,
	})
}

type Index struct {
	I   string `json:"index"`
	Dir string `json:"dir"`
}

func (h *Handler) getImg(c *gin.Context) {
	var index Index
	if err := c.BindJSON(&index); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	i, err := strconv.Atoi(index.I)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	card, err := h.services.OCR.GetImg(i, index.Dir)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, card)

}
