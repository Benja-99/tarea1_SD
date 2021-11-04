package pozo

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetMonto(ctx context.Context, in *Monto) (*Monto, error) {
	log.Printf("Peticion del cliente para obtener monto")
	return &Monto{CantidadTotal: 10000}, nil
}
