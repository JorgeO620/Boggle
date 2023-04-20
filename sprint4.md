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
- Test Restful: List of usernames
- Test Restful: Create account
- Test Restful: Change Password
- Test Restful: Delete Account

**Documentation for Backend APIs**

RESTful API- This RESTful API is used to connect the Angular application to the Go server, which is connected to a MySQL database that will be used to store user profile information and job postings that were made by users of Boggle. With the RESTful API, the angular application will be able to send requests to the Go server for either the retrieval or the manipulation of specified user profiles and/or their job posts that are stored in the MySQL database and the Go server will be able to send a response back, which consists of the resources (user profiles or job post information) that were requested or confirmation that the action that was requested has been carried out. Due to being a RESTful API, it uses the HTTP verbs GET, POST, PUT and DELETE. This means that the actions that can be requested by the Angular application are to retrieve a user profile or post history of user, create a new user profile or post, update the information on a user profile or post or delete a user profile or post. As of now, this API will be used to store a user profile’s username and password or the job post's post ID, the company that needs to fill in the job position, the job post title, the post itself and the username of the user that created the job post. For user profiles, the username will be the primary key that is used to identify a unique record in the MySQL database, which means that the username will be used to determine which record will get updated, deleted or retrieved when a request is sent from the angular application. For job posts, the primary key will be the post ID instead, which means it will serve the same purpose as the username for user profiles. The username in job posts will be used as a foreign key, which will ensure that only people with user profiles can create a job post.
