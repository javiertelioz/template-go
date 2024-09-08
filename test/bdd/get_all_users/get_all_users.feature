@user-management
Feature: User management
  As a user
  I want to manage users
  So that I can retrieve the list of users

  @get-users
  Scenario: Getting the list of users
    When I request the list of users
    Then I should get status code 200
    And the response should be:
    """
    [{"id":"1","name":"Jane Doe","email":"jane.doe@example.com"},{"id":"2","name":"John Doe","email":"john.doe@example.com"}]
    """
