package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	UUID            string `bun:"id"`
	Surname         string `bun:"surname"`         //sn
	GivenName       string `bun:"givenname"`       //givenName
	CreatedAt       string `bun:"createdat"`       //whenCreated
	SAMAccountName  string `bun:"samaccountname"`  //sAMAccountName
	TelephoneNumber string `bun:"telephonenumber"` //telephoneNumber
	Department      string `bun:"department"`      //department
	Title           string `bun:"title"`           //title
	City            string `bun:"city"`            //l
	Mail            string `bun:"mail"`            //mail
}
