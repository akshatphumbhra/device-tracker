# Device Tracker

Welcome to the Device Tracker! This application helps you track devices on a map.

The home view offers a list of all your devices on the left with a map of their locations on the right. This list is updated every minute!
<img width="1728" alt="Screenshot 2023-10-29 at 10 15 28 PM" src="https://github.com/akshatphumbhra/device-tracker/assets/43840267/a530941e-df9a-4a5b-bea9-b6f39dafd5eb">

The settings menu allows users to update the visibility of their devices in the list as well as upload custom icons for each device.
<img width="1728" alt="Screenshot 2023-10-29 at 10 15 13 PM" src="https://github.com/akshatphumbhra/device-tracker/assets/43840267/e9533685-dfdf-45ba-a103-673b7fc6a20a">

## Set up Instructions

1) Clone the repository.

    ```
      git clone git@github.com:akshatphumbhra/device-tracker.git
    ```

2) Add your API KEYS.
   There are two API keys required for this application. The first is the key for the One Step GPS API that
   provides the device data for the app. The `.env` file for this is located at `pkg/config/.env`.

   ```
   ONE_STEP_GPS_API_KEY=API_KEY_HERE
   ```
  
    The other API key required is for Google's Javascript Map API. The `.env` file for this is located at `frontend/.env`.

    ```
    VUE_APP_GOOGLE_MAP_API_KEY=API_KEY_HERE
    ```

    Here is the documentation to generate your own Google Maps Javascript API key.

    https://developers.google.com/maps/documentation/javascript/cloud-setup

3) Run the Go server.

    ```
    cd cmd/main
    go run main.go
    ```

4) In a new terminal window, install dependencies and run the frontend Vue server.

    ```
    cd frontend
    npm install
    npm run serve
    ```
  
    By default this should run on `http://localhost:8080`. However, if there is already an application running on that port, it will default to the next available port.
    In this case, we need to update line 26 in our `main.go` file. 

      ```
      AllowedOrigins:   []string{"http://localhost:8080"}, // Replace with your Vue.js frontend URL
      ```
    
    This is necessary to allow cross-origin resource sharing between the frontend and backend. 
