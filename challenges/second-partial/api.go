package main

import "github.com/gin-gonic/gin"
//import "fmt"
import "time"
import "strconv"

//data
var defUser ="manager"
var defPassword ="mypass123"


var User string

var password string

var token string

var loginBool int

var USERS [3]string


func main() {
	r := gin.Default()
	r.GET("/login/user/:usuario/password/:pass", func(c *gin.Context) {
		User:=c.Param("usuario")
		USERS[0]=User
		password:=c.Param("pass")
		loginBool=check(User, password)

		if loginBool ==1{
		
			token=giveToken(User,password)

			c.JSON(200, gin.H{
				
				"message": "Hi "+ User +", welcome to the DPIP System",
				"Token ": token,
			})

		}else{
			
			c.JSON(200, gin.H{
				"message": "User or password not valid",
				
			})
		}
		
	})
	r.GET("/logout/token/:token", func(c *gin.Context) {
		TokentoDelete:=c.Param("token")
		
		if(TokentoDelete==token){

			token=deleteToken(token)
			
			c.JSON(200, gin.H{
				"Message" : "Bye "+ USERS[0] + " your token has been revoked",
			})
		}else{
			c.JSON(200, gin.H{
			"Message" : "Invalid token",
			})
		}
		
		
	})
	r.POST("/upload/token/:tok", func(c *gin.Context) {
		
		file, _ := c.FormFile("file")
		qrytoken := c.Param("tok")
		//fmt.Println(qrytoken)

		size:=strconv.FormatInt(file.Size,10)

		if(qrytoken==token){

			c.JSON(200, gin.H{

			"Message": "An image has been succesfully uploaded",
			"filename": file.Filename,
			"size": size + " bytes",
			 
		})
		}else{

			c.JSON(200,gin.H{

				"Message": "Invalid token",

			})
		}
		
		
		
	})
	r.GET("/status/token/:token", func(c *gin.Context) {

		TokentoDelete:=c.Param("token")
		//fmt.Println(User)

		if(TokentoDelete==token){
	
			c.JSON(200, gin.H{

				"message": "Hi "+ USERS[0] + ", the DPIP is UP and Running",
				"time": time.Now().Format(time.RFC850),

			})
		}else{

			c.JSON(200, gin.H{

				"Message" : "Invalid token",

			})
		}



		
	})
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//Functions

//generates the token with the data
func giveToken (name string, pass string)string{

	first3 :=name[0:3]
	last3 := pass[len(pass)-3:]
	token = first3+last3
	
	return token
}

//delete the active token from the session
func deleteToken(tokentoDelete string)string{
	
	tokentoDelete = ""
	token = tokentoDelete
	return token
}

//verify that the data is correct
func check(user string, pass string) int{

	if(user==defUser&& pass==defPassword){
		return (1)
	}else{
		return (0)
	}
}
