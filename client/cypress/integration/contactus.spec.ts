import * as cypress from "cypress";
describe('Contact us redirecting',()=>{
    it("Redirectioning",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#contactUs').click()
        cy.get("#name_label_cs").type("Nikhil Annarapu");
        cy.get("#email_label_cs").type("nikhil18@gmail.com");
        cy.get("#message_cs").type("fcgcjhvjhvjh");
        cy.get("#contactussubmit_btn").click();
        cy.url().should('include', '/')
        cy.screenshot() 
       

    })})