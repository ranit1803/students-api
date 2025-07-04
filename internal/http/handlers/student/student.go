package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ranit1803/students-api/internal/types"
	"github.com/ranit1803/students-api/internal/utils/responses"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		slog.Info("Creating a Student Info")
		var student types.Student
		err:= json.NewDecoder(r.Body).Decode(&student)
		//This error is for empty body
		if errors.Is(err, io.EOF){
			responses.WriteJSON(w,http.StatusBadRequest,responses.GeneralError(fmt.Errorf("empty body")))
			return
		}

		//This is for Other Errors
		if err!=nil{
			responses.WriteJSON(w,http.StatusBadRequest,responses.GeneralError(err))
			return
		}

		//Validating the Data
		if err:= validator.New().Struct(student); err!=nil{
			validate_err:= err.(validator.ValidationErrors)
			responses.WriteJSON(w,http.StatusBadRequest,responses.ValidationError(validate_err))
			return
		}
		responses.WriteJSON(w,http.StatusCreated, map[string]string{"Success":"OK"})
	}
}