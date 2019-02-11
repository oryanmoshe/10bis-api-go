package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type TenBisRes struct {
	Errors  []string `json:"Errors"`
	Success bool     `json:"Success"`
	Data    Data     `json:"Data"`
}
type Data struct {
	RestaurantsList []Restaurant `json:"restaurantsList`
	CategoriesList  []Category   `json:"categoriesList`
}
type Restaurant struct {
	Id       int32  `json:"restaurantId"`
	Name     string `json:"restaurantName"`
	Address  string `json:"restaurantAddress"`
	CityName string `json:"restaurantCityName"`
}
type Category struct {
	Id       int32  `json:"categoryId"`
	Desc     string `json:"categoryDesc"`
	Name     string `json:"categoryName"`
	DishList []Dish `json:"dishList"`
}

type Dish struct {
	Id          int32  `json:"dishId"`
	Price       int32  `json:"dishPrice"`
	Description string `json:"dishDescription"`
	Name        string `json:"dishName"`
	ImageUrl    string `json:"dishImageUrl"`
	DishList    []Dish `json:"dishList"`
}

type Login struct {
	Model     LoginModel `json:"model"`
	ReturnURL string     `json:"returnUrl"`
}
type LoginModel struct {
	UserName string `json:"UserName`
	Password string `json:"Password`
}

var host = "https://www.10bis.co.il/"
var api = "NextApi/"
var search = "SearchRestaurants"
var menu = "GetRestaurantMenu"
var login = "Account/LogOnAjax"

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/getRestaurants", func(c *gin.Context) {
		url := host + api + search + "?addressId=3525498&cityId=24&cityName=%D7%AA%D7%9C+%D7%90%D7%91%D7%99%D7%91+%D7%99%D7%A4%D7%95&streetId=23322&streetName=%D7%93%D7%A8%D7%9A+%D7%9E%D7%A0%D7%97%D7%9D+%D7%91%D7%92%D7%99%D7%9F&houseNumber=52&entrance=%D7%91%D7%A0%D7%99%D7%99%D7%9F+%D7%A1%D7%95%D7%A0%D7%95%D7%9C&floor=13&longitude=34.7814426&latitude=32.0637232&phone01=0&isCompanyAddress=true&addressCompanyId=5969&addressKey=24-23322-52-3525498&deliveryMethod=Delivery&enableNewResAndCouponsPromotion=true"
		resp, err := http.Get(url)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		c.Header("Content-Type", "application/json")
		var objmap TenBisRes
		err = json.Unmarshal(body, &objmap)
		restList, err := json.Marshal(objmap.Data.RestaurantsList)
		// fmt.Println(string(restList))
		if err != nil && restList == nil {
			fmt.Println(err)
		}
		c.Data(http.StatusOK, "application/json", restList)
	})

	r.GET("/getDishes", func(c *gin.Context) {
		restId := c.Request.URL.Query()["restaurantId"][0]
		url := host + api + menu + "?restaurantId=" + restId + "&deliveryMethod=Delivery"
		resp, err := http.Get(url)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		c.Header("Content-Type", "application/json")
		var objmap TenBisRes
		err = json.Unmarshal(body, &objmap)
		restList, err := json.Marshal(objmap.Data.CategoriesList)
		// fmt.Println(string(restList))
		if err != nil && restList == nil {
			fmt.Println(err)
		}
		c.Data(http.StatusOK, "application/json", restList)
	})

	r.POST("/wut", func(c *gin.Context) {
		restId := c.Request.URL.Query()["restaurantId"][0]
		url := host + api + menu + "?restaurantId=" + restId + "&deliveryMethod=Delivery"
		resp, err := http.Get(url)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		c.Header("Content-Type", "application/json")
		var objmap TenBisRes
		err = json.Unmarshal(body, &objmap)
		restList, err := json.Marshal(objmap.Data.CategoriesList)
		// fmt.Println(string(restList))
		if err != nil && restList == nil {
			fmt.Println(err)
		}
		c.Data(http.StatusOK, "application/json", restList)
	})

	r.GET("/login", func(c *gin.Context) {
		userName := c.Request.URL.Query()["userName"][0]
		password := c.Request.URL.Query()["password"][0]
		wut := []byte(`{"model": {"UserName":"` + userName + `", "Password":"` + password + `"}}`)
		url := host + login
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(wut))
		fmt.Println("gothere")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		// c.Header("Content-Type", "application/json")
		// var objmap TenBisRes
		// err = json.Unmarshal(body, &objmap)
		// restList, err := json.Marshal(objmap.Data.CategoriesList)
		// fmt.Println(string(restList))
		if err != nil {
			fmt.Println(body, wut)
		}
		uid := resp.Cookies()[1].Value
		tenbisCtx := resp.Cookies()[4].Value
		fmt.Println(uid)
		fmt.Println(tenbisCtx)
		c.Data(http.StatusOK, "application/json", []byte(tenbisCtx))
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
