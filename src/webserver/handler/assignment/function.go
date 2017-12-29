package assignment

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/melodiez14/meiko/src/util/conn"

	asg "github.com/melodiez14/meiko/src/module/assignment"
	cs "github.com/melodiez14/meiko/src/module/course"
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

func handleGradeBySchedule(scheduleID, userID int64) ([]getGradeResponse, int, error) {

	resp := []getGradeResponse{}
	if !cs.IsEnrolled(userID, scheduleID) {
		return nil, http.StatusForbidden, fmt.Errorf("You are not authorized")
	}

	gps, err := cs.SelectGPBySchedule([]int64{scheduleID})
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	if len(gps) < 1 {
		return resp, http.StatusOK, nil
	}

	var gpsID []int64
	for _, gp := range gps {
		gpsID = append(gpsID, gp.ID)
	}

	assignments, err := asg.SelectByGP(gpsID, false)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	var asgID []int64
	var asgGP = make(map[int64]cs.GradeParameter)
	var asgName = make(map[int64]string)
	for _, assignment := range assignments {
		asgID = append(asgID, assignment.ID)
		asgName[assignment.ID] = assignment.Name
		for _, gp := range gps {
			if assignment.GradeParameterID == gp.ID {
				asgGP[assignment.ID] = gp
			}
		}
	}

	submitted, err := asg.SelectSubmittedByUser(asgID, userID)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	for _, val := range submitted {
		if val.Score.Valid {
			resp = append(resp, getGradeResponse{
				Name:  asgName[val.AssignmentID],
				Type:  strings.ToLower(asgGP[val.AssignmentID].Type),
				Score: fmt.Sprintf("%.3g", val.Score.Float64),
			})
		}
	}

	return resp, http.StatusOK, nil
}
