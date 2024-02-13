package server

import (
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/server/routes"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
)

// @title SAC API
// @version 1.0
// @description Backend Server for SAC App

// @contact.name	David Oduneye and Garrett Ladley
// @contact.email	oduneye.d@northeastern.edu and ladley.g@northeastern.edu
// @host 127.0.0.1:8080
// @BasePath /
func Init(db *gorm.DB, settings config.Settings) *fiber.App {
	app := newFiberApp()

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(err)
	}

	middlewareService := middleware.NewMiddlewareService(db, validate, settings.Auth)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(middlewareService.Authenticate)


	eventRoutes(apiv1, &services.EventService{DB: db, Validate: validate})
	routes.Utility(app)

	routes.Auth(apiv1, services.NewAuthService(db, validate), settings.Auth)

	userRouter := routes.User(apiv1, services.NewUserService(db, validate), middlewareService)
	routes.UserTag(userRouter, services.NewUserTagService(db, validate))
	routes.UserFollower(userRouter, services.NewUserFollowerService(db, validate))
	routes.UserMember(userRouter, services.NewUserMemberService(db))

	routes.Contact(apiv1, services.NewContactService(db, validate))

	clubsIDRouter := routes.Club(apiv1, services.NewClubService(db, validate), middlewareService)
	routes.ClubTag(clubsIDRouter, services.NewClubTagService(db, validate))
	routes.ClubFollower(clubsIDRouter, services.NewClubFollowerService(db))
	routes.ClubMember(clubsIDRouter, services.NewClubMemberService(db, validate))
	routes.ClubContact(clubsIDRouter, services.NewClubContactService(db, validate))

	routes.Tag(apiv1, services.NewTagService(db, validate))

	categoryRouter := routes.Category(apiv1, services.NewCategoryService(db, validate))
	routes.CategoryTag(categoryRouter, services.NewCategoryTagService(db, validate))

	return app
}

func newFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	return app
}


	//events := clubs.Group("/:id/events")
	//events.Get("/", clubController.GetClubEvents)



func eventRoutes(router fiber.Router, eventService services.EventServiceInterface) {
	eventController := controllers.NewEventController(eventService)

	events := router.Group("/events")

	events.Get("/:id", eventController.GetEvent)
	events.Get("/:id/series", eventController.GetSeriesByEventId)
	events.Get("/", eventController.GetAllEvents)
	events.Post("/", eventController.CreateEvent)
	events.Patch("/:id", eventController.UpdateEvent)
	events.Delete("/:id", eventController.DeleteEvent)
	events.Delete("/:id/series", eventController.DeleteEventSeries) 
}

