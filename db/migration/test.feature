# Feature: Registration
#     Scenario: Successful Registration
#         Given I am unregisted user
#         When I registered with the following detials
#             | first_name   | middle_name   | last_name   | phone_number   | email   | password   |
#             | <first_name> | <middle_name> | <last_name> | <phone_number> | <email> | <password> |

#         Then I will have new account
#         Examples:
#             | first_name | middle_name | last_name | phone_number | email           | password |
#             | testuser1  | testuser1   | testuser1 | 0925252525   | test1@gmail.com | 1234567  |
#             | testuser1  | testuser1   | testuser1 | 0925252525   | test2@gmail.com | 1234567  |
#     Scenario: Failed Registration
#         Given I am unregisted user
#         When I registered with the following detials
#             | first_name   | middle_name   | last_name   | phone_number   | email   | password   |
#             | <first_name> | <middle_name> | <last_name> | <phone_number> | <email> | <password> |
#         Then I will not have new account
#         And the system should display "<message>" message

#         Examples:
#             | first_name | middle_name | last_name | phone_number | email           | password | message                  |
#             |            | testuser1   | testuser1 | 0925252525   | test1.com       | 1234567  | First name is required   |
#             | testuser1  |             | testuser1 | 0925252525   | test2@gmail.com | 1234567  | Middle name is required  |
#             | testuser1  | testuser1   |           | 0925252525   | test1@gmail.com | 1234567  | Last name is required    |
#             | testuser1  | testuser1   | testuser1 |              | test1@gmail.com | 1234567  | Phone number is required |
#             | testuser1  | testuser1   | testuser1 | 0925252525   |                 | 1234567  | email is required        |
#             | testuser1  | testuser1   | testuser1 | 0925252525   | test1@gmail.com |          | Password is required     |
#             | testuser1  | testuser1   | testuser1 | 0925252525   | test1gmail.com  | 1234567  | Invalid email            |
#             | testuser1  | testuser1   | testuser1 | 0925252525   | test1@gmail.com | 12       | Password to short        |
#             | testuser1  | testuser1   | testuser1 | 52525        | test1@gmail.com | 1234567  | Invalid phone number     |




# Feature: Login

#     Scenario: Successful Login
#         Given I am registed user
#         When I login with the following detials
#             | email   | password   |
#             | <email> | <password> |
#         Then I will be logged in securly to my account
#         Examples:
#             | email          | password |
#             | example.2f.com | 123456   |

#     Scenario: Failed Login
#         Given I am registed user
#         When I login with the following detials
#             | email   | password   |
#             | <email> | <password> |
#         Then I will not logged in to my account
#         And the system should display "<message>" message
#         Examples:
#             | email             | password  | message            |
#             | notexample.2f.com | 123456    | invalid credential |
#             | example.2f.com    | not123456 | invalid credential |


Feature: Redirect to previous session

    Scenario: Successful redirect
        Given I am in sso page
        When I successful logged in
        Then I should redirect back to where I where
        