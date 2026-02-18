package permission

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(name string, description string, resources string, action string) error
	GetByID(id string) (*Permission, error)
	GetAll() ([]*Permission, error)
	Update(id string, name *string, description *string, resources *string, action *string) (string, error)
	SoftDelete(id string) (string, error)
	HardDelete(id string) (string, error)

	GetByName(name string) (*Permission, error)
}

type PermissionRepositoryImpl struct {
	// Add fields for database connection, etc.
	db *gorm.DB
}

func NewPermissionRepository(_db *gorm.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (u *PermissionRepositoryImpl) Create(name string, description string, resources string, action string) error {
	fmt.Println("creating permission in permission repository.")

	// step 0: create a permission instance
	// permission := &models.Permission{
	// 	Name:     name,
	// 	Description:    description,
	// }

	// step 1: prepare the query
	query := "INSERT INTO permissions (name, description, resources, action) VALUES (?, ?, ?, ?)"

	// step 2: execute the query
	result := u.db.Exec(query, name, description, resources, action)

	// step 3: check for errors
	if result.Error != nil {
		// fmt.Printf("Error creating permission: %v\n", result.Error)
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
		fmt.Println("No permission was created.")
		return nil
	}

	fmt.Printf("Created permission (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return nil
}

func (u *PermissionRepositoryImpl) GetByID(id string) (*Permission, error) {
	fmt.Println("Fetching permission by id in permission repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, resources, action, created_at, updated_at FROM permissions WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	row := u.db.Raw(query, id).Row()

	// step 3: process the result
	permission := &Permission{}
	err := row.Scan(&permission.ID, &permission.Name, &permission.Description, &permission.Resources, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Permission not found.")
			return nil, err
		}
		fmt.Printf("Error fetching permission: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched permission: %+v\n", permission)
	return permission, nil
}

func (u *PermissionRepositoryImpl) GetAll() ([]*Permission, error) {
	fmt.Println("Fetching all permissions in permission repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, resources, action, created_at, updated_at FROM permissions WHERE deleted_at IS NULL"

	// step 2: execute the query
	rows, err := u.db.Raw(query).Rows()
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// another way

	// step 2: execute the query with Raw() - NOT Exec()
	// var permissions []*models.Permission
	// result := u.db.Raw(query).Scan(&permissions)

	// step 3: check for errors
	// if result.Error != nil {
	//     fmt.Printf("Error fetching permissions: %v\n", result.Error)
	//     return nil, result.Error
	// }

	// step 4: process the result
	var permissions []*Permission
	for rows.Next() {
		var permission Permission
		err := u.db.ScanRows(rows, &permission)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		permissions = append(permissions, &permission)
	}

	// step 5: return the result
	for _, permission := range permissions {
		fmt.Println(permission)
	}
	return permissions, nil
}

func (u *PermissionRepositoryImpl) Update(id string, name *string, description *string, resources *string, action *string) (string, error) {
	fmt.Println("updating permission in permission repository.")

	// step 1: prepare the query
	query := "UPDATE permissions SET "
	args := []interface{}{}
	if name != nil {
		query += "name = ?, "
		args = append(args, *name)
	}
	if description != nil {
		query += "description = ?, "
		args = append(args, *description)
	}
	if resources != nil {
		query += "resources = ?, "
		args = append(args, *resources)
	}
	if action != nil {
		query += "action = ?, "
		args = append(args, *action)
	}

	query += "updated_at = NOW() "
	query += "WHERE deleted_at IS NULL AND id = ?"
	args = append(args, id)

	// step 2: execute the query
	result := u.db.Exec(query, args...)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error updating permission: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No permission was updated.")
		return "", fmt.Errorf("No permission was updated.")
	}

	fmt.Printf("Updated permission (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Permission updated successfully (rows affected: %d)", rowsAffected), nil
}

func (u *PermissionRepositoryImpl) SoftDelete(id string) (string, error) {
	fmt.Println("deleting permission in permission repository.")

	// step 1: prepare the query
	query := "UPDATE permissions SET deleted_at = NOW() WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting permission: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No permission was deleted.")
		return "", fmt.Errorf("No permission was deleted.")
	}

	fmt.Printf("Deleted permission (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted permission (rows affected: %d)\n", rowsAffected), nil
}

func (u *PermissionRepositoryImpl) HardDelete(id string) (string, error) {
	fmt.Println("deleting permission in permission repository.")

	// step 1: prepare the query
	query := "DELETE FROM permissions WHERE id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting permission: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No permission was deleted.")
		return "", fmt.Errorf("No permission was deleted.")
	}

	fmt.Printf("Deleted permission (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted permission (rows affected: %d)\n", rowsAffected), nil
}

func (u *PermissionRepositoryImpl) GetByName(name string) (*Permission, error) {
	fmt.Println("Fetching permission by id in permission repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, resources, action, created_at, updated_at FROM permissions WHERE deleted_at IS NULL AND name = ?"

	// step 2: execute the query
	row := u.db.Raw(query, name).Row()

	// step 3: process the result
	permission := &Permission{}
	err := row.Scan(&permission.ID, &permission.Name, &permission.Description, &permission.Resources, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Permission not found.")
			return nil, err
		}
		fmt.Printf("Error fetching permission: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched permission: %+v\n", permission)
	return permission, nil
}
