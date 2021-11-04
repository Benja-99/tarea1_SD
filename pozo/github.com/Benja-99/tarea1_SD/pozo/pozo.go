package pozo

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	monto int32
}

func (s *Server) GetMonto(ctx context.Context, in *Monto) (*Monto, error) {
	log.Printf("Peticion del cliente para obtener monto")
	s.monto = 100000
	return &Monto{CantidadTotal: s.monto}, nil
}
