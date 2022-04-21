import * as cypress from "cypress";
describe('Delivery Feature',()=>{
    it("Valid Delivery Component Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email1").type("admin@gmail.com");
        cy.get("#id_Password1").type("admin");
        cy.get("#btn_init_submit1").click(); 
        cy.url().should('include', '/user-homepage')
        cy.get("#deliverySidenav").click(); 
        cy.url().should('include', '/delivery')
        cy.screenshot() 
    })
    it("Invalid Delivery Component Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get("#deliverySidenav").click(); 
        cy.url().should('include', '/')
        cy.screenshot() 
    })
    
});