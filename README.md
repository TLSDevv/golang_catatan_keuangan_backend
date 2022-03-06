# KeuanganKu Backend

## To Do
- [ * ] Refactor handler and repository to be inside a folder called gateways
- [*] Use validator to validate body request from user
- [*] Create mysql helper
- [*] Use https://github.com/cosmtrek/air to have live reloading
- [*] Implement user role
- [*] Remove user_id from transaction payload, we could get it from authorization header
- [*] Protect transactions endpoint with authorization

## Structure Rules
These coding rules below is inspired by simple-bank structure.
- Each method on repository, service and handler should be divided into separated file (look on simple-bank repo), the reason why we do this
is because so that it would be easy for us (the coder) to read the code and also it looks clean, another thing is each method has its own
_test file, isn't that great man? for me yes.
- Filename that represented the method must be the same with its method name.
    For example:
    ```go
    func (r Repository) GetTransactions() ([]entities.Transaction, error) {
        //... the rest of the code
    }
    ```
    the method above should have file name called get_transactions.go
- Each method naming should represent the functionallity
