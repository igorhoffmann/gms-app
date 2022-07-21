package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gym "github.com/igorgofman/gms-app"
)

func (h *Handler) createInfo(c *gin.Context) {
	var (
		input_info       gym.Info
		input_member     gym.Member
		input_instructor gym.Instructor
	)

	if err := c.BindJSON(&input_info); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Info.Create(input_info, input_member, input_instructor)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllInfosResponse struct {
	Data []gym.DataToPrintInfo `json:"data"`
}

type getAllInstructorsResponse struct {
	Data []gym.DataToPrintInstructor `json:"data"`
}

type getAllMembersResponse struct {
	Data []gym.DataToPrintMember `json:"data"`
}

func (h *Handler) getAllInfos(c *gin.Context) {
	infos, err := h.services.Info.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllInfosResponse{
		Data: infos,
	})
}

func (h *Handler) getAllInstructors(c *gin.Context) {
	instructors, err := h.services.Info.GetAllInstructors()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllInstructorsResponse{
		Data: instructors,
	})
}

func (h *Handler) getAllMembers(c *gin.Context) {
	members, err := h.services.Info.GetAllMembers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllMembersResponse{
		Data: members,
	})
}

func (h *Handler) getInfoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	info, err := h.services.Info.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, info)
}

func (h *Handler) updateInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input gym.UpdateInfoInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Info.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Info.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
