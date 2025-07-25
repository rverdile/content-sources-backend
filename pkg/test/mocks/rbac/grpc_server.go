package rbac

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	kesselAPIv2 "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta2"
	"google.golang.org/grpc"
)

// MockKesselInventoryServiceServer implements the KesselInventoryService gRPC interface
type MockKesselInventoryServiceServer struct {
	kesselAPIv2.UnimplementedKesselInventoryServiceServer
}

// Check implements the Check gRPC method - always returns allowed for testing
func (s *MockKesselInventoryServiceServer) Check(ctx context.Context, req *kesselAPIv2.CheckRequest) (*kesselAPIv2.CheckResponse, error) {
	// Always return allowed for testing purposes
	return &kesselAPIv2.CheckResponse{
		Allowed: kesselAPIv2.Allowed_ALLOWED_TRUE,
	}, nil
}

// CheckForUpdate implements the CheckForUpdate gRPC method - always returns allowed for testing
func (s *MockKesselInventoryServiceServer) CheckForUpdate(ctx context.Context, req *kesselAPIv2.CheckForUpdateRequest) (*kesselAPIv2.CheckForUpdateResponse, error) {
	// Always return allowed for testing purposes
	return &kesselAPIv2.CheckForUpdateResponse{
		Allowed: kesselAPIv2.Allowed_ALLOWED_TRUE,
	}, nil
}

// MockGrpcServer represents a Mock gRPC server for testing
type MockGrpcServer struct {
	server   *grpc.Server
	listener net.Listener
	address  string
}

// NewMockGrpcServer creates a new Mock gRPC server for testing
func NewMockGrpcServer() *MockGrpcServer {
	server := grpc.NewServer()
	
	// Register the Mock Kessel inventory service
	inventoryService := &MockKesselInventoryServiceServer{}
	kesselAPIv2.RegisterKesselInventoryServiceServer(server, inventoryService)

	return &MockGrpcServer{
		server: server,
	}
}

// Start starts the Mock gRPC server on the specified address
func (d *MockGrpcServer) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	d.listener = listener
	d.address = listener.Addr().String()

	go func() {
		d.server.Serve(listener)
	}()

	return nil
}

// Stop stops the Mock gRPC server
func (d *MockGrpcServer) Stop() {
	if d.server != nil {
		d.server.Stop()
	}
	if d.listener != nil {
		d.listener.Close()
	}
}

// GetAddress returns the address the server is listening on
func (d *MockGrpcServer) GetAddress() string {
	return d.address
}

// MockWorkspaceServer provides a Mock HTTP server for workspace endpoints
type MockWorkspaceServer struct {
	server *http.Server
	mux    *http.ServeMux
}

// WorkspaceResponse represents the workspace API response structure
type WorkspaceResponse struct {
	Data []WorkspaceData `json:"data"`
}

type WorkspaceData struct {
	ID string `json:"id"`
}

// NewMockWorkspaceServer creates a new Mock workspace HTTP server for testing
func NewMockWorkspaceServer() *MockWorkspaceServer {
	mux := http.NewServeMux()

	// Handle the workspace endpoint used by GetRootWorkspaceID
	mux.HandleFunc("/api/rbac/v2/workspaces", handleWorkspaces)

	return &MockWorkspaceServer{
		mux: mux,
	}
}

// handleWorkspaces handles workspace requests - always returns a Mock workspace for testing
func handleWorkspaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Always return a Mock workspace for testing purposes
	response := WorkspaceResponse{
		Data: []WorkspaceData{
			{ID: "test-workspace-123"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Start starts the Mock workspace HTTP server on the specified address
func (d *MockWorkspaceServer) Start(address string) error {
	d.server = &http.Server{
		Addr:    address,
		Handler: d.mux,
	}

	go func() {
		d.server.ListenAndServe()
	}()

	return nil
}

// Stop stops the Mock workspace HTTP server
func (d *MockWorkspaceServer) Stop() error {
	if d.server != nil {
		return d.server.Close()
	}
	return nil
}

// MockKesselTestServers provides both gRPC and HTTP Mock servers for testing
type MockKesselTestServers struct {
	GrpcServer      *MockGrpcServer
	WorkspaceServer *MockWorkspaceServer
}

// NewMockKesselTestServers creates both Mock servers for comprehensive testing
func NewMockKesselTestServers() *MockKesselTestServers {
	return &MockKesselTestServers{
		GrpcServer:      NewMockGrpcServer(),
		WorkspaceServer: NewMockWorkspaceServer(),
	}
}

// StartAll starts both Mock servers
func (d *MockKesselTestServers) StartAll(grpcAddress, httpAddress string) error {
	if err := d.GrpcServer.Start(grpcAddress); err != nil {
		return err
	}

	if err := d.WorkspaceServer.Start(httpAddress); err != nil {
		d.GrpcServer.Stop()
		return err
	}

	return nil
}

// StopAll stops both Mock servers
func (d *MockKesselTestServers) StopAll() {
	if d.GrpcServer != nil {
		d.GrpcServer.Stop()
	}
	if d.WorkspaceServer != nil {
		d.WorkspaceServer.Stop()
	}
}
