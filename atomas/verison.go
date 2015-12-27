package atomas
import (
	"net/http"
	"fmt"
)

func CreateVersionHandler(version string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, version)
	}
}
