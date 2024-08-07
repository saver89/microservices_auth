package user

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, req model.UserInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.userRepository.Update(ctx, id, req)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userLogRepository.Create(ctx, model.UserLogInfo{
			UserID: id,
			Log:    updateLog,
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
