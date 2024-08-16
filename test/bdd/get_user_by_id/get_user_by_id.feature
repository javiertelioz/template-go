@user-management
Feature: User management
  As a user
  I want to manage users
  So that I can get user by id

  @get-user-by-id
  Scenario Outline: Getting a user by ID
    When I request the user with ID "<user_id>"
    Then I should get status code <status_code>
    And the response should be:
    """
    <response>
    """

    Examples:
      | user_id | status_code | response                                                    |
      | 1       | 200         | {"id":"1","name":"Jane Doe","email":"jane.doe@example.com"}\n |
      | 999     | 404         | User not found\n                                            |


