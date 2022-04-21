import * as cypress from "cypress";
describe('Cart Feature',()=>{
    it("Valid Credential Cart Display Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email1").type("admin@gmail.com");
        cy.get("#id_Password1").type("admin");
        cy.get("#btn_init_submit1").click(); 
        cy.url().should('include', '/user-homepage')
        cy.get("#cartButton").click(); 
        cy.screenshot() 
    })
    it("Invalid Credential Cart DisplayTesting",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email1").type("T01@gmail.com");
        cy.get("#id_Password1").type("Test_Password_001");
        cy.get("#btn_init_submit1").click();
        cy.get("#loginCloseButton").click();
        cy.url().should('include', '/')
        cy.screenshot() 
    })
    
});