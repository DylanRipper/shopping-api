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

	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateProduct(body dto.BodyCreateProducts, userID int) (code int, err error)
	GetAllProducts() (interface{}, error)
	GetProductByID(productID int) (product *dto.GetProduct, err error)
	GetProductsBySubcategoryID(subCategoryID int) (interface{}, error)
	GetProductsByUserID(userID int) (interface{}, error)
	DeleteProductByID(productID, userID int) error
	GetProductsByName(name string) (interface{}, error)
}

type service struct {
	ProductRepository repository.Product
	CityRepository    repository.City
	PhotoRepository   repository.Photo
}

func NewService(f *factory.Factory) Service {
	return &service{
		ProductRepository: f.ProductRepository,
		CityRepository:    f.CityRepository,
		PhotoRepository:   f.PhotoRepository,
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

	// bucket := util.GetEnv("GOOGLE_BUCKET_NAME", "") //your bucket name

	// ctx := appengine.NewContext(c.Request())

	// storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// // Multipart form
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	return err
	// }

	// files := form.File["photos"]
	// if files == nil {
	// 	return http.StatusInternalServerError, errors.New("missing photos")
	// }

	_, err = s.ProductRepository.CreateProduct(&product)
	if err != nil {
		logrus.Error(err)
		return http.StatusInternalServerError, err
	}

	// for _, file := range files {
	// 	if file.Size > constant.MAX_UPLOAD_SIZE {
	// 		databases.DeleteProduct(int(createdProduct.ID))
	// 		return http.StatusBadRequest, errors.New("size image is too big")
	// 	}

	// 	src, err := file.Open()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer src.Close()

	// 	if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
	// 		if file.Filename[len(file.Filename)-4:] != "jpeg" {
	// 			databases.DeleteProduct(int(createdProduct.ID))
	// 			return http.StatusBadRequest, errors.New("format image is invalid")
	// 		}
	// 	}

	// 	sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

	// 	if _, err := io.Copy(sw, src); err != nil {
	// 		return http.StatusInternalServerError, err
	// 	}

	// 	if err := sw.Close(); err != nil {
	// 		return http.StatusInternalServerError, err
	// 	}

	// 	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	// 	if err != nil {
	// 		return http.StatusInternalServerError, err
	// 	}
	// 	photo := models.Photos{
	// 		Photo_Name: sw.Attrs().Name,
	// 		Url:        u.String(),
	// 		ProductsID: uint(createdProduct.ID),
	// 	}
	// 	_, err = databases.InsertPhoto(&photo)
	// 	if err != nil {
	// 		return http.StatusInternalServerError, err
	// 	}

	// }

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"code":       http.StatusOK,
	// 	"product_id": createdProduct.ID,
	// 	"message":    "product created and file uploaded successfully",
	// })

	return http.StatusOK, nil

}

func (s *service) GetAllProducts() (interface{}, error) {
	product, err := s.ProductRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return product, nil
}

func (s *service) GetProductByID(productID int) (product *dto.GetProduct, err error) {
	product, err = s.ProductRepository.GetProductByID(uint(productID))
	if err != nil {
		return nil, err
	}

	product.Url, _ = s.PhotoRepository.GetUrl(uint(productID))

	return product, nil
}

func (s *service) GetProductsBySubcategoryID(subCategoryID int) (interface{}, error) {
	product, err := s.ProductRepository.GetProductsBySubcategoryID(subCategoryID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProductsByUserID(userID int) (interface{}, error) {
	product, err := s.ProductRepository.GetProductsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) DeleteProductByID(productID, userID int) error {

	userId, err := s.ProductRepository.GetProductOwner(productID)
	if err != nil {
		return err
	}

	if userID != userId {
		return errors.New("access forbidden")
	}

	_, err = s.ProductRepository.DeleteProductByID(productID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetProductsByName(name string) (interface{}, error) {
	product, err := s.ProductRepository.GetProductsByName(name)
	if err != nil {
		return nil, err
	}

	return product, nil
}
