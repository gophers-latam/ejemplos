/*
El paquete funciona en 2 tablas en una base de datos SQLite.

Los nombres de las tablas son:

  - users
  - userdata

Las definiciones de las tablas son:

	    CREATE TABLE users (
	        id INTEGER PRIMARY KEY,
	        username TEXT
	    );

	    CREATE TABLE userdata (
	        userid INTEGER NOT NULL,
	        name TEXT,
	        surname TEXT,
	        description TEXT
	    );

			Esto se representa como código

Esto no se representa como código
*/
package document

// BUG(1): Function ListUsers() no funciona como se esperaba
// BUG(2): Function AddUser() es demasiado lento

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

/*
Esta variable global contiene la ruta del archivo de la base de datos SQLite3.

	Filename: En la ruta al archivo de base de datos
*/
var (
	Filename = "" // "database filepath"
)

// Userdata es para contener datos completos del usuario
// de la tabla userdata y el username de la
// tabla users
type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

// openConnection() es para abrir la conexión SQLite3
// para ser utilizado por las otras funciones del paquete.
func openConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", Filename)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// La función devuelve el ID de usuario del username.
// -1 si el usuario no existe
func exists(username string) int {
	username = strings.ToLower(username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := -1
	stmt := fmt.Sprintf(`SELECT ID FROM users where username = '%s'`, username)
	rows, err := db.Query(stmt)

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("exists() Scan", err)
			return -1
		}
		userID = id
	}
	defer rows.Close()
	return userID
}

// AddUser agrega un nuevo usuario a la base de datos
// Devuelve un nuevo ID de usuario
// -1 si hubo un error
func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", d.Username)
		return -1
	}

	stmt := `INSERT INTO users values (NULL,?)`
	_, err = db.Exec(stmt, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	userID = exists(d.Username)
	if userID == -1 {
		return userID
	}

	stmt = `INSERT INTO userdata values (?, ?, ?, ?)`
	_, err = db.Exec(stmt, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}

	return userID
}

/*
DeleteUser elimina un usuario si existe.
Requiere el ID del usuario a ser eliminado.

Retorna error
*/
func DeleteUser(id int) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// el ID existe?
	stmt := fmt.Sprintf(`SELECT username FROM users WHERE id = %d`, id)
	rows, err := db.Query(stmt)

	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return err
		}
	}
	defer rows.Close()

	if exists(username) != id {
		return fmt.Errorf("User with ID %d does not exist", id)
	}

	// eliminar de userdata
	deleteStmt := `DELETE FROM userdata WHERE userid = ?`
	_, err = db.Exec(deleteStmt, id)
	if err != nil {
		return err
	}

	// eliminar de users
	deleteStmt = `DELETE from users where id = ? `
	_, err = db.Exec(deleteStmt, id)
	if err != nil {
		return err
	}

	return nil
}

/*
ListUsers() mostrar todos los usuarios en la base de datos.
Retorna un slice de datos de usuario a llamada de la función.
*/
func ListUsers() ([]Userdata, error) {
	// Data contiene los registros devueltos por la consulta SQL
	Data := []Userdata{}

	db, err := openConnection()
	if err != nil {
		return Data, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, username, name, surname, description
		FROM users, userdata WHERE users.id = userdata.userid`)
	if err != nil {
		return Data, err
	}

	for rows.Next() {
		var (
			id       int
			username string
			name     string
			surname  string
			desc     string
		)
		err = rows.Scan(&id, &username, &name, &surname, &desc)
		temp := Userdata{ID: id, Username: username, Name: name, Surname: surname, Description: desc}
		Data = append(Data, temp)
		if err != nil {
			return Data, err
		}
	}
	defer rows.Close()
	return Data, nil
}

/*
UpdateUser() es para actualizar un usuario existente
dada una estructura Userdata.
Se toma el ID de usuario a encontrar para actualizar
dentro de la función.
*/
func UpdateUser(d Userdata) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID == -1 {
		return errors.New("User does not exist")
	}
	d.ID = userID
	stmt := `UPDATE userdata set name = ?, surname = ?, description = ? where userid = ?`
	_, err = db.Exec(stmt, d.Name, d.Surname, d.Description, d.ID)
	if err != nil {
		return err
	}

	return nil
}
