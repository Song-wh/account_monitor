package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

// DepositRecord represents a deposit transaction record
type DepositRecord struct {
	ID               int     `json:"id"`
	VirtualAccountNo string  `json:"virtual_account_no"`
	RemitterName     string  `json:"remitter_name"`
	RemitterAccount  string  `json:"remitter_account"`
	Amount           float64 `json:"amount"`
	PgSource         string  `json:"pg_source"`
	Payload          string  `json:"payload"`
	CreatedAt        string  `json:"created_at"`
}

var deposits []DepositRecord
var nextID int = 1
var iniFile *ini.File
var configFile = "config.ini"

func InitDB() {
	// Load or create INI file
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Create new INI file
		iniFile = ini.Empty()
		iniFile.Section("").Key("next_id").SetValue("1")
		iniFile.SaveTo(configFile)
		log.Println("Created new config.ini file")
	} else {
		// Load existing INI file
		iniFile, err = ini.Load(configFile)
		if err != nil {
			log.Fatal("Failed to load config.ini:", err)
		}
	}

	// Load next ID
	nextIDKey := iniFile.Section("").Key("next_id")
	if nextIDKey.String() == "" {
		nextID = 1
	} else {
		nextID, _ = strconv.Atoi(nextIDKey.String())
	}

	// Load existing deposits
	loadDeposits()
	log.Println("INI file database initialized")
}

func loadDeposits() {
	deposits = []DepositRecord{}

	// Load deposits from INI file
	for _, section := range iniFile.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}

		// Check if this is a deposit record (starts with "deposit_")
		if len(section.Name()) > 7 && section.Name()[:7] == "deposit" {
			id, _ := strconv.Atoi(section.Name()[7:]) // Extract ID from "deposit_1", "deposit_2", etc.

			deposit := DepositRecord{
				ID:               id,
				VirtualAccountNo: section.Key("virtual_account_no").String(),
				RemitterName:     section.Key("remitter_name").String(),
				RemitterAccount:  section.Key("remitter_account").String(),
				Amount:           section.Key("amount").MustFloat64(0),
				PgSource:         section.Key("pg_source").String(),
				Payload:          section.Key("payload").String(),
				CreatedAt:        section.Key("created_at").String(),
			}
			deposits = append(deposits, deposit)
		}
	}
}

func saveDeposit(deposit DepositRecord) error {
	// Create new section for this deposit
	sectionName := fmt.Sprintf("deposit_%d", deposit.ID)
	section := iniFile.Section(sectionName)

	section.Key("virtual_account_no").SetValue(deposit.VirtualAccountNo)
	section.Key("remitter_name").SetValue(deposit.RemitterName)
	section.Key("remitter_account").SetValue(deposit.RemitterAccount)
	section.Key("amount").SetValue(fmt.Sprintf("%.2f", deposit.Amount))
	section.Key("pg_source").SetValue(deposit.PgSource)
	section.Key("payload").SetValue(deposit.Payload)
	section.Key("created_at").SetValue(deposit.CreatedAt)

	// Update next ID
	iniFile.Section("").Key("next_id").SetValue(fmt.Sprintf("%d", nextID+1))

	// Save to file
	return iniFile.SaveTo(configFile)
}

func getDepositsByVirtualAccount(virtualAccount string) []DepositRecord {
	var result []DepositRecord
	for _, deposit := range deposits {
		if deposit.VirtualAccountNo == virtualAccount {
			result = append(result, deposit)
		}
	}
	return result
}
