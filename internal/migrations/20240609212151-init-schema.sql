
-- +migrate Up
-- Bảng danh mục
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY, -- ID tự tăng cho danh mục
    category_name VARCHAR(255) NOT NULL -- Tên danh mục
);

-- Tạo bảng subcategories
CREATE TABLE IF NOT EXISTS subcategories (
    subcategory_id SERIAL PRIMARY KEY, -- ID tự tăng cho danh mục con
    category_id INT NOT NULL, -- ID danh mục chính
    subcategory_name VARCHAR(255) NOT NULL, -- Tên danh mục con
    FOREIGN KEY (category_id) REFERENCES categories(category_id) -- Khóa ngoại tham chiếu đến bảng categories
);


-- Bảng sản phẩm
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY, -- ID tự tăng cho sản phẩm
    product_name VARCHAR(255) NOT NULL, -- Tên sản phẩm
    description TEXT, -- Mô tả sản phẩm
    category_id INT, -- ID danh mục
    FOREIGN KEY (category_id) REFERENCES categories(category_id) -- Khóa ngoại tham chiếu đến bảng categories
);

-- Bảng thuộc tính sản phẩm
CREATE TABLE product_attributes (
    attribute_id SERIAL PRIMARY KEY, -- ID tự tăng cho thuộc tính
    category_id INT, -- ID danh mục
    attribute_name VARCHAR(255) NOT NULL, -- Tên thuộc tính
    FOREIGN KEY (category_id) REFERENCES categories(category_id) -- Khóa ngoại tham chiếu đến bảng categories
);

-- Bảng biến thể
CREATE TABLE variants (
    variant_id SERIAL PRIMARY KEY, -- ID tự tăng cho biến thể
    product_id INT, -- ID sản phẩm
    FOREIGN KEY (product_id) REFERENCES products(product_id) -- Khóa ngoại tham chiếu đến bảng products
);

-- Bảng giá trị thuộc tính biến thể
CREATE TABLE variant_attributes (
    variant_attribute_id SERIAL PRIMARY KEY, -- ID tự tăng cho giá trị thuộc tính
    variant_id INT, -- ID biến thể
    attribute_id INT, -- ID thuộc tính
    attribute_value VARCHAR(255), -- Giá trị thuộc tính
    FOREIGN KEY (variant_id) REFERENCES variants(variant_id), -- Khóa ngoại tham chiếu đến bảng variants
    FOREIGN KEY (attribute_id) REFERENCES product_attributes(attribute_id) -- Khóa ngoại tham chiếu đến bảng product_attributes
);

-- Bảng SKU
CREATE TABLE sku (
    sku_id SERIAL PRIMARY KEY, -- ID tự tăng cho SKU
    variant_id INT, -- ID biến thể
    sku_code VARCHAR(100) UNIQUE NOT NULL, -- Mã SKU duy nhất
    stock_quantity INT NOT NULL, -- Số lượng tồn kho
    price DECIMAL(10, 2) NOT NULL, -- Giá sản phẩm
    FOREIGN KEY (variant_id) REFERENCES variants(variant_id) -- Khóa ngoại tham chiếu đến bảng variants
);

-- +migrate Down
-- Xóa bảng SKU
DROP TABLE IF EXISTS sku;

-- Xóa bảng giá trị thuộc tính biến thể
DROP TABLE IF EXISTS variant_attributes;

-- Xóa bảng biến thể
DROP TABLE IF EXISTS variants;

-- Xóa bảng thuộc tính sản phẩm
DROP TABLE IF EXISTS product_attributes;

-- Xóa bảng sản phẩm
DROP TABLE IF EXISTS products;

-- Xóa bảng danh mục
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS subcategories;
