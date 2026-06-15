package v1

import "github.com/gogf/gf/v2/frame/g"

type HealthReq struct {
	g.Meta `path:"/health" method:"get"`
}

type HealthRes struct{}
