package userrole

import (
	"errors"
	"fmt"
	"go_project_structure/internal/permission"
	"go_project_structure/internal/role"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(userID string, roleID string) error
	GetByID(id string) (*UserRole, error)
	GetAll() ([]*UserRole, error)
	Update(id string, userID *string, roleID *string) (string, error)
	SoftDelete(id string) (string, error)
	HardDelete(id string) (string, error)

	
	GetUserRoles(userId int64) ([]*role.Role, error)
	AssignRoleToUser(userId int64, roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*permission.Permission, error)
	HasPermission(userId int64, permissionName string) (bool, error)
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)

}

type UserRoleRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRoleRepository(_db *gorm.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func (u *UserRoleRepositoryImpl) Create(userID string, roleID string) error {
	fmt.Println("creating userRole in userRole repository.")

	// step 0: create a userRole instance
	// userRole := &models.UserRole{
	// 	Name:     name,
	// 	Description:    description,
	// }

	// step 1: prepare the query
	query := "INSERT INTO user_roles (name, description, resources, action) VALUES (?, ?, ?, ?)"

	// step 2: execute the query
	result := u.db.Exec(query, userID, roleID)

	// step 3: check for errors
	if result.Error != nil {
		// fmt.Printf("Error creating userRole: %v\n", result.Error)
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
		fmt.Println("No userRole was created.")
		return nil
	}

	fmt.Printf("Created userRole (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return nil
}

func (u *UserRoleRepositoryImpl) GetByID(id string) (*UserRole, error) {
	fmt.Println("Fetching userRole by id in userRole repository.")

	// step 1: prepare the query
	query := "SELECT id, user_id, role_id FROM user_roles WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	row := u.db.Raw(query, id).Row()

	// step 3: process the result
	userRole := &UserRole{}
	err := row.Scan(&userRole.ID, &userRole.UserID, userRole.RoleID, &userRole.CreatedAt, &userRole.UpdatedAt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("UserRole not found.")
			return nil, err
		}
		fmt.Printf("Error fetching userRole: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched userRole: %+v\n", userRole)
	return userRole, nil
}

func (u *UserRoleRepositoryImpl) GetAll() ([]*UserRole, error) {
	fmt.Println("Fetching all userRoles in userRole repository.")

	// step 1: prepare the query
	query := "SELECT id, user_id, role_id FROM user_roles WHERE deleted_at IS NULL"

	// step 2: execute the query
	rows, err := u.db.Raw(query).Rows()
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// another way

	// step 2: execute the query with Raw() - NOT Exec()
	// var userRoles []*models.UserRole
	// result := u.db.Raw(query).Scan(&userRoles)

	// step 3: check for errors
	// if result.Error != nil {
	//     fmt.Printf("Error fetching userRoles: %v\n", result.Error)
	//     return nil, result.Error
	// }

	// step 4: process the result
	var userRoles []*UserRole
	for rows.Next() {
		var userRole UserRole
		err := u.db.ScanRows(rows, &userRole)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		userRoles = append(userRoles, &userRole)
	}

	// step 5: return the result
	for _, userRole := range userRoles {
		fmt.Println(userRole)
	}
	return userRoles, nil
}

func (u *UserRoleRepositoryImpl) Update(id string, userID *string, roleID *string) (string, error) {
	fmt.Println("updating userRole in userRole repository.")

	// step 1: prepare the query
	query := "UPDATE user_roles SET "
	args := []interface{}{}

	if userID != nil {
		query += "user_id = ?, "
		args = append(args, userID)
	}
	if roleID != nil {
		query += "role_id = ?, "
		args = append(args, roleID)
	}


	query += "updated_at = NOW() "
	query += "WHERE deleted_at IS NULL AND id = ?"
	args = append(args, id)

	// step 2: execute the query
	result := u.db.Exec(query, args...)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error updating userRole: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No userRole was updated.")
		return "", fmt.Errorf("No userRole was updated.")
	}

	fmt.Printf("Updated userRole (rows affected: %d)\n",
		rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("UserRole updated successfully (rows affected: %d)", rowsAffected), nil
}

func (u *UserRoleRepositoryImpl) SoftDelete(id string) (string, error) {
	fmt.Println("deleting userRole in userRole repository.")

	// step 1: prepare the query
	query := "UPDATE user_roles SET deleted_at = NOW() WHERE deleted_at IS NULL AND id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting userRole: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No userRole was deleted.")
		return "", fmt.Errorf("No userRole was deleted.")
	}

	fmt.Printf("Deleted userRole (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted userRole (rows affected: %d)\n", rowsAffected), nil
}

func (u *UserRoleRepositoryImpl) HardDelete(id string) (string, error) {
	fmt.Println("deleting userRole in userRole repository.")

	// step 1: prepare the query
	query := "DELETE FROM user_roles WHERE id = ?"

	// step 2: execute the query
	result := u.db.Exec(query, id)

	// step 3: check for errors
	if result.Error != nil {
		fmt.Printf("Error deleting userRole: %v\n", result.Error)
		return "", result.Error
	}

	// step 4: evaluate the result
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		fmt.Println("No userRole was deleted.")
		return "", fmt.Errorf("No userRole was deleted.")
	}

	fmt.Printf("Deleted userRole (rows affected: %d)\n", rowsAffected)

	// step 5: return the result
	return fmt.Sprintf("Deleted userRole (rows affected: %d)\n", rowsAffected), nil
}

func (u *UserRoleRepositoryImpl) GetByName(name string) (*UserRole, error) {
	fmt.Println("Fetching userRole by id in userRole repository.")

	// step 1: prepare the query
	query := "SELECT id, name, description, resources, action, created_at, updated_at FROM userRoles WHERE deleted_at IS NULL AND name = ?"

	// step 2: execute the query
	row := u.db.Raw(query, name).Row()

	// step 3: process the result
	userRole := &UserRole{}
	err := row.Scan(&userRole.ID, &userRole.UserID, userRole.RoleID, &userRole.CreatedAt, &userRole.UpdatedAt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("UserRole not found.")
			return nil, err
		}
		fmt.Printf("Error fetching userRole: %v\n", err)
		return nil, err
	}

	// step 4: return the result
	fmt.Printf("Fetched userRole: %+v\n", userRole)
	return userRole, nil
}



// user role related actions

func (u *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*role.Role, error) {
	
	return nil, nil
}

func (u *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	return nil 
}

func (u *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	return nil 
}

func (u *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*permission.Permission, error) {
	return nil, nil
}

func (u *UserRoleRepositoryImpl) HasPermission(userId int64, permissionName string) (bool, error) {
	return false, nil
}

func (u *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	return false, nil 
}

func (u *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	return false, nil
}

func (u *UserRoleRepositoryImpl) HasAnyRole(userId int64, roleNames []string) (bool, error) {
	return false, nil
}



