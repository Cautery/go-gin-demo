package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Cautery/go-gin-demo/models"
	"github.com/gin-gonic/gin"
)

var tempArticleList []models.Article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func saveList() {
	tempArticleList = models.ArticleList
}

func restoreList() {
	models.ArticleList = tempArticleList
}
