package service

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/mateussilvafreitas/orders/models"
	"github.com/mateussilvafreitas/orders/repository"
)

func PostCreateOrder(po *gin.Context) {
	var createOrder model.CreateOrder
	var orderResponse model.OrderResponse

	if err := po.BindJSON(&createOrder); err != nil {
		return
	}

	client, err := repository.FindClientById(createOrder.ClientId)

	if err != nil {
		po.IndentedJSON(422, gin.H{"message": fmt.Sprintf("The passed client does not exists. ID: %d", createOrder.ClientId)})
		return
	}
	
	orderResponse.ClientName = client.Name


	var products []model.Product
	var totalValue = 0.0
	mapProductQuantity := make(map[int64] int64)

	for _,p := range createOrder.Products {
		product, err := repository.FindProductById(p.ProductId)

		if err != nil {
			po.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error creating order, product %d does not exists", p.ProductId)})
			return
		}

		products = append(products, product)
		totalValue += product.Price * float64(p.Quantity)
		mapProductQuantity[p.ProductId] = p.Quantity
	}

	var order model.Order
	var date = time.Now()
	order.Total = math.Round(totalValue * 100) / 100
	order.ClientID = createOrder.ClientId
	order.DateOrder = date.Format(time.RFC3339)

	orderId, err := repository.SaveOrder(order)

	orderResponse.Id = orderId
	orderResponse.DateOrder = date.Format(time.DateOnly)
	orderResponse.Total = order.Total

	if err != nil {
		po.IndentedJSON(500, gin.H{"message": fmt.Sprintf("Error creating order: %v", err)})
		return
	}


	var productsResponse []model.OrderProductResponse
	for _, p := range products {
		var orderProduct model.OrderProduct
		orderProduct.OrderID = orderId
		orderProduct.ProductID = p.ID
		orderProduct.Quantity = mapProductQuantity[p.ID]
		orderProduct.UnitaryPrice = p.Price
		orderProduct.TotalPrice = p.Price * float64(orderProduct.Quantity)

		_, err := repository.SaveOrderProduct(orderProduct);

		if err != nil {
			po.IndentedJSON(500, gin.H{"message": fmt.Sprintf("Error creating the product order: %v", err)})
			return
		}

		var orderProductResponse model.OrderProductResponse
		orderProductResponse.Name = p.Name
		orderProductResponse.Quantity = orderProduct.Quantity

		productsResponse = append(productsResponse, orderProductResponse)
	}

	orderResponse.Products = productsResponse

	po.IndentedJSON(201, orderResponse)
}

func GetOrderById(o *gin.Context){
	idOrderStr := o.Param("id")
	
	idOrder, err := strconv.ParseInt(idOrderStr, 10, 0)

	if err != nil {
		o.IndentedJSON(400, gin.H{"message": "Invalid id"})
		return
	}

	order, err := repository.FindOrderById(idOrder)

	if err != nil {
		o.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error finding order by id: %v", err)})
		return
	}

	productsOrder, err := repository.FindProductsFromOrder(idOrder)

	if err != nil {
		o.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error finding order by id: %v", err)})
		return
	}

	client, err := repository.FindClientById(order.ClientID)

	if err != nil {
		o.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error finding client from order: %v", err)})
		return
	}

	var orderResponse model.OrderResponse
	orderResponse.Id = order.ID
	orderResponse.DateOrder = order.DateOrder
	orderResponse.Total = order.Total
	orderResponse.ClientName = client.Name

	var orderProductResponse []model.OrderProductResponse

	for _, po := range productsOrder {
		product, err := repository.FindProductById(po.ProductID)

		if err != nil {
			o.IndentedJSON(404, gin.H{"message": fmt.Sprintf("Error finding product from order: %v", err)})
			return
		}

		orderProductResponse = append(orderProductResponse, model.OrderProductResponse{
			Name: product.Name,
			Quantity: po.Quantity,
		})
	}

	orderResponse.Products = orderProductResponse

	o.IndentedJSON(200, orderResponse)
}