Feature: Get User by ID
  As a user
  I want to get a user by their ID
  So that I can view user details

  Scenario Template: Getting user by ID
    Given A repository with user not found with "<id>"
    When I get a user by ID "<id>"
    Then I should get status code <status_code>
    And the response should be "<response>"

    Examples:
      | id                                 | status_code | response                                           |
      | e17d5135-989e-4977-99ef-495c0ab7cd00                    | 500         | user not found\n |

  Scenario Template:
    Given A repository with user found with "<id>"
    When I get a user by ID "<id>"
    Then I should get status code <status_code>
    Examples:
      | id | status_code |
      |e17d5135-989e-4977-99ef-495c0ab7cd01 | 200   |