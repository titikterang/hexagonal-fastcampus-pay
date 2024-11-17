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
2. Topup / Deposit
> a. Check destination account status
> 
> b. Submit deposit via other bank virtual account
> 
> c. Got callback to banking service
> 
> d. Banking service will trigger transfer service to move some amount from corporate account number to user account number

3. Payment
> a. Check source account status
>
> b. Check destination account status
>
> c. validate balance on source account
>
> d. execute payment
> - payment into different bank account, send request to banking service
> - payment within fastcampus pay, handled by payment service
>

4. Settlement Flow
> a. on every transction could have some transaction fee
> 
> b. settlement service will calculate transaction fee if exist
> 
> c. settlement service will also create a cash movement record to be used as a trasnsaction journal
> 
> d. if we need to fetch mutation report, it will be provided by settlement service

---------------------------------