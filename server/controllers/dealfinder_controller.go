package controllers

func SearchProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error decoding product",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			w.WriteHeader(apiErr.StatusCode)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Error Decoding Product")
		return
	}
	log.Println(Product)

}