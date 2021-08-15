CREATE DATABASE IF NOT EXISTS ecommerce;
USE ecommerce;
CREATE TABLE IF NOT EXISTS reviews (
    id varchar(50),
    review_id varchar(50),
    order_id varchar(50),
    score integer,
    comment_title varchar(255),
    comment_message varchar(255),
    created_at datetime,
    answered_at datetime,
    published_at datetime,
    PRIMARY KEY (`id`)
)
ENGINE=InnoDB
AUTO_INCREMENT=1
DEFAULT CHARSET=latin1;
