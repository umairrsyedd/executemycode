# To Build 
# docker build -t umairrsyedd/executer -f executer.dockerfile .

# Use a base image with necessary dependencies
FROM --platform=linux/amd64 ubuntu:latest

# Update package lists
RUN apt-get update

# Install essential packages for Go, JavaScript, Rust, C, C++, Java
RUN apt-get install -y \
    golang \
    nodejs \
    npm \
    rustc \
    cargo \
    build-essential \
    default-jdk \
    default-jre \
    openjdk-11-jdk

# Set up a working directory
WORKDIR /app