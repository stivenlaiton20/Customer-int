package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("secret_key")
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}





//login 
func loginHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Error al decodificar las credenciales", http.StatusBadRequest)
        return
    }

    //verificar las credenciales del usuario y autenticar
    if creds.Username == "usuario" && creds.Password == "contraseña" {
        expirationTime := time.Now().Add(5 * time.Minute) 
        claims := &Claims{
            Username: creds.Username,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: expirationTime.Unix(),
            },
        }
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        tokenString, err := token.SignedString(jwtKey)
        if err != nil {
            http.Error(w, "Error al generar el token", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
    } else {
        http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
    }
}










//Middleware de autenticación
func authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        if !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}


//Controllers de las rutas

func getAllPost(w http.ResponseWriter, r *http.Request) {
    response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
    if err != nil {
        http.Error(w, "Error al realizar la solicitud", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()
    var posts []map[string]interface{}
    err = json.NewDecoder(response.Body).Decode(&posts)
    if err != nil {
        http.Error(w, "Error al decodificar la respuesta", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(response.StatusCode)
    json.NewEncoder(w).Encode(posts)
}

func getOnePost(w http.ResponseWriter, r *http.Request) {
  
    id := mux.Vars(r)["id"]
    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)
    response, err := http.Get(url)
    if err != nil {
        http.Error(w, "Error al realizar la solicitud", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()
    var result map[string]interface{}
    err = json.NewDecoder(response.Body).Decode(&result)
    if err != nil {
        http.Error(w, "Error al decodificar la respuesta", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(response.StatusCode)
    json.NewEncoder(w).Encode(result)
}
func createPost(w http.ResponseWriter, r *http.Request) {
    var newPost Post
    err := json.NewDecoder(r.Body).Decode(&newPost)
    if err != nil {
        http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
        return
    }
    postData, err := json.Marshal(newPost)
    if err != nil {
        http.Error(w, "Error al convertir el nuevo post a JSON", http.StatusInternalServerError)
        return
    }
    response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postData))
    if err != nil {
        http.Error(w, "Error al realizar la solicitud POST", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()
    var createdPost Post
    err = json.NewDecoder(response.Body).Decode(&createdPost)
    if err != nil {
        http.Error(w, "Error al decodificar la respuesta JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(response.StatusCode)
    json.NewEncoder(w).Encode(createdPost)
}


func updatePost(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedPost Post
    err := json.NewDecoder(r.Body).Decode(&updatedPost)
    if err != nil {
        http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
        return
    }
    postData, err := json.Marshal(updatedPost)
    if err != nil {
        http.Error(w, "Error al convertir el post actualizado a JSON", http.StatusInternalServerError)
        return
    }

    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)
    client := &http.Client{}
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(postData))
    if err != nil {
        http.Error(w, "Error al crear la solicitud PUT", http.StatusInternalServerError)
        return
    }
    req.Header.Set("Content-Type", "application/json; charset=UTF-8")

    response, err := client.Do(req)
    if err != nil {
        http.Error(w, "Error al realizar la solicitud PUT", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    var updatedPostResponse Post
    err = json.NewDecoder(response.Body).Decode(&updatedPostResponse)
    if err != nil {
        http.Error(w, "Error al decodificar la respuesta JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(response.StatusCode)
    json.NewEncoder(w).Encode(updatedPostResponse)
}
func deletePostHandler(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)
    client := &http.Client{}
    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        http.Error(w, "Error al crear la solicitud DELETE", http.StatusInternalServerError)
        return
    }

    response, err := client.Do(req)
    if err != nil {
        http.Error(w, "Error al realizar la solicitud DELETE", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "El post con ID %s ha sido eliminado con éxito", id)
}

func main() {
    router := mux.NewRouter()

    // Rutas públicas
    router.HandleFunc("/login", loginHandler).Methods("POST")

    // Rutas privadas (requieren autenticacion)
	//obetenr todos los post
	router.HandleFunc("/posts", authenticate(getAllPost)).Methods("GET") 
	// Obetner un Post
    router.HandleFunc("/posts/{id}", authenticate(getOnePost)).Methods("GET")
	//Crear Post
	router.HandleFunc("/posts", authenticate(createPost)).Methods("POST")
	//actualizar Post
	router.HandleFunc("/posts/{id}", authenticate(updatePost)).Methods("PUT")
	//eliminar post
	router.HandleFunc("/posts/{id}", authenticate(deletePostHandler)).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}