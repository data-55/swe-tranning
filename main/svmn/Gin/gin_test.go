package Gin

import (
	"encoding/json"
	"forbizbe/main/model"
	"forbizbe/main/svmn/Mysql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testRequest struct {
	method      string
	uri         string
	queryString string
}

var r = testSetup()

func testSetup() *gin.Engine {
	Mysql.ConnDB()
	r := SetupRouter()

	return r
}

func testHttpRequest(r *gin.Engine, testRequest testRequest) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, err := http.NewRequest(testRequest.method, testRequest.uri+testRequest.queryString, nil)
	assert.NoError(&testing.T{}, err, "Request fail")

	r.ServeHTTP(res, req)

	return res
}
func dataFetch(user *model.User, posts *[]model.Post, userId string) {
	Mysql.DB.Order("created_at DESC").Find(&posts, "user_id = ?", userId)
	Mysql.DB.Find(&user, userId)
}

func homeGetUserId3(t *testing.T) {
	var posts []model.Post
	var user model.User
	testRequest := testRequest{
		method:      http.MethodGet,
		uri:         "/post/home",
		queryString: "?userId=3",
	}

	res := testHttpRequest(r, testRequest)
	assert.Equal(t, 200, res.Code, "Error - userId=3 : Response Code")

	inx := strings.Index(testRequest.queryString, "=") + 1
	userId := testRequest.queryString[inx:]

	dataFetch(&user, &posts, userId)

	userIdUint, _ := strconv.ParseUint(userId, 10, 0)
	testPosts := make([]model.PostRe, userIdUint)
	for i := range testPosts {
		testPosts[i].ID = 6 - uint(i)
		testPosts[i].Comment = userId + "comment"
		testPosts[i].UpdatedAt = posts[i].UpdatedAt.Format(model.TIME_FORMAT)
		testPosts[i].CreatedAt = posts[i].CreatedAt.Format(model.TIME_FORMAT)
		testPosts[i].Writer.ID = uint(userIdUint)
		testPosts[i].Writer.Name = "3_name"
		testPosts[i].Writer.Email = "3_email@test.com"
		testPosts[i].Writer.UpdatedAt = user.UpdatedAt.Format(model.TIME_FORMAT)
		testPosts[i].Writer.CreatedAt = user.CreatedAt.Format(model.TIME_FORMAT)
	}

	testRes := model.Response{
		Code: "OK",
		Data: testPosts,
	}

	jsonTestRes, _ := json.Marshal(testRes)

	assert.Equal(t, string(jsonTestRes), res.Body.String(), "Error - userId=3 : Response Body")
}

func homeGetUserId100(t *testing.T) {
	testRequest := testRequest{
		method:      http.MethodGet,
		uri:         "/post/home",
		queryString: "?userId=100",
	}

	res := testHttpRequest(r, testRequest)
	assert.Equal(t, 404, res.Code, "Error - userId=100 : Response Code")

	testRes := model.Response{
		Code: model.CODE_USER_ZERO,
		Data: map[string]string{"message": model.MESSAGE_USER_ZERO},
	}

	jsonTestRes, _ := json.Marshal(testRes)

	assert.Equal(t, string(jsonTestRes), res.Body.String(), "Error - userId=3 : Response Body")
}

func homeGetUserIdNill(t *testing.T) {
	testRequest := testRequest{
		method:      http.MethodGet,
		uri:         "/post/home",
		queryString: "?userId=",
	}

	res := testHttpRequest(r, testRequest)
	assert.Equal(t, 404, res.Code, "Error - userId=nil : Response Code")

	testRes := model.Response{
		Code: model.CODE_USER_NIL,
		Data: map[string]string{"message": model.MESSAGE_USER_NIL},
	}

	jsonTestRes, _ := json.Marshal(testRes)

	assert.Equal(t, string(jsonTestRes), res.Body.String(), "Error - userId=nil : Response Body")
}

func homeGetPostCnt(t *testing.T) {
	testRequest := testRequest{
		method:      http.MethodGet,
		uri:         "/post/home",
		queryString: "?userId=1",
	}

	res := testHttpRequest(r, testRequest)
	assert.Equal(t, 200, res.Code, "Error - number of posts : Response Code")

	var testRes model.Response
	var testPosts []model.PostRe
	errUnmar1 := json.Unmarshal(res.Body.Bytes(), &testRes)
	assert.NoError(t, errUnmar1, "Error - number of posts : json unmarshal 1")

	testResData, errMar := json.Marshal(testRes.Data)
	assert.NoError(t, errMar, "Error - number of posts : json marshal")
	errUnmar2 := json.Unmarshal(testResData, &testPosts)
	assert.NoError(t, errUnmar2, "Error - number of posts : json unmarshal 2")

	assert.Equal(t, 20, len(testPosts), "Error - number of posts : Not equal max length")
}

func TestSetupRouter(t *testing.T) {
	if err := t.Run("Test home screen: userId = 3", homeGetUserId3); err == true {
		println("[SUCCES] Test home screen: userId = 3")
	}
	if err := t.Run("Test home screen: userId = 100", homeGetUserId100); err == true {
		println("[SUCCES] Test home screen: userId = 100")
	}
	if err := t.Run("Test home screen: userId = nil", homeGetUserIdNill); err == true {
		println("[SUCCES] Test home screen: userId = nil")
	}
	if err := t.Run("Test home screen: max number of posts", homeGetPostCnt); err == true {
		println("[SUCCES] Test home screen: max number of posts")
	}
}
