package main

import (
	"testing"

	"github.com/syahrulrmdhnn/pendaftaran-coc/backend/config"
	"github.com/syahrulrmdhnn/pendaftaran-coc/backend/models"
)

func TestDatabaseConnection(t *testing.T) {
	// Test database initialization
	config.InitDB()
	
	if config.DB == nil {
		t.Fatal("Database connection failed")
	}
	
	// Test that we can query the database
	var count int64
	result := config.DB.Model(&models.Pendaftar{}).Count(&count)
	if result.Error != nil {
		t.Fatalf("Failed to query database: %v", result.Error)
	}
	
	t.Logf("Database connection successful. Found %d registrants.", count)
}

func TestPendaftarModel(t *testing.T) {
	// Test the Pendaftar model structure
	pendaftar := models.Pendaftar{
		NamaLengkap:   "Test User",
		Email:         "test@example.com",
		NoTelp:        "08123456789",
		BuktiTransfer: "test.jpg",
	}
	
	if pendaftar.NamaLengkap != "Test User" {
		t.Errorf("Expected NamaLengkap to be 'Test User', got %s", pendaftar.NamaLengkap)
	}
	
	if pendaftar.Email != "test@example.com" {
		t.Errorf("Expected Email to be 'test@example.com', got %s", pendaftar.Email)
	}
	
	if pendaftar.NoTelp != "08123456789" {
		t.Errorf("Expected NoTelp to be '08123456789', got %s", pendaftar.NoTelp)
	}
}