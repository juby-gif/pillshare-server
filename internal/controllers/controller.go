package controllers

import (
	"context"
	"database/sql"
	// "fmt"
	"net/http"

	cache "github.com/go-redis/cache/v8"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/internal/repositories"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

type Controller struct {
	db       *sql.DB
	UserRepo models.UserRepo
	cache    *cache.Cache
}

func New(db *sql.DB) *Controller {
	userRepo := repositories.NewUserRepo(db)
	cache := utils.RedisCache()
	return &Controller{
		db:       db,
		UserRepo: userRepo,
		cache:    cache,
	}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		URL         []string
		n           int
		authStatus  bool
		sessionUUID string
		accessToken string
	)

	// Check for the context has key `url_split` and `length` that
	// have values not equal to nil
	if ctx.Value("url_split") != nil && ctx.Value("length") != nil {

		// Get the values of keys `URL` & `n`(length) from the context
		// Assign the values to `URL` and `n`
		URL = ctx.Value("url_split").([]string)
		n = ctx.Value("length").(int)

		// Check for the context has key `is_authorized` and `session_uuid` that
		// have values not equal to nil
		if ctx.Value("is_authorized") != nil && ctx.Value("session_uuid") != nil && ctx.Value("access_token") != nil {

			// Get the values of keys `is_authorized` & `session_uuid` from the context
			// Assign the values to `authStatus` and `sessionUUID`
			authStatus = ctx.Value("is_authorized").(bool)
			sessionUUID = ctx.Value("session_uuid").(string)
			accessToken = ctx.Value("access_token").(string)
		}
	}

	// Get User modal from cache
	var user models.User
	if err := c.cache.Get(ctx, sessionUUID, &user); err == nil {

		// For debugging purpose only
		// fmt.Println(user)

		// Saving the `user.User_id` from cache to context with key `user_id`
		ctx = context.WithValue(ctx, "user_id", user.User_id)
		r = r.WithContext(ctx)
	}

	switch {
	case n == 3 && URL[2] == "version" && r.Method == "GET":
		c.getVersion(w, r)
	case n == 3 && URL[2] == "refresh-token" && r.Method == "GET":
		c.postRefreshToken(w, r, accessToken)
	case n == 3 && URL[2] == "login" && r.Method == "POST":
		c.postLogin(w, r)
	case n == 3 && URL[2] == "register" && r.Method == "POST":
		c.postRegister(w, r)
	case n == 3 && URL[2] == "hello" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getHello(w, r)
		}
	case n == 3 && URL[2] == "dashboard" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getDashboard(w, r)
		}
	case n == 3 && URL[2] == "hello" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getHello(w, r)
		}
	case n == 3 && URL[2] == "user" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getUserProfile(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}
