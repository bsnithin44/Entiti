package entities

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	api "github.com/bsnithin44/entiti/pkg/api"

	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

var entityMetadata = table.Metadata{
	Name:    "fpa.entities",
	Columns: []string{"plural_name", "singular_name", "id", "created_at", "updated_at", "is_active"},
	PartKey: []string{"id"},
}
var entityTable = table.New(entityMetadata)

func GetEntityTypeByName(dbSession gocqlx.Session, entityName string) (*EntityTypeInDb, *api.MyError) {

	records := []EntityTypeInDb{}
	err := api.MyError{}
	q := qb.Select(entityTable.Name()).Where(qb.Eq("plural_name")).Limit(1).Query(dbSession).Bind(entityName)
	rErr := q.SelectRelease(&records)

	if rErr != nil {
		err.AdditionalInfo = err.Error()
		err.Description = "DB error"
		err.StatusCode = http.StatusBadRequest
		return nil, &err
	} else {

		if len(records) > 0 {
			return &records[0], nil
		} else {
			err.Description = "Not found"
			err.StatusCode = http.StatusNotFound
			return nil, &err
		}
	}
}

func GetEntityTypeById(dbSession gocqlx.Session, entityId string) (*EntityTypeInDb, *api.MyError) {

	records := []EntityTypeInDb{}
	err := api.MyError{}
	q := qb.Select(entityTable.Name()).Where(qb.Eq("id")).Limit(1).Query(dbSession).Bind(entityId)
	rErr := q.SelectRelease(&records)

	if rErr != nil {
		err.AdditionalInfo = err.Error()
		err.Description = "DB error"
		err.StatusCode = http.StatusBadRequest
		return nil, &err
	} else {

		if len(records) > 0 {
			return &records[0], nil
		} else {
			err.Description = "Not found"
			err.StatusCode = http.StatusNotFound
			return nil, &err
		}
	}

}
func CreateEntityTypeByName(dbSession gocqlx.Session, dbRecord EntityTypeInDb) (*EntityTypeInDb, *api.MyError) {

	var applied bool
	err := api.MyError{}
	m := make(map[string]interface{})

	q := qb.Insert(entityTable.Name()).Columns(entityMetadata.Columns...).Unique().Query(dbSession).BindStruct(dbRecord)
	fmt.Println(q)
	applied, sErr := q.MapScanCAS(m)

	if sErr != nil {
		err.Description = sErr.Error()
		err.StatusCode = http.StatusInternalServerError
		return nil, &err
	}

	if !applied {

		err.Description = "Conflict, already present in DB"
		err.StatusCode = http.StatusBadRequest
		return nil, &err
	}

	return &dbRecord, nil

}

func GetAllEntityTypes(dbSession gocqlx.Session, isActive string) ([]EntityTypeInDb, *api.MyError) {
	var entities []EntityTypeInDb

	cql := bytes.Buffer{}
	var values []interface{}

	if len(isActive) > 0 {
		if len(cql.String()) > 0 {

			cql.WriteString(" AND is_active = ?")
		} else {
			cql.WriteString(" WHERE is_active = ?")
		}
		isActive, _ := strconv.ParseBool(isActive)
		values = append(values, isActive)
	}

	cql.WriteString(" ALLOW FILTERING")

	q := dbSession.Query("SELECT * FROM fpa.entities"+cql.String(), nil).Bind(values...)
	if rErr := q.SelectRelease(&entities); rErr != nil {
		err := api.MyError{
			AdditionalInfo: rErr.Error(),
		}
		err.Description = "DB error"
		err.StatusCode = http.StatusBadRequest
		return nil, &err
	}
	return entities, nil
}

func UpdateEntityTypeByName(dbSession gocqlx.Session, dbRecord EntityTypeInDb) *api.MyError {
	q := dbSession.Query(entityTable.Insert()).BindStruct(dbRecord)
	if rErr := q.ExecRelease(); rErr != nil {
		err := api.MyError{
			AdditionalInfo: rErr.Error(),
		}
		err.Description = "DB error"
		err.StatusCode = http.StatusBadRequest
		return &err
	}
	return nil
}

// check this
func UpdateEntityTypeById(dbSession gocqlx.Session, dbRecord EntityTypeInDb) (*EntityTypeInDb, *api.MyError) {

	var applied bool
	err := api.MyError{}
	m := make(map[string]interface{})

	q := qb.Update(entityTable.Name()).Existing().Set("updated_at", "is_active").Where(qb.Eq("plural_name")).Query(dbSession).BindStruct(dbRecord)
	applied, sErr := q.MapScanCAS(m)

	if sErr != nil {

		err.Description = err.Error()
		err.StatusCode = http.StatusInternalServerError
		return nil, &err
	}

	if !applied {
		err.Description = "Not found"
		err.StatusCode = http.StatusBadRequest
		return nil, &err
	}

	return &dbRecord, nil
}
