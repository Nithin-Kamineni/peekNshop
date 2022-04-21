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
[Demo Video of PeekNshop](https://www.youtube.com/watch?v=bLQmdahIzR0)  (Too big to upload)

<strong>Cypress test video - </strong>
[Cypress test video]

https://user-images.githubusercontent.com/45183759/164371453-d06eae09-7d30-4060-9d6d-0eda3d9ebe6a.mp4


<strong>Backend unit test video - </strong>
[Backend unit test video]

https://user-images.githubusercontent.com/45183759/164371471-27594764-d700-4af8-ab74-eb358dab1803.mp4


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

### Deployment of the frontendapp of the application is [frontend peekNShop](https://ubiquitous-biscochitos-d3a939.netlify.app)
### Deployment of the backend of the application is [backend peekNShop](https://git.heroku.com/pure-temple-70794.git)

## Frontend screenshots
<img width="1440" alt="Screenshot 2022-04-21 at 12 34 28 AM" src="https://user-images.githubusercontent.com/48962308/164381966-91edf596-5d0b-4957-951f-ce17009ad6c5.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 34 42 AM" src="https://user-images.githubusercontent.com/48962308/164381978-90f23600-16b4-4046-9b78-38b8122be4db.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 34 56 AM" src="https://user-images.githubusercontent.com/48962308/164381993-9e9c7eda-5b29-46a5-9f5b-829f57172f60.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 35 02 AM" src="https://user-images.githubusercontent.com/48962308/164382024-d2665e80-92a2-48fe-a786-887842c939de.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 35 10 AM" src="https://user-images.githubusercontent.com/48962308/164382041-56046a3b-fc81-4f4b-9b8c-a05370355940.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 35 16 AM" src="https://user-images.githubusercontent.com/48962308/164382054-177cb551-5025-4aad-9b1c-d5de5145ce5f.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 35 22 AM" src="https://user-images.githubusercontent.com/48962308/164382069-150061d7-8f71-4cc9-af8f-bb6e59fbb100.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 35 32 AM" src="https://user-images.githubusercontent.com/48962308/164382080-ad3f29bb-3580-4d5d-85ff-003c1d92edd9.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 35 37 AM" src="https://user-images.githubusercontent.com/48962308/164382085-20b3168f-16f9-4755-873b-6cc8c4c49851.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 37 42 AM" src="https://user-images.githubusercontent.com/48962308/164382087-598040c9-f2f2-4ee7-9edc-5194b9af02af.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 37 57 AM" src="https://user-images.githubusercontent.com/48962308/164382093-5d6b8fe1-d88e-40cc-9f72-b6e51aa98f24.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 38 02 AM" src="https://user-images.githubusercontent.com/48962308/164382099-1479c0fc-d972-461c-8744-68a7fa3d0784.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 38 18 AM" src="https://user-images.githubusercontent.com/48962308/164382105-dbe31766-5c87-49e9-bf48-e7e151a7f8ab.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 38 26 AM" src="https://user-images.githubusercontent.com/48962308/164382111-885199e2-370d-4fcc-9307-e76df3a4e193.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 38 42 AM" src="https://user-images.githubusercontent.com/48962308/164382114-b7463322-a0a8-476f-a796-98e3ec566b94.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 38 51 AM" src="https://user-images.githubusercontent.com/48962308/164382119-ace1af1a-3789-4edf-bad3-737ebf679988.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 38 58 AM" src="https://user-images.githubusercontent.com/48962308/164382132-2917193d-ba64-4823-8c7f-f345dc899e0c.png">
<img width="1440" alt="Screenshot 2022-04-21 at 12 39 06 AM" src="https://user-images.githubusercontent.com/48962308/164382135-0e91d390-b65e-40e5-88ab-d7c1a1ef6be1.png">

