package utils

import (
	"fmt"
	"time"
)

type DTime struct {
	time.Time
}

func (dtime DTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, dtime.Format("2006-01-02 15:04:05"))), nil
}

type BaseTime struct {
	CreateTime *DTime `json:"createTime" swaggertype:"primitive,string"`
}
