package retrieval

import (
	"database/sql"
	"net/http"

	"errors"

	"strconv"

	"github.com/fokal/fokal-core/pkg/handler"
	"github.com/fokal/fokal-core/pkg/model"
	"github.com/fokal/fokal-core/pkg/stats"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func UserHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	username := mux.Vars(r)["ID"]

	ref, err := GetUserRef(store.DB, username)
	if err != nil {
		return rsp, err
	}

	user, err := GetUser(store, ref.Id)
	if err != nil {
		return rsp, err
	}
	return handler.Response{
		Code: http.StatusOK,
		Data: user,
	}, nil
}

func UserImagesHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	username := mux.Vars(r)["ID"]

	ref, err := GetUserRef(store.DB, username)
	if err != nil {
		return rsp, err
	}

	user, err := GetUserImages(store, ref.Id)
	if err != nil {
		return handler.Response{}, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: user,
	}, nil
}

func UserFavoritesHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	username := mux.Vars(r)["ID"]

	ref, err := GetUserRef(store.DB, username)
	if err != nil {
		return rsp, err
	}

	user, err := GetUserFavorites(store, ref.Id)
	if err != nil {
		return handler.Response{}, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: user,
	}, nil
}

func LoggedInUserHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response

	val, ok := context.GetOk(r, "auth")
	if !ok {
		return rsp, handler.StatusError{
			Code: http.StatusUnauthorized,
			Err:  errors.New("Must be logged in to use this endpoint")}
	}

	usrRef := val.(model.Ref)
	user, err := GetUser(store, usrRef.Id)
	if err != nil {
		return rsp, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: user,
	}, nil
}

func LoggedInUserImagesHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response

	val, ok := context.GetOk(r, "auth")
	if !ok {
		return rsp, handler.StatusError{
			Code: http.StatusUnauthorized,
			Err:  errors.New("Must be logged in to use this endpoint")}
	}

	usrRef := val.(model.Ref)
	images, err := GetUserImages(store, usrRef.Id)
	if err != nil {
		return rsp, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: images,
	}, nil
}

func ImageHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response

	id := mux.Vars(r)["ID"]

	ref, err := GetImageRef(store.DB, id)
	if err != nil {
		return rsp, err
	}

	img, err := GetImage(store, ref.Id)
	if err != nil {
		return rsp, err
	}

	stats.AddStat(store.DB, ref.Id, "view")

	return handler.Response{
		Code: http.StatusOK,
		Data: img,
	}, nil
}

func TagHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	var err error
	var limit int
	id := mux.Vars(r)["ID"]

	params := r.URL.Query()
	l, ok := params["limit"]
	if ok {
		if len(l) == 1 {
			limit, err = strconv.Atoi(l[0])
			if err != nil {
				limit = 500
			}
		}
	}

	if limit == 0 {
		limit = 500
	}

	var tid int64
	err = store.DB.Get(&tid, "SELECT id FROM content.image_tags as t WHERE t.description = $1;", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return rsp, handler.StatusError{Code: http.StatusNotFound, Err: errors.New("no corresponding Tag found")}
		}
		return rsp, err
	}

	tag := model.Ref{Collection: model.Tags, Id: tid, Shortcode: id}
	images, err := TaggedImages(store, tid, limit)
	if err != nil {
		return rsp, err
	}

	images.Permalink = tag.ToURL(store.Port, store.Local)

	return handler.Response{
		Code: http.StatusOK,
		Data: images,
	}, nil
}

func RecentImageHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	var err error

	params := r.URL.Query()

	var limit int
	l, ok := params["limit"]
	if ok {
		if len(l) == 1 {
			limit, err = strconv.Atoi(l[0])
			if err != nil {
				limit = 500
			}
		}
	}

	if limit == 0 {
		limit = 500
	}

	images, err := RecentImages(store, limit)
	if err != nil {
		return rsp, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: images,
	}, nil
}

func FeaturedImageHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	var err error

	params := r.URL.Query()
	var limit int
	l, ok := params["limit"]
	if ok {
		if len(l) == 1 {
			limit, err = strconv.Atoi(l[0])
			if err != nil {
				limit = 500
			}
		}
	}

	if limit == 0 {
		limit = 500
	}
	images, err := FeaturedImages(store, limit)
	if err != nil {
		return rsp, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: images,
	}, nil
}

func TrendingImagesHander(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var rsp handler.Response
	var err error

	params := r.URL.Query()
	var limit int
	l, ok := params["limit"]
	if ok {
		if len(l) == 1 {
			limit, err = strconv.Atoi(l[0])
			if err != nil {
				limit = 500
			}
		}
	}

	if limit == 0 {
		limit = 500
	}
	images, err := Trending(store, limit)
	if err != nil {
		return rsp, err
	}

	return handler.Response{
		Code: http.StatusOK,
		Data: images,
	}, nil
}
