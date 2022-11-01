package MajorProject

import (
	"bytes"
	"ecommerce/details"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
)

// NewProd :This function adds the new product into the database with product description and image of that product.
func (h Handler) NewProd(c *gin.Context) {
	var prod details.ProductLine
	//a, _ := c.MultipartForm()
	image, _, _ := c.Request.FormFile("image")
	out := new(bytes.Buffer)
	_, err := io.Copy(out, image)
	if err != nil {
		log.Printf("copy file err:%s\n", err)
		return
	}
	prod = details.ProductLine{
		ProductLine:     c.PostForm("ProductLine"),
		TextDescription: c.PostForm("TextDescription"),
		HtmlDescription: c.PostForm("HtmlDescription"),
		Image:           out.Bytes(),
	}
	//if err := c.ShouldBindJSON(&prod); err != nil {
	if err := c.Bind(&prod); err != nil {
		c.JSON(http.StatusNotAcceptable, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
		})
		return
	}
	va := validator.New()
	if err := va.Struct(&prod); err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "CHECK FIELDS REQUIRE",
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   prod,
	})
	h.DB.Create(&prod)
	return
}

// Fproducts : Fetches all the products with the product descriptions.
func (h Handler) Fproducts(c *gin.Context) {
	var proc []details.ProductLine
	if err := h.DB.Model(&details.ProductLine{}).Find(&proc).Error; err != nil {
		c.JSON(http.StatusNotFound, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   proc,
	})
}
