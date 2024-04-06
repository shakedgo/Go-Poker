import NavBar from '../NavBar.vue'

describe('NavBar', () => {
  it('playground', () => {
    cy.mount(NavBar, { props: { msg: 'Hello Cypress' } })
  })

  it('renders properly', () => {
    cy.mount(NavBar, { props: { msg: 'Hello Cypress' } })
    cy.get('h1').should('contain', 'Hello Cypress')
  })
})
