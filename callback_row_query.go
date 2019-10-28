package gorm

import (
	"database/sql"
	"fmt"
)

// Define callbacks for row query
func init() {
	DefaultCallback.RowQuery().Register("gorm:row_query", rowQueryCallback)
}

type RowQueryResult struct {
	Row *sql.Row
}

type RowsQueryResult struct {
	Rows  *sql.Rows
	Error error
}

// queryCallback used to query data from database
func rowQueryCallback(scope *Scope) {
	now := NowFunc()
	if result, ok := scope.InstanceGet("row_query_result"); ok {
		scope.prepareQuerySQL()
		if str, ok := scope.Get("gorm:query_option"); ok {
			scope.SQL += addExtraSpaceIfExist(fmt.Sprint(str))
		}
		for scope.queryRetryTimes = 0; scope.queryRetryTimes <= scope.db.MaxRetryTimes; scope.queryRetryTimes++ {
			if func() error {
				defer func() {
					scope.trace(now)
					now = NowFunc()
				}()
				scope.db.Error = nil
				if rowResult, ok := result.(*RowQueryResult); ok {
					rowResult.Row = scope.SQLDB().QueryRow(scope.SQL, scope.SQLVars...)
				} else if rowsResult, ok := result.(*RowsQueryResult); ok {
					rowsResult.Rows, rowsResult.Error = scope.SQLDB().Query(scope.SQL, scope.SQLVars...)
					if scope.Err(rowsResult.Error) != nil {
						return rowsResult.Error
					}
				}
				return nil
			}() == nil {
				break
			}
		}
	}
}
