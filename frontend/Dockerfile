# Step 1: Build the React app
FROM node:18 AS build

WORKDIR /app

# Install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy source files and build the app
COPY . ./
RUN npm run build

# Step 2: Serve the app using a simple HTTP server
FROM nginx:alpine

# Copy the build folder to the Nginx HTML folder
COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port the app will run on
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
