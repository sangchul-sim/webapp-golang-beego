# webapp-golang-beego

# Quick Start
Download and install
```
go get github.com/astaxie/beego
go get github.com/go-sql-driver/mysql
go get github.com/sangchul-sim/webapp-golang-beego
```

Database schema
```
CREATE TABLE `tb_deal_info` (
  `deal_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `deal_name` varchar(32) NOT NULL DEFAULT '',
  `price` int(10) unsigned NOT NULL,
  `sale_price` int(10) unsigned NOT NULL,
  `sale_start_dt` datetime NOT NULL,
  `sale_end_dt` datetime NOT NULL,
  PRIMARY KEY (`deal_id`),
  KEY `idx__deal_name` (`deal_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

Sample data
```
INSERT INTO `tb_deal_info` (`deal_id`, `deal_name`, `price`, `sale_price`, `sale_start_dt`, `sale_end_dt`)
VALUES
	(1, '고려홍삼정 365 6년근 스틱', 29500, 28020, '2016-08-01 00:00:00', '2016-09-01 00:00:00'),
	(2, '베지밀 비 달콤한 두유 190ml', 15500, 14725, '2016-08-15 00:00:00', '2016-09-15 00:00:00'),
	(3, '에스까다 옴므 하이드라 EX 3종 세트', 20000, 19800, '2016-08-01 00:00:00', '2016-09-01 00:00:00'),
	(4, '삼성농산 백찹쌀떡', 4800, 4560, '2016-08-01 00:00:00', '2016-09-15 00:00:00'),
	(5, '보닌 더 스타일 남성 화장품 2종 기획세트', 25000, 21750, '2016-08-01 00:00:00', '2016-08-15 00:00:00'),
	(6, '맥심 모카골드마일드 커피세트', 22000, 20960, '2016-08-01 00:00:00', '2016-08-18 00:00:00'),
	(7, '농협 껍질째먹는 웰빙 세척사과 3kg', 21500, 19900, '2016-08-01 00:00:00', '2016-08-25 00:00:00'),
	(8, '휴럼 제주감귤파이 선물세트 37g 40개입+종이쇼핑백', 32000, 31000, '2016-08-01 00:00:00', '2016-09-13 00:00:00'),
	(9, '종근당건강 헛개153 골드', 14050, 13348, '2016-08-01 00:00:00', '2016-09-15 00:00:00');

```

Build and run
```
bee run
```
