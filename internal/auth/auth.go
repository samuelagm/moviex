package auth

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Store struct {
	mu     sync.RWMutex
	users  map[string]string
	tokens map[string]string
}

func NewStore() *Store {
	return &Store{
		users:  make(map[string]string),
		tokens: make(map[string]string),
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (s *Store) Register(gctx *gin.Context) {
	var req RegisterRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[req.Email]; exists {
		gctx.JSON(http.StatusConflict, gin.H{"message": "email already registered"})
		return
	}

	s.users[req.Email] = req.Password
	token := uuid.NewString()
	s.tokens[token] = req.Email

	gctx.JSON(http.StatusCreated, TokenResponse{Token: token})
}

func (s *Store) Login(gctx *gin.Context) {
	var req LoginRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	s.mu.RLock()
	password, exists := s.users[req.Email]
	s.mu.RUnlock()

	if !exists || password != req.Password {
		gctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	token := uuid.NewString()

	s.mu.Lock()
	s.tokens[token] = req.Email
	s.mu.Unlock()

	gctx.JSON(http.StatusOK, TokenResponse{Token: token})
}

func (s *Store) Middleware() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		header := gctx.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			gctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "authorization required"})
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		s.mu.RLock()
		email, valid := s.tokens[token]
		s.mu.RUnlock()

		if !valid {
			gctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired token"})
			return
		}

		gctx.Set("user_email", email)
		gctx.Next()
	}
}
