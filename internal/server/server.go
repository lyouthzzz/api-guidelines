package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewClassServer, NewTeacherServer)

func registerClassServer(httpServer *gin.Engine, srv *ClassServer) {
	httpServer.GET("/v1/classes/:id", srv.GetClass)
	httpServer.GET("/v1/classes", srv.ListClass)
	httpServer.POST("/v1/classes", srv.CreateClass)
	httpServer.PATCH("/v1/classes/:id", srv.UpdateClass)
	httpServer.DELETE("/v1/classes/:id", srv.DeleteClass)
	httpServer.POST("/v1/classes/:id/students/:student_id/join", srv.StudentJoinClass)
	httpServer.POST("/v1/classes/:id/students/:student_id/exit", srv.StudentExitClass)
	httpServer.POST("/v1/classes/:id/teachers/:teacher_id/join", srv.TeacherJoinClass)
	httpServer.POST("/v1/classes/:id/teachers/:teacher_id/exit", srv.TeacherExitClass)
}

func registerTeacherServer(httpServer *gin.Engine, srv *TeacherServer) {
	httpServer.GET("/v1/teachers/:id", srv.GetTeacher)
	httpServer.GET("/v1/teachers", srv.ListTeacher)
	httpServer.POST("/v1/teachers", srv.CreateTeacher)
	httpServer.PATCH("/v1/teachers/:id", srv.UpdateTeacher)
	httpServer.DELETE("/v1/teachers/:id", srv.DeleteTeacher)
}

func RegisterServer(httpServer *gin.Engine, teacherServer *TeacherServer, classServer *ClassServer) {
	registerClassServer(httpServer, classServer)
	registerTeacherServer(httpServer, teacherServer)
}
