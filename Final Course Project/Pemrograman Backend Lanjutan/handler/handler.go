package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var UserLogin = make(map[string]model.User)

// DESC: func Auth is a middleware to check user login id, only user that already login can pass this middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_login_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if _, ok := UserLogin[c.Value]; !ok || c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", c.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DESC: func AuthAdmin is a middleware to check user login role, only admin can pass this middleware
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
		c, err := r.Cookie("user_login_role")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if c.Value != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login role not Admin"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userRole", c.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IsUserExist(id string) bool {
	dataList, err := ioutil.ReadFile("data/users.txt") 
	if err != nil {
		panic(err)
	}

	for _, data := range strings.Split(string(dataList), "\n") {
		userId := strings.Split(string(data), "_")[0]
		
		if id == userId {
			return true
		}
	}
	return false
}

func IsStudyCodeExist(studyCode string) bool {
	dataList, err := ioutil.ReadFile("data/list-study.txt") 
	if err != nil {
		panic(err)
	}

	for _, data := range strings.Split(string(dataList), "\n") {
		if  strings.Split(string(data), "_")[0] == studyCode {
			return true
		}
	}
	return false
}

func GetUserRoleAndStudyCode(id string, name string) (string, string) {
	dataList, err := ioutil.ReadFile("data/users.txt") 
	if err != nil {
		panic(err)
	}

	role := "admin"
	studyCode := ""
	for _, data := range strings.Split(string(dataList), "\n") {
		userId := strings.Split(string(data), "_")[0]
		userName := strings.Split(string(data), "_")[1]
		
		if id == userId && name == userName {
			role = strings.Split(string(data), "_")[3]
			studyCode = strings.Split(string(data), "_")[2]
		}
	}
	return role, studyCode
}

func SaveUser(user model.User) {
	file, err := os.Open(filepath.Join("data", "users.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var users []model.User

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "_")
		users = append(users, model.User{ID: data[0], Name: data[1], Role: data[3], StudyCode: data[2]})
	}

	var listUser string
	for _, u := range users {
		listUser += u.ID + "_" + u.Name + "_" + u.StudyCode + "_" + u.Role + "\n"
	}
	if user.Role == "" {
		user.Role = "user"
	}
	listUser += user.ID + "_" + user.Name + "_" + user.StudyCode + "_" + user.Role

	if err = os.WriteFile(filepath.Join("data", "users.txt"), []byte(listUser), 0644); err != nil {
		panic(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	if user.ID == "" || user.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID or name is empty"})
		return
	} else if !IsUserExist(user.ID) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user not found"})
		return
	} else {
		user.Role, user.StudyCode = GetUserRoleAndStudyCode(user.ID, user.Name)
		UserLogin[user.ID] = user
		
		http.SetCookie(w, &http.Cookie{
			Name:    "user_login_id",
			Value:   user.ID,
			Path:  "/",
			Expires: time.Now().Add(365 * 24 * time.Hour),
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "user_login_role",
			Value:   user.Role,
			Path:  "/",
			Expires: time.Now().Add(365 * 24 * time.Hour),
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "login success"})
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	if user.ID == "" || user.Name == "" || user.Role == "" || user.StudyCode == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID, name, study code or role is empty"})
		return
	} else if user.Role != "admin" && user.Role != "user" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "role must be admin or user"})
		return
	} else if !IsStudyCodeExist(user.StudyCode) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "study code not found"})
		return
	} else if IsUserExist(user.ID) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id already exist"})
		return
	} else {
		SaveUser(user)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "register success"})
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	http.SetCookie(w, &http.Cookie{
        Name:    "user_login_id",
        Value:   "",
		Path:  "/",
		Expires: time.Now(),
    })
	http.SetCookie(w, &http.Cookie{
        Name:    "user_login_role",
        Value:   "",
		Path:  "/",
		Expires: time.Now(),
    })

	delete(UserLogin, userID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: userID, Message: "logout success"})
}

func GetStudyProgram(w http.ResponseWriter, r *http.Request) {
	// list study program
	dataList, err := ioutil.ReadFile("data/list-study.txt") 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var studyProgram []model.StudyData
	for _, data := range strings.Split(string(dataList), "\n") {
		value := strings.Split(string(data), "_")
		studyProgram = append(studyProgram, model.StudyData{Code: value[0], Name: value[1]})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) 

	err2 := json.NewEncoder(w).Encode(studyProgram)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	if user.ID == "" || user.Name == "" || user.StudyCode == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID, name, or study code is empty"})
		return
	} else if !IsStudyCodeExist(user.StudyCode) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "study code not found"})
		return
	} else if IsUserExist(user.ID) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id already exist"})
		return
	} else {
		SaveUser(user)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "add user success"})
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id is empty"})
		return
	} else if !IsUserExist(id) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id not found"})
		return
	} else {
		w.WriteHeader(http.StatusOK)

		dataList, err := ioutil.ReadFile("data/users.txt") 
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		file, err := os.Create("data/users.txt") 
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		for i:= 0; i < len(strings.Split(string(dataList), "\n")); i++ {
			
			userId := strings.Split(strings.Split(string(dataList), "\n")[i], "_")[0]
			data := strings.Split(string(dataList), "\n")[i] + "\n"
			
			if id == userId {
				continue
			} else {
				if i == len(strings.Split(string(dataList), "\n"))-1 {
					_, err = file.WriteString(data[0:len(data)-1])
				} else {
					_, err = file.WriteString(data)
				}

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
		file.Close() 

		json.NewEncoder(w).Encode(model.SuccessResponse{Username: id, Message: "delete success"})
		return
	}
}

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeatherByChannel(region string, ch chan model.MainWeather, chErr chan error) {
	weather, err := GetWetherByRegionAPI(region)
	weatherData := &weather
	ch <- *weatherData
	chErr <- err
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	var listRegion = []string{"jakarta", "bandung", "surabaya", "yogyakarta", "medan", "makassar", "manado", "palembang", "semarang", "bali"}

	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	ch := make(chan model.MainWeather, len(listRegion))
	errCh := make(chan error)
	for _, region := range listRegion {
		go GetWeatherByChannel(region, ch, errCh)
	}

	weatherData := []model.MainWeather{}
	var err error
	for i:= 0; i < len(listRegion); i++ {
		err = <-errCh
		if err == nil {
			weather := <-ch
			weatherData = append(weatherData, weather)
		} else {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherData)
}




