package database

// type DB struct {
// 	Collection *sql.DB
// }

// func Connect() (*DB, error) {
// 	db, err := sql.Open("sqlite3", "file:./database/database.db?_auth&_auth_user=admin&_auth_pass=admin&_auth_crypt=sha1")
// 	if err != nil {
// 		return nil, fmt.Errorf("couldn't open database due to %v", err)
// 	}
// 	log.Println("| | preparing database tables...")
// 	if _, err := prepareTables(db); err != nil {
// 		return nil, err
// 	}
// 	log.Println("| | database check done!")

// 	return &DB{
// 		db,
// 	}, nil
// }

// func prepareTables(db *sql.DB) (sql.Result, error) {
// 	arr := []string{users_table, sessions_table, posts_table, categories_table, comment_table, post_reaction_table, comment_reaction_table}

// 	for i := 0; i < len(arr); i++ {
// 		st, err := db.Prepare(arr[i])
// 		if err != nil {
// 			return nil, fmt.Errorf("couldn't create new table due to %v", err)
// 		}
// 		if i == len(arr)-1 {
// 			return st.Exec()
// 		}
// 		st.Exec()
// 	}

// 	return nil, nil
// }
