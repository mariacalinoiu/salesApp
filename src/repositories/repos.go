package repositories

type (
	Partener struct {
		CodPartener  string `json:"CodPartener"`
		NumePartener string `json:"NumePartener"`
		CUI          string `json:"CUI"`
		Email        string `json:"Email"`
		IDAdresa     int    `json:"IDAdresa"`
	}

	Vanzare struct {
		IDIntrare   int     `json:"IDIntrare"`
		CodPartener string  `json:"CodPartener"`
		Status      string  `json:"Status"`
		DataIntrare string  `json:"DataIntrare"`
		DataLivrare string  `json:"DataLivrare"`
		Total       float32 `json:"Total"`
		VAT         float32 `json:"VAT"`
		Discount    float32 `json:"Discount"`
		Moneda      string  `json:"Moneda"`
		Platit      float32 `json:"Platit"`
		Comentarii  string  `json:"Comentarii"`
		CodVanzator int     `json:"CodVanzator"`
		IDSucursala int     `json:"IDSucursala"`
	}

	LinieVanzare struct {
		IDIntrare   int     `json:"IDIntrare"`
		NumarLinie  int     `json:"NumarLinie"`
		CodArticol  string  `json:"CodArticol"`
		Cantitate   int     `json:"Cantitate"`
		Pret        float32 `json:"Pret"`
		Discount    float64 `json:"Discount"`
		VAT         float64 `json:"VAT"`
		TotalLinie  float64 `json:"TotalLinie"`
		NumeProiect string  `json:"NumeProiect"`
	}

	Articol struct {
		CodArticol    string `json:"CodArticol"`
		NumeArticol   string `json:"NumeArticol"`
		CodGrupa      int    `json:"CodGrupa"`
		CantitateStoc int    `json:"CantitateStoc"`
		UnitateMasura string `json:"UnitateMasura"`
		Grupa         string `json:"Grupa"`
	}

	Sucursala struct {
		IDSucursala   int    `json:"IDSucursala"`
		NumeSucursala string `json:"NumeSucursala"`
		IDAdresa      int    `json:"IDAdresa"`
	}

	Vanzator struct {
		CodVanzator int     `json:"CodVanzator"`
		Nume        string  `json:"Nume"`
		Prenume     string  `json:"Prenume"`
		SalariuBaza float32 `json:"SalariuBaza"`
		Comision    float32 `json:"Comision"`
		Email       string  `json:"Email"`
		IDAdresa    int     `json:"IDAdresa"`
	}

	Adresa struct {
		IDAdresa   int    `json:"IDAdresa"`
		NumeAdresa string `json:"NumeAdresa"`
		Oras       string `json:"Oras"`
		Judet      string `json:"Judet"`
		Sector     string `json:"Sector"`
		Strada     string `json:"Strada"`
		Numar      string `json:"Numar"`
		Bloc       string `json:"Bloc"`
		Etaj       int    `json:"Etaj"`
	}
)