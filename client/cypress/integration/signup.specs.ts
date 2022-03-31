import * as cypress from "cypress";
describe('Signup Feature',()=>{
    it("Signup Credential Testing case1",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type("Name 1");
        cy.get("#btn_Submit_signup").should("be.disabled"); 
    })
    it("Signup Credential Testing case2",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type("Name 1");
        cy.get('#id_lastname').type("Name 2")
        cy.get("#btn_Submit_signup").should("be.disabled");; 
    })
    it("Signup Credential Testing case3",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type("Name 1");
        cy.get('#id_lastname').type("Name 2")
        cy.get('#id_email_signup').type("T01@gmail.com");
        cy.get("#btn_Submit_signup").should("be.disabled");; 
    })
    it("Signup Credential Testing case4",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type("Name 1");
        cy.get('#id_lastname').type("Name 2")
        cy.get('#id_email_signup').type("T01@gmail.com");
        cy.get('#id_Password_signup').type("Password1");
        cy.get("#btn_Submit_signup").should("be.disabled");; 
    })

    it("Signup Credential Testing case5",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname').type("Name 1");
        cy.get('#id_lastname').type("Name 2")
        cy.get('#id_email_signup').type("T01@gmail.com");
        cy.get('#id_Password_signup').type("Password1");
        cy.get('#id_confirmPassword').type("Password2");
        cy.get("#btn_Submit_signup").should("not.be.disabled");
        cy.get("#btn_Submit_signup").click();
    })

});