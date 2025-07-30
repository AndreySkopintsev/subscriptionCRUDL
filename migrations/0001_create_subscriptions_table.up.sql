CREATE TABLE IF NOT EXISTS subscriptions (
    service_name VARCHAR(100) NOT NULL,
    price integer,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL
);