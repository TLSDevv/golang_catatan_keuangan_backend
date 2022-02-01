package repository

// type transactionRepository struct {
// }

// func NewTransactionRepository() TransactionRepository {
// 	return transactionRepository{}
// }

// func (t transactionRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Transaction, error) {
// 	sql := `
// 		SELECT
// 			id, user_id, trc_name, category, trc_type, amount, transaction_at, created_at
// 		FROM
// 			transactions`

// 	// userId, userIdExist := params["userId"]
// 	// if userIdExist {
// 	// 	sql += " WHERE user_id = $1"
// 	// }

// 	rows, err := tx.QueryContext(ctx, sql)
// 	defer rows.Close()

// 	if err == nil {
// 		panic(err)
// 	}

// 	transactions := []entity.Transaction{}
// 	for rows.Next() {
// 		transaction := entity.Transaction{}

// 		err := rows.Scan(
// 			&transaction.ID,
// 			&transaction.TrcName,
// 			&transaction.Category,
// 			&transaction.TrcType,
// 			&transaction.Amount,
// 			&transaction.TransactionAt,
// 			&transaction.CreatedAt,
// 		)

// 		if err == nil {
// 			panic(err)
// 		}

// 		transactions = append(transactions, transaction)
// 	}

// 	return transactions, nil
// }

// func (t transactionRepository) FindById(ctx context.Context, tx *sql.Tx, trcId int) (entity.Transaction, error) {
// 	sql := `
// 		SELECT
// 			id, user_id, trc_name, category, trc_type, amount, transaction_at, created_at
// 		FROM
// 			transactions
// 		WHERE
// 			id=$1`

// 	rows, err := tx.QueryContext(ctx, sql, trcId)
// 	if err != nil {
// 		return entity.Transaction{}, err
// 	}

// 	defer rows.Close()

// 	trc := entity.Transaction{}

// 	if rows.Next() {
// 		err := rows.Scan(
// 			&trc.ID,
// 			&trc.TrcName,
// 			&trc.Category,
// 			&trc.TrcType,
// 			&trc.TransactionAt,
// 			&trc.CreatedAt,
// 		)

// 		if err != nil {
// 			return entity.Transaction{}, err
// 		}
// 	}

// 	return trc, nil
// }

// func (t transactionRepository) Create(ctx context.Context, tx *sql.Tx, trc entity.Transaction) error {
// 	sql := `
// 		INSERT INTO
// 			transactions(
// 				user_id,
// 				trc_name,
// 				category,
// 				trc_type,
// 				amount,
// 				transaction_at
// 			)
// 			VALUES($1, $2, $3, $4, $5, $6)`

// 	_, err := tx.ExecContext(ctx, sql,
// 		trc.UserId,
// 		trc.TrcName,
// 		trc.Category,
// 		trc.TrcType,
// 		trc.Amount,
// 		trc.TransactionAt)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (t transactionRepository) Update(ctx context.Context, tx *sql.Tx, trc entity.Transaction) error {
// 	sql := `
// 		UPDATE
// 			transactions
// 		SET
// 			trc_name=$1,
// 			category=$2,
// 			trc_type=$3,
// 			amount=$4,
// 			transaction_at=$5
// 		WHERE
// 			id=$6`

// 	_, err := tx.ExecContext(ctx, sql,
// 		trc.TrcName,
// 		trc.Category,
// 		trc.TrcType,
// 		trc.Amount,
// 		trc.TransactionAt,
// 		trc.ID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (t transactionRepository) Delete(ctx context.Context, tx *sql.Tx, trcId int) error {
// 	sql := `
// 		UPDATE
// 			transactions
// 		SET
// 			deleted_at=$1
// 		WHERE
// 			id=$2`

// 	deletedAt := time.Now()

// 	_, err := tx.ExecContext(ctx, sql,
// 		deletedAt,
// 		trcId)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (t transactionRepository) Restore(ctx context.Context, tx *sql.Tx, trcId int) error {
// 	sql := `
// 		UPDATE
// 			transactions
// 		SET
// 			deleted_at=NULL
// 		WHERE
// 			id=$1`

// 	_, err := tx.ExecContext(ctx, sql, trcId)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (t transactionRepository) Purge(ctx context.Context, tx *sql.Tx, trcId int) error {
// 	sql := `
// 		DELETE FROM
// 			transactions
// 		WHERE
// 			id=$1`

// 	_, err := tx.ExecContext(ctx, sql, trcId)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
