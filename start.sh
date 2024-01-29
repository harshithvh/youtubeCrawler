#!/bin/bash

# Backend build
cd backend
go get -u ./...
cd ..

# Frontend build
cd frontend
# Install dependencies
npm install
# Build the React app
npm run build
cd ..