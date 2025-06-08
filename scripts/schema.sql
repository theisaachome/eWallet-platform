CREATE DATABASE eWallet_db;
-- Enable UUID extension (only needs to be done once per database)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Define ENUM type for status (PostgreSQL doesn't use inline ENUMs)
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
            CREATE TYPE user_status AS ENUM ('active', 'suspended', 'deleted');
        END IF;
    END$$;

-- Create users table
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       public_id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                       phone_number VARCHAR(20) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE,
                       status user_status DEFAULT 'active',
                       created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- Table: profiles
CREATE TABLE      profiles (
                   id UUID PRIMARY KEY,
                   user_id UUID UNIQUE REFERENCES users(id),
                   full_name TEXT,
                   date_of_birth DATE,
                   gender TEXT,
                   profile_picture_url TEXT,
                   created_at TIMESTAMP,
                   updated_at TIMESTAMP
               );

CREATE TABLE users (
                       id BIGINT PRIMARY KEY AUTO_INCREMENT,
                       public_id CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),
                       phone_number VARCHAR(20) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE,
                       status ENUM('active', 'suspended', 'deleted') DEFAULT 'active',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE identity_documents (
                       id UUID PRIMARY KEY,
                       user_id UUID REFERENCES users(id),
                       document_type TEXT NOT NULL CHECK (document_type IN ('NRC', 'PASSPORT')),
                       document_number TEXT NOT NULL,
                       issued_country TEXT, -- for passport
                       issued_date DATE,
                       expiry_date DATE,
                       is_primary BOOLEAN DEFAULT false,
                       UNIQUE (document_type, document_number), -- prevent duplicates
                       UNIQUE (user_id, is_primary) WHERE is_primary = true -- enforce one primary document per user
                   );


CREATE TABLE wallets (
                         id BIGINT PRIMARY KEY AUTO_INCREMENT,
                         public_id CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),

                         user_id BIGINT NOT NULL,
                         balance DECIMAL(18,2) DEFAULT 0.00,
                         currency CHAR(3) DEFAULT 'MYR',
                         status ENUM('active', 'frozen') DEFAULT 'active',

                         FOREIGN KEY (user_id) REFERENCES users(id),
                         UNIQUE(user_id)
);


CREATE TABLE transactions (
                              id BIGINT PRIMARY KEY AUTO_INCREMENT,
                              transaction_id CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),

                              type ENUM('peer_transfer', 'topup', 'withdrawal', 'payment') NOT NULL,
                              status ENUM('pending', 'completed', 'failed', 'reversed') DEFAULT 'pending',

                              sender_wallet_id BIGINT,
                              receiver_wallet_id BIGINT,
                              amount DECIMAL(18,2) NOT NULL,
                              currency CHAR(3) NOT NULL DEFAULT 'MYR',
                              service_fee DECIMAL(18,2) DEFAULT 0.00,
                              total_deducted DECIMAL(18,2) NOT NULL,

                              note VARCHAR(255),
                              reference_code VARCHAR(100),
                              channel VARCHAR(50),

                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                              FOREIGN KEY (sender_wallet_id) REFERENCES wallets(id),
                              FOREIGN KEY (receiver_wallet_id) REFERENCES wallets(id),

                              INDEX idx_sender_wallet (sender_wallet_id),
                              INDEX idx_receiver_wallet (receiver_wallet_id),
                              INDEX idx_created_at (created_at)
);


CREATE TABLE wallet_logs (
                             id BIGINT PRIMARY KEY AUTO_INCREMENT,
                             public_id CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),

                             wallet_id BIGINT NOT NULL,
                             transaction_id CHAR(36),
                             direction ENUM('credit', 'debit') NOT NULL,
                             amount DECIMAL(18,2),
                             balance_before DECIMAL(18,2),
                             balance_after DECIMAL(18,2),
                             description TEXT,

                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                             FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);

CREATE TABLE withdrawals (
                             id BIGINT PRIMARY KEY AUTO_INCREMENT,
                             public_id CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),

                             wallet_id BIGINT,
                             bank_account_no VARCHAR(50),
                             bank_name VARCHAR(100),
                             amount DECIMAL(18,2),
                             status ENUM('pending', 'processing', 'completed', 'failed'),
                             requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                             FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);

CREATE TABLE topups (
                        id BIGINT PRIMARY KEY AUTO_INCREMENT,
                        public_id CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),

                        wallet_id BIGINT,
                        provider VARCHAR(50),
                        external_reference VARCHAR(100),
                        amount DECIMAL(18,2),
                        status ENUM('pending', 'successful', 'failed'),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                        FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);


DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
                             id  int(11) NOT NULL AUTO_INCREMENT,
                             customer_id VARCHAR(20) PRIMARY KEY,
                             name VARCHAR(100) NOT NULL,
                             email VARCHAR(100) UNIQUE NOT NULL,
                             phone VARCHAR(20),
                             address TEXT,
                             date_of_birth DATE,
                             status ENUM('active', 'inactive') DEFAULT 'active'
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `accounts`;
CREATE TABLE accounts (
                          id VARCHAR(36) PRIMARY KEY,
                          account_no VARCHAR(16) NOT NULL UNIQUE,
                          customer_id VARCHAR(20) NOT NULL,
                          opening_date datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          account_type ENUM('savings', 'current') NOT NULL,
                          amount DECIMAL(15,2) DEFAULT 0.00,
                          status ENUM('active', 'inactive') DEFAULT 'active',
    -- Foreign key assuming you have a customers table
                          FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
                              ON DELETE CASCADE
                              ON UPDATE CASCADE
)ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
                                `transaction_id` VARCHAR(36) NOT NULL PRIMARY KEY ,
                                `account_id` int(11) NOT NULL,
                                `amount` decimal(10,2) NOT NULL,
                                `transaction_type` varchar(10) NOT NULL,
                                `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                PRIMARY KEY (`transaction_id`),
                                KEY `transactions_FK` (`account_id`),
                                CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `username` varchar(20) NOT NULL,
                         `password` varchar(20) NOT NULL,
                         `role` varchar(20) NOT NULL,
                         `customer_id` int(11) DEFAULT NULL,
                         `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `users` VALUES
                        ('admin','abc123','admin', NULL, '2020-08-09 10:27:22'),
                        ('2001','abc123','user', 2001, '2020-08-09 10:27:22'),
                        ('2000','abc123','user', 2000, '2020-08-09 10:27:22');

DROP TABLE IF EXISTS `refresh_token_store`;

CREATE TABLE `refresh_token_store` (
                                       `refresh_token` varchar(300) NOT NULL,
                                       created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       PRIMARY KEY (`refresh_token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE transactions (
                              id BIGINT AUTO_INCREMENT PRIMARY KEY,
                              transaction_id VARCHAR(50) NOT NULL UNIQUE, -- e.g., TXN20250608123456
                              type ENUM('peer_transfer', 'topup', 'withdraw', 'merchant_payment') NOT NULL,
                              status ENUM('pending', 'completed', 'failed', 'reversed') NOT NULL DEFAULT 'pending',

                              sender_id BIGINT,         -- FK to users.id
                              receiver_id BIGINT,       -- FK to users.id or merchants.id
                              receiver_type ENUM('user', 'merchant') DEFAULT 'user', -- flexible ownership
                              currency CHAR(3) NOT NULL DEFAULT 'MYR',
                              amount DECIMAL(15,2) NOT NULL,
                              service_fee DECIMAL(15,2) NOT NULL DEFAULT 0.00,

                              wallet_balance_before DECIMAL(15,2),
                              wallet_balance_after DECIMAL(15,2),

                              note VARCHAR(255),
                              reference_code VARCHAR(100), -- useful for reconciliation
                              channel VARCHAR(50),         -- e.g., mobile_app, api, web_portal
                              created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                              INDEX idx_sender_id (sender_id),
                              INDEX idx_receiver_id (receiver_id),
                              INDEX idx_transaction_id (transaction_id),
                              INDEX idx_created_at (created_at)
);

CREATE TABLE transaction_audit_logs (
                                        id BIGINT AUTO_INCREMENT PRIMARY KEY,
                                        transaction_id VARCHAR(50) NOT NULL,
                                        status ENUM('pending', 'completed', 'failed', 'reversed') NOT NULL,
                                        changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        remarks TEXT
);

CREATE TABLE transaction_metadata (
                                      transaction_id VARCHAR(50) PRIMARY KEY,
                                      metadata JSON
);
