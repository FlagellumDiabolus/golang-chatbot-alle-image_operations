# Use the official Rasa SDK image
FROM rasa/rasa:latest-full

# Expose the Rasa webhook port
EXPOSE 5005

# Set the working directory
WORKDIR /app

# Copy the Rasa files into the container
COPY . .

# Install additional dependencies (if any)
RUN pip install -r requirements.txt

# Start the Rasa server
CMD ["run", "-m", "models", "--enable-api", "--cors", "*", "--debug"]
