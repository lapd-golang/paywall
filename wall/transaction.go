package wall

type InvoiceOptions struct {
	Price  int64
	Note   string
	Vendor string
}

var DefaultInvoiceOptions = InvoiceOptions{
	Price: 1,
	Note:  "API call",
}

type Transaction struct {
	// Transaction <- When user charges in
	UserIdentifier string
	Price          int64
	Invoice        InvoiceOptions
}

type RequestQuota struct {
	UserIdentifier string
	TotalQuota     int64
	RemainQuota    int64
}
