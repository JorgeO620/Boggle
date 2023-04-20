**Frontend**

In sprint 4, we changed the UI of the login screen and made it more proffesional looking to make it appeal to a user. We also made changes to the home screen to where it displays recruiter information and a litttle summary about them. We also implemented a box to display current events and news configured to each user and their previous searches and related interests. We also created an about us page where it displays a little bit of background of our company and added an address, phone number, and company email. We made changes to the "my profile" page where it is able to display a users profile information. We also impleneted a "connect" tab where users are able to connect with recruiters and be able to send them messages.

**Frontend Cypress Test**

- login: Input ID and password on the loginpage and press the login button
- register:  be able to register a new user.
- signout: Click the profile icon and signout button

**Frontend Unit Test**

- Display "Personal Info" section
- Display "About me" section
- Creat "Connect" components
- Confirm "Recruiter's List" exists
- Create header components
- Display "Boggle" content
- Display menu item such as "Home", "About Us" and "Connect"
- Display "News" section
- Display "Events" section
- Create "Home" components
- Display home image
- Create layout components
- Create app components
- Create about "About Us" components
- Confirm that "About Us" title exists

**Backend**

The RESTful API is now able to access the MySQL database and manipulate any of the records found in the accounts table (GET, PUT and DELETE), as well as create a new record in the accounts table of the database (POST). The issue that resulted in the POST method only being able to create accounts that have the default username (“John”) that was found in sprint 3 was fixed. All the methods for the accounts table are now working as expected and the RESTful API is working as expected for user profiles. We decided to use the RESTful API for job postings that will be made by users of Boggle. It should be noted that we created a posts table in the MySQL database that will contain every post that currently exists on Boggle, which is separate from the accounts table. We implemented the same methods as the methods used for the accounts table, which include GET, POST, PUT and DELETE. It should be noted that the GET method has two different endpoints, which have different URL paths and functions. The first endpoint for the GET method only includes the posts table in the URL path and uses a function that retrieves every record in the posts table. The second endpoint for the GET method includes a username after the posts table in the URL path and uses a function that searches the posts table for any posts made by that specific user, which is then retrieved. Both endpoints for the GET method for the posts table show the post ID, the company that needs to fill in the job position, the job post title, the post itself and the username of the user that created the job post for every record that was retrieved. The PUT method has a URL path that includes a post ID after the posts table, and it uses a function that replaces the post content of the record with the specified post ID with the post content that was specified in the request. The DELETE method has a URL path that includes a post ID after the posts table and uses a function that deletes the record with the specified post ID from the posts table in the MySQL database. The POST method has a URL path that only includes the posts table and uses a function that creates a new record with the specified company that needs to fill in the job position, the job post title, the post itself and the username of the user that created the job post that was sent in the request (the post ID continuously increments for every job post that was ever added and assigns its current value to a newly created job post). We also attempted to create a third endpoint for the GET method. This endpoint would have included the username after the posts table and then the post ID after the username on the URL path. The function used by this endpoint would have searched for the specified post ID in the group of posts that were made by the specified user, but it caused issues when trying to use the other endpoints of the GET method. Our attempt is currently commented out, but it can be seen in the main.go file.

**Backend Unit Tests**

In Sprint 4 we wrote unit tests for functions using the Restful API. By declaring a variable as httptest.NewRecorder(), I was able to monitor the output of each function locally and compare it to desired outputs. Unit tests were constructed for all methods finalized. In addition, server side unit tests remain as they were in the previous sprint.

- Change First and Last names
- Change current job
- Change current email address
- Change password
- Add contact
- Add past jobs
- Add education
- Add projects
- Test Restful: Get list of all usernames
- Test Restful: Create account
- Test Restful: Get specified username
- Test Restful: Change Password
- Test Restful: Delete Account

The GET, POST, PUT and DELETE methods for the accounts table for the RESTful API were tested on Postman, which is an API platform that allows developers to test their APIs. To test both endpoints of the GET method, I created two records in the accounts table by using the MySQL command line. Afterwards, I tested the endpoint that retrieves every record by setting the request URL to only include the accounts table and sending the request, which returned the two records that were previously created as expected. I then tested the endpoint for the GET method that retrieves the record with a specified username by setting the request URL to include a specific username and sending the request, which returns the record that contains the specified username as expected. The POST method was tested by setting the request URL to only include the accounts table and sending a request that also has body data that includes the username and password for the new account. After sending the request, a new record was added to the accounts table in the MySQL database, which had the custom username and password that was sent in the body data of the request. This means the POST method is working as expected. The PUT method was then tested by setting the request URL to include a specific username and sending a request that also has body data that includes a new password for the account with the specified username. After sending the request, the record in the MySQL database with the specified username has the new password that was sent in the body data of the request instead of the previous one, which is expected. The DELETE method was tested by setting the request URL to include a specific username and sending the request, which resulted in the record that has the specified username to be removed from the MySQL database, which is expected. The GET, POST, PUT and DELETE methods for the posts table for the RESTful API were tested on Postman as well. All the methods for the posts table were tested in a similar way to the methods for the accounts table (they also behave very similarly) but includes the posts table in the request URL rather than the accounts table. The request URL for the PUT and DELETE methods also includes the post ID of a specific post. It should also be noted that the request URL for the GET method that retrieves job posts made by a specific user includes the username after the posts table. When testing the POST method for the posts table, the request contains body data that includes the company interested in filling the job position, the job post title, the job post itself and the username of the user that created the post (the post ID continuously increments for every job post that was ever added and assigns its current value to the newly created job post). When testing the PUT method for the posts table, the request contains body data that includes the updated post since only the job post itself can be edited with the current implementation of our RESTful API. We also tested if it was possible that a person with no user profile can create a post by including a username in the body data of the request that did not exist in the accounts table (we used the POST method), which gave a socket hang up error. This means it is not possible, which is what we wanted.

**Documentation for Backend APIs**

RESTful API- This RESTful API is used to connect the Angular application to the Go server, which is connected to a MySQL database that will be used to store user profile information and job postings that were made by users of Boggle. With the RESTful API, the angular application will be able to send requests to the Go server for either the retrieval or the manipulation of specified user profiles and/or their job posts that are stored in the MySQL database and the Go server will be able to send a response back, which consists of the resources (user profiles or job post information) that were requested or confirmation that the action that was requested has been carried out. Due to being a RESTful API, it uses the HTTP verbs GET, POST, PUT and DELETE. This means that the actions that can be requested by the Angular application are to retrieve a user profile or post history of user, create a new user profile or post, update the information on a user profile or post or delete a user profile or post. As of now, this API will be used to store a user profile’s username and password or the job post's post ID, the company that needs to fill in the job position, the job post title, the post itself and the username of the user that created the job post. For user profiles, the username will be the primary key that is used to identify a unique record in the MySQL database, which means that the username will be used to determine which record will get updated, deleted or retrieved when a request is sent from the angular application. For job posts, the primary key will be the post ID instead, which means it will serve the same purpose as the username for user profiles. The username in job posts will be used as a foreign key, which will ensure that only people with user profiles can create a job post.
