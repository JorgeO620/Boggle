# Project

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15.2.0.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.

## Running the RESTful API

Install Go, MySQL and Postman to run the RESTful API. Make sure to create the same tables as we did in the MySQL database that you will be using. The exact MySQL commands needed to create these tables are found in the MySQLCommands.txt file, which can be found in the sprint4_backend folder (link: https://github.com/mgadit/Boggle/tree/main/sprint4_backend/rest_api_mysql). Afterwards, download the Demo.exe, go.mod, go.sum and main.go files, which can be found in the same location as the MySQLCommands.txt file, and store them in the same folder. Make sure to modify the files to match your circumstances, such as changing the database name in the main.go file (it is currently trying to access a database named "demo"). Once this is done, open a terminal and navigate to the folder with the four downloaded files, then run the following commands (make sure to have Postman open beforehand):

1. go build
2. go run .

Afterwards, the program should be continously running and you can now send requests with Postman to retrieve, manipulate or create records. Example of body data that would be needed in the request for some of the methods can be seen in the PostmanRequestsDemo.txt file, which is in the same location as the other files. Make sure that your URL path matches the pattern that we have for our URL paths when testing, which can be seen in the video presentation for sprint 4.
