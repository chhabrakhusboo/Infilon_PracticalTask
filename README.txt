# FETCH AND CREATE DATA
The API allows users to fetch detailed information about a specific person by their ID and create a new person record in the database.

# Requirements
- Go
- MySQL
- Postman

# API Endpoints
1. Upload Excel File
   Endpoint: GET /person/:personID/info
   Description:  Retrieves the details of a person based on the provided person ID.
   Request:
   	Path Parameters: personID (type: Integer)
        		 The ID of the person whose details are to be retrieved.

2. Create a New Person
   Endpoint: POST /person/create
    Description: Creates a new person record in the database.
    Request:
    	Body: JSON
        