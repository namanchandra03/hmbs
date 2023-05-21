package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"hms/utils"
	"net/http"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		data := c.Request.Header.Get("Authorization")
		fmt.Printf("%v, %v ", data)
		if headerLength := len(data); headerLength == 0 {
			fmt.Fprintf(c.Writer, "auth token is not detected")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Abort()
			return
		}
		authToken := c.Request.Header["Authorization"][0]
		err := utils.ValidateToken(authToken)
		fmt.Println(err)
		if err != nil {
			fmt.Fprintf(c.Writer, "Token is not valid,login again")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Abort()
			return
		}
		userId := utils.SessionData.UserId
		sessionId := utils.SessionData.SessionId
		//currentSessionInfo := models.SessionBody{}
		//db := database.DB
		//tx := db.Model(&models.Session{}).Where("id = ? and end_time is not null", sessionId).Find(&currentSessionInfo)
		//if error := tx.Error; error != nil {
		//	c.Writer.WriteHeader(http.StatusInternalServerError)
		//	fmt.Fprintf(c.Writer, "something went wrong,please try again")
		//	return
		//}
		//if currentSessionInfo.UserId != 0 {
		//	c.Writer.WriteHeader(http.StatusForbidden)
		//	fmt.Fprintf(c.Writer, "Your session has ended, please login again")
		//	c.Abort()
		//	return
		//}
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "UserId", userId)
		ctx = context.WithValue(ctx, "SessionId", sessionId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
