package repository

import (
	"blog/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

func (r *EssayRepoMySQL) insertLabels(tx *sqlx.Tx, eid int, lIDs []int) error {
	// 构建批量插入的SQL和参数
	sqlStr := `
        INSERT INTO eid_lid (essay_id, label_id)
        SELECT ?, l.id 
        FROM e_label l
        WHERE l.id IN (?)
    `

	// 通过sqlx.In来处理IN查询
	query, args, err := sqlx.In(sqlStr, eid, lIDs)
	if err != nil {
		return err
	}

	// 将SQL转换为底层数据库驱动可执行的格式
	query = tx.Rebind(query)

	result, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	// 检查影响的行数是否符合预期
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// 如果影响的行数与传入的标签ID数量不匹配,说明有无效的标签ID
	if int(affected) != len(lIDs) {
		return fmt.Errorf(invalidLabelIds)
	}

	return nil
}

func (r *EssayRepoMySQL) readLabels(eid int) (data []models.EssayLabelData, err error) {
	data = make([]models.EssayLabelData, 0, 10)
	sqlStr := `
		SELECT el.label_id AS id ,l.name as name
		FROM eid_lid el
		LEFT OUTER JOIN e_label l on l.id = el.label_id
		WHERE essay_id = ?
		`
	err = r.db.Select(&data, sqlStr, eid)
	return
}

func (r *EssayRepoMySQL) updateLabel(tx *sqlx.Tx, eid int, oLabelIds []int, labelIds []int) error {
	// 1. 分治：获取需要添加和删除的标签
	var needAddLabelsIds = make([]int, 0, 5)
	var needRemoveLabelsIds = make([]int, 0, 5)
	var oldLabelMap = make(map[int]bool, 5)
	var newLabelMap = make(map[int]bool, 5)

	// 构建新旧标签的map
	for _, id := range oLabelIds {
		oldLabelMap[id] = true
	}

	for _, id := range labelIds {
		newLabelMap[id] = true
	}

	// 找出需要删除和添加的标签
	for id := range oldLabelMap {
		if !newLabelMap[id] {
			needRemoveLabelsIds = append(needRemoveLabelsIds, id)
		}
	}
	for id := range newLabelMap {
		if !oldLabelMap[id] {
			needAddLabelsIds = append(needAddLabelsIds, id)
		}
	}

	// 2. 批量删除标签
	if len(needRemoveLabelsIds) > 0 {
		deleteSql := `DELETE FROM eid_lid WHERE essay_id = ? AND label_id IN (?)`
		query, args, err := sqlx.In(deleteSql, eid, needRemoveLabelsIds)
		if err != nil {
			return fmt.Errorf("construct delete query failed,err: %w", err)
		}
		query = tx.Rebind(query) // 重要：处理不同数据库的占位符
		_, err = tx.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("delete tags in batches failed,err: %w", err)
		}
	}

	// 3. 批量添加标签
	if len(needAddLabelsIds) > 0 {
		// 构造批量插入的VALUES部分
		valueStrings := make([]string, 0, len(needAddLabelsIds))
		valueArgs := make([]interface{}, 0, len(needAddLabelsIds)*2)

		for _, labelID := range needAddLabelsIds {
			valueStrings = append(valueStrings, "(?,?)")
			valueArgs = append(valueArgs, eid, labelID)
		}

		// 使用正确的SQL语法格式（VALUES复数形式）以及通过占位符来构建参数化查询的语句
		addSql := fmt.Sprintf(`
        INSERT INTO eid_lid (essay_id, label_id) 
        VALUES %s`, strings.Join(valueStrings, ","))

		// 使用tx.Prepare来预编译SQL语句，然后通过stmt.Exec结合参数来执行插入操作，确保安全并正确地批量插入数据
		stmt, err := tx.Prepare(addSql)
		if err != nil {
			return fmt.Errorf("prepare insert statement failed: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(valueArgs...)
		if err != nil {
			return fmt.Errorf("add tags in batches failed: %w", err)
		}
	}
	return nil
}

func (r *EssayRepoMySQL) deleteLabels(tx *sqlx.Tx, eid int) error {
	sqlStr := `DELETE FROM eid_lid WHERE essay_id = ?`
	if _, err := tx.Exec(sqlStr, eid); err != nil {
		return err
	}
	return nil
}
