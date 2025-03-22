package handler

import (
	"chat_server/middleware"
	"chat_server/usecase/input_port"
	"net/http"
	"chat_server/api/schema"

	"github.com/labstack/echo/v4"
)

type TweetHandler struct {
	TweetUC input_port.ITweetUsecase
}

func NewTweetHandler(tweetUC input_port.ITweetUsecase) *TweetHandler {
	return &TweetHandler{
		TweetUC: tweetUC,
	}
}

func (h *TweetHandler) CreateTweet(c echo.Context) error {
	user, err := middleware.GetUserFromContext(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	var req schema.CreateTweetReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	err = h.TweetUC.Create(input_port.CreateTweetArgs{
		UserID: user.UserID,
		Content: req.Content,
		ImageUrl: req.ImageUrl,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.NoContent(http.StatusCreated)
}