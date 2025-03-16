package service

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/mateussilvafreitas/orders/models"
	"github.com/mateussilvafreitas/orders/repository"
	"github.com/mateussilvafreitas/orders/utils"
)


func fetchProductsConcurrently(productsIds []int64) (map[int64]model.Product, error) {
	productMap := make(map[int64]model.Product)
	errChan := make(chan error, len(productsIds))
	
	for _, id := range productsIds {
		go func (pid int64) {
			product, err := repository.FindProductById(pid)
			if(err != nil) {
				errChan <- fmt.Errorf("product %d not found", pid)
				return
			}

			productMap[pid] = product
			errChan <- nil
		}(id)
	}

	for range productsIds {
		if err := <-errChan; err != nil {
			return nil, err
		}
	}

	return productMap, nil
}

func PostCreateOrder(po *gin.Context) {
	var createOrder model.CreateOrder

	if err := po.BindJSON(&createOrder); err != nil {
		utils.HandleError(po, 400, "Invalid request", err)
		return
	}

	client, err := repository.FindClientById(createOrder.ClientId)
	if err != nil {
		utils.HandleError(po, 404, fmt.Sprintf("The passed client does not exists. ID: %d", createOrder.ClientId), err)
		return
	}


	var productsIds []int64
	productQuantities := make(map[int64] int64)

	for _, p := range createOrder.Products {
		productsIds = append(productsIds, p.ProductId)
		productQuantities[p.ProductId] = p.Quantity
	}

	productMap, err := fetchProductsConcurrently(productsIds)

	if err != nil {
		utils.HandleError(po, 404, "One or more products not found", err)
		return
	}

	var totalValue float64
	var productsResponse []model.OrderProductResponse

	for _, p := range createOrder.Products {
		product := productMap[p.ProductId]
		quantity := productQuantities[p.ProductId]
		unitTotal := product.Price * float64(quantity)

		totalValue += unitTotal

		productsResponse = append(productsResponse, model.OrderProductResponse{ 
			Name: product.Name,
			Quantity: quantity,
		})
	}

	order := model.Order{
		ClientID: client.ID,
		Total: math.Round(totalValue * 100) / 100,
		DateOrder: time.Now().Format(time.RFC3339),
	}

	orderId, err := repository.SaveOrder(order)

	if err != nil {
		utils.HandleError(po, 500, "Error creating order", err)
		return
	}
	

	for _, p := range createOrder.Products {
		product := productMap[p.ProductId]
		_, err := repository.SaveOrderProduct(model.OrderProduct{
			OrderID: orderId,
			ProductID: p.ProductId,
			Quantity: productQuantities[p.ProductId],
			UnitaryPrice: product.Price,
			TotalPrice: product.Price * float64(productQuantities[p.ProductId]),
		})

		if err != nil {
			utils.HandleError(po, 500, "Error saving product in order", err)
			return
		}

	}

	orderResponse := model.OrderResponse{
		Id: orderId,
		ClientName: client.Name,
		DateOrder: order.DateOrder,
		Total: order.Total,
		Products: productsResponse,
	}

	po.IndentedJSON(201, orderResponse)
}

func GetOrderById(o *gin.Context){
	idOrderStr := o.Param("id")
	
	idOrder, err := strconv.ParseInt(idOrderStr, 10, 0)

	if err != nil {
		utils.HandleError(o, 400, "Invalid id", err)
		return
	}

	order, err := repository.FindOrderById(idOrder)

	if err != nil {
		utils.HandleError(o, 404, "Error finding order by id", err)
		return
	}

	productsOrder, err := repository.FindProductsFromOrder(idOrder)

	if err != nil {
		utils.HandleError(o, 500 , "Error finding products from order", err)
		return
	}

	client, err := repository.FindClientById(order.ClientID)

	if err != nil {
		utils.HandleError(o, 500, "Error finding client from order", err)
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
			utils.HandleError(o, 404, "Error finding product from order", err)
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