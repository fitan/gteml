package restapi

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/rest"
)

type ServiceObj struct {
	rest.BaseRest
}

func (s *ServiceObj) GetTableName() string {
	return "Services"
}

func (s *ServiceObj) GetModelObj() interface{} {
	return &model.Service{}
}

func (s *ServiceObj) GetModelObjs() interface{} {
	return &[]model.Service{}
}

func (s *ServiceObj) GetFirstObj() interface{} {
	return s.GetModelObj()
}

func (s *ServiceObj) GetFindObj() interface{} {
	return s.GetModelObjs()
}

func (s *ServiceObj) RelationsField() map[string]rest.RelationFielder {
	return map[string]rest.RelationFielder{}
}
