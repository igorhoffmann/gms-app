package handler

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	gym "github.com/igorgofman/GMS-app"
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

	fmt.Printf("R: %v\n", input_info.Relationship)
	fmt.Printf("S: %v\n", input_info.Instructor.Salary)
	fmt.Printf("M: %v\n", input_info.Member.MembershipId)
	// fmt.Printf("1: %v\n", input_instructor.Salary)

	id, err := h.services.Info.Create(input_info, input_member, input_instructor)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// type getAllListsResponse struct {
// 	Data []gym.Info `json:"data"`
// }

func (h *Handler) getAllInfos(c *gin.Context) {
	// infoId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }

	// lists, err := h.services.Info.GetAll(infoId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, getAllListsResponse{
	// 	Data: lists,
	// })
}

func (h *Handler) getInfoById(c *gin.Context) {
	// infoId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	// 	return
	// }

	// list, err := h.services.Info.GetById(infoId, id)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, list)
}

func (h *Handler) updateInfo(c *gin.Context) {
	// infoId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	// 	return
	// }

	// var input gym.UpdateListInput
	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if err := h.services.Info.Update(infoId, id, input); err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, statusResponse{
	// 	Status: "ok",
	// })
}

func (h *Handler) deleteInfo(c *gin.Context) {
	// infoId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	// 	return
	// }

	// err = h.services.Info.Delete(infoId, id)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, statusResponse{
	// 	Status: "ok",
	// })
}
