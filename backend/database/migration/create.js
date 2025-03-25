use requestDebugDB

db.createUser(
    {
        user: "requestDebug",
        pwd:  passwordPrompt(),
        roles: [{ role: "readWrite", db: "requestDebug" }]
    }
)


db.createCollection(
    "requestGroups",
    {
        "capped": true,
        "size": 2 * 1024 * 1024 * 1024,
        "max": 5000,
    }
)
