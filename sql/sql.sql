CREATE TABLE shop_product
(
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(100)   NOT NULL,
    product_code VARCHAR(100)   NOT NULL,
    price        DECIMAL(10, 2) NOT NULL DEFAULT 0,
    status       TINYINT        NOT NULL DEFAULT 1,
    remark       VARCHAR(255),
    tenant_id    VARCHAR(64),
    created_by   BIGINT,
    created_at   DATETIME,
    updated_by   BIGINT,
    updated_at   DATETIME
);