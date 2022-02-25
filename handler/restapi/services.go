package restapi

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/restcommon"
)

type ServiceObj struct {
	restcommon.BaseRest
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

func (s *ServiceObj) RelationsField() map[string]restcommon.RelationFielder {
	return map[string]restcommon.RelationFielder{}
}
