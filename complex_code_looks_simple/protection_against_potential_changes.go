package complexcodelookssimple

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"git.dip.pics/dip/platform/go/logger.git"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

// config

func Parse() (*App, error) {
	cfg := &App{}

	_ = godotenv.Load()

	cfgPath, envOnly := getConfigFlags()

	if err := loadConfig(cfg, cfgPath, envOnly); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func main() {
	logger.New()

	logger.Info("starting be-bill", "version", version, "commit", commit, "date", date)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	cfg, err := config.Parse()
	if err != nil {
		panic("config.Parse: " + err.Error())
	}

	app, err := app.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	app.Run()

	<-ctx.Done()
	stop()

	app.Shutdown()
}

// interfaces

type Users interface {
	// CRUD операции
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	ByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
	Update(ctx context.Context, params dto.UpdateUser) error
	List(ctx context.Context, roleNames []domain.Role) ([]domain.User, error)
	Block(ctx context.Context, userIDs []domain.UserID) (map[domain.UserID]error, error)
	Unblock(ctx context.Context, userIDs []domain.UserID) (map[domain.UserID]error, error)
}

// implementation
type Repository struct {
	queries *sqlc.Queries
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		queries: sqlc.New(pool),
	}
}

func (r *Repository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	const op = "repository.users.Create"

	var leaderID *domain.UserID
	if user.Leader != nil {
		leaderID = utils.ToPtr(user.Leader.ID)
	}

	createdUser, err := r.queries.InsertNewUser(ctx, sqlc.InsertNewUserParams{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		MiddleName:  user.MiddleName,
		Department:  user.Department,
		Name:        user.Role.String(), // имя роли для поиска в таблице role
		Position:    user.Position,
		LeaderID:    postgres.UUIDPtrToPgUUID(leaderID),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return nil, postgres.MapError(err, op)
	}

	return &domain.User{
		ID:          createdUser.ID,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.PhoneNumber,
		FirstName:   createdUser.FirstName,
		LastName:    createdUser.LastName,
		MiddleName:  createdUser.MiddleName,
		Department:  createdUser.Department,
		Role:        user.Role,
		Position:    createdUser.Position,
		Leader:      toLeaderSummary(createdUser.LeaderID, &createdUser.FirstName, &createdUser.LastName, createdUser.MiddleName),
	}, nil
}

// .....etc
