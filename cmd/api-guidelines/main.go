package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/api-guidelines/internal/server"
	"github.com/pkg/errors"
)

func main() {
	engine := gin.New()

	classServer := server.NewClassServer()
	teacherServer := server.NewTeacherServer()
	server.RegisterServer(engine, teacherServer, classServer)

	panic(errors.WithStack(engine.Run(":18080")))
}
