package auth

import (
		"github.com/gin-gonic/gin"

)

func (r Routes) Setup (){
	r.group.POST("/signup",r.handler.SignUp)
}