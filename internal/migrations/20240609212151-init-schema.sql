
-- +migrate Up
-- Table for categories and subcategories
CREATE TABLE IF NOT EXISTS categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100),
    category_icon VARCHAR(255),
    parent_id INTEGER REFERENCES categories(category_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Table for products
CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INTEGER REFERENCES categories(category_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Table for product_variants
CREATE TABLE IF NOT EXISTS product_variants (
    variant_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    sku VARCHAR(50) UNIQUE NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);



-- Table for variant_attributes
CREATE TABLE IF NOT EXISTS variant_attributes (
    attribute_id SERIAL PRIMARY KEY,
    variant_id INT NOT NULL,
    FOREIGN KEY (variant_id) REFERENCES product_variants(variant_id) ON DELETE CASCADE,
    attribute_name VARCHAR(100) NOT NULL,
    attribute_value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +migrate Down
-- Drop table variant_attributes
DROP TABLE IF EXISTS variant_attributes;

-- Drop table product_variants
DROP TABLE IF EXISTS product_variants;

-- Drop table products
DROP TABLE IF EXISTS products;

-- Drop table categories
DROP TABLE IF EXISTS categories;
