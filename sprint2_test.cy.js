/// <reference types="cypress" />

describe("simple cypresstest for sprint 2", function() {
  it("input id and password and click the login buttion", function() {
      cy.visit('https://demo.applitools.com/')
      cy.get('#username').type('test123')
      cy.get('#password').type('123')
      cy.get('#log-in').click();
      cy.url().should('include', '/app')
  });
});