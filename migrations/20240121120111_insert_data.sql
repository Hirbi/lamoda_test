-- +goose Up
-- +goose StatementBegin
INSERT INTO storage (name, is_available) VALUES ('warehouse1', true);
INSERT INTO storage (name, is_available) VALUES ('warehouse2', true);
INSERT INTO storage (name, is_available) VALUES ('warehouse3', false);
INSERT INTO storage (name, is_available) VALUES ('warehouse4', false);
INSERT INTO storage (name, is_available) VALUES ('warehouse5', true);

INSERT INTO products (name, code, amount, size, storage_id) VALUES ('keyboard', 'code-1', 2, '10x10', 1);
INSERT INTO products (name, code, amount, size, storage_id) VALUES ('headphones', 'code-2', 4, '10x10', 1);
INSERT INTO products (name, code, amount, size, storage_id) VALUES ('mouse', 'code-3', 3, '10x10', 1);
INSERT INTO products (name, code, amount, size, storage_id) VALUES ('usb-hub', 'code-4', 8, '10x10', 1);

INSERT INTO products (name, code, amount, size, storage_id) VALUES ('keyboard', 'code-5', 2, '10x10', 2);
INSERT INTO products (name, code, amount, size, storage_id) VALUES ('keyboard', 'code-6', 3, '10x10', 3);
INSERT INTO products (name, code, amount, size, storage_id) VALUES ('keyboard', 'code-7', 4, '10x10', 4);
INSERT INTO products (name, code, amount, size, storage_id) VALUES ('keyboard', 'code-8', 5, '10x10', 5);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM products;
DELETE FROM storage;
-- +goose StatementEnd
