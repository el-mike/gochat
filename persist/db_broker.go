package persist

// DBBroker - basic, common database broker interface.
type DBBroker interface {
	// First - finds first record that match given conditions.
	First(dest interface{}, conds ...interface{}) *DBResponse

	// Where - returns records that match given query.
	Where(query interface{}, args ...interface{}) *DBResponse

	// FirstWhere - returns first record that match given query.
	FirstWhere(dest interface{}, query interface{}, args ...interface{}) *DBResponse

	// Find - returns records that match given conditions.
	Find(dest interface{}, conds ...interface{}) *DBResponse

	// Save - update value in the DB or create if it does not exist.
	Save(value interface{}) *DBResponse

	// DeleteByID - deletes a record of given type by passed ID.
	DeleteByID(target interface{}, id interface{}) *DBResponse
}

// DBResponse - basic, unified database response.
type DBResponse struct {
	err error
}

// NewDBResponse - returns DBResponse instance.
func NewDBResponse() *DBResponse {
	return &DBResponse{}
}

// Err - returns an error that occured during DB operation.
func (dr *DBResponse) Err() error {
	return dr.err
}

// SetErr - sets given error on DBResponse instance.
func (dr *DBResponse) SetErr(err error) {
	dr.err = err
}
