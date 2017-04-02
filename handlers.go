package main

import (
	"github.com/mitchellh/mapstructure"
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/linxlad/TraskyEA/models"
)

func addInterest(client *Client, data interface{}) {
	var user models.EarlyAccess

	if err := mapstructure.Decode(data, &user); err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	go func() {
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
			client.send <- Message{"error", "USER_EXISTS"}
			return
		}

		client.send <- Message{"success", user.Email + " added."}
	}()
}