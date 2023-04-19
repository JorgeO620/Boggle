**Frontend**

In sprint 4, we changed We changed the UI and design of the login screen and put our project name on the right. The home screen where you log in and move on is not much different from print 3. However, the number of menus created in sprint 3 was reduced to three, and when each menu was clicked, it was moved to the corresponding screen.
In addition, information on the website and information such as e-mails and phone numbers that can be contacted were added to the website. User can also click on the "Contact" menu to send messages to people user can contact. Finally, User can view the user's profile by clicking the profile that exists in the upper right corner, and user can also log out by clicking the Logout button.

**Fronted Cypress Test**

- login: Input ID and password on the loginpage and press the login button
- register:  
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

**Unit Tests**

In Sprint 4 we wrote unit tests for functions using the Restful API. By declaring a variable as httptest.NewRecorder(), I was able to monitor the output of each function locally and compare it to desired outputs. Unit tests were constructed for all methods finalized. In addition, server side unit tests remain as they were in the previous sprint.
