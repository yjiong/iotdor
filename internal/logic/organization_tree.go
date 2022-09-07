package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/organizationposition"
	"github.com/yjiong/iotdor/ent/organizationtree"
)

// OrganizationTree ....
func (m *Manage) OrganizationTree() (eos []*ent.OrganizationTree, err error) {
	return m.entC.OrganizationTree.Query().
		Order(ent.Asc(organizationtree.FieldLeft)).
		All(m.ctx)
}

// ListOrganizationTreePositions ....
func (m *Manage) ListOrganizationTreePositions(eo *ent.OrganizationTree) (eos []*ent.OrganizationPosition, err error) {
	return eo.QueryOrganizationPositions().All(m.ctx)
}

// SubOrganizationFromID ....
func (m *Manage) SubOrganizationFromID(id int) (eos []*ent.OrganizationTree, err error) {
	var left, right int
	if op, err := m.entC.OrganizationTree.Query().
		Where(organizationtree.ID(id)).
		Only(m.ctx); err == nil {
		left = op.Left
		right = op.Right
	} else {
		return nil, err
	}
	return m.entC.OrganizationTree.Query().
		Where(organizationtree.LeftGT(left), organizationtree.RightLT(right)).
		Order(ent.Asc(organizationtree.FieldLeft)).
		All(m.ctx)
}

// AddOrganizationNode  add for same level at courrent node left or right
func (m *Manage) AddOrganizationNode(id int, name, leftOrRight string) error {
	tx, err := m.entC.Tx(m.ctx)
	if err != nil {
		return err
	}
	var left, right, handleLeft int
	if op, err := tx.OrganizationTree.Query().
		Where(organizationtree.ID(id)).
		Only(m.ctx); err == nil {
		left = op.Left
		right = op.Right
		handleLeft = left - 1
		if leftOrRight == "right" {
			left = right + 1
			handleLeft = right
		} else if leftOrRight != "left" {
			return errors.New("argument must = 'left' or 'right'")
		}
		if err := handleLeftRight(m.ctx, handleLeft, 2, tx); err != nil {
			return err
		}
		if err := tx.OrganizationTree.Create().
			SetParentID(op.ParentID).
			SetLeft(left).
			SetRight(left + 1).
			SetLevel(op.Level).
			SetName(name).Exec(m.ctx); err != nil {
			return rollback(tx, err)
		}
	} else {
		return err
	}
	return tx.Commit()
}

// CreateOrganizationSubNode .....
func (m *Manage) CreateOrganizationSubNode(o ent.OrganizationTree) error {
	tx, err := m.entC.Tx(m.ctx)
	if err != nil {
		return err
	}
	var left, level int
	if o.ParentID == 0 {
		if exist, _ := tx.OrganizationTree.Query().
			Where(organizationtree.Level(1)).Exist(m.ctx); exist {
			return errors.New("level=1 aleady exist")
		}
		level = 0
	} else {
		if op, err := tx.OrganizationTree.Query().
			Where(organizationtree.ID(o.ParentID)).
			Only(m.ctx); err == nil {
			left = op.Left
			level = op.Level
			if op.Right > (op.Left + 1) {
				return fmt.Errorf("subnode level=%d aleady exist, use add api to add same level node", level+1)
			}
		} else {
			return err
		}
		if err := handleLeftRight(m.ctx, left, 2, tx); err != nil {
			return err
		}
	}
	if err := tx.OrganizationTree.Create().
		SetName(o.Name).
		SetParentID(o.ParentID).
		SetLeft(left + 1).
		SetRight(left + 2).
		SetLevel(level + 1).
		Exec(m.ctx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func handleLeftRight(ctx context.Context, left, amount int, tx *ent.Tx) error {
	if lerr := tx.OrganizationTree.Update().
		Where(organizationtree.LeftGT(left)).
		AddLeft(amount).
		Exec(ctx); lerr != nil {
		return rollback(tx, lerr)
	}
	if rerr := tx.OrganizationTree.Update().
		Where(organizationtree.
			RightGT(left)).
		AddRight(amount).
		Exec(ctx); rerr != nil {
		return rollback(tx, rerr)
	}
	return nil
}

// UpdateOrganizationTree ...
func (m *Manage) UpdateOrganizationTree(o ent.OrganizationTree) error {
	return m.entC.OrganizationTree.Update().
		Where(organizationtree.ID(o.ID)).SetName(o.Name).Exec(m.ctx)
}

// DeleteOrganizationTree ...
func (m *Manage) DeleteOrganizationTree(id int) error {
	tx, err := m.entC.Tx(m.ctx)
	var left, right int
	if err != nil {
		return err
	}
	if op, err := tx.OrganizationTree.Query().
		Where(organizationtree.ID(id)).
		Only(m.ctx); err == nil {
		left = op.Left
		right = op.Right
	} else {
		return err
	}
	if i, err := tx.OrganizationTree.Delete().
		Where(organizationtree.LeftGTE(left), organizationtree.RightLTE(right)).
		Exec(m.ctx); err != nil && i != (right-left+1)/2 {
		return rollback(tx, err)
	}
	if err := tx.OrganizationTree.Update().
		Where(organizationtree.LeftGTE(left)).
		AddLeft(left - right - 1).
		Exec(m.ctx); err != nil {
		return rollback(tx, err)
	}
	if err := tx.OrganizationTree.Update().
		Where(organizationtree.RightGTE(left)).
		AddRight(left - right - 1).
		Exec(m.ctx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// RelatePositioinToOranizationTree .....
func (m *Manage) RelatePositioinToOranizationTree(id, posid string) error {
	idi, _ := strconv.Atoi(id)
	po, err := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	if po != nil && err == nil {
		return m.entC.OrganizationTree.Update().
			Where(organizationtree.ID(idi)).
			SetOrganizationPositionsID(po.ID).
			Exec(m.ctx)
	}
	return errors.Wrap(err, fmt.Sprintf("%s not exist", posid))
}

// RemovePositioinFromOranizationTree .....
func (m *Manage) RemovePositioinFromOranizationTree(id, posid string) error {
	idi, _ := strconv.Atoi(id)
	po, err := m.entC.OrganizationPosition.Query().Where(organizationposition.PositionID(posid)).Only(m.ctx)
	if po != nil && err == nil {
		return m.entC.OrganizationPosition.Update().
			Where(organizationposition.ID(po.ID)).
			RemoveOrganizationTreeIDs(idi).
			Exec(m.ctx)
	}
	return errors.Wrap(err, fmt.Sprintf("%s not exist", posid))
}

// official rollback example func
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
