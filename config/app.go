package config

type Application struct {
	Env      Env
	Postgres *Database
}

func App() (*Application, error) {
	app := &Application{}
	app.Env = NewEnv()

	db, err := NewDatabase(app.Env, GetLogger())
	if err != nil {
		return nil, err
	}
	app.Postgres = db

	// Migrar modelos aquí
	// TODO: Lleva a otra parte
	err = app.Postgres.AutoMigrate(&Entrepreneur{}, &Entrepreneurship{}, &Resources{}, &Tags{}, &ResourcesTags{}, &Costs{}, &HistoricalCost{})
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *Application) CloseDBConnection() {
	if app.Postgres != nil {
		app.Postgres.CloseDBConnection()
	}
}
