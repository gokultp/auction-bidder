# Design For Auction Bidding

## High Level Design




## Low Level Design

### APIs

#### 1. Create User

**Endpoint**    : /v1/users

**Method**      : POST

**Body**

```json
{
    "first_name"    : "Gokul",
    "last_name"     : "TP",
    "email"         : "tp.gokul@gmail.com",
    "is_admmin"     : true
}
```

**Response**

```json
{   
    "id"            : 1,
    "first_name"    : "Gokul",
    "last_name"     : "TP",
    "email"         : "tp.gokul@gmail.com",
    "is_admmin"     : true,
    "token"         : "xxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

