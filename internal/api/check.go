package api

import "github.com/pkg/errors"

func (r *TestReq) Verify() error {
	if r.Page <= 0 {
		return errors.Errorf("Page 字段必须＞0")
	}
	return nil
}
