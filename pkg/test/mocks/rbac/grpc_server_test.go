package rbac

import (
	"context"
	"fmt"
	"testing"
	"time"

	kesselAPIv2 "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestDummyGrpcServer demonstrates how to use the dummy gRPC server for testing
func TestDummyGrpcServer(t *testing.T) {
	// Create and start the dummy gRPC server
	dummyServer := NewDummyGrpcServer()
	err := dummyServer.Start("localhost:0") // Use port 0 to get a random available port
	if err != nil {
		t.Fatalf("Failed to start dummy gRPC server: %v", err)
	}
	defer dummyServer.Stop()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Create a gRPC client connection to test against the dummy server
	conn, err := grpc.NewClient(dummyServer.GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to dummy server: %v", err)
	}
	defer conn.Close()

	// Create the Kessel inventory service client
	client := kesselAPIv2.NewKesselInventoryServiceClient(conn)

	// Test the Check method
	t.Run("Check method", func(t *testing.T) {
		req := &kesselAPIv2.CheckRequest{
			Object: &kesselAPIv2.ResourceReference{
				ResourceType: "workspace",
				ResourceId:   "test-workspace-123",
				Reporter: &kesselAPIv2.ReporterReference{
					Type: "rbac",
				},
			},
			Relation: "content_sources_repositories_view",
			Subject: &kesselAPIv2.SubjectReference{
				Resource: &kesselAPIv2.ResourceReference{
					ResourceType: "principal",
					ResourceId:   "redhat/test-user",
					Reporter: &kesselAPIv2.ReporterReference{
						Type: "rbac",
					},
				},
			},
		}

		resp, err := client.Check(context.Background(), req)
		if err != nil {
			t.Fatalf("Check call failed: %v", err)
		}

		if resp.GetAllowed() != kesselAPIv2.Allowed_ALLOWED_TRUE {
			t.Errorf("Expected ALLOWED_TRUE, got %v", resp.GetAllowed())
		}
	})

	// Test the CheckForUpdate method
	t.Run("CheckForUpdate method", func(t *testing.T) {
		req := &kesselAPIv2.CheckForUpdateRequest{
			Object: &kesselAPIv2.ResourceReference{
				ResourceType: "workspace",
				ResourceId:   "test-workspace-123",
				Reporter: &kesselAPIv2.ReporterReference{
					Type: "rbac",
				},
			},
			Relation: "content_sources_repositories_edit",
			Subject: &kesselAPIv2.SubjectReference{
				Resource: &kesselAPIv2.ResourceReference{
					ResourceType: "principal",
					ResourceId:   "redhat/test-user",
					Reporter: &kesselAPIv2.ReporterReference{
						Type: "rbac",
					},
				},
			},
		}

		resp, err := client.CheckForUpdate(context.Background(), req)
		if err != nil {
			t.Fatalf("CheckForUpdate call failed: %v", err)
		}

		if resp.GetAllowed() != kesselAPIv2.Allowed_ALLOWED_TRUE {
			t.Errorf("Expected ALLOWED_TRUE, got %v", resp.GetAllowed())
		}
	})
}

// TestDummyWorkspaceServer demonstrates how to use the dummy workspace server for testing
func TestDummyWorkspaceServer(t *testing.T) {
	// Create and start the dummy workspace server
	workspaceServer := NewDummyWorkspaceServer()
	err := workspaceServer.Start("localhost:0") // Start on any available port
	if err != nil {
		t.Fatalf("Failed to start dummy workspace server: %v", err)
	}
	defer workspaceServer.Stop()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Note: In real tests, you would configure your kessel client to point to this server
	// and then test the GetRootWorkspaceID method. The dummy server will always return
	// a test workspace with ID "test-workspace-123".
	
	t.Log("Dummy workspace server is running and ready for testing")
	t.Log("Configure your kessel client to use this server for GetRootWorkspaceID tests")
}

// TestDummyKesselTestServers demonstrates how to use both services together
func TestDummyKesselTestServers(t *testing.T) {
	// Create both dummy servers
	testServers := NewDummyKesselTestServers()
	
	// Start both servers
	err := testServers.StartAll("localhost:0", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to start dummy test servers: %v", err)
	}
	defer testServers.StopAll()

	// Give the servers a moment to start
	time.Sleep(100 * time.Millisecond)

	t.Log("Both dummy servers (gRPC and HTTP) are running")
	t.Logf("gRPC server address: %s", testServers.GrpcServer.GetAddress())
	t.Log("HTTP workspace server is ready for workspace endpoint testing")
	
	// You can now configure your kessel client to use these addresses for comprehensive testing
	// The gRPC server handles Check and CheckForUpdate calls
	// The HTTP server handles GetRootWorkspaceID calls
}

// ExampleUsage demonstrates how to use the dummy servers in your own tests
func ExampleUsage() {
	// This is how you would use the dummy servers in your actual tests:
	
	// 1. Start the dummy servers
	testServers := NewDummyKesselTestServers()
	testServers.StartAll("localhost:9090", "localhost:8080")
	defer testServers.StopAll()
	
	// 2. Configure your kessel client to use these addresses
	// (You would modify your kessel client configuration here)
	
	// 3. Run your tests that use the kessel client
	// All Check, CheckForUpdate, and GetRootWorkspaceID calls will now
	// go to the dummy servers which always return "allowed" responses
	
	fmt.Println("Dummy servers configured for testing")
} 