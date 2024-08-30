# Desafio_BrandMonitor
Aplicação console testada via
<a href="https://www.postman.com/" target="_blank">Postman</a>

## Operações CRUD

- ✅ Create: CreateUser( )
- ✅ Read: FindUserByID( ), FindUserByEmail( )
- ✅ Update: UpdateUser( )
- ✅ Delete: DeleteUser( )

## Extra (Opcional)
- ✅ Utilizar um banco de dados MongoDB

## Exemplo de Requisição

### POST /createUser
**Requisição:**
```json
{
  "email": "teste10@gmail.com",
  "password": "123456@",
  "name": "Novo Nome",
  "age": 32
}
```

### PUT /updateUser/id
**Requisição**
```json
{
  "name": "Muda nome",
  "age": 29
}
```
### GET /getUserById/id

### GET /getUserByEmail/email

### DELETE /deleteUser/id

