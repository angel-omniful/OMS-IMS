CREATE TABLE IF NOT EXISTS skus (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    seller_id UUID NOT NULL,
    sku_code TEXT NOT NULL,
    name TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (tenant_id, seller_id, sku_code),

    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE,
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_skus_tenant_id ON skus (tenant_id);
CREATE INDEX IF NOT EXISTS idx_skus_seller_id ON skus (seller_id);      