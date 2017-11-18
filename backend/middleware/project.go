package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"repo.letscode.sii.pl/wroclaw/three/backend/model"
	"repo.letscode.sii.pl/wroclaw/three/backend/utils"
)

const (
	ContextProjectKey ContextKey = "project"
)

func Project(db *gorm.DB, h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//user := req.Context().Value(ContextUserKey).(*model.User)
		vars := mux.Vars(req)

		if _, ok := vars["projectid"]; !ok {
			utils.NewErrorResponse(http.StatusBadRequest, model.ErrProjectInternal).AppendDebug(fmt.Errorf("middleware used without projectid key")).Write(rw)
			return
		}

		projectID, err := strconv.Atoi(vars["projectid"])
		if err != nil {
			utils.NewErrorResponse(http.StatusBadRequest, model.ErrProjectIDInvalid).AppendDebug(err).Write(rw)
			return
		}

		project := &model.Project{}
		project.ID = uint(projectID)
		if err := project.Find(db); err != nil {
			utils.NewErrorResponse(http.StatusInternalServerError, err).AppendDebug(err).Write(rw)
			return
		}
		ctx := context.WithValue(req.Context(), ContextProjectKey, project)
		h.ServeHTTP(rw, req.WithContext(ctx))
	})
}
