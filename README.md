# Project description:
The application will display items and inventories of local stores. The user can view the stores near their vicinity and observe the items available in that store. The users can also order the things they want to be delivered to their homes or office. The items that need to be paid will be done by freelance delivery personnel, like students who pass by the exact location and want to earn some quick money to provide the items. The user can observe the store traffic from the application. The newly available inventory can be added, and not available stock can be updated or deleted. This application helps local stores deliver their items by implementing a freelance delivery feature.

# Problem Statement

People go to the store to buy things, but sometimes those things may not be available in that store. This waste lot of time for people who want to buy things and not find them in the store they went to.

# Solutions

By using peekNShop users can look at the store's inventory and know whether the items they are looking for are available or not. They can also find the best deal in terms of cost and quality for the items they are looking for with our application.

# Features
- finding the items avilable in the stores near the user
- freelance delivary option for getting delivary for grosseries at low cost
- order at your favorate local stores with best deals
- holding items that user want in the store through our application
- finding the latest offers that currently are avilable in the stores near user

## Stack Used

```bash
Frontend: Angular, Typescript
Database: SQLite #using gorm
Backend: Golang #using Mux
```

<strong>Demo Video of the project - </strong>
[Demo Video of PeekNshop](https://www.youtube.com/watch?v=bLQmdahIzR0)

<strong>Cypress test video - </strong>
[Cypress test video](https://youtu.be/RTktAoyVcZU)

<strong>Backend unit test video - </strong>
[Backend unit test video](https://youtu.be/fgYtMIBUJJQ)


### API Documentation is [here](https://github.com/Nithin-Kamineni/peekNshop/wiki/REST-API-Documentation)

### Project Sprint Board is [here](https://github.com/Nithin-Kamineni/peekNshop/projects/1)

### Sprint 4 delivarables are [here](https://github.com/Nithin-Kamineni/peekNshop/blob/main/sprint.md)

### File structure is [here](https://github.com/Nithin-Kamineni/peekNshop/wiki/File-Structure)

# Running the project
Step 1: [Introduction and Environment Setup for GoLang (Windows & Mac)](https://www.youtube.com/watch?v=dgIh-VYcWYw "Introduction and Environment Setup for GoLang (Windows & Mac)")

Step 2: [Angular Project Setup in Visual Studio Code](https://www.youtube.com/watch?v=ZJejjL1Iev0 "Angular Project Setup in Visual Studio Code")

Step 3: Arrange the files according to the file paths given below 
- peekNshop
  - client
    > go to client and run "npm i" t install all npm libraries
    - api
    > run the cleint using command "ng serve" to launch cleint side webite
    - src
    > visit http://localhost:4200/ to view the project
  - server
    > install go
    - bin
    - controllers
    - models
    - pkg
    - src
    - utils
    - main.go
    > run main.go file using "go run main.go" in termianal in the directory of server
    - Users.db
  
  
## Team Members
### BackEnd
Venkata Nithin Kamineni: vkamineni@ufl.edu github - Nithin-Kamineni

Saieswar Reddy Vaka: vaka.s@ufl.edu github - sai6221
### FrontEnd
Vamsi Pachamatla: vamsi.pachamatla@ufl.edu github - vamsi3379

Mattaparthi Nitish Veera: nitish.veera@ufl.edu github - VeeraNitish7


### Deployment of the backend of thee application is [Back end peekNShop](https://git.heroku.com/pure-temple-70794.git)
