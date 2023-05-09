package context

import (
	"github.com/lphhh/douyin/open/config"
	"github.com/lphhh/douyin/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
