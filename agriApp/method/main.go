package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	
)

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Stock    int
}

type Sale struct {
	ID          int
	ProductName string
	Quantity    int
	Buyer       string
	Date        string
	Total       float64
}

type Schedule struct {
	ID         int
	FarmerName string
	CropName   string
	LandArea   string
	Task       string
	Date       string
	Status     string
}

type PageData struct {
	Products  []Product
	Sales     []Sale
	Schedules []Schedule
}

var (
	db   *sql.DB
	tmpl *template.Template
)

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./agri.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTables()

	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-product", addProductHandler)
	http.HandleFunc("/add-sale", addSaleHandler)
	http.HandleFunc("/add-schedule", addScheduleHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTables() {
	productTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		category TEXT,
		price REAL,
		stock INTEGER
	);`

	saleTable := `
	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT,
		quantity INTEGER,
		buyer TEXT,
		date TEXT,
		total REAL
	);`

	scheduleTable := `
	CREATE TABLE IF NOT EXISTS schedules (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		farmer_name TEXT,
		crop_name TEXT,
		land_area TEXT,
		task TEXT,
		date TEXT,
		status TEXT
	);`

	_, err := db.Exec(productTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(saleTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(scheduleTable)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Products:  getProducts(),
		Sales:     getSales(),
		Schedules: getSchedules(),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	category := r.FormValue("category")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	_, err := db.Exec("INSERT INTO products(name, category, price, stock) VALUES(?, ?, ?, ?)", name, category, price, stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addSaleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	productName := r.FormValue("product_name")
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	buyer := r.FormValue("buyer")
	date := r.FormValue("date")
	total, _ := strconv.ParseFloat(r.FormValue("total"), 64)

	_, err := db.Exec("INSERT INTO sales(product_name, quantity, buyer, date, total) VALUES(?, ?, ?, ?, ?)", productName, quantity, buyer, date, total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	farmerName := r.FormValue("farmer_name")
	cropName := r.FormValue("crop_name")
	landArea := r.FormValue("land_area")
	task := r.FormValue("task")
	date := r.FormValue("date")
	status := r.FormValue("status")

	_, err := db.Exec("INSERT INTO schedules(farmer_name, crop_name, land_area, task, date, status) VALUES(?, ?, ?, ?, ?, ?)",
		farmerName, cropName, landArea, task, date, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getProducts() []Product {
	rows, err := db.Query("SELECT id, name, category, price, stock FROM products ORDER BY id DESC")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock)
		products = append(products, p)
	}
	return products
}

func getSales() []Sale {
	rows, err := db.Query("SELECT id, product_name, quantity, buyer, date, total FROM sales ORDER BY id DESC")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var sales []Sale
	for rows.Next() {
		var s Sale
		rows.Scan(&s.ID, &s.ProductName, &s.Quantity, &s.Buyer, &s.Date, &s.Total)
		sales = append(sales, s)
	}
	return sales
}

func getSchedules() []Schedule {
	rows, err := db.Query("SELECT id, farmer_name, crop_name, land_area, task, date, status FROM schedules ORDER BY id DESC")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var schedules []Schedule
	for rows.Next() {
		var s Schedule
		rows.Scan(&s.ID, &s.FarmerName, &s.CropName, &s.LandArea, &s.Task, &s.Date, &s.Status)
		schedules = append(schedules, s)
	}
	return schedules
}
