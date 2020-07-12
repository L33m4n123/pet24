# Pet24
A sideproject for me to improve in
- Api driven development
- Golang
- Mongo
- Vue.js (frontend-part)
- Elasticsearch
- Caching
- ...

The idea behind this project is, to give animal shelters a common plattform where they can put their inquiries and people looking to rescue a pet have a common website to find them. Imagine it like your typical appartment sites that offer appartments for rent. However with animal shelters displaying their animals in need of a new home

# ToDo

## General
- [ ] split up all 3 in seperate repositories?
  - still debating about this myself
- [ ] creating seperate docker containers for each service
- [ ] Come up with a better name

## API
- [x] setup basic Server
- [x] add middleware so all endpoints serve json
- [x] setup connection to mongoDB
- [x] setup basic models
- [x] Shelter
    - [x] get all
    - [x] get specific by Id
    - [x] add new
    - [x] (soft-) delete 
    - [x] update
- [ ] Restructure CRUD Operations for Shelter so it can be used for every other model aswell
- [ ] Users
    - [ ] get all
    - [ ] get specific by Id
    - [ ] add new
    - [ ] (soft-) delete 
    - [ ] update
- [ ] Pets
    - [ ] get all
    - [ ] get specific by Id
    - [ ] add new
    - [ ] (soft-) delete
    - [ ] update
- [ ] Add Caching (Redis?)
- [ ] automatic cache invalidation on entity change (PDP)
- [ ] automatic populating of elastic on CRUD operations
    - like a job that runs in the background
- [ ] Search / uncached Listingpages from Elastic
- [ ] authorization middleware
    - only SuperUsers/Owner can delete Shelters
    - only SuperUsers/Owner can delete Pets
- [ ] thinking about changing to SQL Database to use schematics

## Administration

## Frontend