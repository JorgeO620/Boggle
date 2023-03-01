***Work Completed for Sprint 2***  

**Frontend**   

As far as the frontend we were able to make a better login screen that is more visually appealing and another welcome page after you log in. We also added to our previous code and implemented a routing module to be able to route a path for the website. We also make better icons for the login screen to make a suit well for our project.

**Backend**  

Furthered progress on storing and retrieving user profiles by starting implementation of a RESTful API that uses Golang, the gorilla/mux package and a MySQL database that will be consumed by angular with the HttpClient Module so it can display any necessary information that pertains to user profiles, such as the search results when trying to search for a specific username of a profile, as well as add and delete user profiles in the MySQL database. For now, we are currently working on retrieving the more basic elements of a user profile, such as its name and current job title, and eventually work on retrieving more complicated items in a user profile, such as their resume. Once this RESTful API is able to properly create, retrieve, update and delete the basic elements of the user profiles from the MySQL database when testing with Postman (currently in the progress of debugging the API and working out some errors), we will begin to integrate it with Angular so that the Angular application can manipulate and retrieve data through the Go server, which retrieved the data from the MySQL database, 

---

**Unit Tests and Cypress Test for Frontend:**  

For the cypress and unit tests for the frontend we tested the login screen, tested if the application is created, if it contains the title, and if the application should render the title. We successfully were able to pass it through 21 tests.

**Unit Tests for Backend:**  

For the backend unit tests, we tested the helper functions that will be used to modify various account and profile elements, such as the user’s registered email address or their current occupation. We additionally began tests for the contacts list, although the implementation of the contact list may change in the future.

---

**Documentation for Backend APIs:**  

RESTful API- This RESTful API is used to connect the Angular application to the Go server, which is connected to a MySQL database that will be used to store user profile information. With the RESTful API, the angular application will be able to send requests to the Go server for either  the retrieval or the manipulation of specified User profiles that are stored in the MySQL database and the Go server will be able to send a response back, which consists of the resources (user profiles) that were requested or confirmation that the action that was requested has been carried out. Due to being a RESTful API, it uses the HTTP verbs GET, POST, PUT and DELETE. This means that the actions that can be requested by the Angular application are to retrieve a user profile, create a new user profile, update the information on a user profile or delete a user profile.

