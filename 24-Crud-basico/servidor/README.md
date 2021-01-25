insere um usuário
curl -X POST localhost:5000/usuarios --header "Content-Type: application/json" --data '{"id":123,"nome":"Fulano","email":"fulano@gmail.com"}'

atualiza um usuário
curl -X PUT localhost:5000/usuarios/1 --header "Content-Type: application/json" --data '{"nome":"Joãozinho","email":"joaozinho@gmail.com"}'

consulta todos os usuários
curl -X GET localhost:5000/usuarios --header "Content-Type: application/json"                                                                

deleta um usuário
curl -X DELETE localhost:5000/usuarios/1 --header "Content-Type: application/json"