package mongodb

var stdout = struct {
}{}

var stderr = struct {
	CannotConnect     string
	CannotInsertData  string
	CannotLoadSession string
	CannotSaveSession string
	CannotUpsertData  string
	EnvVarUnset       string
}{
	CannotConnect:     "could not connect to the database: %v",
	CannotInsertData:  "could not insert data into %v.%v; %v",
	CannotLoadSession: "could not load session data; %v",
	CannotSaveSession: "could not save session data; %v",
	CannotUpsertData:  "could not upsert data, %v",
	EnvVarUnset:       "environment variable %v has not been set",
}
