# Vinance DB Migration Tool

## How To

### _Generate_
Run the below script to generate a new migration file
```
go run main.go generate <migration_name>
```
Example:
```
go run main.go generate create_users_table
```

What it will do:
- Create 2 new files; `<timestamp>_<name>_up.sql` for migration & `<timestamp>_<name>_down.sql` for rollback

Constraints:
- `<migration_name>` is mandatory
- `<migration_name>` needs to be underscore (_) not space ( ) based

### _Migrate_
Run the below script to run the latest migration version
```
go run main.go migrate
```

What it will do:
- Run the latest un-ran version migration file
- After successfully ran, the db will record the migration by version
