package db

type MockDB struct {
	db map[string]string
}

func NewMockDB() *MockDB {
	mockDB := &MockDB{
		db: make(map[string]string),
	}

	mockDB.addEntry("refund", "You can get the information and issue a refund at www.example.com/refund")
	mockDB.addEntry("orderStatus", "You can track your order on www.example.com/track")
	mockDB.addEntry("greeting", "Hello! How can I assist you today?")
	mockDB.addEntry("goodbye", "Goodbye! Have a great day!")
	mockDB.addEntry("cancelOrder", "To cancel your order, please visit your order page")
	mockDB.addEntry("changeAddress", "To change your address, please go to your account settings")
	mockDB.addEntry("technicalIssue", "We are sorry to hear that you're having an issue. Please contact support")
	mockDB.addEntry("paymentIssue", "Please check your payment method or contact support if the issue persists")
	mockDB.addEntry("speakToHuman", "Please wait while we connect you to a human agent")
	mockDB.addEntry("productInfo", "You can find out the current information on the product page")
	mockDB.addEntry("businessHours", "Our business hours are from 9 AM to 5 PM, Monday to Friday")
	mockDB.addEntry("supportContact", "You can contact our support team at support@example.com")

	return mockDB
}

func (mdb *MockDB) addEntry(intent, answer string) {
	mdb.db[intent] = answer
}

func (mdb *MockDB) GetAnswer(intent string) (string, bool) {
	answer, exists := mdb.db[intent]
	return answer, exists
}
