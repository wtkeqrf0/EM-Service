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

	return &Postgres{cl: cl}
}

// InitSchema generates and applies migrations.
func (p *Postgres) InitSchema(ctx context.Context) error {
	return p.cl.Schema.Create(ctx)
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
func (p *Postgres) Users(ctx context.Context, f Filter) (ent.Users, error) {
	q := p.cl.User.Query().Limit(f.Limit).Offset(f.Offset)

	if f.Order != nil && *f.Order == "DESC" {
		q.Order(ent.Desc(user.FieldSurname, user.FieldName, user.FieldPatronymic))
	} else {
		q.Order(ent.Asc(user.FieldSurname, user.FieldName, user.FieldPatronymic))
	}

	if f.MaxAge != nil {
		q.Where(user.AgeLT(*f.MaxAge))
	}

	if f.MinAge != nil {
		q.Where(user.AgeGTE(*f.MinAge))
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
func (p *Postgres) SaveUser(ctx context.Context, fio EnrichedFIOWithCreationTime) error {
	return p.cl.User.Create().SetSurname(fio.Surname).
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
func (p *Postgres) UpdateUser(ctx context.Context, fio UpdateEnrichedFIO) (*ent.User, error) {
	q := p.cl.User.UpdateOneID(fio.ID).SetNillablePatronymic(fio.Patronymic)

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
func (p *Postgres) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	u, err := p.cl.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, p.cl.User.DeleteOne(u).Exec(ctx)
}

func (p *Postgres) Close() error {
	return p.cl.Close()
}
