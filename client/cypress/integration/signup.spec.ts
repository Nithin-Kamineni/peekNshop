import * as cypress from "cypress";
describe('Signup Feature',()=>{
    it("Valid Credential Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type('dsfsdf')
        cy.get('#id_lastname').type('fffff')
        cy.get("#id_signup_email").type("nitin5@gmail.com");
        cy.get("#id_Signup_Password").type("sai");
        cy.get("#id_ConfPassword").type("sai");
        cy.get("#btn_Signup_Submit").click(); 
        cy.screenshot() 
    })
    it("Invalid Credential Testing",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type('sai')
        cy.get('#id_lastname').type('reddy')
        cy.get("#id_signup_email").type("saireddy@gmail.com");
        cy.get("#id_Signup_Password").type("saireddy");
        cy.get("#id_ConfPassword").type("saireddy");
        cy.get("#btn_Signup_Submit").click(); 
        cy.screenshot() 
    })
    
});