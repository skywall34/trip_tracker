package handlers

import (
	"net/http"

	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/templates"
)


type GetWorldMap3dHandler struct {}

type GetWorldMap3dHandlerParams struct {}

func NewGetWorldMap3dHandlerHandler() *GetWorldMap3dHandler {
	return &GetWorldMap3dHandler{}
}

func (t *GetWorldMap3dHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
    _, ok := ctx.Value(m.UserKey).(int)

	if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
	
	c := templates.WorldMap3D()
    templates.Layout(c, "3D World Map").Render(r.Context(), w)
}