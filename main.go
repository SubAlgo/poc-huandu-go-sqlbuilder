package main

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

func main() {

	flavor := sqlbuilder.PostgreSQL
	sb := flavor.NewSelectBuilder()

	sb.From("order as o").Select("o.id, o.name")
	sb.Where(
		sb.Or(
			sb.GT("o.amount", 1000),
			sb.Equal("o.isActive", true),
		),
	)

	showOnlyPaid := true

	if showOnlyPaid {
		sb.Where(
			sb.EQ("o.paid_status", "paid"),
		)
	}

	sql, args := sb.Build()

	fmt.Println(sql)
	fmt.Println(args)

	ub := flavor.NewUpdateBuilder()

	ub = ub.Update("user as u").
		Set(
			ub.Assign("u.name", "test"),
			ub.Assign("u.age", 30),
		).Where(
		ub.EQ("u.id", "0001"),
	)

	sql, args = ub.Build()
	fmt.Println(sql)
	fmt.Println(args)

	ib := flavor.NewInsertBuilder()
	sql, args = ib.InsertInto("user as u").
		Cols("u.id", "u.name", "u.age").Values(
		"xxx-01",
		"name-01",
		30,
	).Build()
	fmt.Println(sql)
	fmt.Println(args)

	d := flavor.NewDeleteBuilder()
	sql, args = d.DeleteFrom("user as u").Where(d.Between("u.point", 20, 40)).Build()
	fmt.Println(sql)
	fmt.Println(args)
}
