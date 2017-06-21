package service

import (
	"database/sql"
	"log"
	"time"

	"github.com/charakoba-com/fault_info/backend/db"
	"github.com/charakoba-com/fault_info/backend/model"
	"github.com/pkg/errors"
)

var _infoSvc InfoService

// InfoService type
type InfoService struct{}

// Create a new info
func (svc *InfoService) Create(tx *sql.Tx, request model.PostInfoHandlerRequest) (int, error) {
	log.Printf("service.Info.Create")
	begin, err := time.Parse("2006-01-02 15:04:05", request.Begin)
	if err != nil {
		return 0, errors.Wrap(err, `parsing begin time`)
	}

	di := db.Info{
		Type:    request.InfoType,
		Service: request.Service,
		Begin: begin,
		Detail:  request.Detail,
	}
	if request.End != "" {
		di.End.Valid = true
		end, err := time.Parse("2006-01-02 15:04:05", request.End)
		if err != nil {
			return 0, errors.Wrap(err, `parsing end time`)
		}
		di.End.Time = end
	}
	if err := di.Create(tx); err != nil {
		return 0, errors.Wrap(err, `creating new record`)
	}
	return di.ID, nil
}

// Listup informations
func (svc *InfoService) Listup(tx *sql.Tx) ([]model.Info, error) {
	log.Printf("service.Info.Listup")

	var infoList db.InfoList
	if err := infoList.Listup(tx); err != nil {
		return nil, errors.Wrap(err, `loading info list`)
	}
	l := make(model.InfoList, len(infoList))
	for i, info := range infoList {
		if err := l[i].FromDB(&info); err != nil {
			return nil, errors.Wrap(err, `converting db.Info to model.Info`)
		}
	}

	return l, nil
}

// Update info
func (svc *InfoService) Update(tx *sql.Tx, id int, request model.UpdateInfoHandlerRequest) error {
	log.Printf("service.Info.Update")

	var di db.Info
	if err := di.Load(tx, id); err != nil {
		return errors.Wrap(err, `loading info`)
	}
	if request.InfoType != "" {
		di.Type = request.InfoType
	}
	if request.Service != "" {
		di.Service = request.Service
	}
	if request.Begin != "" {
		begin, err := time.Parse("2006-01-02 15:04:05", request.Begin)
		if err != nil {
			return errors.Wrap(err, `parsing begin time`)
		}
		di.Begin = begin
	}
	if request.End != "" {
		end, err := time.Parse("2006-01-02:15:04:05", request.End)
		if err != nil {
			return errors.Wrap(err, `parsing end time`)
		}
		di.End.Valid = true
		di.End.Time = end
	}
	if request.Detail != "" {
		di.Detail = request.Detail
	}
	if err := di.Update(tx); err != nil {
		return errors.Wrap(err, `updating database`)
	}

	return nil
}
