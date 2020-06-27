########################################
# MySQL Crash Course
# http://www.forta.com/books/0672327120/
# Example table creation scripts
########################################

###############################
# Create `crashcourse` database
###############################
CREATE DATABASE IF NOT EXISTS `crashcourse`
DEFAULT CHARACTER SET `utf8`            # 设置默认字符集
DEFAULT COLLATE `utf8_general_ci`;      # 设置默认字符集校对规则

############################
# Use `crashcourse` database
############################
USE crashcourse;


##########################
# Create `customers` table
##########################
CREATE TABLE IF NOT EXISTS `customers`
(
`cust_id`      INT          NOT NULL AUTO_INCREMENT,
`cust_name`    VARCHAR(50)  NOT NULL ,
`cust_address` VARCHAR(50)  NULL ,
`cust_city`    VARCHAR(50)  NULL ,
`cust_state`   VARCHAR(5)   NULL ,
`cust_zip`     VARCHAR(10)  NULL ,
`cust_country` VARCHAR(50)  NULL ,
`cust_contact` VARCHAR(50)  NULL ,
`cust_email`   VARCHAR(255) NULL ,
CONSTRAINT `pk_customers` PRIMARY KEY (`cust_id`)
) ENGINE=InnoDB;


###########################
# Create `orderitems` table
###########################
CREATE TABLE IF NOT EXISTS `orderitems`
(
`order_num`  INT          NOT NULL ,
`order_item` INT          NOT NULL ,
`prod_id`    VARCHAR(10)  NOT NULL ,
`quantity`   INT          NOT NULL ,
`item_price` DECIMAL(8,2) NOT NULL ,
CONSTRAINT `pk_orderitems` PRIMARY KEY (`order_num`, `order_item`)
) ENGINE=InnoDB;


#######################
# Create `orders` table
#######################
CREATE TABLE IF NOT EXISTS `orders`
(
`order_num`  INT      NOT NULL AUTO_INCREMENT,
`order_date` DATETIME NOT NULL ,
`cust_id`    INT      NOT NULL ,
CONSTRAINT `pk_orders` PRIMARY KEY (`order_num`)
) ENGINE=InnoDB;


#########################
# Create `products` table
#########################
CREATE TABLE IF NOT EXISTS `products`
(
`prod_id`    VARCHAR(10)   NOT NULL,
`vend_id`    INT           NOT NULL ,
`prod_name`  VARCHAR(255)  NOT NULL ,
`prod_price` DECIMAL(8,2)  NOT NULL ,
`prod_desc`  TEXT          NULL ,
CONSTRAINT `pk_products` PRIMARY KEY(`prod_id`)
) ENGINE=InnoDB;


########################
# Create `vendors` table
########################
CREATE TABLE IF NOT EXISTS `vendors`
(
`vend_id`      INT         NOT NULL AUTO_INCREMENT,
`vend_name`    VARCHAR(50) NOT NULL ,
`vend_address` VARCHAR(50) NULL ,
`vend_city`    VARCHAR(50) NULL ,
`vend_state`   VARCHAR(5)  NULL ,
`vend_zip`     VARCHAR(10) NULL ,
`vend_country` VARCHAR(50) NULL ,
CONSTRAINT `pk_vendors` PRIMARY KEY (`vend_id`)
) ENGINE=InnoDB;


#############################
# Create `productnotes` table
#############################
CREATE TABLE IF NOT EXISTS `productnotes`
(
`note_id`    INT           NOT NULL AUTO_INCREMENT,
`prod_id`    VARCHAR(10)   NOT NULL,
`note_date`  DATETIME      NOT NULL,
`note_text`  TEXT          NULL ,
CONSTRAINT `pk_productnotes` PRIMARY KEY(`note_id`),
FULLTEXT(`note_text`)
) ENGINE=MyISAM;


#######################
# Define `foreign` keys
#######################
ALTER TABLE `orderitems` ADD CONSTRAINT `fk_orderitems_orders` FOREIGN KEY (`order_num`) REFERENCES `orders` (`order_num`);
ALTER TABLE `orderitems` ADD CONSTRAINT `fk_orderitems_products` FOREIGN KEY (`prod_id`) REFERENCES `products` (`prod_id`);
ALTER TABLE `orders` ADD CONSTRAINT `fk_orders_customers` FOREIGN KEY (`cust_id`) REFERENCES `customers` (`cust_id`);
ALTER TABLE `products` ADD CONSTRAINT `fk_products_vendors` FOREIGN KEY (`vend_id`) REFERENCES `vendors` (`vend_id`);
