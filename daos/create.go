package daos

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serjbibox/FSTR/models"
)

func (dao *PassDAO) Create(f *models.Flow, r *http.Request) *models.Flow {
	if f.Err != nil {
		return f
	}
	if err := json.NewDecoder(r.Body).Decode(&f.Pass); err != nil {
		f.Err = fmt.Errorf("%w", err)
		return f
	}
	return f
}
