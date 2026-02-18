package role

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(name string, description string) error
	GetByID(id string) (*Role, error)
	GetAll() ([]*Role, error)
	Update(id string, name *string, description *string) (string, error)
	SoftDelete(id string) (string, error)
	HardDelete(id string) (string, error)

	GetByName(name string) (*Role, error)
}

type RoleRepositoryImpl struct {
	// Add fields for database connection, etc.
	db *gorm.DB
}

func NewRoleRepository(_db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (u *RoleRepositoryImpl) Create(name string, description string) error {
	fmt.Println("creating role in role repository.")

	// step 0: create a role instance
	// role := &models.Role{
	// 	Name:     name,
	// 	Description:    description,
	// }

	// step 1: prepare the query
	query := "INSERT INTO roles (name, description) VALUES (?, ?)"

	// step 2: execute the query
	result := u.db.Exec(query, name, description)

	// step 3: check for errors
	if result.Error != nil {
		// fmt.Printf("Error creating role: %v\n", result.Error)
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			switch pgErr.Code {
			case "23505": // unique_violation
				return fmt.Errorf("unique constraint violation")
			case "23503": // foreign_key_violation
				return fmt.Errorf("foreign key violation.")
			case "23502": // not_null_violation
				return fmt.Errorf("not null violation.")
			default:
				return fmt.Errorf("database error: %v", pgErr.Message)
			}
		}
		return result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No role was created.")
		return nil
	}

	fmt.Printf("Created role (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return nil
}

func (u *RoleRepositoryImpl) GetByID(id string) (*Role, error) {
	fmt.Println("Fetching role by id in role repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	row := u.db.Raw(query, id).Row()

	// step 3: process the result
	role := &Role{}
	err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Role not found.")
			return nil, err
		}
		fmt.Printf("Error fetching role: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched role: %+v\n", role)
	return role, nil
}

func (u *RoleRepositoryImpl) GetAll() ([]*Role, error) {
	fmt.Println("Fetching all roles in role repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE deleted_at IS NULL"

	// step 2: execute the query
	rows, err := u.db.Raw(query).Rows()
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// another way

	// step 2: execute the query with Raw() - NOT Exec()
	// var roles []*models.Role
	// result := u.db.Raw(query).Scan(&roles)

	// step 3: check for errors
	// if result.Error != nil {
	//     fmt.Printf("Error fetching roles: %v\n", result.Error)
	//     return nil, result.Error
	// }

	// step 4: process the result
	var roles []*Role
	for rows.Next() {
		var role Role
		err := u.db.ScanRows(rows, &role)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		roles = append(roles, &role)
	}

	// step 5: return the result
	for _, role := range roles {
		fmt.Println(role)
	}
	return roles, nil
}

func (u *RoleRepositoryImpl) Update(id string, name *string, description *string) (string, error) {
	fmt.Println("updating role in role repository.")

	// step 1: prepare the query
	query := "UPDATE roles SET "
	args := []interface{}{}
	if name != nil {
		query += "name = ?, "
		args = append(args, *name)
	}
	if description != nil {
		query += "description = ?, "
		args = append(args, *description)
	}
	query += "updated_at = NOW() "
	query += "WHERE deleted_at IS NULL AND id = ?"
	args = append(args, id)

	// step 2: execute the query
	result := u.db.Exec(query, args...)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error updating role: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No role was updated.")
		return "", fmt.Errorf("No role was updated.")
	}

	fmt.Printf("Updated role (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Role updated successfully (rows affected: %d)", rowsAffected), nil
}

func (u *RoleRepositoryImpl) SoftDelete(id string) (string, error) {
	fmt.Println("deleting role in role repository.")

	// step 1: prepare the query
	query := "UPDATE roles SET deleted_at = NOW() WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting role: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No role was deleted.")
		return "", fmt.Errorf("No role was deleted.")
	}

	fmt.Printf("Deleted role (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted role (rows affected: %d)\n", rowsAffected), nil
}

func (u *RoleRepositoryImpl) HardDelete(id string) (string, error) {
	fmt.Println("deleting role in role repository.")

	// step 1: prepare the query
	query := "DELETE FROM roles WHERE id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting role: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No role was deleted.")
		return "", fmt.Errorf("No role was deleted.")
	}

	fmt.Printf("Deleted role (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted role (rows affected: %d)\n", rowsAffected), nil
}


func (u *RoleRepositoryImpl) GetByName(name string) (*Role, error) {
	fmt.Println("Fetching role by id in role repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE deleted_at IS NULL AND name = ?"

	// step 2: execute the query
	row := u.db.Raw(query, name).Row()

	// step 3: process the result
	role := &Role{}
	err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Role not found.")
			return nil, err
		}
		fmt.Printf("Error fetching role: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched role: %+v\n", role)
	return role, nil
}