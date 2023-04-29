package controllers_test

import (
	"pollapalooza/src/controllers"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestSignup(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := controllers.Signup(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignin(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := controllers.Signin(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Signin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
