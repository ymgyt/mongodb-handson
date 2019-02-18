// https://docs.mongodb.com/manual/tutorial/write-scripts-for-the-mongo-shell/#differences-between-interactive-and-scripted-mongo

db = db.getSiblingDB('tests')

db.createUser({
    user: "gopher",
    pwd:  "secret",
    roles: [
        {role: "dbOwner", db: "tests"}
    ]
})

db.getUsers()

db.createCollection('users')