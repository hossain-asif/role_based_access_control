package rolepermission

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	Create(roleID string, permissionID string) error
	GetByID(id string) (*RolePermission, error)
	GetAll() ([]*RolePermission, error)
	Update(id string, roleID *string, permissionID *string) (string, error)
	SoftDelete(id string) (string, error)
	HardDelete(id string) (string, error)

	GetRolePermissionById(id int64) (*RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*RolePermission, error)
}

type RolePermissionRepositoryImpl struct {
	// Add fields for database connection, etc.
	db *gorm.DB
}

func NewRolePermissionRepository(_db *gorm.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{
		db: _db,
	}
}

func (u *RolePermissionRepositoryImpl) Create(roleID string, permissionID string) error {
	fmt.Println("creating rolePermission in rolePermission repository.")

	// step 0: create a rolePermission instance
	// rolePermission := &models.RolePermission{
	// 	Name:     name,
	// 	Description:    description,
	// }

	// step 1: prepare the query
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)"

	// step 2: execute the query
	result := u.db.Exec(query, roleID, permissionID)

	// step 3: check for errors
	if result.Error != nil {
		// fmt.Printf("Error creating rolePermission: %v\n", result.Error)
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
		fmt.Println("No rolePermission was created.")
		return nil
	}

	fmt.Printf("Created rolePermission (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return nil
}

func (u *RolePermissionRepositoryImpl) GetByID(id string) (*RolePermission, error) {
	fmt.Println("Fetching rolePermission by id in rolePermission repository.")

	// step 1: prepare the query
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	row := u.db.Raw(query, id).Row()

	// step 3: process the result
	rolePermission := &RolePermission{}
	err := row.Scan(&rolePermission.ID, &rolePermission.PermissionID, &rolePermission.RoleID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("RolePermission not found.")
			return nil, err
		}
		fmt.Printf("Error fetching rolePermission: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched rolePermission: %+v\n", rolePermission)
	return rolePermission, nil
}

func (u *RolePermissionRepositoryImpl) GetAll() ([]*RolePermission, error) {
	fmt.Println("Fetching all rolePermissions in rolePermission repository.")

	// step 1: prepare the query
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE deleted_at IS NULL"

	// step 2: execute the query
	rows, err := u.db.Raw(query).Rows()
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// another way

	// step 2: execute the query with Raw() - NOT Exec()
	// var rolePermissions []*models.RolePermission
	// result := u.db.Raw(query).Scan(&rolePermissions)

	// step 3: check for errors
	// if result.Error != nil {
	//     fmt.Printf("Error fetching rolePermissions: %v\n", result.Error)
	//     return nil, result.Error
	// }

	// step 4: process the result
	var rolePermissions []*RolePermission
	for rows.Next() {
		var rolePermission RolePermission
		err := u.db.ScanRows(rows, &rolePermission)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		rolePermissions = append(rolePermissions, &rolePermission)
	}

	// step 5: return the result
	for _, rolePermission := range rolePermissions {
		fmt.Println(rolePermission)
	}
	return rolePermissions, nil
}

func (u *RolePermissionRepositoryImpl) Update(id string, roleID *string, permissionID *string) (string, error) {
	fmt.Println("updating rolePermission in rolePermission repository.")

	// step 1: prepare the query
	query := "UPDATE role_permissions SET "
	args := []interface{}{}
	if roleID != nil {
		query += "role_id = ?, "
		args = append(args, *roleID)
	}
	if permissionID != nil {
		query += "permission_id = ?, "
		args = append(args, *permissionID)
	}

	query += "updated_at = NOW() "
	query += "WHERE deleted_at IS NULL AND id = ?"
	args = append(args, id)

	// step 2: execute the query
	result := u.db.Exec(query, args...)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error updating rolePermission: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No rolePermission was updated.")
		return "", fmt.Errorf("No rolePermission was updated.")
	}

	fmt.Printf("Updated rolePermission (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("RolePermission updated successfully (rows affected: %d)", rowsAffected), nil
}

func (u *RolePermissionRepositoryImpl) SoftDelete(id string) (string, error) {
	fmt.Println("deleting rolePermission in rolePermission repository.")

	// step 1: prepare the query
	query := "UPDATE role_permissions SET deleted_at = NOW() WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting rolePermission: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No rolePermission was deleted.")
		return "", fmt.Errorf("No rolePermission was deleted.")
	}

	fmt.Printf("Deleted rolePermission (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted rolePermission (rows affected: %d)\n", rowsAffected), nil
}

func (u *RolePermissionRepositoryImpl) HardDelete(id string) (string, error) {
	fmt.Println("deleting rolePermission in rolePermission repository.")

	// step 1: prepare the query
	query := "DELETE FROM role_permissions WHERE id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting rolePermission: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No rolePermission was deleted.")
		return "", fmt.Errorf("No rolePermission was deleted.")
	}

	fmt.Printf("Deleted rolePermission (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted rolePermission (rows affected: %d)\n", rowsAffected), nil
}



// role-permission related actions
func (u *RolePermissionRepositoryImpl) GetRolePermissionById(id int64) (*RolePermission, error) {
	return nil, nil
}

func (u *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*RolePermission, error) {
	return nil, nil
}

func (u *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) (*RolePermission, error) {
	return nil, nil
}

func (u *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return nil
}

func (u *RolePermissionRepositoryImpl) GetAllRolePermissions() ([]*RolePermission, error) {
	return nil, nil
}
