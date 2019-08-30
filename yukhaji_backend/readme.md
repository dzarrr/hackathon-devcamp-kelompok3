


// Get user's saving data
GET: /saving/:userID
RESULT:
[
    {
        "id": 3,
        "user_id": 2,
        "balance": 5000000,
        "target": 1000000,
        "start_date": "2019-08-29T00:00:00Z",
        "end_date": "2021-08-29T00:00:00Z"
    }
]

// Insert new saving record
POST: /saving/
BODY:
{
    "user_id": 2,
    "balance": 5000000,
    "target": 1000000,
    "start_date": "2019-08-29T00:00:00Z",
    "end_date": "2021-08-29T00:00:00Z"
}

// Add balance to user's saving
PUT: /addbalance/:userID
BODY:
{
    "balance": 1000
}
RESULT:
{
		status: "success",
		message: "Update balance success!",
		balance: 9000,
		daily_pay: 241.82,
		end_date: 2030-11-18 23: 01: 29.864506 +0700 WIB
}

// Edit user's saving end date
PUT: /editenddate/:userID
BODY:
{
	"end_date":"2030-11-18T00:00:00Z"
}
RESULT:
{
		status: "success",
		message: "Update End Date success!"
		daily_pay 241.81,
}