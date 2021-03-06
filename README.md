# Audition
Project for Golang

##Objective
Create a simple application that exposes a simple REST API that allows users to perform the following.
* Allows users to submit/post messages
* Lists received messages
* Retrieves a specific message on demand, and determines if it is a palindrome.
* Allows users to delete specific messages
* Build capability into your code so that it can be “observed”(Monitoring/Traceability/metrics)
* The application is deployed to a Cloud Provider (AWS is preferred)
* A Compute instance is provisioned programmatically o Application can be reached via public DNS record
 

###(Extra points) Provides a simple UI to interact with the service:
- Shows the list of messages posted by the users
- Allows to post new messages
- Allows to select a given message to see extra details
- (Extra points) Functional and System Tests
- (Extra points) Application is deployed in a container
####Additional Info / Restrictions
- Extra information can be requested by email.
- Code should have a permissive license.
- Any programming language can be used.
- Use of proprietary tools/libraries is discouraged.


##Submission
Code should be committed to a public GitHub/Bitbucket repository:
- The repo includes service, UI, and provisioning code.
###README.md Includes:
- Brief description of the implementation architecture. 
- Sequence diagram of the usecases
- How to: build, deploy and access the app.
- REST API documentation.
A URI to the app should be provided.Be prepared to discuss the architectural design and implementation details of your
project.

##Evaluation Criteria
- Code quality, ie: Style, complexity, good practices.
- Application architecture, ie: Design patterns, modularity.
- API design quality, ie: Follows standards and good practices.
- Documentation quality, ie: Content quality, completeness and accuracy.
- Deployment automation, ie: Reduced number of manual steps.


##Running Instructions for Running on Local System

- Pull the repository https://github.com/elkarto91/audition
- Execute Command -> docker-compose up -d
- The app will be now running on http://localhost:8080/

##Flow of Use

- Open app at http://localhost:8080/login
- Click on Register
- Enter Admin Username as admin & Admin Password as password
- Enter a unique user name and password
- Upon registration , you will redirected to login page.
- Login with the credentials
- Dashboard will have user username displayed as logged in user and an option to logout.
- Dashboard will host an option to enter a comment and submit
- Dashboard will let you check palindrome validity and delete a comment


##Other Information

Application uses basic authentication to handle REST API calls from outside.

The examples of REST API calls can be found in Documentation for Audition Document