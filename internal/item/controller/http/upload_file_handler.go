package item_http_handler

import (
	item_dto "Food-Delivery/entity/dto/item"
	"Food-Delivery/pkg/common"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
)

func (handler *itemHandler) Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const maxUploadSize = 10 << 20 // 10MB

		// Get file
		file, fileHeader, err := getUploadedFile(ctx, maxUploadSize)
		if err != nil {
			panic(common.ErrBadRequest(err))
			return
		}
		defer file.Close()

		// Get restaurant_id
		restaurantId, err := getIntFormField(ctx, "restaurant_id")
		if err != nil {
			panic(common.ErrBadRequest(err))
			return
		}

		// Get vendor_category_id
		vendorCategoryId, err := getIntFormField(ctx, "vendor_category_id")
		if err != nil {
			panic(common.ErrBadRequest(err))
			return
		}
		fmt.Println("restaurantId:", restaurantId)
		fmt.Println("vendorCategoryId:", vendorCategoryId)

		rows, err := getRows(ctx, file, fileHeader)

		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var nameIdx int
		var priceIdx int
		var descriptionIdx int
		var dtos []item_dto.CreateDTO
		for i, row := range rows {

			if i == 0 {
				for j, cell := range rows[0] {

					switch cell {
					case "name":
						nameIdx = j
						break
					case "price":
						priceIdx = j
						break
					case "description":
						descriptionIdx = j
						break
					}
				}
				continue
			}

			priceFloat, err := strconv.ParseFloat(row[priceIdx], 32)

			if err != nil {
				log.Printf("invalid price on row %d: %v", i, err)
				continue
			}

			dtos = append(dtos, item_dto.CreateDTO{
				RestaurantId:     restaurantId,
				VendorCategoryId: vendorCategoryId,
				Name:             row[nameIdx], // assuming name is in first column
				Price:            float32(priceFloat),
				Description:      &row[descriptionIdx],
				// Add more fields here as needed
			})
		}

		// Call service (uncomment this when ready)
		// handler.itemService.BatchCreate(ctx, dtos)
		ctx.JSON(http.StatusOK, common.Response(dtos))
	}
}

func getUploadedFile(ctx *gin.Context, maxUploadSize int64) (multipart.File, *multipart.FileHeader, error) {
	if err := ctx.Request.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, nil, err
	}

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return nil, nil, fmt.Errorf("file not found in request")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, nil, err
	}

	return file, fileHeader, nil
}

func getIntFormField(ctx *gin.Context, fieldName string) (int, error) {
	valueStr := ctx.PostForm(fieldName)
	if valueStr == "" {
		return 0, fmt.Errorf("%s is required", fieldName)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", fieldName)
	}

	return value, nil
}

func getRows(ctx *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) ([][]string, error) {
	var rows [][]string
	var err error
	ext := filepath.Ext(fileHeader.Filename)

	switch ext {
	case ".csv":
		reader := csv.NewReader(file)
		rows, err = reader.ReadAll()
		if err != nil {
			return nil, err
		}
	case ".xlsx", ".xls":
		f, err := excelize.OpenReader(file)
		if err != nil {
			return nil, err
		}
		sheetName := f.GetSheetName(0)
		rows, err = f.GetRows(sheetName)
		if err != nil {
			return nil, err
		}
	default:
		return nil, common.ErrBadRequest(fmt.Errorf("unsupported file type: %s", ext))
	}
	return rows, nil
}
