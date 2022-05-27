package home

import (
	. "forbizbe/main/model"
	. "forbizbe/main/svmn/Mysql"

	"sort"

	"github.com/gin-gonic/gin"
)

func HomeUserList(c *gin.Context) {
	var posts []Post
	var users []User

	userId := c.Query("userId")

	if len(userId) == 0 {
		message := map[string]string{"message": MESSAGE_USER_NIL}
		response := Response{
			Code: CODE_USER_NIL,
			Data: message,
		}
		c.JSON(404, response)
		return
	}

	/* Data fetch */
	DB.Preload("Posts.Writer").Preload("Follows.Posts.Writer").Find(&users, userId)
	if len(users) == 0 {
		message := map[string]string{"message": MESSAGE_USER_ZERO}
		response := Response{
			Code: CODE_USER_ZERO,
			Data: message,
		}
		c.JSON(404, response)
		return
	}

	posts = append(posts, users[0].Posts...)
	for _, follow := range users[0].Follows {
		posts = append(posts, follow.Posts...)
	}

	/* Sorting */
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(*posts[j].CreatedAt)
	})

	/* Formating */
	data := make([]PostRe, len(posts))
	for i := 0; i < len(posts); i++ {
		data[i].ID = posts[i].ID
		data[i].Comment = posts[i].Comment
		data[i].UpdatedAt = posts[i].UpdatedAt.Format(TIME_FORMAT)
		data[i].CreatedAt = posts[i].CreatedAt.Format(TIME_FORMAT)
		data[i].Writer.ID = posts[i].Writer.ID
		data[i].Writer.Name = posts[i].Writer.Name
		data[i].Writer.Email = posts[i].Writer.Email
		data[i].Writer.CreatedAt = posts[i].Writer.CreatedAt.Format(TIME_FORMAT)
		data[i].Writer.UpdatedAt = posts[i].Writer.UpdatedAt.Format(TIME_FORMAT)
	}

	/* Cut to Top 20 */
	maxSize := len(data)
	if maxSize > 20 {
		maxSize = 20
	}

	data = data[:maxSize]

	result := Response{
		Code: "OK",
		Data: data,
	}

	c.JSON(200, result)
}
