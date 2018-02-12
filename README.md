# twitter_db
The database is running in a docker container on my DigitalOcean server. The queries are really slow!

## Installation

### Using Go
In order to run the program you can pull this repo and run the program with *go run ./src/main.go* from the root of the project. Helge often use examples in Go so I highly recommend installing Go even though it might seem annoying. 

### Using Docker
If you do not want to use Go you can use Docker instead. Simply run the command below on your computer:

```
docker run --name dbapp -d -p 8085:8080 xnoga/twitter_db
```
This will start a container running on port 8085 on your local machine. You can then either use a browser or a program like PostMan to get information from the system via a REST-API.

### Using the server
If none of the above works for you, you can also just call the REST-API on my server: http://138.68.69.47:8085/users

## REST-API

### How many Twitter users are in the database?
Route: /users
```json
{
  "numberOfUsers": 659774
}
```

### Which Twitter users link the most to other Twitter users? (Provide the top ten.)
Route: /mentions

```json
[
  {
    "_id": {
      "user": "lost_dog"
    },
    "mentions": 549
  },
  {
    "_id": {
      "user": "tweetpet"
    },
    "mentions": 310
  },
  {
    "_id": {
      "user": "VioletsCRUK"
    },
    "mentions": 251
  },
  {
    "_id": {
      "user": "what_bugs_u"
    },
    "mentions": 246
  },
  {
    "_id": {
      "user": "tsarnick"
    },
    "mentions": 245
  },
  {
    "_id": {
      "user": "SallytheShizzle"
    },
    "mentions": 229
  },
  {
    "_id": {
      "user": "mcraddictal"
    },
    "mentions": 217
  },
  {
    "_id": {
      "user": "Karen230683"
    },
    "mentions": 216
  },
  {
    "_id": {
      "user": "keza34"
    },
    "mentions": 211
  },
  {
    "_id": {
      "user": "TraceyHewins"
    },
    "mentions": 202
  }
]
```

### Who are the most active Twitter users (top ten)?
Route: /active

```json
[
  {
    "_id": {
      "user": "lost_dog"
    },
    "mentions": 549
  },
  {
    "_id": {
      "user": "webwoke"
    },
    "mentions": 345
  },
  {
    "_id": {
      "user": "tweetpet"
    },
    "mentions": 310
  },
  {
    "_id": {
      "user": "SallytheShizzle"
    },
    "mentions": 281
  },
  {
    "_id": {
      "user": "VioletsCRUK"
    },
    "mentions": 279
  },
  {
    "_id": {
      "user": "mcraddictal"
    },
    "mentions": 276
  },
  {
    "_id": {
      "user": "tsarnick"
    },
    "mentions": 248
  },
  {
    "_id": {
      "user": "what_bugs_u"
    },
    "mentions": 246
  },
  {
    "_id": {
      "user": "Karen230683"
    },
    "mentions": 238
  },
  {
    "_id": {
      "user": "DarkPiano"
    },
    "mentions": 236
  }
]
```

### Who are the five most grumpy (most negative tweets) and the most happy (most positive tweets)? (Provide five users for each group)
Route: /negative (I seach the database for the keyword 'shit' in this query)
```json
[
  {
    "_id": {
      "user": "Spidersamm"
    },
    "mentions": 13
  },
  {
    "_id": {
      "user": "Dutchrudder"
    },
    "mentions": 12
  },
  {
    "_id": {
      "user": "D_AMAZIN"
    },
    "mentions": 11
  },
  {
    "_id": {
      "user": "risha_"
    },
    "mentions": 11
  },
  {
    "_id": {
      "user": "coolshite"
    },
    "mentions": 9
  },
  {
    "_id": {
      "user": "wowshaggy"
    },
    "mentions": 9
  },
  {
    "_id": {
      "user": "mr_apollo"
    },
    "mentions": 8
  },
  {
    "_id": {
      "user": "lesley007"
    },
    "mentions": 8
  },
  {
    "_id": {
      "user": "robynsweeney"
    },
    "mentions": 6
  },
  {
    "_id": {
      "user": "original_one"
    },
    "mentions": 6
  }
]
```

Route: /positive (I search for the keyword 'awesome' in this query)
```json
[
  {
    "_id": {
      "user": "ruhanirabin"
    },
    "mentions": 16
  },
  {
    "_id": {
      "user": "thisgoeshere"
    },
    "mentions": 15
  },
  {
    "_id": {
      "user": "tsarnick"
    },
    "mentions": 13
  },
  {
    "_id": {
      "user": "Djalfy"
    },
    "mentions": 12
  },
  {
    "_id": {
      "user": "kjgriffin18"
    },
    "mentions": 11
  },
  {
    "_id": {
      "user": "jbfanforever94"
    },
    "mentions": 11
  },
  {
    "_id": {
      "user": "patriciaco"
    },
    "mentions": 10
  },
  {
    "_id": {
      "user": "yaseminx3"
    },
    "mentions": 9
  },
  {
    "_id": {
      "user": "JBnVFCLover786"
    },
    "mentions": 8
  },
  {
    "_id": {
      "user": "WTFJAY"
    },
    "mentions": 8
  }
]
```
