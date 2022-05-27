package Gin

import (
	"encoding/json"
	"forbizbe/main/model"
	"forbizbe/main/svmn/Mysql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func homeGetUserId3(t *testing.T) {
	Mysql.ConnDB()
	r := SetupRouter()
	userId := "3"

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/post/home?userId="+userId, nil)
	assert.NoError(t, err, "Request fail")

	r.ServeHTTP(res, req)

	var posts []model.Post
	var user model.User
	Mysql.DB.Order("created_At DESC").Find(&posts, "user_id = ?", userId)
	Mysql.DB.Find(&user, userId)

	userIdUint, _ := strconv.ParseUint(userId, 10, 0)

	testPosts := make([]model.PostRe, 3)
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

	resUserid3 := model.Response{
		Code: "OK",
		Data: testPosts,
	}

	jsonResUserId3, _ := json.Marshal(resUserid3)

	assert.Equal(t, 200, res.Code, "Error:1")
	assert.Equal(t, string(jsonResUserId3), res.Body.String(), "Error:2")
}
func TestSetupRouter(t *testing.T) {

	t.Run("userId = 3", homeGetUserId3)
}
