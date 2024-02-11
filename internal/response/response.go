package response

import (
	"net/http"
	"shopping-api/internal/dto"
)

func StatusOK(msg string) map[string]interface{} {
	result := map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": msg,
	}
	return result
}

func StatusOKWithData(msg string, data interface{}) dto.CommonResponse {
	result := dto.CommonResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}
	return result
}

func StatusInternalServerError(msg string, data interface{}) dto.CommonResponse {
	result := dto.CommonResponse{
		Status:  "failed",
		Code:    http.StatusInternalServerError,
		Message: msg,
		Data:    data,
	}
	return result
}

func StatusBadRequest(msg string, data interface{}) dto.CommonResponse {
	result := dto.CommonResponse{
		Status:  "failed",
		Code:    http.StatusBadRequest,
		Message: msg,
		Data:    data,
	}
	return result
}

func StatusNotFound(msg string, data interface{}) dto.CommonResponse {
	result := dto.CommonResponse{
		Status:  "failed",
		Code:    http.StatusNotFound,
		Message: msg,
		Data:    data,
	}
	return result
}

func StatusForbidden(msg string, data interface{}) dto.CommonResponse {
	result := dto.CommonResponse{
		Status:  "failed",
		Code:    http.StatusForbidden,
		Message: msg,
		Data:    data,
	}
	return result
}

func StatusUnauthorized(msg string, data interface{}) dto.CommonResponse {
	result := dto.CommonResponse{
		Status:  "failed",
		Code:    http.StatusUnauthorized,
		Message: msg,
		Data:    data,
	}
	return result
}

func FailedPublish(err error) map[string]interface{} {
	errStr := err.Error()
	result := map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": errStr,
	}
	return result
}

func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "False Param",
	}
	return result
}
