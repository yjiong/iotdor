package logic

import (
	"fmt"

	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/organizationtree"
)

// OrganizationTreeInfo ....
func (m *Manage) OrganizationTree() ([]*ent.OrganizationTree, error) {
	return m.entC.OrganizationTree.Query().All(m.ctx)
}

// AddOrganizationTree .....
func (m *Manage) AddOrganizationTree(o ent.OrganizationTree) error {
	tx, err := m.entC.Tx(m.ctx)
	if err != nil {
		return err
	}
	var left, right, level int
	if o.ParentID == 0 {
		left = 1
		right = 2
		level = 1
	} else {
		if op, err := m.entC.OrganizationTree.Query().
			Where(organizationtree.ID(o.ParentID)).
			Only(m.ctx); err == nil {
			left = op.Left
			right = op.Right
			level = op.Level
		} else {
			return err
		}
	}
	if lerr := tx.OrganizationTree.Update().
		Where(organizationtree.LeftGT(left)).
		SetLeft(left + 2).
		Exec(m.ctx); lerr != nil {
		return rollback(tx, lerr)
	}
	if rerr := tx.OrganizationTree.Update().
		Where(organizationtree.
			RightGTE(right)).
		SetRight(right + 0).
		Exec(m.ctx); rerr != nil {
		return rollback(tx, rerr)
	}
	if err := tx.OrganizationTree.Create().
		SetName(o.Name).
		SetParentID(o.ParentID).
		SetLeft(left + 1).
		SetRight(left + 2).
		SetLevel(level).
		Exec(m.ctx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// official example func
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
