package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/internal/kafka"
	"sync"
)

type CreateFIORequest struct {
	FIOs []*Fio `json:"fios"`
}

type CreateFIOResponse struct {
	FailedFIOs []*FailedFio `json:"failed_fios"`
}

func (s *Server) CreateFIO(ctx context.Context, r CreateFIORequest) (CreateFIOResponse, error) {
	var gw sync.WaitGroup
	var m sync.Mutex
	var res []*FailedFio
	for _, fio := range r.FIOs {
		go func(fio *Fio) {
			defer gw.Done()

			var msg map[string]string
			if err := vr.Struct(fio); err != nil {
				msg = newValidationError(err.(validator.ValidationErrors)).Fields
			}

			if err := s.ctrl.Produce(ctx, kafka.FIO(*fio), msg); err != nil || msg != nil {
				f := FailedFio(*fio)
				m.Lock()
				defer m.Unlock()
				res = append(res, &f)
			}
		}(fio)
	}
	gw.Add(len(r.FIOs))
	gw.Wait()
	return CreateFIOResponse{FailedFIOs: res}, nil
}
