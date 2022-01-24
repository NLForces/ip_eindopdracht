USE ip_eindopdracht;

INSERT INTO `accounts` (`id`, `name`, `code`, `maxcredit`, `pincode`, `created_at`, `updated_at`) 
VALUES (1, 'Bob Mulder', 'IBAN 001', '-100', '1111', NOW(), NOW());

INSERT INTO `accounts` (`id`, `name`, `code`, `maxcredit`, `pincode`, `created_at`, `updated_at`) 
VALUES (2, 'Gerrit Vogelzang', 'IBAN 002', '-120', '2222',  NOW(), NOW());

INSERT INTO `accounts` (`id`, `name`, `code`, `maxcredit`, `pincode`, `created_at`, `updated_at`) 
VALUES (3, 'Rene Zeeders', 'IBAN 003', '-150', '3333', NOW(), NOW());

INSERT INTO `accounts` (`id`, `name`, `code`, `maxcredit`, `pincode`, `created_at`, `updated_at`) 
VALUES (4, 'David Doorn', 'IBAN 004', '-200', '4444', NOW(), NOW());


-- Transactions

INSERT INTO `transactions` (`id`, `account_id`, `amount`, `type`, `created_at`, `updated_at`) 
VALUES (1, 1, 20.00, 'debet', NOW(), NOW());

INSERT INTO `transactions` (`id`, `account_id`, `amount`, `type`, `created_at`, `updated_at`) 
VALUES (2, 1, 40.20, 'debet', NOW(), NOW());

INSERT INTO `transactions` (`id`, `account_id`, `amount`, `type`, `created_at`, `updated_at`) 
VALUES (3, 2, 20.16, 'debet', NOW(), NOW());