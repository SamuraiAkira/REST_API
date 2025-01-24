package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var mutex = &sync.Mutex{}

func saveCarsToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(cars)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newCar Car
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &newCar)

	newCar.ID = nextID
	nextID++

	cars = append(cars, newCar)

	// Сохраняем изменения в файл
	saveCarsToFile("cars.json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func GetCarsList(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

func GetCarByID(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	for _, car := range cars {
		if car.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func UpdateCarPut(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var updatedCar Car
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &updatedCar)

	for i, car := range cars {
		if car.ID == id {
			updatedCar.ID = car.ID
			cars[i] = updatedCar

			// Сохраняем изменения в файл
			saveCarsToFile("cars.json")

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func UpdateCarPatch(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var updates map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &updates)

	for i, car := range cars {
		if car.ID == id {
			for key, value := range updates {
				switch key {
				case "brandName":
					cars[i].BrandName = value.(string)
				case "modelName":
					cars[i].ModelName = value.(string)
				case "mileage":
					cars[i].Mileage = int(value.(float64))
				case "numberOfOwners":
					cars[i].NumberOfOwners = int(value.(float64))
				}
			}
			// Сохраняем изменения в файл
			saveCarsToFile("cars.json")

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	for i, car := range cars {
		if car.ID == id {
			cars = append(cars[:i], cars[i+1:]...)

			// Сохраняем изменения в файл
			saveCarsToFile("cars.json")

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
