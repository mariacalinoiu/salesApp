package datasources

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/sijms/go-ora"

	"github.com/mariacalinoiu/salesApp/src/repositories"
)

type DBClient struct {
	db *sql.DB
}

func GetClient(user string, password string, hostname string, dbName string) DBClient {
	db, err := sql.Open(
		"oracle",
		fmt.Sprintf("oracle://%s:%s@%s/%s", user, password, hostname, dbName),
	)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3000)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	return DBClient{db: db}
}

func (client DBClient) GetParteneri() ([]repositories.Partener, error) {
	var (
		parteneri []repositories.Partener
		cod       string
		nume      string
		cui       string
		email     string
		IDAdresa  int
	)

	rows, err := client.db.Query(
		`SELECT "CodPartener", "NumePartener", "CUI", "EMail", "IdAdresa" FROM "Parteneri"`,
	)
	if err != nil {
		return []repositories.Partener{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cod, &nume, &cui, &email, &IDAdresa)
		if err != nil {
			return []repositories.Partener{}, err
		}

		parteneri = append(
			parteneri,
			repositories.Partener{
				CodPartener:  cod,
				NumePartener: nume,
				CUI:          cui,
				Email:        email,
				IDAdresa:     IDAdresa,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.Partener{}, err
	}

	return parteneri, nil
}

func (client DBClient) InsertPartener(partener repositories.Partener) error {
	stmt, err := client.db.Prepare("INSERT INTO Parteneri(CodPartener, NumePartener, CUI, Email, IdAdresa) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		partener.CodPartener,
		partener.NumePartener,
		partener.CUI,
		partener.Email,
		partener.IDAdresa,
	)

	return err
}

func (client DBClient) GetVanzari() ([]repositories.Vanzare, error) {
	var (
		vanzari     []repositories.Vanzare
		id          int
		codPartener string
		status      string
		data        string
		dataLivrare string
		total       float32
		vat         float32
		discount    float32
		moneda      string
		platit      float32
		comentarii  string
		codVanzator int
		IDSucursala int
	)

	rows, err := client.db.Query(
		"SELECT IdIntrare, CodPartener, Status, DataIntrare, DataLivrare, Total, Vat, Discount, Moneda, Platit, Comentarii, CodVanzator, IdSucursala FROM Vanzari",
	)
	if err != nil {
		return []repositories.Vanzare{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &codPartener, &status, &data, &dataLivrare, &total, &vat, &discount, &moneda, &platit, &comentarii, &codVanzator, &IDSucursala)
		if err != nil {
			return []repositories.Vanzare{}, err
		}

		vanzari = append(
			vanzari,
			repositories.Vanzare{
				IDIntrare:   id,
				CodPartener: codPartener,
				Status:      status,
				Data:        data,
				DataLivrare: dataLivrare,
				Total:       total,
				VAT:         vat,
				Discount:    discount,
				Moneda:      moneda,
				Platit:      platit,
				Comentarii:  comentarii,
				CodVanzator: codVanzator,
				IDSucursala: IDSucursala,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.Vanzare{}, err
	}

	return vanzari, nil
}

func (client DBClient) GetLiniiVanzare(IDIntrareVanzari int) ([]repositories.LinieVanzare, error) {
	var (
		liniiVanzare []repositories.LinieVanzare
		IDIntrare    int
		numarLinie   int
		codArticol   string
		cantitate    float32
		pret         float32
		discount     float32
		VAT          float32
		totalLinie   float32
		IDProiect    string
	)

	rows, err := client.db.Query(
		"SELECT IdIntrare, NumarLinie, CodArticol, Cantitate, Pret, Discount, Vat, TotalLinie, IdProiect FROM LiniiVanzari WHERE IdIntrare = ?",
		IDIntrareVanzari,
	)
	if err != nil {
		return []repositories.LinieVanzare{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&IDIntrare, &numarLinie, &codArticol, &cantitate, &pret, &discount, &VAT, &totalLinie, &IDProiect)
		if err != nil {
			return []repositories.LinieVanzare{}, err
		}

		liniiVanzare = append(
			liniiVanzare,
			repositories.LinieVanzare{
				IDIntrare:  IDIntrare,
				NumarLinie: numarLinie,
				CodArticol: codArticol,
				Cantitate:  cantitate,
				Pret:       pret,
				Discount:   discount,
				VAT:        VAT,
				TotalLinie: totalLinie,
				IDProiect:  IDProiect,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.LinieVanzare{}, err
	}

	return liniiVanzare, nil
}

func (client DBClient) GetArticole() ([]repositories.Articol, error) {
	var (
		articole        []repositories.Articol
		cod             string
		nume            string
		codGrupa        int
		cantitateStoc   int
		IDUnitateMasura int
	)

	rows, err := client.db.Query(
		"SELECT CodArticol, NumeArticol, CodGrupa, CantitateStoc, IdUnitateDeMasura FROM Articole",
	)
	if err != nil {
		return []repositories.Articol{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cod, &nume, &codGrupa, &cantitateStoc, &IDUnitateMasura)
		if err != nil {
			return []repositories.Articol{}, err
		}

		articole = append(
			articole,
			repositories.Articol{
				CodArticol:      cod,
				NumeArticol:     nume,
				CodGrupa:        codGrupa,
				CantitateStoc:   cantitateStoc,
				IDUnitateMasura: IDUnitateMasura,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.Articol{}, err
	}

	return articole, nil
}

func (client DBClient) InsertArticol(articol repositories.Articol) error {
	stmt, err := client.db.Prepare("INSERT INTO Articole(CodArticol, NumeArticol, CodGrupa, CantitateStoc, IdUnitateMasura) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		articol.CodArticol,
		articol.NumeArticol,
		articol.CodGrupa,
		articol.CantitateStoc,
		articol.IDUnitateMasura,
	)

	return err
}

func (client DBClient) GetVanzatori() ([]repositories.Vanzator, error) {
	var (
		vanzatori   []repositories.Vanzator
		codVanzator int
		nume        string
		prenume     string
		salariuBaza float32
		comision    float32
		email       string
		IDAdresa    int
	)

	rows, err := client.db.Query(
		"SELECT CodVanzator, Nume, Prenume, SalariuBaza, Comision, Email, IdAdresa FROM Vanzatori",
	)
	if err != nil {
		return []repositories.Vanzator{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&codVanzator, &nume, &prenume, &salariuBaza, &comision, &email, &IDAdresa)
		if err != nil {
			return []repositories.Vanzator{}, err
		}

		vanzatori = append(
			vanzatori,
			repositories.Vanzator{
				CodVanzator: codVanzator,
				Nume:        nume,
				Prenume:     prenume,
				SalariuBaza: salariuBaza,
				Comision:    comision,
				Email:       email,
				IDAdresa:    IDAdresa,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.Vanzator{}, err
	}

	return vanzatori, nil
}

func (client DBClient) GetSucursale() ([]repositories.Sucursala, error) {
	var (
		sucursale   []repositories.Sucursala
		IDSucursala int
		nume        string
		IDAdresa    int
	)

	rows, err := client.db.Query(
		"SELECT IdSucursala, NumeSucursala, IdAdresa FROM Sucursale",
	)
	if err != nil {
		return []repositories.Sucursala{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&IDSucursala, &nume, &IDAdresa)
		if err != nil {
			return []repositories.Sucursala{}, err
		}

		sucursale = append(
			sucursale,
			repositories.Sucursala{
				IDSucursala:   IDSucursala,
				NumeSucursala: nume,
				IDAdresa:      IDAdresa,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.Sucursala{}, err
	}

	return sucursale, nil
}

func (client DBClient) GetProiecte() ([]repositories.Proiect, error) {
	var (
		proiecte    []repositories.Proiect
		IDProiect   string
		nume        string
		validDeLa   string
		validPanaLa string
		activ       string
	)

	rows, err := client.db.Query(
		"SELECT IDProiect, NumeProiect, ValidDeLa, ValidPanaLa, Activ FROM Proiecte",
	)
	if err != nil {
		return []repositories.Proiect{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&IDProiect, &nume, &validDeLa, &validPanaLa, &activ)
		if err != nil {
			return []repositories.Proiect{}, err
		}

		proiecte = append(
			proiecte,
			repositories.Proiect{
				IDProiect:   IDProiect,
				NumeProiect: nume,
				ValidDeLa:   validDeLa,
				ValidPanaLa: validPanaLa,
				Activ:       activ,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.Proiect{}, err
	}

	return proiecte, nil
}

func (client DBClient) InsertProiect(proiect repositories.Proiect) error {
	stmt, err := client.db.Prepare("INSERT INTO Proiecte(IdProiect, NumeProiect, ValidDeLa, ValidPanaLa, Activ) VALUES(?, ?, TO_DATE(?, 'MM/DD/YYYY'), TO_DATE(?, 'MM/DD/YYYY'), ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		proiect.IDProiect,
		proiect.NumeProiect,
		proiect.ValidDeLa,
		proiect.ValidPanaLa,
		proiect.Activ,
	)

	return err
}

func (client DBClient) GetGrupeArticole() ([]repositories.GrupaArticole, error) {
	var (
		grupe []repositories.GrupaArticole
		cod   int
		nume  string
	)

	rows, err := client.db.Query(
		"SELECT CodGrupa, NumeGrupa FROM GrupaArticole",
	)
	if err != nil {
		return []repositories.GrupaArticole{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cod, &nume)
		if err != nil {
			return []repositories.GrupaArticole{}, err
		}

		grupe = append(
			grupe,
			repositories.GrupaArticole{
				CodGrupa:  cod,
				NumeGrupa: nume,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.GrupaArticole{}, err
	}

	return grupe, nil
}

func (client DBClient) GetUnitatiDeMasura() ([]repositories.UnitateDeMasura, error) {
	var (
		um       []repositories.UnitateDeMasura
		id       int
		nume     string
		inaltime float32
		latime   float32
		lungime  float32
	)

	rows, err := client.db.Query(
		"SELECT IdUnitateMasura, NumeUnitateDeMasura, Inaltime, Latime, Lungime FROM UnitatiDeMasura",
	)
	if err != nil {
		return []repositories.UnitateDeMasura{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &nume, &inaltime, &latime, &lungime)
		if err != nil {
			return []repositories.UnitateDeMasura{}, err
		}

		um = append(
			um,
			repositories.UnitateDeMasura{
				IDUnitateMasura:     id,
				NumeUnitateDeMasura: nume,
				Inaltime:            inaltime,
				Latime:              latime,
				Lungime:             lungime,
			},
		)
	}

	err = rows.Err()
	if err != nil {
		return []repositories.UnitateDeMasura{}, err
	}

	return um, nil
}

//func (client DBClient) GetProductsByCategoryID(categoryID int) (repositories.ProductsJSON, error) {
//	var (
//		products    []repositories.Product
//		id          int
//		name        string
//		imageURL    string
//		description string
//		price       float32
//	)
//
//	rows, err := client.db.Query(
//		"SELECT ID, name, imageURL, description, price FROM Products WHERE categoryID = ?",
//		categoryID,
//	)
//	if err != nil {
//		return repositories.ProductsJSON{Products: products}, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&id, &name, &imageURL, &description, &price)
//		if err != nil {
//			return repositories.ProductsJSON{Products: products}, err
//		}
//
//		products = append(
//			products,
//			repositories.Product{
//				ID:          id,
//				Name:        name,
//				ImageURL:    imageURL,
//				Description: description,
//				Price:       price,
//				CategoryID:  categoryID,
//			},
//		)
//	}
//
//	err = rows.Err()
//	if err != nil {
//		return repositories.ProductsJSON{Products: products}, err
//	}
//
//	return repositories.ProductsJSON{Products: products}, nil
//}
//
//func (client DBClient) GetCategoriesByDepartmentID(departmentID int) (repositories.CategoriesJSON, error) {
//	var (
//		categories []repositories.Category
//		id         int
//		name       string
//	)
//
//	rows, err := client.db.Query(
//		"SELECT ID, name FROM Categories WHERE departmentID = ?",
//		departmentID,
//	)
//	if err != nil {
//		return repositories.CategoriesJSON{Categories: categories}, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&id, &name)
//		if err != nil {
//			return repositories.CategoriesJSON{Categories: categories}, err
//		}
//
//		categories = append(
//			categories,
//			repositories.Category{
//				ID:           id,
//				Name:         name,
//				DepartmentId: departmentID,
//			},
//		)
//	}
//
//	err = rows.Err()
//	if err != nil {
//		return repositories.CategoriesJSON{Categories: categories}, err
//	}
//
//	return repositories.CategoriesJSON{Categories: categories}, nil
//}
//
//func (client DBClient) GetDepartments() (repositories.DepartmentsJSON, error) {
//	var (
//		departments []repositories.Department
//		id          int
//		name        string
//	)
//
//	rows, err := client.db.Query(
//		"SELECT ID, name FROM Departments",
//	)
//	if err != nil {
//		return repositories.DepartmentsJSON{Departments: departments}, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&id, &name)
//		if err != nil {
//			return repositories.DepartmentsJSON{Departments: departments}, err
//		}
//
//		departments = append(
//			departments,
//			repositories.Department{
//				ID:   id,
//				Name: name,
//			},
//		)
//	}
//
//	err = rows.Err()
//	if err != nil {
//		return repositories.DepartmentsJSON{Departments: departments}, err
//	}
//
//	return repositories.DepartmentsJSON{Departments: departments}, nil
//}
//
//func (client DBClient) InsertOrder(order repositories.Order) (repositories.OrderIDResponse, error) {
//	var res driver.Result
//	var err error
//
//	if len(order.VoucherCode) > 0 {
//		if !client.isVoucherValid(order.VoucherCode) {
//			return repositories.OrderIDResponse{OrderID: 0}, errors.New("the voucher code provided is invalid")
//		}
//
//		stmt, err := client.db.Prepare("INSERT INTO Orders(firstName, lastName, email, phoneNumber, city, address, voucherCode, paymentMethod, status, timestamp) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
//		if err != nil {
//			return repositories.OrderIDResponse{OrderID: 0}, err
//		}
//		res, err = stmt.Exec(
//			order.FirstName,
//			order.LastName,
//			order.Email,
//			order.PhoneNumber,
//			order.City,
//			order.Address,
//			order.VoucherCode,
//			order.PaymentMethod,
//			order.Status,
//			int(time.Now().UnixNano()/1000000000),
//		)
//	} else {
//		stmt, err := client.db.Prepare("INSERT INTO Orders(firstName, lastName, email, phoneNumber, city, address, paymentMethod, status, timestamp) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
//		if err != nil {
//			return repositories.OrderIDResponse{OrderID: 0}, err
//		}
//		res, err = stmt.Exec(
//			order.FirstName,
//			order.LastName,
//			order.Email,
//			order.PhoneNumber,
//			order.City,
//			order.Address,
//			order.PaymentMethod,
//			order.Status,
//			int(time.Now().UnixNano()/1000000000),
//		)
//	}
//
//	if err != nil {
//		return repositories.OrderIDResponse{OrderID: 0}, err
//	}
//	orderID, err := res.LastInsertId()
//
//	stmt, err := client.db.Prepare("INSERT INTO ProductOrders(orderID, productID, quantity) VALUES(?, ?, ?)")
//	if err != nil {
//		return repositories.OrderIDResponse{OrderID: 0}, err
//	}
//
//	for _, product := range order.ProductsOrdered {
//		_, err = stmt.Exec(
//			orderID,
//			product.ProductID,
//			product.Quantity,
//		)
//		if err != nil {
//			return repositories.OrderIDResponse{OrderID: 0}, err
//		}
//	}
//
//	return repositories.OrderIDResponse{OrderID: int(orderID)}, nil
//}
//
//func (client DBClient) EditOrder(order repositories.Order) error {
//	isVoucherValid := len(order.VoucherCode) == 0 || client.isVoucherValid(order.VoucherCode)
//	if !isVoucherValid {
//		return errors.New("the voucher code provided is invalid")
//	}
//
//	stmt, err := client.db.Prepare("UPDATE Orders SET firstName = ?, lastName = ?, email = ?, phoneNumber = ?, city = ?, address = ?, voucherCode = ?, paymentMethod = ?, status = ? WHERE ID = ?")
//	if err != nil {
//		return err
//	}
//
//	_, err = stmt.Exec(
//		order.FirstName,
//		order.LastName,
//		order.Email,
//		order.PhoneNumber,
//		order.City,
//		order.Address,
//		order.VoucherCode,
//		order.PaymentMethod,
//		order.Status,
//		order.ID,
//	)
//
//	return err
//}
//
//func (client DBClient) DeleteOrder(orderID int) error {
//	_, err := client.db.Exec(
//		"DELETE FROM ProductOrders WHERE orderID = ?",
//		orderID,
//	)
//	if err != nil {
//		return err
//	}
//
//	_, err = client.db.Exec(
//		"DELETE FROM Orders WHERE ID = ?",
//		orderID,
//	)
//
//	return err
//}
//
//func (client DBClient) GetOrders(orderIDProvided ...int) (repositories.OrdersJSON, error) {
//	var (
//		orderRows *sql.Rows
//		err       error
//
//		orders             []repositories.Order
//		orderID            int
//		firstName          string
//		lastName           string
//		email              string
//		phoneNumber        string
//		city               string
//		address            string
//		voucherCode        *string
//		paymentMethod      string
//		status             string
//		timestamp          int
//		discountPercentage *int
//	)
//
//	query := `
//		SELECT o.ID, o.firstName, o.lastName, o.email, o.phoneNumber, o.city, o.address, o.voucherCode, o.paymentMethod, o.status, o.timestamp, v.discountPercentage
//		FROM Orders o
//		LEFT JOIN Vouchers v
//		ON o.voucherCode = v.code
//	`
//
//	if len(orderIDProvided) == 1 {
//		orderRows, err = client.db.Query(query+" WHERE ID = ?", orderIDProvided[0])
//	} else {
//		orderRows, err = client.db.Query(query)
//	}
//	if err != nil {
//		return repositories.OrdersJSON{Orders: orders}, err
//	}
//
//	defer orderRows.Close()
//	for orderRows.Next() {
//		err := orderRows.Scan(&orderID, &firstName, &lastName, &email, &phoneNumber, &city, &address, &voucherCode, &paymentMethod, &status, &timestamp, &discountPercentage)
//		if err != nil {
//			return repositories.OrdersJSON{Orders: orders}, err
//		}
//		products, totalValue, err := client.getOrderedProducts(orderID)
//		if err != nil {
//			return repositories.OrdersJSON{Orders: orders}, err
//		}
//
//		code := ""
//		discount := 0
//		if voucherCode != nil {
//			code = *voucherCode
//			discount = *discountPercentage
//		}
//
//		orders = append(
//			orders,
//			repositories.Order{
//				ID:                 orderID,
//				FirstName:          firstName,
//				LastName:           lastName,
//				Email:              email,
//				PhoneNumber:        phoneNumber,
//				City:               city,
//				Address:            address,
//				VoucherCode:        code,
//				DiscountPercentage: discount,
//				PaymentMethod:      paymentMethod,
//				Status:             status,
//				Timestamp:          timestamp,
//				Date:               ParseTimestamp(timestamp),
//				Value:              totalValue * 100 / (100 + float32(discount)),
//				ProductsOrdered:    products,
//			},
//		)
//	}
//
//	err = orderRows.Err()
//	if err != nil {
//		return repositories.OrdersJSON{Orders: orders}, err
//	}
//
//	return repositories.OrdersJSON{Orders: orders}, nil
//}
//
//func (client DBClient) getOrderedProducts(orderID int) ([]repositories.OrderedProduct, float32, error) {
//	var (
//		products    []repositories.OrderedProduct
//		productID   int
//		quantity    int
//		name        string
//		imageURL    string
//		description string
//		price       float32
//		categoryID  int
//	)
//
//	totalValue := float32(0)
//	productOrderRows, err := client.db.Query(`
//			SELECT po.productID, po.quantity, p.name, p.imageURL, p.description, p.price, p.categoryID
//			FROM ProductOrders po, Products p
//			WHERE po.productID = p.ID AND orderID = ?
//		`,
//		orderID,
//	)
//	if err != nil {
//		return products, totalValue, err
//	}
//
//	for productOrderRows.Next() {
//		err := productOrderRows.Scan(&productID, &quantity, &name, &imageURL, &description, &price, &categoryID)
//		if err != nil {
//			fmt.Println(err.Error())
//			return products, totalValue, err
//		}
//
//		totalValue += price
//
//		products = append(
//			products,
//			repositories.OrderedProduct{
//				ProductID: productID,
//				OrderID:   orderID,
//				Quantity:  quantity,
//				Product: repositories.Product{
//					ID:          productID,
//					Name:        name,
//					ImageURL:    imageURL,
//					Description: description,
//					Price:       price,
//					CategoryID:  categoryID,
//				},
//			},
//		)
//	}
//
//	productOrderRows.Close()
//
//	err = productOrderRows.Err()
//	if err != nil {
//		return products, totalValue, err
//	}
//
//	return products, totalValue, nil
//}
//
//func (client DBClient) isVoucherValid(voucherCode string) bool {
//	rows, err := client.db.Query(
//		"SELECT discountPercentage FROM Vouchers WHERE code = ?",
//		voucherCode,
//	)
//	if err != nil {
//		return false
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		return true
//	}
//
//	return false
//}
