package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/igorgofman/gms-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIndentity)
	{
		info := api.Group("/info")
		{
			info.POST("/", h.createInfo)
			info.GET("/", h.getAllInfos)
			info.GET("/instructors", h.getAllInstructors)
			info.GET("/members", h.getAllMembers)
			info.GET("/:id", h.getInfoById)
			info.PUT("/:id", h.updateInfo)
			info.DELETE("/:id", h.deleteInfo)

			visits := info.Group(":id/visits")
			{
				visits.POST("/", h.createVisit)
				visits.GET("/", h.getAllVisitsById)
			}
		}

		memberships := api.Group("/memberships")
		{
			memberships.POST("/", h.createMembership)
			memberships.GET("/", h.getAllMemberships)
			memberships.GET("/:id", h.getMembershipById)
			memberships.PUT("/:id", h.updateMembership)
			memberships.DELETE("/:id", h.deleteMembership)
		}

		visits := api.Group("/visits")
		{
			visits.GET("/", h.getAllVisits)
			visits.GET("/:id", h.getVisitById)
			visits.PUT("/:id", h.updateVisit)
			visits.DELETE("/:id", h.deleteVisit)
		}

	}

	return router
}
