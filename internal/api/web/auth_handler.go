package web

import (
	"net/http"
	"strconv"

	"wtg_desktop/internal/domain/auth"
	"wtg_desktop/internal/middleware"
	token "wtg_desktop/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service auth.Service
}

func NewAuthHandler(service auth.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth/users")
	authGroup.POST("/", middleware.AuthRequired(), middleware.RequireRoles("ADMIN"), h.CreateUser)
	authGroup.GET("/", h.GetAllUsers)
	authGroup.GET("/:id", h.GetUserByID)
	authGroup.PUT("/:id", h.UpdateUser)
	authGroup.DELETE("/:id", h.DeleteUser)
	// authGroup.POST("/:id/roles", h.AssignRole)
	authGroup.POST("/refresh", h.RefreshToken)
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		RoleID   int    `json:"role_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &auth.AuthUser{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := h.service.CreateUserWithRole(user, input.Password, input.RoleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Identifier string `json:"identifier" binding:"required"`
		Password   string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, roles, err := h.service.Authenticate(input.Identifier, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := token.GenerateTokens(user.ID, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := token.ParseToken(input.RefreshToken, true)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	accessToken, refreshToken, err := token.GenerateTokens(claims.UserID, claims.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not refresh tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// // POST /auth/users/:id/roles
// func (h *AuthHandler) AssignRole(c *gin.Context) {
// 	idStr := c.Param("id")
// 	userID, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 		return
// 	}

// 	var input struct {
// 		RoleID int `json:"role_id" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.service.AssignRole(userID, input.RoleID); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// GET /auth/users/:id
func (h *AuthHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GET /auth/users
func (h *AuthHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// PUT /auth/users/:id
func (h *AuthHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user auth.AuthUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DELETE /auth/users/:id
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := h.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
