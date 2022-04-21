import * as cypress from "cypress";
describe('About us redirecting',()=>{
    it("Redirectioning",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#aboutUsSidenav').click()
        cy.screenshot()      
    })})