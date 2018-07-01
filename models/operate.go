package models

type OperateType int

const (
	INSERT OperateType = iota
	UPDATE
	DELETE
)

type Operate struct {
	Type    OperateType `bson: "type" json:"type"`
	Data    Item        `bson: "data" json:"data"`
	Version int         `bson: "version" json:"version"`
}
