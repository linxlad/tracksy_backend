package main

import (
	"github.com/mitchellh/mapstructure"
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/linxlad/TraskyEA/models"
	"github.com/asaskevich/govalidator"
	"time"
)

func addInterest(client *Client, data interface{}) {
	var user models.EarlyAccess

	if err := mapstructure.Decode(data, &user); err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	// TODO: Move into own function.
	// Create an error array with all invalid struct fields.
	//var structErrors []error
	//if _, err := govalidator.ValidateStruct(user); err != nil {
	//	// split string for ';'
	//	errParse := strings.Split(err.Error(), ";")
	//	// remove last index empty
	//	removeLastEmpty := errParse[:len(errParse) -1]
	//
	//	for _, element := range removeLastEmpty {
	//		structErrors = append(structErrors, errors.New(element))
	//	}
	//}

	if isEmail := govalidator.IsEmail(user.Email); isEmail == false {
		client.send <- Message{"error", "INVALID_EMAIL"}
		return
	}

	go func() {
		user.CreatedAt = time.Now()
		response, err := r.Branch(
			r.Table("sign_up").GetAll(user.Email).OptArgs(r.GetAllOpts{
				Index: "email",
			}).IsEmpty(),
			r.Table("sign_up").Insert(user),
			nil,
		).RunWrite(client.session)

		if response.Errors != 0 && err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}

		if response.Inserted != 1 {
			client.send <- Message{"error", "EMAIL_EXISTS"}
			return
		}

		client.send <- Message{"success", user.Email + " added."}
	}()
}