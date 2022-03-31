package convert

import (
	`path/filepath`

	`github.com/goexl/gfx`
	`github.com/storezhang/transfer`
	`github.com/ziyunix/core`
)

var _ = NewTask

// NewTask 创建转码任务
func NewTask(id int64, file *transfer.File, converts ...*core.Convert) *Task {
	return &Task{
		Id:       id,
		File:     file,
		Converts: converts,
	}
}

func (t *Task) SrcFilename(home string) string {
	return filepath.Join(t.SrcDir(home), filepath.Base(t.File.Name))
}
func (t *Task) SrcDir(home string) string {
	return filepath.Join(home, `src`, gfx.Name(t.File.Name))
}
