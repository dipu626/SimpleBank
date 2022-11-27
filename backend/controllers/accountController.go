package controllers

import "github.com/gin-gonic/gin"

func GetAccounts() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func SingUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return false, ""
}
