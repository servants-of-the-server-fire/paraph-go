# CreateTemplateRequest

Upload a PDF template with named form fields.


## Supported Types

### RequestBody1

```go
createTemplateRequest := operations.CreateCreateTemplateRequestRequestBody1(operations.RequestBody1{/* values here */})
```

### RequestBody2

```go
createTemplateRequest := operations.CreateCreateTemplateRequestRequestBody2(operations.RequestBody2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createTemplateRequest.Type {
	case operations.CreateTemplateRequestTypeRequestBody1:
		// createTemplateRequest.RequestBody1 is populated
	case operations.CreateTemplateRequestTypeRequestBody2:
		// createTemplateRequest.RequestBody2 is populated
}
```
