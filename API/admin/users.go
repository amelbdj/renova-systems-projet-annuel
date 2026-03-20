package admin

import (
	"encoding/json"
	"fmt"

	"upcycleconnect/auth"
	"upcycleconnect/bdd"

	"net/http"
	"strconv"
	"upcycleconnect/models"
)
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	fmt.Println("hello from login")
	
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Erreur décodage :", err)
		http.Error(w, "Impossible de décoder le JSON", http.StatusBadRequest)
		return
	}

	userBdd, err := bdd.LoginUser(user.Email, user.MotDePasse)
	if err != nil {
		fmt.Println("Erreur login :", err)
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(userBdd.Id, userBdd.Role)
	if err != nil {
		fmt.Println("Erreur génération token :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from GetAllUsers")

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	users, err := bdd.GetUsers()

	if err != nil {
		http.Error(w, "erreur de récupération des utilisateurs", http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

    fmt.Println("hello from createUser")

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        fmt.Println("Erreur décodage :", err)
        http.Error(w, "Impossible de décoder le JSON", http.StatusBadRequest)
        return
    }
		user.MotDePasse, err = auth.HashPassword(user.MotDePasse)
		if err != nil {
			fmt.Println("Erreur hashage mot de passe :", err)
			http.Error(w, "Erreur lors du hashage du mot de passe", http.StatusInternalServerError)
			return
		}
    
    err = bdd.CreateUser(user) 
    if err != nil {
        fmt.Println("Erreur lors de l'insertion en BDD :", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, `user creer`)
    
}

func DeletedUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("hello from DeletedUser")

	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	err = bdd.DeletedUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "utilisateur suppr")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

		fmt.Println("hello from updateUser")


	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"ID invalide"}`, http.StatusBadRequest)
		return
	}

	var userDto models.User
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		fmt.Println("Erreur décodage :", err)
		return
		
	}


	userDto.Id = id

	if err := bdd.UpdateUserById(userDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	user, err := bdd.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	role := r.PathValue("role")

	users, err := bdd.GetUserByRole(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	nameQuery := r.URL.Query().Get("name") // recup les valeurs de la query string(diff de path variable)
	roleQuery := r.URL.Query().Get("role")

	users, err := bdd.GetUserByName(nameQuery, roleQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}