package handlers

r := mux.NewRouter()

	r.HandleFunc("/", mongoapi.HelloWorld)