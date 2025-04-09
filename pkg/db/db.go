package db

type MockDB struct {
	db map[string]string
}

func NewMockDB() *MockDB {
	mockDB := &MockDB{
		db: make(map[string]string),
	}

	mockDB.addEntry("order status", "Your order is being processed.")
	mockDB.addEntry("order status for product", "Your order for product XYZ is on the way.")
	mockDB.addEntry("product availability", "Product is available in stock.")
	mockDB.addEntry("refund policy", "Refunds are processed within 7-10 business days.")
	mockDB.addEntry("business hours", "Our business hours are from 9 AM to 5 PM, Monday to Friday.")
	mockDB.addEntry("support contact", "You can contact our support team at support@example.com.")

	return mockDB
}

func (mdb *MockDB) addEntry(query, answer string) {
	mdb.db[query] = answer
}

func (mdb *MockDB) GetAnswer(query string) string {
	return mdb.db[query]
}
