package user

import (
	"context"

	"github.com/saver89/microservices_auth/internal/model"
)

func (s *serv) Create(ctx context.Context, req model.CreateUserRequest) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, req)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userLogRepository.Create(ctx, model.UserLogInfo{
			UserId: id,
			Log:    createLog,
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
