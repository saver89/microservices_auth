package user

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.userRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userLogRepository.Create(ctx, model.UserLogInfo{
			UserID: id,
			Log:    deleteLog,
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
