# PollaPalooza
## Get ready to rock the vote with PollaPalooza.
![images-removebg-preview](https://user-images.githubusercontent.com/78376735/234934199-b1aa3390-947a-4df2-9c0c-7376ae45ff35.png)

Pollapalooza gives you the power to design interactive polls, allowing you to capture insights in real time.
Polling your audience can be the most effective way to increase engagement and allow your users to feel empowered.

Oh, and did we mention it'll mAkE U rIch?!?!?!

The current objective of this repository is to create a nice clean RESTful API using golang with some yummy unit tests.
It should provide all the endpoints for the Pollapalooza app to consume.


Oh, and did we mention it'll mAkE U rIch?!?!?!

So what are you waiting for? Give me your money already!

#### Foundational work
- [x] Connect to local Postgres with docker-compoise.
- [x] Build the Go container. 
- [x] Live Reload for container.

#### Implement the API endpoints

- [ ] Create the necessary handler functions to handle requests to each API endpoint.
- [ ] Implement the logic for each handler function.
- [ ] Route functionality to map each endpoint to its corresponding handler function.
- [ ] Implement the necessary middleware functions to perform tasks such as authentication, logging, and error handling
- [ ] Write unit tests for each API endpoint to ensure that the functionality works as expected
- [ ] Use integration tests to test the API as a whole and to ensure that all endpoints work together correctly
- [ ] Perform load testing to determine the performance of the API under heavy traffic conditions.

##### Endpoints
- [ ] `POST /auth/signup` - Create a user account to create polls under.
- [ ] `POST /auth/signin` - Sign into an account created by a user.
- [ ] `POST /auth/user` - Get the information of the current signed in user.

- [ ] `GET /polls/:id` - retrieves a specific poll by its ID
- [ ] `GET /polls` - retrieves all the available polls
- [ ] `GET /polls/:id` - retrieves a specific poll by its ID
- [ ] `POST /polls` - creates a new poll
- [ ] `PUT /polls/:id` - updates an existing poll by its ID
- [ ] `DELETE /polls/:id` - deletes an existing poll by its ID
- [ ] `POST /polls/:id/vote` - allows a user to cast a vote on a specific poll by its ID
.... probably more



