package rbac

import (
	"net/http"

	"github.com/content-services/content-sources-backend/pkg/config"
)

func server() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/inventory/v1beta2/check", serveCheck)
	mux.HandleFunc("/api/inventory/v1beta2/checkForUpdate/", serveCheckForUpdate)
	mux.HandleFunc("/api/rbac/v2/workspaces", serveGetWorkspaces)
	return mux
}
func serveCheck(w http.ResponseWriter, r *http.Request) {
	kesselMock := config.Get().Mocks.Kessel
	usersWithRead := kesselMock.UserRead
	usersWithRead = append(kesselMock.UserReadWrite, usersWithRead...)
}

func serveCheckForUpdate(w http.ResponseWriter, r *http.Request) {}

func serveGetWorkspaces(w http.ResponseWriter, r *http.Request) {}
