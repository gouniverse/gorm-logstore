package logstore

import (
	"gorm.io/gorm"
	"testing"
)

func Test_Log_BeforeCreate(t *testing.T) {
	var tx gorm.DB
	var _log Log
	_l := &_log

	err := _l.BeforeCreate(&tx)
	if err != nil {
		t.Errorf("Error BeforeCreate %v", err.Error())
	}
}
