import * as cypress from "cypress";
describe('Location Feature',()=>{
    it("Valid Location feature Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email1").type("admin@gmail.com");
        cy.get("#id_Password1").type("admin");
        cy.get("#btn_init_submit1").click(); 
        cy.url().should('include', '/user-homepage')
        cy.get("#locationbutton").click();
        cy.get("#locationStreet").type("3800 SW 34th St");
        cy.get("#city").type("Gainesville");
        cy.get("#state").type("Florida");
        cy.get("#zipcode").type("32608");
        cy.get("#locationSubmit").click();
        cy.url().should('include', '/')
        cy.screenshot() 
    })
    
});