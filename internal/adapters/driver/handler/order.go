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
