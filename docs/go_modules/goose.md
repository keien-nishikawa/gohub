# pressly/goose

## Commands

### Generate file

```bash
goose -dir db/migrations/ create AddSomeColumns sql

```

Some options

```bash
Options:
  -allow-missing
    	applies missing (out-of-order) migrations
  -certfile string
    	file path to root CA's certificates in pem format (only support on mysql)
  -dir string
    	directory with migration files (default ".")
  -h	print help
  -no-versioning
    	apply migration commands with no versioning, in file order, from directory pointed to
  -s	use sequential numbering for new migrations
  -ssl-cert string
    	file path to SSL certificates in pem format (only support on mysql)
  -ssl-key string
    	file path to SSL key in pem format (only support on mysql)
  -table string
    	migrations table name (default "goose_db_version")
  -v	enable verbose mode
  -version
    	print version
```

### FYI

https://github.com/pressly/goose
