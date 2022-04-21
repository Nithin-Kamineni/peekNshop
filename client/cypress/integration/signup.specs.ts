import * as cypress from "cypress";
describe('Signup Feature',()=>{
    it("Signup Credential Testing case1",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname2').type("Name 1");
        cy.get("#btn_Signup_Submit2").should("be.disabled"); 
    })
    it("Signup Credential Testing case2",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname2').type("Name 1");
        cy.get('#id_lastname2').type("Name 2")
        cy.get("#btn_Signup_Submit2").should("be.disabled");; 
    })
    it("Signup Credential Testing case3",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname2').type("Name 1");
        cy.get('#id_lastname2').type("Name 2")
        cy.get('#id_signup_email2').type("T01@gmail.com");
        cy.get("#btn_Signup_Submit2").should("be.disabled");; 
    })
    it("Signup Credential Testing case4",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname2').type("Name 1");
        cy.get('#id_lastname2').type("Name 2")
        cy.get('#id_signup_email2').type("T01@gmail.com");
        cy.get('#id_Signup_Password2').type("Password1");
        cy.get("#btn_Signup_Submit2").should("be.disabled");; 
    })

    it("Signup Valid Credential Testing case5",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname2').type("Name 1");
        cy.get('#id_lastname2').type("Name 2")
        cy.get('#id_signup_email2').type("T017@gmail.com");
        cy.get('#id_Signup_Password2').type("Password1");
        cy.get('#id_ConfPassword2').type("Password1");
        cy.get("#btn_Signup_Submit2").should("not.be.disabled");
        cy.get("#btn_Signup_Submit2").click();
        cy.url().should('include', '/user-homepage')
    })

    
    it("Signup Invalid Credential Testing case6",()=>{
        cy.visit('http://localhost:4200/');
        cy.get('#signupbutton').click()
        cy.get('#id_firstname2').type("Admin");
        cy.get('#id_lastname2').type("User")
        cy.get('#id_signup_email2').type("admin@gmail.com");
        cy.get('#id_Signup_Password2').type("admin");
        cy.get('#id_ConfPassword2').type("admin");
        cy.get("#btn_Signup_Submit2").should("not.be.disabled");
        cy.get("#btn_Signup_Submit2").click();
        cy.url().should('include', '/')
    })

});