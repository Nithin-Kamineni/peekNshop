import * as cypress from "cypress";
describe('Login Feature',()=>{
    it("Valid Credential Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email").type("nitin1@gmail.com");
        cy.get("#id_Password").type("sai");
        cy.get("#btn_init_submit").click(); 
        cy.screenshot() 
    })
    it("Invalid Credential Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#loginbutton').click()
        cy.get("#id_email").type("T01@gmail.com");
        cy.get("#id_Password").type("Test_Password_001");
        cy.get("#btn_init_submit").click();
        cy.screenshot() 
    })
    
});