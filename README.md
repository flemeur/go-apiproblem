go-apiproblem
=============

Go package implementing API Problem

For more info see: http://tools.ietf.org/html/draft-nottingham-http-problem

Usage example

	func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		var (
			// Values returned from handler
			status   int
			resource interface{}
			err      error
			// Holds the response data
			response []byte
		)

		// Custom handler of some kind
		status, resource, err = myHandler(w, r)
		if err != nil {
			apiErr, ok := err.(*apiproblem.APIProblem)
			if !ok {
				log.Print(err)
				status = http.StatusInternalServerError
				apiErr = apiproblem.New(status, http.StatusText(status), "An unknown error happened")
			}
			resource = apiErr
		}

		// ... then later in your API handler
		if err != nil {
			w.Header().Set("Content-Type", apiproblem.JSONMimeType)
		} else if h := w.Header().Get("Content-Type"); h == "" {
			w.Header().Set("Content-Type", "application/json")
		}

		response, err = json.Marshal(resource)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(status)
		w.Write(response)
	}
