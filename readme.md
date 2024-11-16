# Membership Service - Setup Project & Repository
1. Setup Monorepo
2. 2 Public API, Login, Registrasi
3. 1 Private API, get user profile
4. Prerequisites
   - protoc , generate protobuf
   - golang CI/lint, untuk code quality
   - Setup postgreSQL
   - Setup redis


# Overall Business Flow
---------------------------------
1. Transfer
> a. Check source account status
>
> b. Check destination account status
>
> c. validate balance on source account
>
> d. execute transfer
> - transfer into different bank account, send request to banking service
> - transfer within fastcampus pay, handled by transfer service
>
3. Topup / Deposit
   a.
3. Payment

4.

# Money Service
---------------------------------