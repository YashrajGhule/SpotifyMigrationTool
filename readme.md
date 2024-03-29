
# Spotify Account Migration Tool

This Python application assists in migrating playlists, including songs, descriptions, and cover images , between two Spotify accounts.

## Features

-   Transfer playlists and their contents from one Spotify account to another.
-   Leverage the Spotify Web API for data retrieval and manipulation.
-   Secure user authorization through the OAuth 2.0 authorization code flow (requires separate authorization for both accounts).

## Requirements

-   Python 3.x
-   A Spotify developer account and app credentials for both source and destination accounts ([https://developer.spotify.com/](https://developer.spotify.com/))
-   Go (for the callback server)

## Setup

1.  **Install Dependencies:**
    
    Bash
    
    ```
    pip install requests
    pip install colorama
    pip install webbrowser
    
    ```
    
2.  **Create a Spotify Developer App:**
    
    -   Visit  [https://developer.spotify.com/](https://developer.spotify.com/)  and create a developer account.
    -   Create a new app and note down the Client ID and Client Secret for both the primary(source) and secondary(destination) Spotify accounts.
3.  **Configure Credentials:**
    
    -   Update the following variables in  `app.py`  with your actual credentials:
        
        Python
        
        ```
        # Primary(Source) Account
        primary_client_id = ""
        primary_client_secret = ""

		# Secondary(Destination) Account
        secondary_client_id = ""
        secondary_client_secret = ""
        
        ```
4.  **Start the Go Callback Server:**
    
    -   **This step is mandatory.**  The callback server written in Go is crucial for handling the authorization flow. Build and run the Go callback server following the instructions specific to your implementation.
5.  **Execute the App:**
    
    -   Run  `python app.py`  to start the migration process.

## Usage

1.  The application will initiate the OAuth 2.0 authorization flow for the  **source account**. Grant access to your playlists and data.
2.  After successful authorization, you'll be redirected back to the app. Repeat this authorization process for the  **destination account**.
3.  Once both authorizations are complete, the app will retrieve playlists from the source account, create corresponding playlists in the destination account (with descriptions if available), and transfer song data (URIs) to populate the new playlists.
4.  (Optional) Cover image transfer might require additional libraries or workarounds due to limitations in the Spotify API.

## Notes

-   Ensure the redirect URI in your app's configuration matches the one used during app creation in Spotify Developer Dashboard.
-   Consider implementing a more user-friendly approach to handling authorization codes instead of manual entry.

## Contributing

Feel free to submit pull requests for improvements or additional features. Make sure to follow code style conventions and include proper documentation for your changes.