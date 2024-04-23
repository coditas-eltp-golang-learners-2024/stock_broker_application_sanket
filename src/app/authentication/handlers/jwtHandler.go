package handlers

import (
	"authentication/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

// @Summary JWT Authentication Middleware
// @Description Authenticate incoming requests using JSON Web Tokens (JWT)
// @ID jwt-auth-middleware
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth ApiKeyAuth
// @In header
// @Name Authorization
// @Description JWT authorization header
// @Success 200 {object} string "Authentication successful"
// @Failure 401 {object} string "Unauthorized: missing or invalid token"
// @Router /customer-changepassword [patch]
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.ErrGeneric.Error(): constants.ErrMissingToken.Error()})
			ctx.Abort()
			return
		}
		// Check if the token starts with "Bearer "
		const bearerPrefix = "Bearer "
		if strings.HasPrefix(tokenString, bearerPrefix) {
			// Remove the "Bearer " prefix
			tokenString = tokenString[len(bearerPrefix):]
		}
		// Trim any leading or trailing whitespace from the token string
		tokenString = strings.TrimSpace(tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return constants.SecretKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.ErrGeneric.Error(): constants.ErrPrasingToken.Error()})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
			ctx.Set("email", claims.Email)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.ErrGeneric.Error(): constants.ErrInvalidToken.Error()})
			ctx.Abort()
			return
		}

	}
}
