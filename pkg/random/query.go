package random

import (
	"github.com/devinmcgloin/fokal/pkg/handler"
	"github.com/devinmcgloin/fokal/pkg/model"
	"github.com/devinmcgloin/fokal/pkg/retrieval"
)

func Image(state *handler.State, u *int64) (model.Image, error) {
	var imageId int64
	var err error
	if u != nil {

		err = state.DB.Get(&imageId, "SELECT random_image($1);", *u)
	} else {
		err = state.DB.Get(&imageId, "SELECT random_image();")
	}

	if err != nil {
		return model.Image{}, err
	}
	return retrieval.GetImage(state, imageId)
}