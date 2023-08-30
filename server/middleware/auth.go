package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print(c.Request.Header["Email"])
		if len(c.Request.Header["Email"]) == 0 || len(c.Request.Header["Password"]) == 0 {
			fmt.Print("aborting")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": false, "message": "username or password missing@"})
			return
		}

		email := c.Request.Header["Email"][0]

		password := c.Request.Header["Password"][0]

		// before request

		supabaseUrl := os.Getenv("SUPABASE_URL")

		supabaseAnonKey := os.Getenv("SUPABASE_KEY")

		supabase := supa.CreateClient(supabaseUrl, supabaseAnonKey)

		ctx := context.Background()
		// ctx, _ = context.WithTimeout(ctx, 60000*time.Millisecond)
		//fmt.Println(ctx)

		//user, err := supabase.Auth.User(ctx, token)

		user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
			Email:    email,
			Password: password,
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": false, "message": err.Error()})
			return
			//fmt.Printf(err.Error())

		}
		c.Request.Header.Add("x-request-id", user.AccessToken)

		c.Next()

		// after request

	}
}
