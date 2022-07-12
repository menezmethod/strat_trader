ALTER TABLE IF EXISTS currency_rate
    DROP CONSTRAINT currency_rate_base_currency;

ALTER TABLE IF EXISTS currency_rate
    DROP CONSTRAINT currency_rate_currency;

ALTER TABLE IF EXISTS currency_used
    DROP CONSTRAINT currency_used_country;

ALTER TABLE IF EXISTS currency_used
    DROP CONSTRAINT currency_used_currency;

ALTER TABLE IF EXISTS current_inventory
    DROP CONSTRAINT current_inventory_item;

ALTER TABLE IF EXISTS current_inventory
    DROP CONSTRAINT current_inventory_trader;

ALTER TABLE IF EXISTS item
    DROP CONSTRAINT item_currency;

ALTER TABLE IF EXISTS offer
    DROP CONSTRAINT offer_item;

ALTER TABLE IF EXISTS offer
    DROP CONSTRAINT offer_trader;

ALTER TABLE IF EXISTS price
    DROP CONSTRAINT price_currency;

ALTER TABLE IF EXISTS price
    DROP CONSTRAINT price_item;

ALTER TABLE IF EXISTS report
    DROP CONSTRAINT report_currency;

ALTER TABLE IF EXISTS report
    DROP CONSTRAINT report_item;

ALTER TABLE IF EXISTS trade
    DROP CONSTRAINT trade_buyer;

ALTER TABLE IF EXISTS trade
    DROP CONSTRAINT trade_item;

ALTER TABLE IF EXISTS trade
    DROP CONSTRAINT trade_offer;

ALTER TABLE IF EXISTS trade
    DROP CONSTRAINT trade_seller;

ALTER TABLE IF EXISTS trader
    DROP CONSTRAINT trader_country;

ALTER TABLE IF EXISTS trader
    DROP CONSTRAINT trader_currency;

-- tables
DROP TABLE IF EXISTS country;

DROP TABLE IF EXISTS currency;

DROP TABLE IF EXISTS currency_rate;

DROP TABLE IF EXISTS currency_used;

DROP TABLE IF EXISTS current_inventory;

DROP TABLE IF EXISTS item;

DROP TABLE IF EXISTS offer;

DROP TABLE IF EXISTS price;

DROP TABLE IF EXISTS report;

DROP TABLE IF EXISTS trade;

DROP TABLE IF EXISTS trader;
