package module

import (
	"github.com/qingniusoft/muraena/module/crawler"
	"github.com/qingniusoft/muraena/module/necrobrowser"
	"github.com/qingniusoft/muraena/module/statichttp"
	"github.com/qingniusoft/muraena/module/tracking"
	"github.com/qingniusoft/muraenateam/session"
)

// LoadModules load modules
func LoadModules(s *session.Session) {
	s.Register(crawler.Load(s))
	s.Register(statichttp.Load(s))
	s.Register(tracking.Load(s))
	s.Register(necrobrowser.Load(s))
}
