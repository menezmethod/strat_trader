-- name: GetTrader :one
SELECT * FROM trader
WHERE id = $1;

-- name: ListTraders :one
SELECT * FROM trader
ORDER BY last_name;

-- name: CreateTrader :one
INSERT INTO trader (first_name, last_name, user_name, password, email, confirmation_code, time_registered,
                    time_confirmed, country_id, preferred_currency_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: DeleteTrader :exec
DELETE FROM trader
WHERE id = $1;

-- name: UpdateTrader :exec
UPDATE trader
set first_name = $2, last_name = $3, user_name = $4, password = $5, email = $6, confirmation_code = $7, time_registered = $8,
    time_confirmed = $9, country_id = $10, preferred_currency_id = $11
WHERE id = $1;

-- name: GetCountry :one
SELECT * FROM country
WHERE id = $1
LIMIT 1;

-- name: ListCountries :many
SELECT * FROM country
ORDER BY name;

-- name: CreateCountry :one
INSERT INTO country (name)
VALUES ($1)
RETURNING *;

-- name: DeleteCountry :exec
DELETE FROM country
WHERE id = $1;

-- name: UpdateCountry :exec
UPDATE country
set name = $2
WHERE id = $1;

-- name: GetCurrency :one
SELECT * FROM currency
WHERE id = $1
LIMIT 1;

-- name: ListCurrencies :many
SELECT * FROM currency
ORDER BY name;

-- name: CreateCurrency :one
INSERT INTO currency (code, name, is_active, is_base_currency) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteCurrency :exec
DELETE FROM currency
WHERE id = $1;

-- name: UpdateCurrency :exec
UPDATE currency
set name = $2, code = $3, is_active = $4, is_base_currency = $5
WHERE id = $1;

-- name: GetCurrencyRate :one
SELECT * FROM currency_rate
WHERE id = $1
LIMIT 1;

-- name: ListCurrencyRates :many
SELECT * FROM currency_rate
ORDER BY rate;

-- name: CreateCurrencyRate :one
INSERT INTO currency_rate (currency_id, base_currency_id, rate, ts) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteCurrencyRate :exec
DELETE FROM currency_rate
WHERE id = $1;

-- name: UpdateCurrencyRate :exec
UPDATE currency_rate
set currency_id = $2, base_currency_id = $3, rate = $4, ts = $5
WHERE id = $1;

-- name: GetCurrencyUsed :one
SELECT * FROM currency_used
WHERE id = $1
LIMIT 1;

-- name: ListCurrenciesUsed :many
SELECT * FROM currency_used
ORDER BY date_from;

-- name: CreateCurrencyUsed :one
INSERT INTO currency_used (country_id, currency_id, date_from, date_to) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteCurrencyUsed :exec
DELETE FROM currency_used
WHERE id = $1;

-- name: UpdateCurrencyUsed :exec
UPDATE currency_used
set country_id = $2, currency_id = $3, date_from = $4, date_to = $5
WHERE id = $1;

-- name: GetCurrentInventory :one
SELECT * FROM current_inventory
WHERE id = $1
LIMIT 1;

-- name: ListCurrentInventory :many
SELECT * FROM current_inventory
ORDER BY date_from;

-- name: CreateCurrentInventory :one
INSERT INTO current_inventory (trader_id, item_id, quantity) VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteCurrentInventory :exec
DELETE FROM current_inventory
WHERE id = $1;

-- name: UpdateCurrentInventory :exec
UPDATE current_inventory
set trader_id = $2, item_id = $3, quantity = $4
WHERE id = $1;

-- name: GetItem :one
SELECT * FROM item
WHERE id = $1
LIMIT 1;

-- name: ListItems :many
SELECT * FROM item
ORDER BY name;

-- name: CreateItem :one
INSERT INTO item (code, name, is_active, currency_id, details) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM item
WHERE id = $1;

-- name: UpdateItem :exec
UPDATE item
set code = $2, name = $3, is_active = $4, currency_id = $5, details = $6
WHERE id = $1;

-- name: GetOffer :one
SELECT * FROM offer
WHERE id = $1
LIMIT 1;

-- name: ListOffers :many
SELECT * FROM offer
ORDER BY quantity;

-- name: CreateOffer :one
INSERT INTO offer (trader_id, item_id, quantity, buy, sell, price, ts, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: DeleteOffer :exec
DELETE FROM offer
WHERE id = $1;

-- name: UpdateOffer :exec
UPDATE offer
set trader_id = $2, item_id = $3, quantity = $4, buy = $5, sell = $6, price = $7, ts = $8, is_active = $9
WHERE id = $1;

-- name: GetPrice :one
SELECT * FROM price
WHERE id = $1
LIMIT 1;

-- name: ListPrices :many
SELECT * FROM price
ORDER BY id;

-- name: CreatePrice :one
INSERT INTO price (item_id, currency_id, buy, sell, ts) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeletePrice :exec
DELETE FROM price
WHERE id = $1;

-- name: UpdatePrice :exec
UPDATE price
set item_id = $1, currency_id = $2, buy = $3, sell = $4, ts = $5
WHERE id = $1;

-- name: GetReport :one
SELECT * FROM report
WHERE id = $1
LIMIT 1;

-- name: ListReports :many
SELECT * FROM report
ORDER BY id;

-- name: CreateReport :one
INSERT INTO report (trading_date, item_id, currency_id, first_price, last_price, min_price, max_price, avg_price, total_amount, quantity) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: DeleteReport :exec
DELETE FROM report
WHERE id = $1;

-- name: UpdateReport :exec
UPDATE report
set trading_date = $2, item_id = $3, currency_id = $4, first_price = $5, last_price = $6, min_price = $7, max_price = $8, avg_price = $9, total_amount = $10, quantity = $11
WHERE id = $1;

-- name: GetTrade :one
SELECT * FROM trade
WHERE id = $1
LIMIT 1;

-- name: ListTrades :many
SELECT * FROM trade
ORDER BY id;

-- name: CreateTrade :one
INSERT INTO trade (item_id, seller_id, buyer_id, quantity, unit_price, description, offer_id) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteTrade :exec
DELETE FROM trade
WHERE id = $1;

-- name: UpdateTrade :exec
UPDATE trade
set item_id = $2, seller_id = $3, buyer_id = $4, quantity = $5, unit_price = $6, description = $7, offer_id = $8
WHERE id = $1;