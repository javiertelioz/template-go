@user-management
Feature: User management
  As a user
  I want to manage users
  So that I can create users

  @create-user
  Scenario Outline: Creating users
    When I create a user with payload:
    """
    <payload>
    """
    Then I should get status code <status_code>
    And the response should be "<response>"

    Examples:
      | payload                                               | status_code | response                     |
      | {"email": "jane.doe@example.com"}                     | 500         | Name cannot be empty\n       |
      | {"name": "John Doe", "email": "john.doeexample.com"}  | 500         | Email format is invalid\n    |
      | {"name": "Jane Doe", "email": "jane.doe@example.com"} | 201         |                              |
