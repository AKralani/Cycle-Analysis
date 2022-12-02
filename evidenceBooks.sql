USE analysisdb;

CREATE TABLE `analysisdb`.`evidenceBooks` (
  `evidenceBookId` INT NOT NULL AUTO_INCREMENT,
  `date` VARCHAR(255) NOT NULL,
  `time` VARCHAR(255) NOT NULL,
  `coin` VARCHAR(255) NOT NULL,
  `price` FLOAT NOT NULL,
  `usdValue` FLOAT NOT NULL,
  `quantity` FLOAT NOT NULL,
  `fee` FLOAT NOT NULL,
  `buySell` VARCHAR(255) NOT NULL,
  `profitLoss` VARCHAR(255),
  `percentProfitLoss` FLOAT,
  `usdValueProfitLoss` FLOAT,
  `snapshot` BLOB,
  PRIMARY KEY (`evidenceBookId`));