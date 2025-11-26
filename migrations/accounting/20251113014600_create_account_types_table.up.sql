CREATE TABLE accounting.account_types
(
    id         SERIAL PRIMARY KEY,
    code       VARCHAR(10) UNIQUE NOT NULL,
    name       VARCHAR(50)        NOT NULL,
    dc_pattern CHAR(1)            NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(255)
);

CREATE INDEX idx_account_types_code ON accounting.account_types(code);
CREATE INDEX idx_account_types_deleted ON accounting.account_types(deleted_at);