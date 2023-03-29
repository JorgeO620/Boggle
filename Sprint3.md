**Frontend:**
 
As far as the frontend, we were able to continue on with the same login page. However, we were able to make changes in the home page. We made our home page more visually appealing by adding all the text to one side of the home page screen. A short catchphrase was added to the homepage to express what purpose this homepage can be used for. We were also able to implement a menu bar and home button that will lead to the corresponding instruction one may need to execute. If we take the mouse cursor to the menu located at the top of the homepage, we can see that the color of the menu changes to gray and is recognized. In addition, if we put the cursor on four other menus except for the home menu, the submenus of each menu can be seen.

**Unit tests Frontend:**

We tested for valid login credentials with different usernames and passwords so a user is able to put in a username and password unique to them  and be able to login into their own account. We tested for a working home button in the main menu. We tested for a menu button at the home page. We tested for invalid login credentials, so any user can't just login to another person’s account. We tested for the correct text to be displayed when one enters the website.

**Backend:**

The structure of user profiles at this point is largely complete, with some room for minor tweaking. Excess and unused code have been removed from the file, and the rest of the helper functions for updating individual structs are finished. The code has been fully commented, with the exception of the test file, as the naming convention for test functions makes them somewhat self explanatory. 
The remaining unit tests were constructed in the same manner as Sprint 2. Two stucts are created with on small difference, for example the usernames. The function being tested is run, and the structs are compared using deep equal. The primary focus on tests for this sprint were focused on the resume structure and the blog structure, which was ultimately removed as it did not align with our concept of Boggle. Each of the values in the resume that are able to be edited by users has been tested and works as intended.
