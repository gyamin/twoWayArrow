CREATE TABLE IF NOT EXISTS `stock_codes` (
    `code` int primary key,
    `name` varchar(255) not null,
    `market` varchar(255) not null
)