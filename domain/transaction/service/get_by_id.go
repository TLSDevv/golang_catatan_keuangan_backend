package service

func (s Service) GetByID(ctx context.Context, transactionID int) (*entities.Transaction, error) {
    transaction, err := s.repo.GetByID(ctx, transactionID)
    if err != nil {
        return nil, err
    }

    return transaction, nil
}