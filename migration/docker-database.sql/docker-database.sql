CREATE TABLE banks(
    bank_id integer PRIMARY KEY, 
    bank_code text UNIQUE, 
    bank_name text
    
);

CREATE TABLE accounts(
    account_id integer PRIMARY KEY,
    fk_customer_id bigint REFERENCES customers(customer_id),
    account_number numeric  UNIQUE NOT NULL,
    is_active boolean,
    account_type text NOT NULL,
    current_balance numeric(10,2) NOT NULL
   
);

CREATE TABLE customers(
    customer_id integer PRIMARY KEY,
    name text NOT NULL,
    CPF_CNPJ text NOT NULL,
    phone_number text NOT NULL,
    email text NOT NULL,
    fk_bank_id integer REFERENCES banks(bank_id)
   
);

CREATE TABLE transactions(
    transaction_id integer PRIMARY KEY,
    fk_account_id integer REFERENCES accounts(account_id),
    credited_amount numeric(10,2),
    debited_amount numeric(10,2),
    running_balance numeric(10,2),
    other_party_account_number numeric(10)
        
);

