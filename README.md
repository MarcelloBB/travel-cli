## Travel CLI

Lightweight command-line interface built in Go that helps you manage workspaces and collections, and perform customizable HTTP GET requests with ease.

Structure your API test environment directly from the terminal.

### Features
- Create and list workspaces and collections
- Send HTTP requests with custom headers and verbose output
- Save the formatted reponse body in a JSON file (--output [file.json])
- Send saved requests with ease
- Clean and extensible CLI structure with [Cobra](https://github.com/spf13/cobra)

### Example usage
```bash
travel create -w my-workspace
travel create -c my-collection
travel list -w
travel use -w my-workspace
travel get https://foo.bar.com -H "Accept: application/json" -o output.json
travel get https://... -c my_collection -t my_request
travel req -c my_collection -r my_request
```

### Database schema (demo)
<img width="702" height="611" alt="image" src="https://github.com/user-attachments/assets/929e2029-7515-491e-ac03-016f2193be8e" />


