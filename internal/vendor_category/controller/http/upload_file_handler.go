package vendor_category_http_handler

import (
	"Food-Delivery/pkg/common"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"path/filepath"
)

func (handler *vendorCategoryHandler) Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const maxUploadSize = 10 << 20 // 10MB is safer for Excel/CSV

		if err := ctx.Request.ParseMultipartForm(maxUploadSize); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(err))
			return
		}

		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(fmt.Errorf("file not found in request")))
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(err))
			return
		}
		defer file.Close()

		ext := filepath.Ext(fileHeader.Filename)
		switch ext {

		case ".csv":
			reader := csv.NewReader(file)
			records, err := reader.ReadAll()

			if err != nil {
				ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(err))
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"type":    "csv",
				"records": records,
			})

		case ".xlsx", ".xls":
			f, err := excelize.OpenReader(file)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(err))
				return
			}

			sheetName := f.GetSheetName(0)
			rows, err := f.GetRows(sheetName)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(err))
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"type":    "excel",
				"records": rows,
			})

		default:
			ctx.JSON(http.StatusBadRequest, common.ErrBadRequest(fmt.Errorf("unsupported file type: %s", ext)))
		}
	}
}
