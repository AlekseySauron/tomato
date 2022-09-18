package dbpkg

import ()

// type Service interface {
// 	Check(ctx context.Context, inn string) ([]models.Order, error)
// 	Pay(ctx context.Context, orderNumber string, amount int) error
// }

type DBRepository interface {
	//InsertDB(ctx context.Context, user string, status string) ([]models.Order, error)
	// GetOrder(ctx context.Context, orderNumber string) (models.Order, error)
	// Pay(ctx context.Context, orderNumber string, amount int) error
	Close() error
}
