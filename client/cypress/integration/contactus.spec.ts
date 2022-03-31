import * as cypress from "cypress";
describe('Contact us redirecting',()=>{
    it("Redirectioning",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#aboutUsBtn').click()       
    })})