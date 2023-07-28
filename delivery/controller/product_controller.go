package controller

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/usecase"
	"enigma-laundry-apps/utils/exceptions"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type ProductController struct {
	productUC usecase.ProductUseCase
}

func (p *ProductController) HandlerMainForm() {
	fmt.Println(`
| ++++ Master Product ++++ |
| 1. Tambah Data           |
| 2. Lihat Data            |
| 3. Detail Data           |
| 4. Update Data           |
| 5. Hapus Data            |
| 6. Kembali ke Menu       |
	`)
	fmt.Print("Pilih Menu (1-6): ")

	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			uom := p.createHandlerForm()
			err := p.productUC.RegisterNewProduct(uom)
			exceptions.CheckErr(err)
			return
		case "2":
			requestPaging := dto.PaginationParam{
				Page: 1,
			}
			products, paging, err := p.productUC.FindAllProduct(requestPaging)
			exceptions.CheckErr(err)
			p.findAllHandlerForm(products, paging)
			return
		// case "3":
		// 	p.uomGetForm()
		// 	return
		// case "4":
		// 	uom := p.uomUpdateForm()
		// 	err := p.uomUC.UpdateUom(uom)
		// 	exceptions.CheckErr(err)
		// case "5":
		// 	id := p.uomDeleteForm()
		// 	err := p.uomUC.DeleteUom(id)
		// 	exceptions.CheckErr(err)
		case "6":
			return
		default:
			fmt.Println("Menu tidak ditemukan")
		}
	}
}

func (p *ProductController) createHandlerForm() model.Product {
	var (
		id, name, price, uomId, saveConfirmation string
	)

	fmt.Print("Product Name: ")
	fmt.Scanln(&name)

	fmt.Print("Product Price: ")
	fmt.Scanln(&price)

	fmt.Print("Product Uom ID: ")
	fmt.Scanln(&uomId)

	id = uuid.New().String()
	priceConv, _ := strconv.Atoi(price)

	fmt.Printf("Product Id: %s, Name: %s, Price: %d, UOM ID: %s akan disimpan? (y/t)", id, name, priceConv, uomId)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		product := model.Product{
			Id:    id,
			Name:  name,
			Price: priceConv,
			Uom:   model.Uom{Id: uomId},
		}
		return product
	}
	return model.Product{}
}

func (p *ProductController) findAllHandlerForm(products []model.Product, paging dto.Paging) {
	for _, product := range products {
		fmt.Println("Product List")
		fmt.Printf("ID: %s \n", product.Id)
		fmt.Printf("Name: %s \n", product.Name)
		fmt.Printf("Price: %d \n", product.Price)
		fmt.Printf("UOM ID: %s \n", product.Uom.Id)
		fmt.Printf("UOM Name: %s \n", product.Uom.Name)
		fmt.Println()
		fmt.Println("Paging: ")
		fmt.Printf("Page: %d \n", paging.Page)
		fmt.Printf("RowsPerPage: %d \n", paging.RowsPerPage)
		fmt.Printf("TotalPages: %d \n", paging.TotalPages)
		fmt.Printf("TotalRows: %d \n", paging.TotalRows)

	}
}

// func (p *ProductController) updateHandlerForm() model.Product {
// 	var (
// 		uomId, uomName, saveConfirmation string
// 	)

// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&uomId)
// 	fmt.Print("UOM Name: ")
// 	fmt.Scanln(&uomName)
// 	fmt.Printf("UOM Id: %s, Name: %s akan disimpan? (y/t)", uomId, uomName)
// 	fmt.Scanln(&saveConfirmation)

// 	if saveConfirmation == "y" {
// 		uom := model.Product{
// 			Id:   uomId,
// 			Name: uomName,
// 		}
// 		return uom
// 	}
// 	return model.Product{}
// }

// func (p *ProductController) uomDeleteForm() string {
// 	var id string
// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&id)
// 	return id
// }

// func (p *ProductController) uomGetForm() {
// 	var id string
// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&id)
// 	uom, err := p.uomUC.FindByIdUom(id)
// 	exceptions.CheckErr(err)
// 	fmt.Printf("UOM ID %s \n", id)
// 	fmt.Println(strings.Repeat("=", 15))
// 	fmt.Printf("UOM ID: %s \n", uom.Id)
// 	fmt.Printf("UOM Name: %s \n", uom.Name)
// }

func NewProductController(usecase usecase.ProductUseCase) *ProductController {
	return &ProductController{productUC: usecase}
}
