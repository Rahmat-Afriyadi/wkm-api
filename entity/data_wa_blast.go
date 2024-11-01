package entity

type DataWaBlast struct {
	NoMsn            string
	NmCustomer       string
	KdUser           *string
	NoYgDiHubRenewal *string
	NoWa             string
	TglAkhirTenor    string
}

type SearchNoMsnByWa struct {
	NoMsn      string
	NmCustomer string
	NoWa       string
}
