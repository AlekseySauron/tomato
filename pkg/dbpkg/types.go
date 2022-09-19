package dbpkg

import "github.com/AlekseySauron/tomato/models"

// type Service interface {
// 	Check(ctx context.Context, inn string) ([]models.Order, error)
// 	Pay(ctx context.Context, orderNumber string, amount int) error
// }

type DBRepository interface {
	//InsertDB(ctx context.Context, user string, status string) ([]models.Order, error)
	// GetOrder(ctx context.Context, orderNumber string) (models.Order, error)
	// Pay(ctx context.Context, orderNumber string, amount int) error
	GetUser(chatID int64) (*models.User, error)
	Save(*models.User) error
	Close() error
}
