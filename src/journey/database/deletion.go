package database

const stmtDeletePostTagsByPostId = "DELETE FROM posts_tags WHERE post_id = ?"
const stmtDeletePostById = "DELETE FROM posts WHERE id = ?"
const stmtDeletePostAuthorById = "DELETE FROM post_authors WHERE post_id = ? AND author_id = ?"
const stmtDeleteUserById = "DELETE FROM users WHERE id = ?"

func DeletePostTagsForPostId(post_id int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostTagsByPostId, post_id)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

func DeletePostById(id int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostById, id)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

func DeleteUserById(id int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeleteUserById, id)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

func DeletePostAuthor(post_id int64, author_id int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostAuthorById, post_id, author_id)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}
