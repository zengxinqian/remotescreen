package vnc

import "github.com/zengxinqian/remotescreen"

type Server struct {
	screen *remotescreen.Screen
}

func (s *Server) ListenAndServ(addr string) error {
	return nil
}

func ListenAndServ(addr string, screen *remotescreen.Screen) error {

	s := &Server{screen: screen}
	return s.ListenAndServ(addr)

}
