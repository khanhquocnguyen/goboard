# Goboard

### Simple React Frontend

    http://test.khanhquocnguyen.com:3000

### Backend API Endpoint:

**Get all task**

    GET: http://test.khanhquocnguyen.com:8080/tasks

**Create a task:**

    POST: http://test.khanhquocnguyen.com:8080/tasks

Request body:

    {
        "description" : "Learn something new",
        "status" : "todo"
    }

**Get a task:**

    GET: http://test.khanhquocnguyen.com:8080/tasks/{id}

**Update a task:**

    PUT: http://test.khanhquocnguyen.com:8080/tasks/{id}

Request body:

    {
        "description" : "It was changed",
        "status" : "todo"
    }

**Delete a task:**

    DELETE: http://test.khanhquocnguyen.com:8080/tasks/{id}
