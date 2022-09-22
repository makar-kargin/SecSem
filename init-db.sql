create table if not exists users (
    id integer unique not null,
    login text not null,
    money_amount integer not null,
    card_number text not null,
    status integer not null
);

create table if not exists pswd (
    id integer unique not null,
    pswd text not null
);

replace into users values
    (1, 'admin', 10000000, '5504581490119087', 1),
    (2, 'PoorSasha', 10, '5578714387970509', 1),
    (3, 'RichMasha', 10000000, '5222601280321445', 1),
    (4, 'AverageDima', 1000, '5415778272955872', 0),
    (5, 'BogdanBogomDan', 228, '5382279802385861', 0);

replace into pswd values 
    (1, 'admin'),
    (2, 'Sashaisthebest2007'),
    (3, 'kmiujbtgvcde'),
    (4, 'Dimasamiykrutoioioioi'),
    (5, 'Olegcaresaboutsecurity712531678631Yjnbsdq_4');

