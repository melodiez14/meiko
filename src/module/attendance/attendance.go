package attendance

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/melodiez14/meiko/src/util/helper"

	"github.com/jmoiron/sqlx"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetMeeting(meetingNumber uint8, scheduleID int64) (Meeting, error) {

	var meeting Meeting
	query := fmt.Sprintf(`
		SELECT
			id,
			description,
			date
		FROM
			meetings
		WHERE
			number = (%d) AND
			schedules_id = (%d)
		LIMIT 1;
		`, meetingNumber, scheduleID)
	err := conn.DB.Get(&meeting, query)
	if err != nil {
		return meeting, err
	}

	return meeting, nil
}

func GetMeetingByID(id uint64) (Meeting, error) {
	var meeting Meeting
	query := fmt.Sprintf(`
		SELECT
			id,
			subject,
			number,
			description,
			date,
			schedules_id
		FROM
			meetings
		WHERE id = (%d)
		LIMIT 1`, id)
	err := conn.DB.Get(&meeting, query)
	if err != nil {
		return meeting, err
	}
	return meeting, nil
}

func IsExistMeeting(number uint8, scheduleID int64) bool {
	var x string
	query := fmt.Sprintf(`SELECT 'x' FROM meetings WHERE number = (%d) AND schedules_id = (%d) LIMIT 1`, number, scheduleID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func IsExistByMeetingID(meetingID uint64) bool {
	var x string
	query := fmt.Sprintf(`SELECT 'x' FROM attendances WHERE meetings_id = (%d) LIMIT 1`, meetingID)
	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}
	return true
}

func Insert(usersID []int64, meetingID uint64, tx *sqlx.Tx) error {

	if len(usersID) < 1 {
		return fmt.Errorf("User ID cannot be empty")
	}

	var value []string
	for _, val := range usersID {
		value = append(value, fmt.Sprintf("(%d, %d, NOW(), NOW())", meetingID, val))
	}

	queryValue := strings.Join(value, ", ")
	query := fmt.Sprintf(`
		INSERT INTO
			attendances (
				meetings_id,
				users_id,
				created_at,
				updated_at
			) VALUES %s;
	`, queryValue)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func Delete(usersID []int64, meetingID uint64, tx *sqlx.Tx) error {

	if len(usersID) < 1 {
		return fmt.Errorf("User ID cannot be empty")
	}

	queryUsers := strings.Join(helper.Int64ToStringSlice(usersID), ", ")
	query := fmt.Sprintf(`
		DELETE FROM attendances WHERE meetings_id = (%d) AND users_id IN (%s);
	`, meetingID, queryUsers)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	return nil
}

// InsertMeeting ...
func InsertMeeting(subject string, number uint8, description sql.NullString, date time.Time, scheduleID int64, tx *sqlx.Tx) (uint64, error) {

	desc := "(NULL)"
	if description.Valid {
		desc = fmt.Sprintf(`('%s')`, description.String)
	}

	query := fmt.Sprintf(`
		INSERT INTO
			meetings (
				subject,
				number,
				description,
				date,
				schedules_id,
				created_at,
				updated_at
			) VALUES (
				('%s'),
				(%d),
				%s,
				('%v'),
				(%d),
				NOW(),
				NOW()
			);
	`, subject, number, desc, date, scheduleID)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return 0, err
	}

	meetingID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Cannot get last insert id")
	}

	return uint64(meetingID), nil
}

func UpdateMeeting(id uint64, subject string, number uint8, description sql.NullString, date time.Time, tx *sqlx.Tx) error {

	desc := "(NULL)"
	if description.Valid {
		desc = fmt.Sprintf(`('%s')`, description.String)
	}

	query := fmt.Sprintf(`
		UPDATE
			meetings
		SET
			subject = ('%s'),
			number = (%d),
			description = %s,
			date = ('%v'),
			updated_at = NOW()
		WHERE
			id = (%d);
	`, subject, number, desc, date, id)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows < 1 || err != nil {
		return fmt.Errorf("No rows affected")
	}

	return nil
}

func DeleteByMeetingID(meetingID uint64, tx *sqlx.Tx) error {
	query := fmt.Sprintf(`
		DELETE FROM
			attendances
		WHERE
			meetings_id = (%d);
	`, meetingID)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows < 1 || err != nil {
		return fmt.Errorf("No rows affected")
	}

	return nil
}

func DeleteMeeting(id uint64, tx *sqlx.Tx) error {

	query := fmt.Sprintf(`
		DELETE FROM
			meetings
		WHERE
			id = (%d);
	`, id)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows < 1 || err != nil {
		return fmt.Errorf("No rows affected")
	}

	return nil
}

// SelectMeetingByPage ...
func SelectMeetingByPage(scheduleID int64, limit, offset int, isCount bool) ([]Meeting, int, error) {

	meetings := []Meeting{}
	var count int
	query := fmt.Sprintf(`
		SELECT
			m.id,
			m.number,
			m.description,
			m.subject,
			m.date,
			count(a.meetings_id) as total
		FROM attendances a
		RIGHT JOIN meetings m ON m.id = a.meetings_id
		WHERE m.schedules_id = (%d)
		GROUP BY m.id
		ORDER BY m.number ASC
		LIMIT %d OFFSET %d
	`, scheduleID, limit, offset)

	err := conn.DB.Select(&meetings, query)
	if err != nil {
		return meetings, count, err
	}

	if !isCount {
		return meetings, count, nil
	}

	query = fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM
			meetings
		WHERE
			schedules_id = (%d)
	`, scheduleID)

	err = conn.DB.Get(&count, query)
	if err != nil {
		return meetings, count, err
	}

	return meetings, count, nil
}

func SelectUserIDByMeetingID(meetingID uint64) ([]int64, error) {
	usersID := []int64{}
	query := fmt.Sprintf(`
		SELECT
			users_id
		FROM
			attendances
		WHERE
			meetings_id = (%d);	
	`, meetingID)
	err := conn.DB.Select(&usersID, query)
	if err != nil {
		return usersID, err
	}
	return usersID, nil
}

func SelectMeetingIDByScheduleID(userID, scheduleID int64) ([]int64, error) {

	var meetingsID []int64
	query := fmt.Sprintf(`
		SELECT
			id
		FROM
			meetings
		WHERE
			schedules_id = (%d);
	`, scheduleID)

	err := conn.DB.Select(&meetingsID, query)
	if err != nil {
		return meetingsID, err
	}

	return meetingsID, nil
}

func CountByUserMeeting(userID int64, meetingsID []int64) (int, error) {

	var count int
	meetingsString := helper.Int64ToStringSlice(meetingsID)
	queryMeetingsID := strings.Join(meetingsString, ", ")
	query := fmt.Sprintf(`
		SELECT
			COUNT(*)
		FROM
			attendances
		WHERE
			users_id = (%d) AND
			meetings_id IN (%s);
	`, userID, queryMeetingsID)

	err := conn.DB.Get(&count, query)
	if err != nil {
		return count, err
	}

	return count, nil
}

func CountByUserSchedule(userID int64, schedulesID []int64) (map[int64]AttendanceReport, error) {

	report := map[int64]AttendanceReport{}
	querySchID := strings.Join(helper.Int64ToStringSlice(schedulesID), ", ")
	query := fmt.Sprintf(`
		SELECT
			m.schedules_id,
			count(m.id) as meeting_total,
			count(a.meetings_id) as attendance_total
		FROM meetings m
		LEFT JOIN attendances a ON m.id = a.meetings_id AND users_id = (%d)
		WHERE m.schedules_id IN (%s)
		GROUP BY m.schedules_id
	`, userID, querySchID)
	rows, err := conn.DB.Queryx(query)
	if err != nil {
		return report, err
	}
	defer rows.Close()

	var scheduleID int64
	var meetTotal int
	var attTotal int
	for rows.Next() {
		err = rows.Scan(&scheduleID, &meetTotal, &attTotal)
		if err != nil {
			return report, err
		}
		report[scheduleID] = AttendanceReport{AttendanceTotal: attTotal, MeetingTotal: meetTotal}
	}

	return report, nil
}
