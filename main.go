package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

func main() {
	r := gin.Default()
	os.Setenv("PORT", "9090")
	userRepository := NewUserRepository()
	userController := NewUserController(userRepository)
	userRoute := NewUserRoute(userController)
	allUserEndpoint := userRoute.RegisterUserRoute(r)
	allEnpointsForSwag := combine(allUserEndpoint)
	swagRoute := NewSwagRoute(r)
	swagRoute.RegisterRoutes(allEnpointsForSwag)
	r.Run(":9090")
}

// SWAG ROUTE
type SwagRoute struct {
	router *gin.Engine
}

func NewSwagRoute(router *gin.Engine) *SwagRoute {
	return &SwagRoute{router: router}
}

func (s *SwagRoute) RegisterRoutes(endpoints []*swagger.Endpoint) {
	api := swag.New(
		swag.Endpoints(endpoints...),
		swag.Description("THis is the test description"),
		swag.Version("1.0.0"),
		swag.Title("Test Title"),
	)
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = swag.ColonPath(path)

		s.router.Handle(endpoint.Method, path, h)
	})

	enableCors := true
	s.router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))
	s.router.GET("/docs", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		scheme := "http://"
		if ctx.Request.TLS != nil {
			scheme = "https://"
		}
		content := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Scalar API Reference</title>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1" />
		</head>
		<body>
			<!-- Need a Custom Header? Check out this example https://codepen.io/scalarorg/pen/VwOXqam -->
			<script
			id="api-reference"
			type="application/json"
			data-url="%s"
			></script>
			<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
		</body>
		</html>
		`, scheme+ctx.Request.Host+"/swagger")
		ctx.String(http.StatusOK, content)
	})
}

func combine(endpoints []*swagger.Endpoint) []*swagger.Endpoint {
	return append(endpoints, endpoints...)
}

// ROUTE LAYER
type UserRoute struct {
	userController *UserController
}

func NewUserRoute(userController *UserController) *UserRoute {
	return &UserRoute{userController: userController}
}

func (u *UserRoute) RegisterUserRoute(r *gin.Engine) []*swagger.Endpoint {
	endpoints := []*swagger.Endpoint{}

	getAllUsers := endpoint.New(
		http.MethodGet,
		"/users",
		"Get all users",
		endpoint.Handler(u.userController.getUsers),
		endpoint.Response(http.StatusOK, []User{}, "Get all users Response"),
	)
	getUserByID := endpoint.New(
		http.MethodGet,
		"/users/:id",
		"Get User By ID",
		endpoint.Path("id", "string", "User ID", true),
		endpoint.Handler(u.userController.getUserByID),
		endpoint.Response(http.StatusOK, User{}, "Get User By ID Response"),
	)
	addUser := endpoint.New(
		http.MethodPost,
		"/users",
		"Add User",
		endpoint.Handler(u.userController.addUser),
		endpoint.Body(User{}, "Add User Request Payload", true),
		endpoint.Response(http.StatusCreated, gin.H{}, "Add User Response"),
	)
	updateUserByID := endpoint.New(
		http.MethodPut,
		"/users/:id",
		"Update User By ID",
		endpoint.Path("id", "string", "User ID", true),
		endpoint.Handler(u.userController.updateUser),
		endpoint.Response(http.StatusOK, User{}, "Update User By ID Response"),
	)
	deleteUserByID := endpoint.New(
		http.MethodDelete,
		"/users/:id",
		"Delete User By ID",
		endpoint.Path("id", "string", "User ID", true),
		endpoint.Handler(u.userController.deleteUser),
		endpoint.Response(http.StatusOK, gin.H{}, "Delete User By ID Response"),
	)

	endpoints = append(endpoints, getAllUsers, getUserByID, addUser, updateUserByID, deleteUserByID)
	return endpoints
}

// MODEL LAYER
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CONTROLLER LAYER
type UserController struct {
	router         *gin.Engine
	userRepository *UserRepository
}

func NewUserController(userRepository *UserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}

func (h *UserController) getUsers(c *gin.Context) {
	users := h.userRepository.GetUsers()
	c.JSON(200, users)
}

func (h *UserController) getUserByID(c *gin.Context) {
	id := c.Param("id")
	user := h.userRepository.GetUserByID(id)
	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func (h *UserController) addUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.userRepository.AddUser(&user)
	c.JSON(201, user)
}

func (h *UserController) updateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.userRepository.UpdateUser(&user)
	c.JSON(200, user)
}

func (h *UserController) deleteUser(c *gin.Context) {
	id := c.Param("id")
	h.userRepository.DeleteUser(id)
	c.JSON(204, gin.H{
		"message": "User deleted",
	})
}

// REPOSTITORY LAYER
type UserRepository struct {
	users []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []User{
			{ID: "1", Name: "Alice", Age: 25},
			{ID: "2", Name: "Bob", Age: 30},
		},
	}
}

func (r *UserRepository) GetUsers() []User {
	return r.users
}

func (r *UserRepository) GetUserByID(id string) *User {
	for _, user := range r.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (h *UserRepository) AddUser(user *User) {
	h.users = append(h.users, *user)
}

func (h *UserRepository) UpdateUser(user *User) {
	for i, u := range h.users {
		if u.ID == user.ID {
			h.users[i] = *user
			return
		}
	}
}

func (h *UserRepository) DeleteUser(id string) {
	for i, user := range h.users {
		if user.ID == id {
			h.users = append(h.users[:i], h.users[i+1:]...)
			return
		}
	}
}
