package evidenceBook

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/akralani/cycleanalysis/database"
)

func getEvidenceBook(evidenceBookID int) (*EvidenceBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT 
	evidenceBookId, 
	date, 
	time, 
	coin, 
	price, 
	usdValue, 
	quantity, 
	fee, 
	buySell, 
	profitLoss, 
	percentProfitLoss, 
	usdValueProfitLoss, 
	snapshot 
	FROM evidenceBooks 
	WHERE evidenceBookId = ?`, evidenceBookID)

	evidenceBook := &EvidenceBook{}
	err := row.Scan(
		&evidenceBook.EvidenceBookID,
		&evidenceBook.Date,
		&evidenceBook.Time,
		&evidenceBook.Coin,
		&evidenceBook.Price,
		&evidenceBook.USDValue,
		&evidenceBook.Quantity,
		&evidenceBook.Fee,
		&evidenceBook.BuySell,
		&evidenceBook.ProfitLoss,
		&evidenceBook.PercentProfitLoss,
		&evidenceBook.USDValueProfitLoss,
		&evidenceBook.Snapshot,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return evidenceBook, nil
}

func GetTopTenEvidenceBooks() ([]EvidenceBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT 
	evidenceBookId, 
	date, 
	time, 
	coin, 
	price, 
	usdValue, 
	quantity, 
	fee, 
	buySell, 
	profitLoss, 
	percentProfitLoss, 
	usdValueProfitLoss, 
	snapshot 
	FROM evidenceBooks ORDER BY evidenceBookId DESC LIMIT 10
	`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	evidenceBooks := make([]EvidenceBook, 0)
	for results.Next() {
		var evidenceBook EvidenceBook
		results.Scan(&evidenceBook.EvidenceBookID,
			&evidenceBook.Date,
			&evidenceBook.Time,
			&evidenceBook.Coin,
			&evidenceBook.Price,
			&evidenceBook.USDValue,
			&evidenceBook.Quantity,
			&evidenceBook.Fee,
			&evidenceBook.BuySell,
			&evidenceBook.ProfitLoss,
			&evidenceBook.PercentProfitLoss,
			&evidenceBook.USDValueProfitLoss,
			&evidenceBook.Snapshot)

		evidenceBooks = append(evidenceBooks, evidenceBook)
	}
	return evidenceBooks, nil
}

func removeEvidenceBook(evidenceBookID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM evidenceBooks where evidenceBookId = ?`, evidenceBookID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func getEvidenceBookList() ([]EvidenceBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT 
	evidenceBookId, 
	date, 
	time, 
	coin, 
	price, 
	usdValue, 
	quantity, 
	fee, 
	buySell, 
	profitLoss, 
	percentProfitLoss, 
	usdValueProfitLoss, 
	snapshot 
	FROM evidenceBooks`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	evidenceBooks := make([]EvidenceBook, 0)
	for results.Next() {
		var evidenceBook EvidenceBook
		results.Scan(&evidenceBook.EvidenceBookID,
			&evidenceBook.Date,
			&evidenceBook.Time,
			&evidenceBook.Coin,
			&evidenceBook.Price,
			&evidenceBook.USDValue,
			&evidenceBook.Quantity,
			&evidenceBook.Fee,
			&evidenceBook.BuySell,
			&evidenceBook.ProfitLoss,
			&evidenceBook.PercentProfitLoss,
			&evidenceBook.USDValueProfitLoss,
			&evidenceBook.Snapshot)

		evidenceBooks = append(evidenceBooks, evidenceBook)
	}
	return evidenceBooks, nil
}

func updateEvidenceBook(evidenceBook EvidenceBook) error {
	// if the evidenceBook id is set, update, otherwise add
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if evidenceBook.EvidenceBookID == nil || *evidenceBook.EvidenceBookID == 0 {
		return errors.New("evidenceBook has invalid ID")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE evidenceBooks SET 
		date=?, 
		time=?, 
		coin=?, 
		price=?, 
		usdValue=?, 
		quantity=?, 
		fee=?, 
		buySell=?, 
		profitLoss=?, 
		percentProfitLoss=?, 
		usdValueProfitLoss=?, 
		snapshot=?
		WHERE evidenceBookId=?`,
		evidenceBook.Date,
		evidenceBook.Time,
		evidenceBook.Coin,
		evidenceBook.Price,
		evidenceBook.USDValue,
		evidenceBook.Quantity,
		evidenceBook.Fee,
		evidenceBook.BuySell,
		evidenceBook.ProfitLoss,
		evidenceBook.PercentProfitLoss,
		evidenceBook.USDValueProfitLoss,
		evidenceBook.Snapshot,
		evidenceBook.EvidenceBookID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func insertEvidenceBook(evidenceBook EvidenceBook) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO evidenceBooks  
	(date, 
	time, 
	coin, 
	price, 
	usdValue, 
	quantity, 
	fee, 
	buySell, 
	profitLoss, 
	percentProfitLoss, 
	usdValueProfitLoss, 
	snapshot) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		evidenceBook.Date,
		evidenceBook.Time,
		evidenceBook.Coin,
		evidenceBook.Price,
		evidenceBook.USDValue,
		evidenceBook.Quantity,
		evidenceBook.Fee,
		evidenceBook.BuySell,
		evidenceBook.ProfitLoss,
		evidenceBook.PercentProfitLoss,
		evidenceBook.USDValueProfitLoss,
		evidenceBook.Snapshot)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil
}
