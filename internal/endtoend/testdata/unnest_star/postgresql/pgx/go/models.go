// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import ()

type Item struct {
	ItemID int64
}

type Plan struct {
	PlanID int64
}

type PlanItem struct {
	PlanItemID int64
	PlanID     int64
	ItemID     int64
}
