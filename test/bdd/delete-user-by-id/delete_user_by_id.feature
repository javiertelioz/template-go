@user-management
Feature: User management
  As a user
  I want to manage users
  So that I can delete a user by ID

  @delete-user-by-id
  Scenario Outline: Deleting a user by ID
    When I delete the user with ID "<user_id>"
    Then I should get status code <status_code>

    Examples:
      | user_id | status_code |
      | 1       | 200         |
      | 999     | 404         |
