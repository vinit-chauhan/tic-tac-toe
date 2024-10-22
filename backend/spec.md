# **API Spec:**

## **1. `/game`** -> Webhook Connection

This endpoint manages the creation, fetching, and updating of games.

### Endpoints:

**GET `/game/:id`** -> Fetch the state of a particular game.

**Response:**
```json
{
  "game_id": "<game_id>",
  "users": ["<user_1_id>", "<user_2_id>"],
  "state": [<game_board_state>]
}
```

**POST `/game/:id/start`** -> Starts a game with users in the lobby.

**Response:**
```json
{
  "game_id": "<game_id>",
  "users": ["<user_1_id>", "<user_2_id>"],
  "state": [<initial_game_board_state>]
}
```

**POST `/game/:id/update`** -> Updates the game state after a move.

**Request:**
```json
{
  "move": {
    "index": <move_index>,
    "user_id": "<user_id>"
  }
}
```

**Response:**
```json
{
  "winner": "<user_id>",  // Optional, if game is won
  "completed": <true|false>,
  "state": [<updated_game_board_state>]
}
```

## **2. `/lobby`** -> Manages lobbies and invites users to join.

### Endpoints:

**GET `/lobby/:lobby_id`** -> Fetch information about the lobby.

**Response:**
```json
{
  "lobby_id": "<lobby_id>",
  "users": ["<user_1_id>", "<user_2_id>"], // Current users in the lobby
  "timestamp": "<timestamp>"
}
```

**GET `/lobby/requests`** -> Lists all incoming join requests for the lobby.

**Response:**
```json
{
  "requests": [
    {
      "request_id": "<req_id>",
      "user_id": "<user_id>",
      "timestamp": "<timestamp>"
    }
  ]
}
```

**POST `/lobby/requests/:user_id`** -> Send an invite to a user to join the lobby.

**Response:**
```json
{
  "game_id": "<game_id>",
  "status": "invite_sent"
}
```

**POST `/lobby/requests/:req_id/join`** -> User joins the requested lobby.

**Response:**
```json
{
  "status": "joined"
}
```

## **3. `/user`** -> Manages user creation, updates, and login.

### Endpoints:

**POST `/user/create`** -> Create a new user.

**Request:**
```json
{
  "username": "<username>",
  "password": "<password>"
}
```

**Response:**
```json
{
  "user_id": "<user_id>",
  "username": "<username>"
}
```

**GET `/user/:id`** -> Fetch the user information by ID.

**Response:**
```json
{
  "user_id": "<user_id>",
  "username": "<username>"
}
```

**PUT `/user/:id/update`** -> Update user information by ID.

**Request:**
```json
{
  "username": "<new_username>",
  "password": "<new_password>"
}
```

**Response:**
```json
{
  "status": "updated"
}
```

**DELETE `/user/:id/delete`** -> Delete the user by ID.

**Response:**
```json
{
  "status": "deleted"
}
```

**POST `/user/login`** -> Logs a user in.

**Request:**
```json
{
  "username": "<username>",
  "password": "<password>"
}
```

**Response:**
```json
{
  "user_id": "<user_id>",
  "token": "<auth_token>"
}
```
