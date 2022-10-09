package controllers

import (
	"assignment2/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		orderItem structs.RequestCreate
		items     structs.Items
		order     structs.Orders
		item      []structs.ItemRequestCreate
		response  structs.ResponseOrder
		respItem  structs.ResponseItem
		result    gin.H
	)
	err := c.Bind(&orderItem)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	item = orderItem.Items
	order.OrderedAt = orderItem.OrderedAt
	order.CustomerName = orderItem.CustomerName
	err = idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	response.OrderID = order.OrderID
	response.OrderedAt = order.OrderedAt
	response.CustomerName = order.CustomerName
	for value := range item {
		items.ItemCode = item[value].ItemCode
		items.Description = item[value].Description
		items.Quantity = item[value].Quantity
		items.OrderId = order.OrderID
		items.Orders = order
		err := idb.DB.Create(&items).Error
		if err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
			return
		}
		respItem.ItemId = items.ItemId
		respItem.ItemCode = items.ItemCode
		respItem.Description = items.Description
		respItem.Quantity = items.Quantity
		response.Items = append(response.Items, respItem)
	}
	result = gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"data":    response,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders    []structs.Orders
		items     []structs.Items
		response  []structs.ResponseOrder
		respOrder structs.ResponseOrder
		respItem  structs.ResponseItem
		result    gin.H
	)
	err := idb.DB.Find(&orders).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	for value := range orders {
		err := idb.DB.Table("items").Where("order_id = ?", orders[value].OrderID).Find(&items).Error
		if err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
			return
		}
		respOrder.OrderID = orders[value].OrderID
		respOrder.OrderedAt = orders[value].OrderedAt
		respOrder.CustomerName = orders[value].CustomerName
		for value := range items {
			respItem.ItemId = items[value].ItemId
			respItem.ItemCode = items[value].ItemCode
			respItem.Description = items[value].Description
			respItem.Quantity = items[value].Quantity
			respOrder.Items = append(respOrder.Items, respItem)
		}
		response = append(response, respOrder)
	}
	if len(orders) <= 0 {
		result = gin.H{
			"result": "Data is empty",
		}
		c.JSON(http.StatusOK, result)
		return
	} else {
		result = gin.H{
			"message": "success",
			"status":  http.StatusOK,
			"data":    response,
		}
		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		orderItem structs.RequestUpdate
		items     structs.Items
		order     structs.Orders
		item      []structs.ItemRequestUpdate
		response  structs.ResponseOrder
		respItem  structs.ResponseItem
		result    gin.H
		orderId   = c.Param("orderId")
	)
	err := c.Bind(&orderItem)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	item = orderItem.Items
	order.OrderedAt = orderItem.OrderedAt
	order.CustomerName = orderItem.CustomerName
	err = idb.DB.Table("orders").Where("order_id = ?", orderId).Updates(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	response.OrderID, _ = strconv.Atoi(orderId)
	response.OrderedAt = order.OrderedAt
	response.CustomerName = order.CustomerName
	for value := range item {
		items.ItemCode = item[value].ItemCode
		items.Description = item[value].Description
		items.Quantity = item[value].Quantity
		items.OrderId, _ = strconv.Atoi(orderId)
		items.Orders = order
		err := idb.DB.Table("items").Where("item_id = ?", item[value].LineItemId).Updates(&items).Error
		if err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
			return
		}
		respItem.ItemId = items.ItemId
		respItem.ItemCode = items.ItemCode
		respItem.Description = items.Description
		respItem.Quantity = items.Quantity
		response.Items = append(response.Items, respItem)
	}
	result = gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"data":    response,
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order   structs.Orders
		item    structs.Items
		result  gin.H
		orderId = c.Param("orderId")
	)
	err := idb.DB.Table("items").Where("order_id = ?", orderId).Delete(&item).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = idb.DB.Table("orders").Where("order_id = ?", orderId).Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"message": "success",
		"status":  http.StatusOK,
	}
	c.JSON(http.StatusOK, result)
}
