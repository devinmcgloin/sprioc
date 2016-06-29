package core

import (
	"net/http"
	"strings"

	"github.com/sprioc/sprioc-core/pkg/model"
	"github.com/sprioc/sprioc-core/pkg/rsp"
	"github.com/sprioc/sprioc-core/pkg/store"
)

func GetUser(ref model.DBRef) (model.User, rsp.Response) {

	if strings.Compare(ref.Collection, "users") != 0 {
		return model.User{}, rsp.Response{Message: "Ref is of the wrong collection type",
			Code: http.StatusNotFound}
	}

	var user = model.User{}

	err := store.Get(ref, &user)
	if err != nil {
		return model.User{}, rsp.Response{Message: "User not found",
			Code: http.StatusNotFound}
	}

	return user, rsp.Response{Code: http.StatusOK}
}

func GetImage(ref model.DBRef) (model.Image, rsp.Response) {
	if strings.Compare(ref.Collection, "images") != 0 {
		return model.Image{}, rsp.Response{Message: "Ref is of the wrong collection type",
			Code: http.StatusNotFound}
	}

	var image model.Image

	err := store.Get(ref, &image)
	if err != nil {
		return model.Image{}, rsp.Response{Message: "Image not found",
			Code: http.StatusNotFound}
	}

	return image, rsp.Response{Code: http.StatusOK}
}

func GetCollection(ref model.DBRef) (model.Collection, rsp.Response) {
	if strings.Compare(ref.Collection, "collections") != 0 {
		return model.Collection{}, rsp.Response{Message: "Ref is of the wrong collection type",
			Code: http.StatusNotFound}
	}

	var col model.Collection

	err := store.Get(ref, &col)
	if err != nil {
		return model.Collection{}, rsp.Response{Message: "Collection not found",
			Code: http.StatusNotFound}
	}

	return col, rsp.Response{Code: http.StatusOK}
}