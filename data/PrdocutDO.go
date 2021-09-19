package data

type ProductDO struct {
	id      uint64
	barcode string
	name    string
	unit    string
	amount  uint64
}

func (this *ProductDO) GetId() uint64 {
	return this.id
}

func (this *ProductDO) SetId(id uint64) {
	this.id = id
}

func (this *ProductDO) GetBarcode() string {
	return this.barcode
}

func (this *ProductDO) SetBarcode(barcode string) {
	this.barcode = barcode
}

func (this *ProductDO) GetName() string {
	return this.name
}

func (this *ProductDO) SetName(name string) {
	this.name = name
}

func (this *ProductDO) GetUnit() string {
	return this.unit
}

func (this *ProductDO) SetUnit(unit string) {
	this.unit = unit
}

func (this *ProductDO) GetAmount() uint64 {
	return this.amount
}

func (this *ProductDO) SetAmount(amount uint64) {
	this.amount = amount
}

func CreateProductDO(id uint64, barcode string, name string, unit string, amount uint64) *ProductDO {
	return &ProductDO{id, barcode, name, unit, amount}
}
