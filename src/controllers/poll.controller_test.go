package controllers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestListPolls(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListPolls(tt.args.c)
		})
	}
}

func TestCreatePoll(t *testing.T) {
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
			if err := CreatePoll(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreatePoll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeletePoll(t *testing.T) {
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
			if err := DeletePoll(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeletePoll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdatePoll(t *testing.T) {
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
			if err := UpdatePoll(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePoll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPoll(t *testing.T) {
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
			if err := GetPoll(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetPoll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegisterVote(t *testing.T) {
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
			if err := RegisterVote(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RegisterVote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
