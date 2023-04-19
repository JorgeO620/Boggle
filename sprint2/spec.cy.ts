describe("test for login", function() {
  it("input id and password and click the login buttion", function() {
      cy.visit('http://localhost:4200/')
      cy.get('#id').type('Mahamil')
      cy.get('#password').type('1234')
      cy.get('#log-in').click();
      cy.url().should('include', '/app')
  });
});


