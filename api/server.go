package api

type Server struct {
 DB  *gorm.DB
 Gin *gin.Engine
}