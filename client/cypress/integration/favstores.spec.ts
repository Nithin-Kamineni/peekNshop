import * as cypress from "cypress";
describe('Favorite Stores Feature',()=>{
    it("Valid Favorite Component Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email1").type("admin@gmail.com");
        cy.get("#id_Password1").type("admin");
        cy.get("#btn_init_submit1").click(); 
        cy.url().should('include', '/user-homepage')
        cy.get("#favoriteSidenav").click(); 
        cy.url().should('include', '/user/favorate-stores')
        cy.screenshot() 
    })
    it("Invalid Favorite Component Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get("#favoriteSidenav").click(); 
        cy.url().should('include', '/')
        cy.screenshot() 
    })
    
});