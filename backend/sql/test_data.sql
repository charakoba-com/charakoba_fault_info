INSERT INTO infotype VALUES ('maintenance');
INSERT INTO services VALUES ('www');

INSERT INTO faultinfo (infotype, service, begin, end, detail) VALUES (
        'maintenance',
        'www',
        '2017-01-01 00:00:00',
        '2017-01-01 12:00:00',
        'test data'
);
