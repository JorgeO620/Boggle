**Intro to Software Engineering Sprint 1**

***User stories***

As a founder, I'd like to be able to build a list of contacts (other users), so that I can strengthen business relationships and have them more readily available to reach on Boggle than other users. 

As a developer, I want to know my account and information are secure and password protected, so that I’m not at risk of being impersonated.

As a CEO, I want to check the applicants' final graduation school, credits, and graduation certificate, so that I can better gauge their knowledge on their respective field.

As an applicant, I want to know that employers are being incentivized to read my resume and messages, so I don’t feel like I’m being ignored.

As an investor, I want to know how many users have searched about the specific companies so that I can determine whether it is appropriate to invest.

As a recent college graduate, I would like to personalize my profile page, so that I can stand out among the crowd and better introduce myself.

As an employer, I want to be able to search for potential job candidates while either filtering out candidates with undesirable qualities or allowing candidates with desirable qualities to filter through, so that I can spend more time and money focusing on my product and company.

As an expert in my field, I would like  to share my knowledge, so I can strengthen another user’s skills while making some extra money.

---

***What issues our team planned to address?***


Get Angular links and site set up (Frontend)

Begin construction of user profiles (Backend)

Set up contacts list (prior employees, recruiters, etc.) (Backend)

---

***Which ones were successfully completed***


We were succefully able to modify two files "app.component.html" and "app.component.ts" to display what we needed.For frontend work we were able to setup angular and were able to make a mock-up welcome page displaying the title of the project and with a login screen. We were also able to use the login screen using a username and password that was already represented in the code.

For backend work, we were able to successfully create the structures needed to store various pieces of information of a user and have the code needed to successfully store new information into a user profile. We also have started to gain familiarity with GORM in order to properly connect to MySQL and create a database with GoLang so we can store and retrieve user profiles.


***Which ones didn’t and why?***

As mentioned previously our code shows a potential username and password that is fixed to what we had coded therefore, a user can not make up a username or password that is not already embedded into our code. We currently dont have the right database to extract user profiles and need to modify the code to be able to take saved usernames and passwords that is unique to a user.

In terms of backend problems, the primary obstacle encountered was with the Couchbase database. Our initial plan was to use Couchbase as our means of data storage, but issues with the implementation took up a significant amount of time. We eventually decided to switch to GORM. This was the major roadblock that stopped us from completing the friends list functionality. With the switch to GORM, the friends list will be one of our primary goals for Sprint 2.

---
