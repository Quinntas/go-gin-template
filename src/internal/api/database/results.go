package database

type InsertResult struct {
	LastInsertId int64
}

func NewInsertResult(lastInsertId int64) InsertResult {
	return InsertResult{
		LastInsertId: lastInsertId,
	}
}

type UpdateResult struct {
	RowsAffected int64
}

func NewUpdateResult(rowsAffected int64) UpdateResult {
	return UpdateResult{
		RowsAffected: rowsAffected,
	}
}
