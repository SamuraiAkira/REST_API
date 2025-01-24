package pkg

type Car struct {
    ID             int64  `json:"id"`
    BrandName      string `json:"brandName"`
    ModelName      string `json:"modelName"`
    Mileage        int    `json:"mileage"`
    NumberOfOwners int    `json:"numberOfOwners"`
}

var (
    cars    []Car
    nextID  int64 = 1
)
