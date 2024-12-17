package handler

// import (
// 	partnerdto "kredit-plus/internal/dto/partner"
// 	dto "kredit-plus/internal/dto/result"
// 	transactiondetaildto "kredit-plus/internal/dto/transaction-detail.go"
// 	"kredit-plus/internal/models"
// 	repositories "kredit-plus/internal/repository"
// 	errorhandler "kredit-plus/pkg/error"
// 	"strconv"

// 	// "log"
// 	"net/http"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/labstack/echo/v4"
// )

// type handlerTransactionDetail struct {
// 	TransactionDetailRepository repositories.TransactionDetail
// 	AdminAuthRepository         repositories.AdminAuth
// 	CustomerAuthRepository      repositories.CustomerAuth
// }

// func HandlerTransactionDetail(
// 	TransactionDetailRepository repositories.TransactionDetail,
// 	AdminAuthRepository repositories.AdminAuth,
// 	CustomerAuthRepository repositories.CustomerAuth) *handlerTransactionDetail {
// 	return &handlerTransactionDetail{
// 		TransactionDetailRepository: TransactionDetailRepository,
// 		AdminAuthRepository:         AdminAuthRepository,
// 		CustomerAuthRepository:      CustomerAuthRepository}
// }

// func (h *handlerTransactionDetail) CreateTransactionDetail(c echo.Context) error {
// 	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
// 	accessLoginID := accessLogin["id"].(float64)

// 	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
// 	if err != nil {
// 		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
// 	}
// 	// catch json  type
// 	request := new(transactiondetaildto.RequestTransactionDetail)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	// Step 2: Bind the incoming JSON payload to the.
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	error := c.Validate(request)

// 	if error != nil {
// 		return errorhandler.ErrorHandler(c, error, error.Error(), http.StatusBadRequest)
// 	}

// 	partner := models.TransactionDetail{
// 		// 		ContractNumber
// 		// OTR
// 		// AdminFee
// 		// InstallmentAmount
// 		// InterestAmount
// 		// Status
// 		// CreditLimitID
// 		// PartnerBankID
// 		// AssetID
// 		// PartnerID
// 		// CustomerID

// 	}

// 	creditlimit, err := h.TransactionDetailRepository.CreateTransactionDetail(partner)
// 	if err != nil {
// 		return errorhandler.ErrorHandler(c, err, "Failed Create Limit", http.StatusBadRequest)
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
// }

// func (h *handlerTransactionDetail) ListTransactionDetail(c echo.Context) error {
// 	accessLogin := c.Get("adminLogin").(jwt.MapClaims)
// 	accessLoginID := accessLogin["id"].(float64)

// 	_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))
// 	if err != nil {
// 		return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
// 	}

// 	creditlimit, err := h.TransactionDetailRepository.ListTransactionDetail()
// 	if err != nil {
// 		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
// }
// func (h *handlerTransactionDetail) TransactionDetailByID(c echo.Context) error {
// 	accessLogin, ok := c.Get("adminLogin").(jwt.MapClaims)
// 	if !ok || accessLogin == nil {
// 		userLogin, _ := c.Get("customerLogin").(jwt.MapClaims)

// 			_, err := h.CustomerAuthRepository.ReauthCustomer(uint(accessLoginID))
// 			}
// 		} else {
// 			// If neither adminLogin nor customerLogin exists, return an error
// 		}
// 	} else {
// 		// Use adminLogin if it exists
// 		accessLoginID := accessLogin["id"].(float64)
// 		_, err := h.AdminAuthRepository.Reauth(uint(accessLoginID))

// 		if err != nil {
// 			return errorhandler.ErrorHandler(c, err, "Admin Not Found", http.StatusInternalServerError)
// 		}
// 	}

// 	ID := c.Param("id")
// 	LimitID, _ := strconv.Atoi(ID)

// 	creditlimit, err := h.TransactionDetailRepository.TransactionDetailByID(uint(LimitID))
// 	if err != nil {
// 		return errorhandler.ErrorHandler(c, err, "User Not Found", http.StatusBadRequest)
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: creditlimit})
// }
