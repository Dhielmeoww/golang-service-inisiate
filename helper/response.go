package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type Meta struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CoreMeta struct {
	Data any `json:"data"`
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"meta"`
}

// func Response(ctx *gin.Context, message string, status int, data interface{}) Meta {

// 	var username string = "-"
// 	getusername, err := ctx.Get("username")

// 	if err {
// 		username = getusername.(string)
// 	}

// 	var bo interface{}

// 	ctx.ShouldBindJSON(&bo)

// 	fmt.Println(bo)

// 	jsonData, _ := io.ReadAll(ctx.Request.Body)

// 	jsonMap := make(map[string]interface{})
// 	json.Unmarshal(jsonData, &jsonMap)
// 	fmt.Print("body:", jsonMap)

// 	errors := database.InitialDB().DB.Debug().Table("tr_logs").Create(&core_entity.Log{
// 		CreatedBy:  username,
// 		Endpoint:   ctx.Request.RequestURI,
// 		StatusCode: status,
// 		Payload:    string(jsonData),
// 		Message:    message,
// 		IP:         ctx.ClientIP(),
// 		CreatedAt:  time.Now(),
// 	}).Error

// 	fmt.Println(errors)

// 	return Meta{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}

// }

func ResponsePaginate(message string, status int, request *http.Request, model *gorm.DB, data any) Meta {
	pg := paginate.New()

	return Meta{
		Status:  status,
		Message: message,
		Data:    pg.With(model).Request(request).Response(data),
	}
}

func SuccessJSON(c *gin.Context, message string, statusCode int, data interface{}) gin.H {
	return gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	}
}

func ErrorJSON(c *gin.Context, message string, statusCode int, data interface{}) gin.H {
	return gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	}
}
