It's a test for building a user management application with beego.

role:
Management: an administrator who manages all users
Users: ordinary users, manage their own information

function:
Commonly used functions:
The user login
Optional stay logged in (7 days)
After login, the user name and login time of the current user need to be displayed
Personal information viewing and modification
UID (unmodifiable)
User name (not modifiable)
Role (unmodifiable)
email
Password (not visible)

Management functions:
View a list of all users
Add user (user information filled by administrator)
Delete user

Other requirements:
Italic section has low priority.
The database is MySQL or mangodb.
Passwords in the database cannot be stored in clear text.
There is no limit to the way the user interface can be displayed simply.
If you're using an API form to provide a service, you need to provide an API document, not implement the UI.
Code specifications and naming styles need to be noted during coding.
Create your own project in the gitlab group: http://10.204.28.137/intern. All codes need to be submitted and updated according to the gitlab workflow.
You can provide simple flow, module diagrams.

time:
No more than August 18, 2018