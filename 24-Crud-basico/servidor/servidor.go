package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"crud/banco"
)

type usuario struct {
	ID    int32  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if error = json.Unmarshal(body, &usuario); error != nil {
		w.Write([]byte("Erro ao converter o usuário"))
		return
	}

	db, error := banco.Conect()
	if error != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("insert into usuarios(nome, email) values (?, ?)")
	if error != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, error := statement.Exec(usuario.Nome, usuario.Email)
	if error != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, error := insercao.LastInsertId()
	if error != nil {
		w.Write([]byte("Erro ao recuperar o id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso!!! Id inserido %d\n", idInserido)))
}

// 	fmt.Println(db)
