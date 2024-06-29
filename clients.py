import socket
import threading

# Define the server address and port
SERVER_ADDRESS = '127.0.0.1'
# SERVER_ADDRESS = '8.tcp.ngrok.io'
# SERVER_ADDRESS = '3.142.129.56'
SERVER_PORT = 8080
# SERVER_PORT = 16648
NUM_CONNECTIONS = 5

# Function to handle each client connection
def handle_client(sock, message):
    try:
        # Send data to the server
        sock.sendall(message.encode())

        # Receive response from the server
        response = sock.recv(1024).decode()
        print(f"Received from server: {response}")

    except Exception as e:
        print(f"Error with client: {e}")

    finally:
        sock.close()

# Create multiple client connections
def create_clients():
    clients = []
    try:
        for _ in range(NUM_CONNECTIONS):
            # Create a new socket
            client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

            # Connect to the server
            client_socket.connect((SERVER_ADDRESS, SERVER_PORT))

            # Define the message to send
            message = f"Hello from client {_ + 1}"

            # Create a thread to handle each client connection
            thread = threading.Thread(target=handle_client, args=(client_socket, message))
            thread.start()
            clients.append(thread)

        # Wait for all threads to complete
        for thread in clients:
            thread.join()

    except Exception as e:
        print(f"Error creating clients: {e}")

# Main execution
if __name__ == "__main__":
    create_clients()
