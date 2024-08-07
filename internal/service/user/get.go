package user

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userLogRepository.Create(ctx, model.UserLogInfo{
			UserId: id,
			Log:    updateLog,
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
