# Backend Go Masterclass Notes

---

This file contains all the notes/tips/guides taken from the masterclass as a reference for future projects.
It include commands, workflows, advanced topics etc.

### Create a Makefile

By using a makefile, we can store commonly used commands and reuse them easily.

1. Download postgres image from docker hub
2. Create DB
3. Create migration folder with files inside db folder, and migrate up
4. Run the sqlc docker init command to generate a sqlc.yaml file
5. Create a query folder inside db folder, and create schemas sql file. This file will function-ize the queries so that sqlc can use it.
6. Run the sqlc docker generate command to generate the sqlc go files. It will create db.go, models.go, and the respective ${schema}.go files.

### For testing, we use the testify/require package

### TestMain

It is a special function which takes a testing.M object as input. It is the main entry point of all unit tests inside a specific golang package.

=="Cannot connect to the database:sql: unknown driver "postgres" (forgotten import?)"==
[Answer](https://stackoverflow.com/questions/52789531/how-do-i-solve-panic-sql-unknown-driver-postgres-forgotten-import)

==To avoid values of db row from getting used by multiple different operations in different transactions, use FOR NO KEY UPDATE in the SQL statement.Only once the transaction is committed or rolled back, can these db rows get read by other operations==

When tables are linked to each other via reference/foreign keys, any insert/select statements on those tables will result in deadlock if the linked table is in an uncommitted transaction. Basically, DEADLOCKS can be caused due to foreign key constraints.

#### How to avoid DB deadlocks?

When working with db transactions, locking and handling deadlocks is a tricky thing. We should try to fine tune our queries in the transaction to avoid/minimize the occurrence of deadlocks.

---

There are several ways a transaction can be interfered with by other transactions that run simultaneously. This is known as **read phenomenon**.

- **Dirty Read**
  A transaction reads data written by other concurrent uncommitted transactions. This is bad because we do not know if the other transactions would commit or rollback and hence might use incorrect data.

- **Non-Repeatable Read**
  A transaction reads the same row twice and sees different value because it has been modified by other committed transactions.

- **Phantom Read**
  A transaction re-executes a query to find rows that satisfy a condition and sees a different set of rows, due to changes by other committed transactions.
  Similar to Non-Repeatable Read but affects queries that search for multiple rows instead of one.

- **Serialization Anomaly**
  The result of a group of concurrent committed transactions is impossible to achieve if we try to run them sequentially in any order without overlapping.'

### Standard Isolation Levels

In order of lowest to highest:

- Read Uncommitted
- Read Committed
- Repeatable Read
- Serializable

==Watch video009 of working with database module for explanations==
