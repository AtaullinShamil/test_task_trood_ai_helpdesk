package db

type MockDB struct {
	db map[string]string
}

func NewMockDB() *MockDB {
	mockDB := &MockDB{
		db: make(map[string]string),
	}

	mockDB.addEntry("refund", "Refunds are processed within 7-10 business days.")
	mockDB.addEntry("orderStatus", "Your order is being processed.")
	mockDB.addEntry("orderStatusForProduct", "Your order is on the way.")
	mockDB.addEntry("productInfo", "Product is available in stock.")
	mockDB.addEntry("businessHours", "Our business hours are from 9 AM to 5 PM, Monday to Friday.")
	mockDB.addEntry("supportContact", "You can contact our support team at support@example.com.")
	mockDB.addEntry("greeting", "Hello! How can I assist you today?")
	mockDB.addEntry("goodbye", "Goodbye! Have a great day!")
	mockDB.addEntry("cancelOrder", "To cancel your order, please visit your order page.")
	mockDB.addEntry("changeAddress", "To change your address, please go to your account settings.")
	mockDB.addEntry("technicalIssue", "We are sorry to hear that you're having an issue. Please contact support.")
	mockDB.addEntry("paymentIssue", "Please check your payment method or contact support if the issue persists.")
	mockDB.addEntry("speakToHuman", "Please wait while we connect you to a human agent.")

	return mockDB
}

func (mdb *MockDB) addEntry(intent, answer string) {
	mdb.db[intent] = answer
}

func (mdb *MockDB) GetAnswer(intent string) (string, bool) {
	answer, exists := mdb.db[intent]
	return answer, exists
}
