package data

type ProductDTO struct {
	Barcode string
	Name    string
	Unit    string
	Amount  uint64
}

func (this *ProductDTO) GetBarcode() string {
	return this.Barcode
}

func (this *ProductDTO) SetBarcode(barcode string) {
	this.Barcode = barcode
}

func (this *ProductDTO) GetName() string {
	return this.Name
}

func (this *ProductDTO) SetName(name string) {
	this.Name = name
}

func (this *ProductDTO) GetUnit() string {
	return this.Unit
}

func (this *ProductDTO) SetUnit(unit string) {
	this.Unit = unit
}

func (this *ProductDTO) GetAmount() uint64 {
	return this.Amount
}

func (this *ProductDTO) SetAmount(amount uint64) {
	this.Amount = amount
}

func CreateProductDTO(barcode string, name string, unit string, amount uint64) *ProductDTO {
	return &ProductDTO{barcode, name, unit, amount}
}
