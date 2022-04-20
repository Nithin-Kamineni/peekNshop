import * as cypress from "cypress";
describe('About us redirecting',()=>{
    it("Redirectioning",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#sidenavleftbutton').click()
        cy.get('#about_us_btn').click()
        cy.screenshot()      
    })})