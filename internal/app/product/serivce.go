package product

import (
	"errors"
	"net/http"
	"shopping-api/internal/dto"
	"shopping-api/internal/factory"
	"shopping-api/internal/model"
	"shopping-api/internal/repository"
	"shopping-api/pkg/helper"
	"strings"
)

type Service interface {
}

type service struct {
	ProductRepository repository.Product
	CityRepository    repository.City
}

func NewService(f *factory.Factory) Service {
	return &service{
		ProductRepository: f.ProductRepository,
		CityRepository:    f.CityRepository,
	}
}

func (s *service) CreateProduct(body dto.BodyCreateProducts, userID int) (code int, err error) {
	var product model.Products
	spaceEmptyName := strings.TrimSpace(body.Name)
	product.Name = strings.TrimRight(spaceEmptyName, "\r\n")
	if product.Name == "" {
		return http.StatusBadRequest, errors.New("product name is empty")
	}
	product.SubcategoryID = body.SubcategoryID
	product.CityID = body.CityID
	if body.Price <= 0 {
		return http.StatusBadRequest, errors.New("price is must be greater than 0")
	}
	product.Price = body.Price
	product.Description = strings.TrimSpace(body.Description)
	if product.Description == "" {
		return http.StatusBadRequest, errors.New("description is empty")
	}
	product.Stock = body.Stock
	if body.Stock <= 0 {
		return http.StatusBadRequest, errors.New("stock is must be greater than 0")
	}
	product.UsersID = uint(userID)
	getCity, _ := s.CityRepository.GetOneCity("name", "id = ?", product.CityID)
	lat, long, _ := helper.Geocode(getCity.City)
	product.Longitude = long
	product.Latitude = lat

	// bucket := "rentz-id" //your bucket name

	// ctx := appengine.NewContext(c.Request())

	// storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
	// }

	// // Multipart form
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	return err
	// }

	// files := form.File["photos"]
	// if files == nil {
	// 	return c.JSON(http.StatusBadRequest, response.ProductsBadGatewayResponse("must add photo"))
	// }

	// createdProduct, err := databases.CreateProduct(&product)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	// }

	// for _, file := range files {
	// 	if file.Size > constant.MAX_UPLOAD_SIZE {
	// 		databases.DeleteProduct(int(createdProduct.ID))
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 			"code":    http.StatusBadRequest,
	// 			"message": "The uploaded image is too big. Please use an image less than 1MB in size",
	// 		})
	// 	}

	// 	src, err := file.Open()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer src.Close()

	// 	if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
	// 		if file.Filename[len(file.Filename)-4:] != "jpeg" {
	// 			databases.DeleteProduct(int(createdProduct.ID))
	// 			return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 				"code":    http.StatusBadRequest,
	// 				"message": "The provided file format is not allowed. Please upload a JPG or JPEG or PNG image",
	// 			})
	// 		}
	// 	}

	// 	sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

	// 	if _, err := io.Copy(sw, src); err != nil {
	// 		return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
	// 	}

	// 	if err := sw.Close(); err != nil {
	// 		return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
	// 	}

	// 	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
	// 	}
	// 	photo := models.Photos{
	// 		Photo_Name: sw.Attrs().Name,
	// 		Url:        u.String(),
	// 		ProductsID: uint(createdProduct.ID),
	// 	}
	// 	_, err = databases.InsertPhoto(&photo)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	// 	}

	// }

	// // add product guarantee
	// for _, guarantee := range body.Guarantee {
	// 	var input = models.ProductsGuarantee{
	// 		ProductsID:  createdProduct.ID,
	// 		GuaranteeID: uint(guarantee),
	// 	}
	// 	_, err := databases.InsertGuarantee(&input)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	// 	}
	// }

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"code":       http.StatusOK,
	// 	"product_id": createdProduct.ID,
	// 	"message":    "product created and file uploaded successfully",
	// })

	return http.StatusOK, nil

}
