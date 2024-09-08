@user-management
Feature: User management
  As a user
  I want to manage users
  So that I can update user details by ID

  @update-user-by-id
  Scenario Outline: Updating a user by ID
    When I update the user with ID "<user_id>" with payload:
    """
    <payload>
    """
    Then I should get status code <status_code>
    And the response should be "<response>"

    Examples:
      | user_id | payload                                                                            | status_code | response       |
      | 1       | {"name": "Jane Doe Updated", "email": "jane.doe@example.com", "dob": "1990-01-01"} | 200         |                |
      | 999     | {"name": "John Doe", "email": "john.doe@example.com"}                              | 404         | User not found |
