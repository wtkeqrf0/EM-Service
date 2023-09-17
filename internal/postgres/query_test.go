package postgres

import (
	"context"
	"github.com/wtkeqrf0/restService/pkg/ent/user"
	"testing"
	"time"
)

func TestPostgres_Users(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	cl := New(postgresURL)
	defer cl.Close()

	users, err := cl.Users(ctx, Filter{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", users)
}

func TestPostgres_SaveUser(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	cl := New(postgresURL)
	defer cl.Close()

	data := EnrichedFIO{
		Name:      "Matvey",
		Surname:   "Sizov",
		Age:       18,
		Gender:    "male",
		CountryID: "RU",
	}

	now := time.Now()

	if err := cl.SaveUser(ctx, EnrichedFIOWithCreationTime{
		EnrichedFIO:  data,
		CreationTime: now,
	}); err != nil {
		t.Fatal(err)
	}

	u := cl.cl.User.Query().Where(user.CreateTimeEQ(now)).OnlyX(ctx)
	defer cl.cl.User.DeleteOne(u).ExecX(ctx)

	t.Logf("%+v", u)
}

func TestPostgres_UpdateUser(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	cl := New(postgresURL)
	defer cl.Close()

	data := EnrichedFIO{
		Name:      "Matvey",
		Surname:   "Sizov",
		Age:       18,
		Gender:    "male",
		CountryID: "RU",
	}

	u := cl.cl.User.Create().SetName(data.Name).SetSurname(data.Surname).
		SetAge(data.Age).SetGender(data.Gender).SetCountry(data.CountryID).SaveX(ctx)
	defer cl.cl.User.DeleteOneID(u.ID)

	newName := "Sasha"

	u, err := cl.UpdateUser(ctx, UpdateEnrichedFIO{
		ID:   u.ID,
		Name: &newName,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", u)
}

func TestPostgres_DeleteUser(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	cl := New(postgresURL)
	defer cl.Close()

	data := EnrichedFIO{
		Name:      "Matvey",
		Surname:   "Sizov",
		Age:       18,
		Gender:    "male",
		CountryID: "RU",
	}

	u := cl.cl.User.Create().SetName(data.Name).SetSurname(data.Surname).
		SetAge(data.Age).SetGender(data.Gender).SetCountry(data.CountryID).SaveX(ctx)

	u, err := cl.DeleteUser(ctx, u.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cl.cl.User.Get(ctx, u.ID)
	if err == nil {
		t.Fatal("user is not deleted")
	}

	t.Logf("%+v", u)
}
