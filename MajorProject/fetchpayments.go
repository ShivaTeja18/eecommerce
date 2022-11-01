package MajorProject

import (
	"ecommerce/details"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// FromtoTill : This function fetches all the payments by dates provided by client.
func (h Handler) FromtoTill(c *gin.Context) {
	var fpay []details.Payment
	from := c.Request.FormValue("payment_date")
	parse, err := time.Parse("2006-01-02", from)
	if err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "check requires in parsing time parse",
			Error:  err.Error(),
		})
		return
	}
	toDate := c.Request.FormValue("to")
	to, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "check requires in parsing time to",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}
	if from == "" || toDate == "" {
		c.JSON(http.StatusNotAcceptable, details.Response{
			Status: "FAILURE",
			Error:  "Fields should not be empty",
			Code:   http.StatusNotAcceptable,
			Data:   nil,
		})
		return
	}

	if err := h.DB.Model(&details.Payment{}).Where("payment_date BETWEEN ? AND ?", parse, to).Find(&fpay).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "check require",
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &fpay,
	})
	return
}
