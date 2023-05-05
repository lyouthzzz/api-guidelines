package server

import "github.com/gin-gonic/gin"

func NewClassServer() *ClassServer { return &ClassServer{} }

type ClassServer struct{}

func (srv *ClassServer) GetClass(c *gin.Context) {}

func (srv *ClassServer) ListClass(c *gin.Context) {}

func (srv *ClassServer) CreateClass(c *gin.Context) {}

func (srv *ClassServer) UpdateClass(c *gin.Context) {}

func (srv *ClassServer) DeleteClass(c *gin.Context) {}

// StudentJoinClass 学生加入班级
func (srv *ClassServer) StudentJoinClass(c *gin.Context) {}

// StudentExitClass 学生退出班级
func (srv *ClassServer) StudentExitClass(c *gin.Context) {}

func (srv *ClassServer) TeacherJoinClass(c *gin.Context) {}

func (srv *ClassServer) TeacherExitClass(c *gin.Context) {}
