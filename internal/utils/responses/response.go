package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

//sending the data encoded
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

//This is for the error, from this all the error will be of this kind
type Response struct{
	Status string
	Error string
}

const(
	StatusOK = "OK"
	StatusError = "Error" 
)

func GeneralError(err error) Response{
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

//validating error response
func ValidationError(errs validator.ValidationErrors) Response{
	var errmsg []string
	for _,err:= range errs{
		switch err.ActualTag(){
		case "required":
			errmsg = append(errmsg, fmt.Sprintf("Field %s is Required Field",err.Field()))
		default:
			errmsg = append(errmsg, fmt.Sprintf("Field %s is Invalid Field",err.Field()))
		}
	}
	
	return Response{
		Status: StatusError,
		Error: strings.Join(errmsg, ","),
	}
}