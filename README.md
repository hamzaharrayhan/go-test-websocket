In this project, I utilize google for authentication. The authentication itself is still not perfectly implemented as the main purpose is only to understand the flow to integrate authentication using google APIs. The web can add data for phonenumbers alone, the providers data must be inputed manually through database. We can also delete and update the phonenumber data. When there is change in the database, the website will automatically update the data in dashboard page in realtime manner, because it has utilizing websocket.
This project is still far from perfect but the main objective of the requirements has been fulfilled, and can still be improved.

To start using, first initialize the database with name "phonenumber" in postgresql, after that, try to run the app to automatically migrate the data from the models. Then, copy this to the Init() function in init.go file in folder /config to populate the initial value for provider data:
	providerRepo.Save(models.Provider{Provider: "telkomsel"})
	providerRepo.Save(models.Provider{Provider: "XL"})
	providerRepo.Save(models.Provider{Provider: "Tri"})
	providerRepo.Save(models.Provider{Provider: "Indosat"})
run the program again.
After that, open URL: http://localhost:9090/
