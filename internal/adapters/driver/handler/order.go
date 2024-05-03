package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new order
// @Description Create new order. To make an order the payment needs to be completed
// @Tags Order
// @Accept json
// @Produce json
// @Param product body domain.Order true "order"
// @Success 200 {object} domain.OrderResponse
// @Failure 400 "Order has required fields"
// @Router /api/orders [post]
func CreateOrderHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order domain.Order

		err := httpserver.DecodeJSONBody(w, r, &order)

		if err != nil {
			log.Print("decoding order body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		response, err := orderService.CreateOrder(context.Background(), order)

		if err != nil {
			log.Print("create order", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

// @Summary Get order by Id
// @Description Get an order by Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Success 200 {object} domain.OrderResponse
// @Failure 400 "Order has required fields"
// @Router /api/orders/{id} [get]
func GetOrderByIdHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("get order by id path", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		orderId, err := strconv.Atoi(orderIdStr)

		if err != nil {
			log.Print("get order by id path", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		response, err := orderService.GetOrderById(context.Background(), uint(orderId))

		if err != nil {
			log.Print("get order by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

// @Summary Get all orders to prepare
// @Description Get all orders already payed that needs to be prepared. This endpoint will be used by the kitchen
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Router /api/orders/to-prepare [get]
func GetOrdersToPrepareHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := orderService.GetOrdersToPrepare(context.Background())

		if err != nil {
			log.Print("get orders to prepare", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

// @Summary Update an order to PREPARING
// @Description Update an order. This service wil be used by the kitchen to notify a customer that the order is being prepared
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Success 204
// @Failure 404 "Order not found"
// @Router /api/orders/{id}/preparing [put]
func UpdateOrderPreparingHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		id, err := GetOrderId(idStr)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		err = orderService.UpdateToPreparing(context.Background(), id)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Update an order to DONE
// @Description Update an order. This service wil be used by the kitchen to notify a customer and the waiter that the order is done
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Success 204
// @Failure 404 "Order not found"
// @Router /api/orders/{id}/done [put]
func UpdateOrderDoneHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		id, err := GetOrderId(idStr)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		err = orderService.UpdateToDone(context.Background(), id)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Update an order to DELIVERED
// @Description Update an order. This service wil be used by the waiter to close the order informing that user got its order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Success 204
// @Failure 404 "Order not found"
// @Router /api/orders/{id}/delivered [put]
func UpdateOrderDeliveredHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		id, err := GetOrderId(idStr)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		err = orderService.UpdateToDelivered(context.Background(), id)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Update an order to NOT_DELIVERED
// @Description Update an order. This service wil be used by the waiter to close the order informing that user didn't get the order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Success 204
// @Failure 404 "Order not found"
// @Router /api/orders/{id}/not-delivered [put]
func UpdateOrderNotDeliveredandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		id, err := GetOrderId(idStr)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		err = orderService.UpdateToNotDelivered(context.Background(), id)

		if err != nil {
			log.Print("update order status", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

func GetOrderId(orderdStr string) (uint, error) {
	orderId, err := strconv.Atoi(orderdStr)

	if err != nil {
		return 0, err
	}

	return uint(orderId), nil
}
