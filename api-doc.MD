## **Test Interview Dompet Kilat for Backend Developer**

---

## **_Register_**

Add new user to database.

- **URL**

  /api/v1/signup

- **Method:**

  `POST`

- **URL Params**

  _None_

- **Data Params**

  `body:`

  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }
  ```

- **Success Response:**

  - **Code:** 200 OK<br />
    **Content:**

    ```json
    {
      "email": "string",
      "msg": "Register successfully",
      "username": "string"
    }
    ```

- **Error Response:**

  - **Code:** 503 Service Unavailable <br />
    **Content:**

    ```json
    {
      "error": "Content-Type: application/json is required"
    }
    ```

---

## **_Login_**

User authentication.

- **URL**

  /api/v1/signin

- **Method:**

  `POST`

- **URL Params**

  _None_

- **Data Params**

  `body:`

  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

- **Success Response:**

  - **Code:** 201 <br />
    **Content:**

    ```json
    {
      "token": "_token_"
    }
    ```

- **Error Response:**

  - **Code:** 401 Unauthorized <br />
    **Content:**
    ```json
    {
      "error": "Bad credentials"
    }
    ```

---

## **_GetOne_**

Get username and email from token.

- **URL**

  /api/v1/getone

- **Method:**

  `GET`

- **URL Params**

  _None_

- **Data Params**

  `headers:`

  ```json
  {
    "Authorization": "Bearer _token_"
  }
  ```

- **Success Response:**

  - **Code:** 200 OK <br />
    **Content:**
    ```json
    {
      "email": "string",
      "username": "string"
    }
    ```

- **Error Response:**

  - **Code:** 401 Unathorized <br />
    **Content:**

    ```json
    {
      "error": "Unauthorized"
    }
    ```
