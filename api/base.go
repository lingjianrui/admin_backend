package api

import (
	"backend/middleware"
	"backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type Context struct {
	*gin.Context
}

//Server 服务模型
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

//Ping 连通性测试
func (server *Server) Ping(c *gin.Context) {
	//nameParam := c.PostForm("name")
	/* middleware.Logger().WithFields(logrus.Fields{*/
	//"name": "hanyun",
	/*}).Info("记录一下日志", "Info")*/
	middleware.Logger().Info("记录一下日志", "Info")

	res := make(map[string]interface{})
	res["status"] = "pong"
	c.JSON(http.StatusOK, res)
}

type File struct {
	Fl string `json:"file"`
}

func (server *Server) ImagePost(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		//ignore
	}
	uid := uuid.NewV4()
	dst := "upload/" + uid.String()
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {
		// ignore
	}
	res := make(map[string]interface{})
	res["fileuid"] = uid.String()
	c.JSON(http.StatusOK, res)
}

func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}

//Initialize 初始化数据库
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		fmt.Println(DBURL)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		server.DB.SingularTable(true)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	} else {
		fmt.Println("Unknown Driver")
	}
	//数据库初始化修改
	server.DB.Debug().AutoMigrate(
		&model.Vendor{},
		&model.User{},
	)
	server.Router = gin.Default()
	server.initializeRoutes()
}

//Run 系统运行入口方法
func (s *Server) Run(addr string) {
	fmt.Println("service is runing")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
