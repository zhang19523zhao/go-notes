create database cmdb default charset utf8mb4;

create table user(
    id bigint primary key auto_increment,
    staff_id varchar(32) not null default '',
    name varchar(64) not null default  '',
    nickname varchar(64) not null default '',
    password varchar(1024) not null default '',
    gender int not null default 0 comment '0: 正常 1: 男',
    tel varchar(64) not null default '',
    email varchar(64) not null default '',
    addr varchar(128) not null default '',
    department varchar(128) not null default '',
    status int not null default 0 comment '0: 正常, 1: 锁定',
    created_at datetime not null ,
    updated_at datetime not null ,
    deleted_at datetime
)engine=innodb default charset utf8mb4;

insert into user(staff_id, name,nickname, password,gender,tel,email,addr,department,status,created_at,updated_at) values
 ("Z00001", "tt", "t", md5("123"), 1, "1241441442145", "13313@qq.com", "广州", "研发一部",0, now(), now()),
 ("Z00001", "lisi", "l4", md5("123"), 0, "1241445734742145", "13534313@qq.com", "广州", "研发二部",0, now(), now()),
 ("Z00001", "zhagnsan", "z3", md5("123"), 0, "1241441442145", "1331343@qq.com", "广州", "研发二部",0, now(), now()),
 ("Z00001", "zh2", "zh2", md5("123"), 1, "242421442145", "1331223@qq.com", "广州", "研发三部",0, now(), now()),
 ("Z00001", "zh3", "zh3", md5("123"), 1, "1242441442145", "13313@164.com", "广州", "研发运维部",0, now(), now()),
 ("Z00001", "zh4", "zh4", md5("123"), 1, "12635464541442145", "1@qq.com", "广州", "研发测试部",0, now(), now()),
 ("Z00001", "zh5", "zh5", md5("123"), 1, "124176522145", "1@163.com", "广州", "运营",0, now(), now());
                                                                                                                                   ;


insert into user(staff_id, name,nickname, password,gender,tel,email,addr,department,status,created_at,updated_at) values
 ("Z00009", "tt9", "t9", md5("123"), 1, "1241441442145", "13313@qq.com", '广_州', "研发一部",0, now(), now());

select id, staff_id, name,nickname,gender,tel,email,addr,department,status,created_at,updated_at,deleted_at from user WHERE staff_id like  '%/%%'  ESCAPE '/' or name like  '%/%%'  ESCAPE '/' or nickname like '%/%%'  ESCAPE '/'  or email like ? ESCAPE '/'  or tel like '%/%%'  ESCAPE  '/'   or department like '%/%%' ESCAPE '/' or addr like '%/%%'  ESCAPE  '/';
