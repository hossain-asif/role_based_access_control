-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO permissions (name, description, resource, action) VALUES
('user:create', 'Create user', 'user', 'create'),
('user:read', 'Read user', 'user', 'read'),
('user:update', 'Update user', 'user', 'update'),
('user:delete', 'Delete user', 'user', 'delete'),
('role:create', 'Create role', 'role', 'create'),
('role:read', 'Read role', 'role', 'read'),
('role:update', 'Update role', 'role', 'update'),
('role:delete', 'Delete role', 'role', 'delete'),
('permission:create', 'Create permission', 'permission', 'create'),
('permission:read', 'Read permission', 'permission', 'read'),
('permission:update', 'Update permission', 'permission', 'update'),
('permission:delete', 'Delete permission', 'permission', 'delete');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
