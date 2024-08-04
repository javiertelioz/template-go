Feature: User management
  As a user
  I want to manage users
  So that I can create, read, update, and delete users

  Scenario Template: Creating users
    When I create a user with name "<name>" and email "<email>"
    Then I should get status code <status_code>
    And the response should be "<response>"

    Examples:
      | name     | email                | status_code | response                    |
      |          | jane.doe@example.com | 500         | Name cannot be empty\n    |
      | John Doe | john.doeexample.com  | 500         | Email format is invalid\n |
      | Jane Doe | jane.doe@example.com | 201         |                             |
