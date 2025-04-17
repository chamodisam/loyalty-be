# Backend Service

This is the backend service for handling user loyalty operations. It provides endpoints for earning and redeeming loyalty points and integrates with both Square and Google APIs.

## Prerequisites

Before running the backend service locally, you will need to:

1. Set up a **Square Developer App**.
2. Set up a **Google Developer App**.
3. Create a `.env` file to store your API credentials.

### 1. Set up Square Developer App

To use Square's API for loyalty points, follow these steps:

1. Go to the [Square Developer Dashboard](https://developer.squareup.com/apps).
2. Create a new application.
3. After the app is created, you'll get the **Access Token** that will be required for API calls to Square's services.

### 2. Set up Google Developer App

To integrate Google OAuth for authentication, follow these steps:

1. Go to the [Google Developer Console](https://console.developers.google.com/).
2. Create a new project (or use an existing one).
3. Enable the **Google OAuth API**.
4. Set up the **OAuth consent screen** and add the required scopes.
5. Create OAuth credentials (Client ID and Client Secret).
6. You'll need the **Google Client ID** and **Client Secret** to authenticate users via Google OAuth.

### 3. Create a `.env` file

In the root of your project, create a `.env` file and add the following environment variables:

```env
# Square API credentials
SQUARE_ACCESS_TOKEN=your_square_access_token_here

# Google OAuth credentials
GOOGLE_CLIENT_ID=your_google_client_id_here
GOOGLE_CLIENT_SECRET=your_google_client_secret_here
```

### 4. Install Dependencies
Install the required dependencies for the backend service:

```bash
npm install
```
### 5. Run the Backend Locally
To run the backend service locally, use the following command:
```bash
npm run dev
```
This will start the server on `http://localhost:8080` (or the port you have configured).

### 6. API Endpoints
Here are the API endpoints available in the backend:

#### POST `/api/earn`
**Description**: This endpoint allows users to earn loyalty points.

**Request Body**:

```json
{
  "account_id": "user_account_id",
  "points": "amount_of_points"
}
```

**Response**:

Success:

```json
{"message": "Points earned successfully!"}
```
Failure:

```json
{"message": "Failed to earn points."}
```
#### POST `/api/redeem`
Description: This endpoint allows users to redeem loyalty points.

**Request Body**:

```json
{
  "account_id": "user_account_id",
  "points": "amount_of_points_to_redeem"
}
```
**Response**:

Success:

```json
{"message": "Points redeemed successfully!"}
```
Failure:

```json
{"message": "Failed to redeem points."}
```
#### GET `/api/balance?account_id=`
Description: This endpoint returns the current loyalty balance of an account.

**Response**:

```json
{
  "balance": 100
}
```
### 7. Error Handling
All errors will be returned with a proper status code and an error message. Example error response:

```json
{
  "error": "Invalid input"
}
```
### 8. Backend Dependencies
The backend service uses the following libraries:

express – for the HTTP server.

axios – for making HTTP requests.

dotenv – for managing environment variables.

square-connect – for interacting with Square’s API.

google-auth-library – for Google OAuth integration.

### 9. Testing the Backend
To test the backend service, you can use tools like Postman or cURL to make API requests to the endpoints.

Example cURL to earn points:

```bash
curl -X POST http://localhost:8080/api/earn \
  -H "Content-Type: application/json" \
  -d '{"account_id": "user123", "points": "50"}'
```
