
# **Chatbot Application in Golang using Rasa**

This is a simple chatbot application written in Go, with natural language processing powered by Rasa. The application allows users to interact with the chatbot through a web interface, sending messages and receiving responses. It also supports operations for saving and retrieving images.

### **How it Works**

User Interaction: Users can interact with the chatbot by typing messages into the input field on the web interface and clicking the "Send" button. The chatbot processes the message and responds accordingly.
Message Processing: When a message is sent to the server, it is processed by the ChatHandler function. If the message is identified as a file operation (e.g., saving or retrieving images), it is handled directly. Otherwise, the message is sent to the Rasa server for natural language processing.
Rasa Integration: The SendToRasa function sends the message to the Rasa server using an HTTP POST request. Rasa processes the message using its trained machine learning models and returns a response, which is then sent back to the client.
Image Operations: If the message is identified as a file operation (e.g., "save image" or "retrieve image"), the appropriate operation is performed. For saving images, the message is processed to extract the image URL and locally given name, and the image is saved in the database. For retrieving images, the locally given name is used to retrieve the corresponding image URL from the database and display the image in the chat.
How to Run

### **To run the application locally:**

Clone the repository: git clone https://github.com/yourusername/chatbot.git
Navigate to the project directory: cd chatbot
Install dependencies: go mod tidy
Start the application: go run cmd/chatbot-api/main.go
Open a web browser and navigate to http://localhost:8080 to access the chat interface.
Ensure that you have Go installed on your system and the Rasa server is running and accessible at http://localhost:5005 for the natural language processing to work properly.