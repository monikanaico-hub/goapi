CREATE TABLE 'test'.`user` (
  `iduser` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(500) DEFAULT NULL,
  `password` varchar(500) DEFAULT NULL,
  `firstname` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`iduser`),
  UNIQUE KEY `iduser` (`iduser`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
