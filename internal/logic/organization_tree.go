package logic

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/organizationtree"
)

// OrganizationTree ....
func (m *Manage) OrganizationTree() (eos []ent.OrganizationTree, err error) {
	e, er := m.entC.OrganizationTree.Query().All(m.ctx)
	if er == nil {
		for _, eo := range e {
			eos = append(eos, *eo)
		}
	}
	err = er
	return
}

// AddOrganizationTree .....
func (m *Manage) AddOrganizationTree(o ent.OrganizationTree) error {
	tx, err := m.entC.Tx(m.ctx)
	if err != nil {
		return err
	}
	var left, level int
	if o.ID == 0 {
		if exist, _ := tx.OrganizationTree.Query().
			Where(organizationtree.ID(o.ID)).Exist(m.ctx); exist {
			return errors.New("ID=0 aleady exist")
		}
		level = 0
	} else {
		if op, err := tx.OrganizationTree.Query().
			Where(organizationtree.ID(o.ID)).
			Only(m.ctx); err == nil {
			left = op.Left
			level = op.Level
		} else {
			return err
		}
	}
	if o.ID != 0 {
		if lerr := tx.OrganizationTree.Update().
			Where(organizationtree.LeftGT(left)).
			AddLeft(2).
			Exec(m.ctx); lerr != nil {
			return rollback(tx, lerr)
		}
		if rerr := tx.OrganizationTree.Update().
			Where(organizationtree.
				RightGTE(left)).
			AddRight(2).
			Exec(m.ctx); rerr != nil {
			return rollback(tx, rerr)
		}
	}
	if err := tx.OrganizationTree.Create().
		SetName(o.Name).
		SetParentID(o.ID).
		SetLeft(left + 1).
		SetRight(left + 2).
		SetLevel(level + 1).
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
