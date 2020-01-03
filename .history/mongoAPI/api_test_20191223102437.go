func Test_Login(t *testing.T) {

	var dataAccess = data.DummyUserDataAccess{}
	dataAccess.New()
	var handler UserHandler = UserHandler{&amp;dataAccess}
	
	user := &amp;models.User{Username: "JohnSmith", Password: "Password"}
	jsonPerson, _ := json.Marshal(user)
	
	request, _ := http.NewRequest("POST", "/User/Login", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	decoder := json.NewDecoder(response.Body)
	
	handler.Login(response, request)
	
	var u models.User
	if err := decoder.Decode(&amp;u); err != nil {
		t.Errorf("User does not exist.")
	}
	
	if u.Address != "200 Place Zero" || u.City != "City Zero" {
		t.Errorf("User's information doesnt match.")
	}
	}
	
	