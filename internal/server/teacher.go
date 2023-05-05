package server

import "github.com/gin-gonic/gin"

func NewTeacherServer() *TeacherServer { return &TeacherServer{} }

type TeacherServer struct{}

func (srv *TeacherServer) GetTeacher(c *gin.Context) {}

func (srv *TeacherServer) ListTeacher(c *gin.Context) {}

func (srv *TeacherServer) CreateTeacher(c *gin.Context) {}

func (srv *TeacherServer) UpdateTeacher(c *gin.Context) {}

func (srv *TeacherServer) DeleteTeacher(c *gin.Context) {}
