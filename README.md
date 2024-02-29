# Simple CRUD application 
## Technologies:
* PostgreSQL as main database
* Redis as cache
* minIO S3 as storage

### List of endpoints
#### Authorization
1. [Register](#register)
2. [SignIn](#sign-in)
#### Book
1. [Create book](#create-book)
2. [Get book genres](#get-book-genres)
3. [Get all books](#get-books)
4. [Get book by id](#get-book-by-id)
5. [Search book](#search-books)
6. [Update book](#update-book)
7. [Delete book](#delete-book)
8. [Download book IMG](#download-book-img)

### Register

* **URL**

  /v1/register

* **Method**

  `POST`

* **Headers**
    * Accept-Language

* **URL Params**

  None

* **Data Params**

      ```  
        {
            "username":     string
            "password":     string
        }
      ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            "message":     "success"
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Sign in

* **URL**

  /v1/signIn

* **Method**

  `POST`

* **Headers**
    * Accept-Language

* **URL Params**

  None

* **Data Params**

      ```  
        {
            "username":     string
            "password":     string
        }
      ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImlhdCI6MTcwODgzNDkxOX0.ScsV2jjuhXmiEpTx2lFcU5jF4MNAdPrY5_gxA2Mo5to"
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Create book

* **URL**

  /v1/books

* **Method**

  `POST`

* **Headers**
    * Authorization (Bearer token)
    * Accept-Language

* **URL Params**

  None

* **Data Params**

      ```  
        {
            "genreId":           uint32
            "title":             string
            "description":       string
        }
      ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            "message":     "success"
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Get book genres

* **URL**

  /v1/genres

* **Method**

  `GET`

* **Headers**
    * Accept-Language

* **URL Params**

    None

* **Data Params**

    None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            []{
                id    uint32 
                title string 
              }
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Get books

* **URL**

  /v1/books

* **Method**

  `GET`

* **Headers**
    * Accept-Language

* **URL Params**

  * filter
    * genreId uint32
    * authorId uint32
  * sort
    * sort string (ASC | DESC) // it will sort result by release date
  * pagination
    * page uint32
    * limit uint32

* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            []{
                id uint32 
                author {
                         id       uint32
                         username string
                       }
                genre {
                         id       uint32
                         title string
                       }
                title string 
                imgURL string
                createdAt timestamp
                }
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Get book by id

* **URL**

  /v1/books/:id

* **Method**

  `GET`

* **Headers**
    * Accept-Language

* **URL Params**

    * id string

* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            []{
                id        uint32 
                author    string
                genre {
                         id    uint32
                         title string
                       }
                title       string 
                description string 
                imgURL      string
                createdAt   timestamp
                }
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Search books

* **URL**

  /v1/books/search

* **Method**

  `GET`

* **Headers**
    * Accept-Language

* **URL Params**

    * filter
        * genreId uint32
        * authorId uint32
    * sort
        * sort string (ASC | DESC) // it will sort result by release date
    * pagination
        * page uint32
        * limit uint32

* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            []{
                id uint32 
                author {
                         id       uint32
                         username string
                       }
                genre {
                         id       uint32
                         title string
                       }
                title string 
                imgURL string
                createdAt timestamp
                }
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Update book

* **URL**

  /v1/books/:id

* **Method**

  `PATCH`

* **Headers**
    * Authorization (Bearer token)
    * Accept-Language

* **URL Params**

  None

* **Data Params**

      ```  
        {
            "genreId":           uint32
            "title":             string
            "description":       string
        }
      ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            "message":     "success"
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```
   * **Code**: 401
   * **Description**: invalid access token or empty access token   
     **Content**:
     ``` 
     {
       "message": "unauthorized
     }
     ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Delete book

* **URL**

  /v1/books/:id

* **Method**

  `DELETE`

* **Headers**
    * Authorization (Bearer token)
    * Accept-Language

* **URL Params**

  * id uint32

* **Data Params**

    None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        {
            "message":     "success"
        }
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```
    * **Code**: 401
    * **Description**: invalid access token or empty access token   
      **Content**:
      ``` 
      {
        "message": "unauthorized
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```

### Download book IMG

* **URL**

  /v1/img/:path/download

* **Method**

  `GET`

* **Headers**
    * Authorization (Bearer token)
    * Accept-Language

* **URL Params**

    * path string 

* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```  
        image file
      ```

* **Error Response**

    * **Code**: 400
    * **Description**: invalid input param   
      **Content**:
      ``` 
      {
        "message": "invalid input param"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server   
      **Content**:
      ``` 
      {
        "message": "internal server error"
      }
      ```