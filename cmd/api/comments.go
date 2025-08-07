package main

import (
	"net/http"
	"strconv"

	"github.com/ardhiqii/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreateCommentPayload struct {
  Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request){
  idParam := chi.URLParam(r, "postID")
  postID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

  var payload CreateCommentPayload
  if err:= readJSON(w,r,&payload); err != nil{
    app.badRequestResponse(w,r,err)
    return
  }

  comment := &store.Comment{
    UserID: 1,
    PostID: postID,
    Content: payload.Content,
  }

  ctx := r.Context()
  if err := app.store.Comments.Create(ctx,comment); err != nil{
    app.internalServerError(w,r,err)
    return
  }

  if err := app.jsonResponse(w,http.StatusCreated,comment); err != nil{
    app.internalServerError(w,r,err)
    return;
  }
  
}