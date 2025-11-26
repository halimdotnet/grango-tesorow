CREATE TABLE accounting.account_categories
(
    id              SERIAL PRIMARY KEY,
    account_type_id INT                NOT NULL REFERENCES accounting.account_types (id) ON DELETE CASCADE, -- !!!
    code            VARCHAR(20) UNIQUE NOT NULL,
    name            VARCHAR(255)       NOT NULL,
    classification  VARCHAR(50),
    is_active bool NOT NULL DEFAULT true,
    created_at      TIMESTAMP DEFAULT NOW(),
    created_by      VARCHAR(255),
    updated_at      TIMESTAMP DEFAULT NOW(),
    updated_by      VARCHAR(255),
    deleted_at      TIMESTAMP,
    deleted_by      VARCHAR(255)
);

CREATE INDEX idx_account_categories_type ON accounting.account_categories(account_type_id);
CREATE INDEX idx_account_categories_code ON accounting.account_categories(code);
CREATE INDEX idx_account_categories_classification ON accounting.account_categories(classification);
CREATE INDEX idx_account_categories_deleted ON accounting.account_categories(deleted_at);
CREATE INDEX idx_account_categories_is_active ON accounting.account_categories(is_active);