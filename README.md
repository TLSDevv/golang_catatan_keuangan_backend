# golang_catatan_keuangan_backend
Backend untuk aplikasi catatan keuangan

# System Requirements
- Go
- PostgresDB

## Tecnology Stack
* Http Router: github.com/go-chi/chi
* JWT library: github.com/dgrijalva/jwt-go
* Ansible: https://docs.ansible.com/ansible/latest/installation_guide/index.html

# Initial Setup
- Download and install dependenies
```
make install
```
- Manual migrate table, 

# Run
- Create config files
```
make config
```

## Deployment

To deploy this project run

```bash
  go run main.go
```