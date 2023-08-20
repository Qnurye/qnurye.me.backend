package controllers

import (
	"github.com/gin-gonic/gin"
	"qnurye/qnurye.me/models"
	"qnurye/qnurye.me/pkg/response"
	"strconv"
)

type PostController struct{}

var postModel = new(models.Post)

func (receiver PostController) Index(c *gin.Context) {
	var (
		limit, offset int
		err           error
	)

	if c.Query("limit") != "" {
		if limit, err = strconv.Atoi(c.Query("limit")); err != nil {
			response.Failure(c, response.INVALID_PARAMS, "").Abort()
		}
	} else {
		limit = 0
	}
	if c.Param("offset") != "" {
		if offset, err = strconv.Atoi(c.Query("offset")); err != nil {
			response.Failure(c, response.INVALID_PARAMS, "").Abort()
		}
	} else {
		offset = 0
	}

	if posts, err := postModel.Take(limit, offset); err != nil {
		response.Failure(c, response.UNAVAILABLE, err.Error())
	} else {
		response.Success(c, *posts)
	}

	return
}
