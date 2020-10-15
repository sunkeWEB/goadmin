package article

import (
	"github.com/gin-gonic/gin"
	"go-admin/tools/app"
	"net/http"
)

func GetArticleList(c *gin.Context) {
	var res app.Response
	res.Data = "hello world"
	c.JSON(http.StatusOK, res.ReturnOK())
}
