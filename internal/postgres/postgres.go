package postgres

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/restService/pkg/ent"
	"github.com/wtkeqrf0/restService/pkg/ent/user"
	"time"
)

// Postgres struct provides the ability to interact with the database.
//
// controller.Postgres interface and mock.MockPostgres
// are generated and based on Postgres implementation.
//
//go:generate ifacemaker -f postgres.go -o controller/postgres.go -i Postgres -s Postgres -p controller -y "Controller describes methods, implemented by the postgres package."
//go:generate mockgen -package mock -source controller/postgres.go -destination controller/mock/mock_postgres.go
type Postgres struct {
	cl *ent.Client
}

// New open new connection, start stats recorder and create the tables.
func New(url string) *Postgres {
	db, _ := sql.Open("pgx", url)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	cl := ent.NewClient(ent.Driver(drv))

	if log.GetLevel() == log.DebugLevel {
		// Setup `debug` client.
		cl = cl.Debug()
	}

	if err := cl.Schema.Create(context.Background()); err != nil {
		log.WithError(err).Fatal("tables initialization failed")
	}

	return &Postgres{cl: cl}
}

// Filter contains Pagination, and also the maximum and minimum age.
// The structure was created as indicative and can be supplemented by other filters.
type Filter struct {
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Order  *string `json:"order,omitempty"`
	MinAge *int    `json:"minAge,omitempty"`
	MaxAge *int    `json:"maxAge,omitempty"`
}

// Users method gets users with the specified limit, offset, order, maxAge and minAge.
func (c *Postgres) Users(ctx context.Context, p Filter) (ent.Users, error) {
	q := c.cl.User.Query().Limit(p.Limit).Offset(p.Offset)

	if p.Order != nil && *p.Order == "DESC" {
		q.Order(ent.Desc(user.FieldSurname, user.FieldName, user.FieldPatronymic))
	} else {
		q.Order(ent.Asc(user.FieldSurname, user.FieldName, user.FieldPatronymic))
	}

	if p.MaxAge != nil {
		q.Where(user.AgeLT(*p.MaxAge))
	}

	if p.MinAge != nil {
		q.Where(user.AgeGTE(*p.MinAge))
	}

	return q.All(ctx)
}

// EnrichedFIO represents enriched full name.
type EnrichedFIO struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
	Age        int     `json:"age"`
	Gender     string  `json:"gender"`
	CountryID  string  `json:"countryId"`
}

type EnrichedFIOWithCreationTime struct {
	EnrichedFIO
	CreationTime time.Time
}

// SaveUser to database.
func (c *Postgres) SaveUser(ctx context.Context, fio EnrichedFIOWithCreationTime) error {
	return c.cl.User.Create().SetSurname(fio.Surname).
		SetName(fio.Name).SetNillablePatronymic(fio.Patronymic).
		SetCountry(fio.CountryID).SetAge(fio.Age).SetGender(fio.Gender).
		SetCreateTime(fio.CreationTime).Exec(ctx)
}

type UpdateEnrichedFIO struct {
	ID         int     `json:"id"`
	Name       *string `json:"name,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Patronymic *string `json:"patronymic,omitempty"`
	Age        *int    `json:"age,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	Country    *string `json:"country,omitempty"`
}

// UpdateUser in the database by id.
func (c *Postgres) UpdateUser(ctx context.Context, fio UpdateEnrichedFIO) (*ent.User, error) {
	q := c.cl.User.UpdateOneID(fio.ID).SetNillablePatronymic(fio.Patronymic)

	if fio.Gender != nil {
		q.SetGender(*fio.Gender)
	}

	if fio.Name != nil {
		q.SetName(*fio.Name)
	}

	if fio.Surname != nil {
		q.SetSurname(*fio.Surname)
	}

	if fio.Age != nil {
		q.SetAge(*fio.Age)
	}

	if fio.Country != nil {
		q.SetCountry(*fio.Country)
	}

	return q.Save(ctx)
}

// DeleteUser in the database by id.
func (c *Postgres) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	u, err := c.cl.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, c.cl.User.DeleteOne(u).Exec(ctx)
}

func (c *Postgres) Close() error {
	return c.cl.Close()
}
