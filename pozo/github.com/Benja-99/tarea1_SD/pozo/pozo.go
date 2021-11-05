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
	return &Monto{CantidadTotal: s.monto}, nil
}

func (s *Server) SetMonto(ctx context.Context, in *Monto) (*Monto, error) {
	log.Printf("Peticion del cliente para obtener monto")
	s.monto = in.Monto2
	return &Monto{CantidadTotal: s.monto}, nil
}
