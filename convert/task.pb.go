package convert

import (
	`github.com/storezhang/transfer`
	`github.com/ziyunix/core`
)

// Task 不要删除此文件，只是占位
type Task struct {
	Id       int64
	File     *transfer.File
	Converts []*core.Convert
}
