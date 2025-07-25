# Dummy Kessel gRPC Services for Testing

This package provides dummy gRPC and HTTP services for testing the Kessel client calls without needing a real Kessel instance.

## Components

### 1. DummyKesselInventoryServiceServer (gRPC)
Implements the `KesselInventoryService` gRPC interface with dummy responses:
- **Check()** - Always returns `ALLOWED_TRUE`
- **CheckForUpdate()** - Always returns `ALLOWED_TRUE`

### 2. DummyWorkspaceServer (HTTP)
Provides an HTTP server for workspace-related endpoints:
- **GET /api/rbac/v2/workspaces** - Returns a test workspace with ID `test-workspace-123`

## Usage

### Quick Start
```go
// Start both servers for comprehensive testing
testServers := NewDummyKesselTestServers()
testServers.StartAll("localhost:9090", "localhost:8080")
defer testServers.StopAll()

// Configure your kessel client to use these addresses
// - gRPC server: localhost:9090 (for Check/CheckForUpdate)  
// - HTTP server: localhost:8080 (for GetRootWorkspaceID)
```

### Individual Services

#### gRPC Service Only
```go
grpcServer := NewDummyGrpcServer()
grpcServer.Start("localhost:9090")
defer grpcServer.Stop()

// Use grpcServer.GetAddress() to get the actual address
```

#### HTTP Workspace Service Only
```go
httpServer := NewDummyWorkspaceServer()
httpServer.Start("localhost:8080")
defer httpServer.Stop()
```

## What Gets Tested

### gRPC Calls (via DummyKesselInventoryServiceServer)
- `kessel_client.CheckRead()` → calls `Check()` → returns `ALLOWED_TRUE`
- `kessel_client.CheckWrite()` → calls `CheckForUpdate()` → returns `ALLOWED_TRUE`

### HTTP Calls (via DummyWorkspaceServer)
- `kessel_client.GetRootWorkspaceID()` → calls `/api/rbac/v2/workspaces` → returns `test-workspace-123`

## Benefits

1. **No Dependencies**: No need for a real Kessel instance
2. **Predictable**: Always returns "allowed" responses for testing happy paths
3. **Fast**: Lightweight dummy implementations
4. **Isolated**: Each test can start its own server instances
5. **Flexible**: Can run gRPC and HTTP servers independently or together

## Example Test

```go
func TestKesselClientWithDummyServer(t *testing.T) {
    // Start dummy servers
    testServers := NewDummyKesselTestServers()
    testServers.StartAll("localhost:0", "localhost:0") // Use random ports
    defer testServers.StopAll()
    
    // Configure your kessel client to use the dummy servers
    // (modify config to point to testServers.GrpcServer.GetAddress())
    
    // Your actual test code here
    // All kessel client calls will now use the dummy servers
}
```

## Notes

- The dummy services always return successful/allowed responses
- They don't perform any actual permission checking logic
- Perfect for testing the client integration without business logic
- Use `localhost:0` to get random available ports for parallel testing 