CREATE TABLE accounting.accounts
(
    id                  SERIAL PRIMARY KEY,
    account_category_id INT                NOT NULL REFERENCES accounting.account_categories (id) ON DELETE CASCADE, -- !!!
    code                VARCHAR(20) UNIQUE NOT NULL,
    name                VARCHAR(255)       NOT NULL,
    init_balance        DECIMAL(19, 4)     NOT NULL DEFAULT 0,
    balance             DECIMAL(19, 4)     NOT NULL DEFAULT 0,
    impacted_account    INT,
    impacted_account2   INT,
    impacted_account3   INT,
    is_active           BOOLEAN                     DEFAULT true,
    created_at          TIMESTAMP                   DEFAULT NOW(),
    created_by          VARCHAR(255),
    updated_at          TIMESTAMP                   DEFAULT NOW(),
    updated_by          VARCHAR(255),
    deleted_at          TIMESTAMP,
    deleted_by          VARCHAR(255)
);

CREATE INDEX idx_accounts_category ON accounting.accounts(account_category_id);
CREATE INDEX idx_accounts_code ON accounting.accounts(code);
CREATE INDEX idx_accounts_active ON accounting.accounts(is_active);
CREATE INDEX idx_accounts_deleted ON accounting.accounts(deleted_at);