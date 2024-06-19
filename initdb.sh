#!/bin/bash

# Thông tin kết nối đến cơ sở dữ liệu PostgreSQL
DB_HOST="localhost"
DB_PORT="54321"
DB_NAME="productdb"
DB_USER="root"
DB_PASSWORD="my-secret-pw"

# Đường dẫn đầy đủ của psql
PSQL_PATH="/usr/bin/psql"
export PATH="/usr/bin:$PATH"
# Các câu lệnh SQL để khởi tạo dữ liệu
$PSQL_PATH -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
    -- Thêm dữ liệu vào bảng categories
    INSERT INTO categories (category_name, category_icon, parent_id) VALUES
        ('Điện thoại', 'phone_icon.png', NULL),
        ('Màn hình', 'monitor_icon.png', NULL),
        ('Bàn phím', 'keyboard_icon.png', NULL),
        ('Tủ lạnh', 'fridge_icon.png', NULL),
        ('TV', 'tv_icon.png', NULL);

    -- Thêm dữ liệu vào bảng products
    INSERT INTO products (product_name, description, category_id) VALUES
        ('iPhone 12', 'Latest iPhone model', 1),
        ('Samsung Galaxy S21', 'Flagship Android phone', 1),
        ('Dell U2720Q', '27-inch 4K monitor', 2),
        ('Logitech G Pro X', 'Mechanical gaming keyboard', 3),
        ('Samsung RT18M6213SG', 'Top-freezer refrigerator', 4),
        ('Samsung QN85A', '4K QLED TV', 5);

    -- Thêm dữ liệu vào bảng product_variants và variant_attributes
    INSERT INTO product_variants (product_id, sku, price) VALUES
        (1, 'IP12-B64', 999.99),
        (1, 'IP12-W128', 1099.99),
        (2, 'S21-B128', 899.99),
        (2, 'S21-P256', 999.99);

    INSERT INTO variant_attributes (variant_id, attribute_name, attribute_value) VALUES
        (1, 'Color', 'Black'),
        (1, 'Storage', '64GB'),
        (2, 'Color', 'White'),
        (2, 'Storage', '128GB'),
        (3, 'Color', 'Black'),
        (3, 'Storage', '128GB'),
        (4, 'Color', 'Phantom Violet'),
        (4, 'Storage', '256GB');
"

echo "Dữ liệu đã được khởi tạo vào cơ sở dữ liệu."
