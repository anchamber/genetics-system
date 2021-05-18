package db

import (
	"fmt"
	apiModel "github.com/anchamber/genetics-api/model"
)

func (o *Options) createPaginationClause() string {
	if o.Pageination == nil {
		return ""
	}
	var limit int64 = int64(o.Pageination.Limit)
	if limit <= 0 {
		limit = -1
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", limit, o.Pageination.Offset)
}

func getOperatorAsString(operator apiModel.Operator) string {
	switch operator {
	case apiModel.EQ:
		return "="
	case apiModel.GREATER:
		return ">"
	case apiModel.GREATER_EQ:
		return ">="
	case apiModel.SMALLER:
		return "<"
	case apiModel.SMALLER_EQ:
		return "<="
	case apiModel.CONTAINS:
		return "LIKE"
	default:
		return "="
	}
}

func (o *Options) createFilterClause() string {
	if len(o.Filters) == 0 {
		return ""
	}
	whereClause := "WHERE "

	for index, filter := range o.Filters {
		if index > 0 {
			whereClause += " AND "
		}

		if filter.Operator == apiModel.CONTAINS {
			whereClause += fmt.Sprintf("instr(%s, :%s) > 0", filter.Key, filter.Key)
		} else {
			whereClause += fmt.Sprintf("%s %v :%s", filter.Key, getOperatorAsString(filter.Operator), filter.Key)
		}
	}
	// fmt.Println(whereClause)
	return whereClause
}

func (o *Options) createFilterMap() map[string]interface{} {
	values := make(map[string]interface{})
	for _, filter := range o.Filters {
		values[filter.Key] = filter.Value
	}
	return values
}
