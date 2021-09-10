package entities

import (
	"net/http"
	"time"

	api "github.com/bsnithin44/entiti/pkg/api"
	db "github.com/bsnithin44/entiti/pkg/database"

	"github.com/google/uuid"
)

func GetEntityTypes(isActive string) ([]EntityTypeInDb, *api.MyError) {

	dbRecord, err := GetAllEntityTypes(db.GetDbSession(), isActive)
	if err != nil {
		return nil, err
	} else {
		return dbRecord, nil
	}
}

func UpdateEntityType(entityId string, entity EntityTypeBase) (*EntityTypeInDb, *api.MyError) {

	dbRecord, err := GetEntityTypeById(db.GetDbSession(), entityId)
	if err != nil {
		return nil, err
	} else {
		newdbRecord := EntityTypeInDb{

			EntityTypeBase: entity,
			EntityTypeSystem: EntityTypeSystem{
				Id:        dbRecord.Id,
				UpdatedAt: time.Now().UTC(),
			},
		}
		newdbRecord.PluralName = dbRecord.PluralName
		record, err := UpdateEntityTypeById(db.GetDbSession(), newdbRecord)
		if err != nil {
			return nil, err
		} else {
			return record, nil
		}

	}
}

// Validate all the attributes in payload
func CreateEntityType(entity CreateEntityTypeStruct) (*EntityTypeInDb, *api.MyError) {

	record, _ := GetEntityTypeByName(db.GetDbSession(), *entity.PluralName)
	if record != nil {
		err2 := api.MyError{}
		err2.Description = "Conflict, already present in DB"
		err2.StatusCode = http.StatusBadRequest
		return nil, &err2
	}

	Id := *entity.SingularName + ":" + uuid.New().String()
	createdAt := time.Now().UTC()
	updatedAt := createdAt
	dbRecord := &EntityTypeInDb{
		EntityTypeBase:     entity.EntityTypeBase,
		EntityTypeReadOnly: entity.EntityTypeReadOnly,
	}
	dbRecord.Id = Id
	dbRecord.CreatedAt = createdAt
	dbRecord.UpdatedAt = updatedAt
	dbRecord, err := CreateEntityTypeByName(db.GetDbSession(), *dbRecord)
	if err != nil {
		return nil, err
	} else {
		return dbRecord, nil
	}

}
func GetEntityType(entityId string) (*EntityTypeInDb, *api.MyError) {

	dbRecord, err := GetEntityTypeById(db.GetDbSession(), entityId)
	if err != nil {
		return nil, err
	} else {
		return dbRecord, nil
	}

}

// Bulk
func CreateEntityTypeBulk(entities []CreateEntityTypeStruct) (*int, *api.MyError) {
	count := 0
	for _, entity := range entities {
		_, err := CreateEntityType(entity)
		if err == nil {
			count += 1
		}
	}
	return &count, nil

}
