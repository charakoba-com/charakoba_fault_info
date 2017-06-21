CREATE TABLE infotype (
        infotype VARCHAR(64) NOT NULL,
        PRIMARY KEY (infotype)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE services (
        name VARCHAR(64) NOT NULL,
        PRIMARY KEY(name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE faultinfo (
        id INT NOT NULL AUTO_INCREMENT,
        infotype VARCHAR(64) NOT NULL,
        service VARCHAR(64) NOT NULL,
        begin DATETIME NOT NULL,
        end DATETIME,
        detail VARCHAR(1024) NOT NULL DEFAULT '',
        PRIMARY KEY(id),
        FOREIGN KEY (infotype) REFERENCES infotype (infotype),
        FOREIGN KEY (service) REFERENCES services (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
