CREATE TABLE `Account` (
  `Id` int PRIMARY KEY AUTO_INCREMENT,
  `FirstName` varchar(255) NOT NULL,
  `Ballance` bigint NOT NULL,
  `Currency` varchar(255) NOT NULL,
  `CreatedAt` timestamptz NOT NULL DEFAULT "now()"
);

CREATE TABLE `Entries` (
  `Id` int PRIMARY KEY AUTO_INCREMENT,
  `AccountNumber` bigint,
  `Amount` bigint NOT NULL COMMENT 'Can be negative or positive',
  `CreatedAt` timestamptz NOT NULL DEFAULT "now()"
);

CREATE TABLE `Transfers` (
  `Id` bigint PRIMARY KEY AUTO_INCREMENT,
  `FromAccount` int NOT NULL,
  `ToAccount` int NOT NULL,
  `Amount` bigint NOT NULL COMMENT 'Must be positive',
  `CreatedAt` timestamptz NOT NULL DEFAULT "now()"
);

CREATE INDEX `Account_index_0` ON `Account` (`Id`);

CREATE INDEX `Entries_index_1` ON `Entries` (`AccountNumber`);

CREATE INDEX `Transfers_index_2` ON `Transfers` (`FromAccount`);

CREATE INDEX `Transfers_index_3` ON `Transfers` (`ToAccount`);

CREATE INDEX `Transfers_index_4` ON `Transfers` (`FromAccount`, `ToAccount`);

ALTER TABLE `Entries` ADD FOREIGN KEY (`AccountNumber`) REFERENCES `Account` (`Id`);

ALTER TABLE `Transfers` ADD FOREIGN KEY (`FromAccount`) REFERENCES `Account` (`Id`);

ALTER TABLE `Transfers` ADD FOREIGN KEY (`ToAccount`) REFERENCES `Account` (`Id`);
