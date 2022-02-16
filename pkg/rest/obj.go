package rest

type Objer interface {
	ModelObjer
	GetFirstObj() interface{}
	GetFindObj() interface{}
}

type ModelObjer interface {
	GetTableName() string
	GetModelObj() interface{}
	GetModelObjs() interface{}
}

type FieldConfer interface {
	CreateField() (s []string, o []string)
	UpdateField() (s []string, o []string)
	RelationField() map[string]RelationConf
}

type BaseModelObj struct {
}

func (b *BaseModelObj) GetTableName() string {
	return ""
}

func (b *BaseModelObj) GetModelObj() interface{} {
	i := make(map[string]interface{}, 0)
	return &i
}

func (b *BaseModelObj) GetModelObjs() interface{} {
	i := make([]map[string]interface{}, 0, 0)
	return &i
}

type BaseObj struct {
	ModelObjer
}

func (b *BaseObj) GetFirstObj() interface{} {
	return b.GetModelObj()
}

func (b *BaseObj) GetFindObj() interface{} {
	return b.GetModelObjs()
}

type BaseFieldConf struct {
}

func (f *BaseFieldConf) CreateField() (s []string, o []string) {
	return []string{"*"}, []string{"id", "updated_at", "deleted_at", "created_at"}
}

func (f *BaseFieldConf) UpdateField() (s []string, o []string) {
	return []string{"*"}, []string{"id", "updated_at", "deleted_at", "created_at"}
}

func (f *BaseFieldConf) RelationField() map[string]RelationConf {
	return nil
}
