package router

import (
	"yoga/clean/app/middlewares"

	userdata "yoga/clean/features/user/data"
	userhandler "yoga/clean/features/user/handler"
	userservice "yoga/clean/features/user/service"

	projectdata "yoga/clean/features/project/data"
	projecthandler "yoga/clean/features/project/handler"
	projectservice "yoga/clean/features/project/service"

	taskdata "yoga/clean/features/task/data"
	taskhandler "yoga/clean/features/task/handler"
	taskservice "yoga/clean/features/task/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	userData := userdata.New(db)
	userService := userservice.New(userData)
	userHandlerAPI := userhandler.New(userService)

	projectData := projectdata.New(db)
	projectService := projectservice.New(projectData)
	projectHandlerAPI := projecthandler.New(projectService)

	taskData := taskdata.New(db)
	taskService := taskservice.New(taskData)
	taskHandlerAPI := taskhandler.New(taskService)

	//e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users", userHandlerAPI.FindUserByID, middlewares.JWTMiddleware())
	e.POST("/users", userHandlerAPI.CreateUser)
	e.DELETE("/users", userHandlerAPI.DeleteUserByID, middlewares.JWTMiddleware())
	e.POST("/login", userHandlerAPI.Login)
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	//e.GET("/projects", projectHandlerAPI.GetAllProject, middlewares.JWTMiddleware())
	e.GET("/projects/:project_id", projectHandlerAPI.GetProjectByID, middlewares.JWTMiddleware())
	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.PUT("/projects/:project_id", projectHandlerAPI.UpdateProjectByID, middlewares.JWTMiddleware())
	e.DELETE("/projects/:project_id", projectHandlerAPI.DeleteProjectByID, middlewares.JWTMiddleware())

	//e.GET("/tasks", taskHandlerAPI.GetAllTask, middlewares.JWTMiddleware())
	//e.GET("/tasks/:task_id", taskHandlerAPI.GetTaskByID, middlewares.JWTMiddleware())
	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:task_id", taskHandlerAPI.UpdateTaskByID, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:task_id", taskHandlerAPI.DeleteTaskByID, middlewares.JWTMiddleware())
}
