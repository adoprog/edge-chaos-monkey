# edge-chaos-monkey

Simple proxy server for Sitecore Edge built in Go. Inspect Sitecore JSS GraphQL queries and their responses, or simulate throttling, server errors, and slow responses.

![image](https://github.com/user-attachments/assets/4f26110c-cfd6-48b4-ae71-12f2fde96ed2)

## Build Instructions
You can also download a Windows executable from the 'Releases' tab and run that one instead.

1. **Ensure Go is Installed**  
   Make sure you have [Go](https://golang.org/dl/) installed on your system (version 1.23 or later).

2. **Build and run the Project**

   Run the app with:
   ```bash
   go run .
   ```

   The app will start listening on port 8080, and you'll be prompted to switch between different modes by pressing keys (1-9). 

## Using the app

Update the .env file of your Sitecore JSS application - add the following line:

```bash
SITECORE_EDGE_URL=http://localhost:8080
```

Then build or start the JSS application and browse it as usually.
