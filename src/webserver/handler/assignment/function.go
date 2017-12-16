package assignment

import (
	"database/sql"
	"strconv"

	"github.com/melodiez14/meiko/src/util/conn"

	asg "github.com/melodiez14/meiko/src/module/assignment"
	fl "github.com/melodiez14/meiko/src/module/file"
	"github.com/melodiez14/meiko/src/util/helper"
)

func handleSubmitInsert(id, userID int64, desc sql.NullString, fileID []string) error {
	tableID := strconv.FormatInt(id, 10)

	tx := conn.DB.MustBegin()

	// update assignment
	err := asg.InsertSubmit(id, userID, desc, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert new file
	for _, val := range fileID {
		err = fl.UpdateRelation(val, fl.TypAssignmentUpload, tableID, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func handleSubmitUpdate(id, userID int64, desc sql.NullString, fileID []string) error {
	tableID := strconv.FormatInt(id, 10)
	oldFile, err := fl.SelectIDByRelation(fl.TypAssignmentUpload, tableID, userID)
	if err != nil {
		return err
	}

	var insert []string
	for _, val := range fileID {
		if !helper.IsStringInSlice(val, oldFile) {
			insert = append(insert, val)
		}
	}

	var delete []string
	for _, val := range oldFile {
		if !helper.IsStringInSlice(val, fileID) {
			delete = append(delete, val)
		}
	}

	tx := conn.DB.MustBegin()

	// update assignment
	err = asg.UpdateSubmit(id, userID, desc, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert new file
	for _, val := range fileID {
		if !helper.IsStringInSlice(val, oldFile) {
			if err = fl.UpdateRelation(val, fl.TypAssignmentUpload, tableID, tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// delete old file
	for _, val := range oldFile {
		if !helper.IsStringInSlice(val, fileID) {
			if err = fl.Delete(val, tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	return nil
}
