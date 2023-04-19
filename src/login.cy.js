describe('Login Test cases', () => {
  it('should visits', () => {
    cy.visit('http://localhost:4200/auth/login')
    .wait(1000)
  })
  it('should give Invalid/ username/Password error', () => {
    cy.visit('http://localhost:4200/auth/login')
    .wait(1000)
    cy.get('#username').type('Mahamil')
    .wait(1000)
    .get('#password').type('password')
    .wait(1000)
    .get('#login-submit').click()
    .wait(2000)
    .get('#username-error').should('contain', 'Invalid username/password')
    .wait(1000)
  })
  it('Login is successfull and navigate to main portal', () => {
    cy.visit('http://localhost:4200/auth/login')
    .wait(1000)
    cy.get('#username').type('Mahamil')
    .wait(1000)
    .get('#password').type('123456')
    .wait(1000)
    .get('#login-submit').click()
    .wait(2000)
    .wait(1000)
    .url().should('eq', 'http://localhost:4200/main/home')
    .wait(1000)
  })
})