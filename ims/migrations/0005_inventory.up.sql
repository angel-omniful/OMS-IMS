CREATE TABLE IF NOT EXISTS inventory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    seller_id UUID NOT NULL,
    hub_id UUID NOT NULL,
    sku_id UUID NOT NULL,
    quantity INT DEFAULT 0,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (tenant_id, seller_id, hub_id, sku_id),

    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE,
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE,
    FOREIGN KEY (hub_id) REFERENCES hubs(id) ON DELETE CASCADE,
    FOREIGN KEY (sku_id) REFERENCES skus(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_inventory_tenant_id ON inventory (tenant_id);
CREATE INDEX IF NOT EXISTS idx_inventory_seller_id ON inventory (seller_id);        