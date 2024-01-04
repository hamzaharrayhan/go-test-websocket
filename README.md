In this project, I utilize google for authentication. The authentication itself is still not perfectly implemented as the main purpose is only to understand the flow to integrate authentication using google APIs. The web can add data for phonenumbers alone, the providers data must be inputed manually through database. We can also delete and update the phonenumber data. When there is change in the database, the website will automatically update the data in dashboard page in realtime manner, because it has utilizing websocket.
This project is still far from perfect but the main objective of the requirements has been fulfilled, and can still be improved.

To start using, first initialize the database with name "phonenumber" in postgresql, after that, try to run the app to automatically migrate the data from the models. Then, use this query to populate the initial value for provider data:

INSERT INTO providers
(id, provider)
VALUES('07c3ea86-38b1-47b7-b8cc-f93b65011e33', 'indosat');

INSERT INTO providers
(id, provider)
VALUES('d19035db-d804-41db-ae6b-89d5a2b4b1ef', 'tri');

INSERT INTO providers
(id, provider)
VALUES('2fddb3ea-b9a2-4427-b3d9-0cadf7e2ff57', 'telkomsel');

run the program again.
After that, open URL: http://localhost:9090/
