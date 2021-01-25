package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"crud/banco"

	"github.com/gorilla/mux"
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

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, error := banco.Conect()
	if error != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	lines, error := db.Query("select * from usuarios")
	if error != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
		return
	}
	defer lines.Close()

	var usuarios []usuario

	for lines.Next() {
		var usuario usuario
		if error := lines.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); error != nil {
			w.Write([]byte("Erro ao ler informações do banco de dados"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(usuarios); error != nil {
		w.Write([]byte("Erro ao converter os usuários"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, error := strconv.ParseUint(parametros["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Falha au ler o parâmetro para inteiro"))
		return
	}

	db, error := banco.Conect()
	if error != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	lines, error := db.Query("select * from usuarios where id = ?", ID)
	if error != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	var usuario usuario
	if lines.Next() {
		lines.Scan(&usuario.ID, &usuario.Nome, &usuario.Email)
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Uauário não encontrado!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(usuario); error != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, error := strconv.ParseUint(parametros["id"], 10, 32)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao ler o corpo do request"))
		return
	}
	var usuario usuario
	if error := json.Unmarshal(body, &usuario); error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os dados a serem atualizados"))
		return
	}

	db, error := banco.Conect()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("update usuarios set nome=?, email=? where id=?")
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao preparar o statement"))
		return
	}
	defer statement.Close()

	atualizacao, error := statement.Exec(usuario.Nome, usuario.Email, ID)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	idInserido, error := atualizacao.LastInsertId()
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Falha ao recuperar o usuário atualizado"))
		return
	}

	if idInserido == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
