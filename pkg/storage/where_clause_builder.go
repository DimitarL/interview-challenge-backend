package storage

import (
	"fmt"
	"strings"

	"github.com/DimitarL/interview-challenge-backend/common"
)

type whereClauseBuilder struct {
	where strings.Builder
}

func newWhereClauseBuilder() *whereClauseBuilder {
	return &whereClauseBuilder{}
}

func (b *whereClauseBuilder) AddParamsIfPresent(params *common.RentalQueryParameters) string {
	var paramName string
	var operator string
	var isWhereClauseElementFirst bool = true

	if params.PriceMin != nil {
		b.addConnectingPart(&isWhereClauseElementFirst)

		paramName = "price_per_day"
		operator = ">="
		b.where.WriteString(fmt.Sprintf("%s %s %d", paramName, operator, *params.PriceMin))
	}
	if params.PriceMax != nil {
		b.addConnectingPart(&isWhereClauseElementFirst)

		paramName = "price_per_day"
		operator = "<="
		b.where.WriteString(fmt.Sprintf("%s %s %d", paramName, operator, *params.PriceMax))
	}
	if params.Ids != "" {
		b.addConnectingPart(&isWhereClauseElementFirst)

		paramName = "rentals.id"
		operator = "IN"
		b.where.WriteString(fmt.Sprintf("%s %s (%s)", paramName, operator, params.Ids))
	}
	if params.Near != nil {
		b.addConnectingPart(&isWhereClauseElementFirst)

		b.where.WriteString(fmt.Sprintf(`acos(
			sin(radians(%f))
				* sin(radians(lat))
			+ cos(radians(%f))
				* cos(radians(lat))
				* cos( radians(%f)
					- radians(lng))
			) * 6371 <= 100`, params.Near[0], params.Near[0], params.Near[1]))
	}
	if params.Sort != "" {
		switch params.Sort {
		case "price":
			params.Sort = "price_per_day"
		}

		b.where.WriteString(fmt.Sprintf(" ORDER BY %s", params.Sort))
	}
	if params.Limit != nil {
		b.where.WriteString(fmt.Sprintf(" LIMIT %d", *params.Limit))
	}
	if params.Offset != nil {
		b.where.WriteString(fmt.Sprintf(" OFFSET %d", *params.Offset))
	}

	return b.where.String()
}

func (b *whereClauseBuilder) addConnectingPart(isWhereClauseElementFirst *bool) {
	if !*isWhereClauseElementFirst {
		b.where.WriteString(" AND ")
	} else {
		*isWhereClauseElementFirst = false
		b.where.WriteString(" WHERE ")
	}
}
