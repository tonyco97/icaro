insert into accounts values (null, 1, "uuid-gen-12345", "reseller", "My Name", "firstuser",md5("password"),"firstuser@email.org",now());
insert into hotspots values (null,1,"HSTest","HS for development",now());
insert into accounts values (null, 1, "uuid-gen-99999", "customer", "Customer Account", "cust1",md5("cust1"),"cust1@example.com",now());
insert into accounts_hotspots values (null, 3, 1);
insert into units values (null,1,'00-00-00-00-00-00','testunit','1234-uuid-aaaa','secret',NOW());
insert into users values (NULL, 1, 'First User', 'firstuser', 'firstpassword', 'first.user@nethserver.org', 'test', 0, 0, NOW(), NOW() + INTERVAL 3600 DAY, NOW());
